// internal/discovery/services/discovery_service.go
package services

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	ent "github.com/UnoraApp/be/ent/generated"
	"github.com/UnoraApp/be/ent/generated/discoverybatch"
	"github.com/UnoraApp/be/ent/generated/discoverycard"
	"github.com/UnoraApp/be/ent/generated/filter"
	"github.com/UnoraApp/be/ent/generated/hobby"
	"github.com/UnoraApp/be/ent/generated/photo"
	"github.com/UnoraApp/be/ent/generated/profile"
	"github.com/UnoraApp/be/ent/generated/server"
	"github.com/UnoraApp/be/ent/generated/user"
	"github.com/UnoraApp/be/internal/discovery/dto"
	"github.com/UnoraApp/be/pkg/storage"
)

// DiscoveryService handles discovery-related business logic
type DiscoveryService struct {
	entClient     *ent.Client
	storageClient storage.Client
}

// NewDiscoveryService creates a new discovery service
func NewDiscoveryService(entClient *ent.Client, storageClient storage.Client) *DiscoveryService {
	return &DiscoveryService{
		entClient:     entClient,
		storageClient: storageClient,
	}
}

// GetServers returns all available server types
func (s *DiscoveryService) GetServers(ctx context.Context) ([]*dto.ServerResponse, error) {
	servers, err := s.entClient.Server.
		Query().
		Order(ent.Asc(server.FieldSortOrder)).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get servers: %w", err)
	}

	result := make([]*dto.ServerResponse, len(servers))
	for i, srv := range servers {
		result[i] = &dto.ServerResponse{
			ID:          srv.ID,
			ServerType:  string(srv.ServerType),
			DisplayName: srv.DisplayName,
			IconName:    srv.IconName,
			SortOrder:   srv.SortOrder,
		}
	}

	return result, nil
}

// GetFilter gets user's filter for a server type
func (s *DiscoveryService) GetFilter(ctx context.Context, userID, serverType string) (*dto.FilterResponse, error) {
	f, err := s.entClient.Filter.
		Query().
		Where(filter.UserIDEQ(userID)).
		Where(filter.ServerTypeEQ(filter.ServerType(serverType))).
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			// Return default empty filter
			return &dto.FilterResponse{
				ServerType: serverType,
				Config:     dto.FilterConfig{},
				IsPending:  false,
				UpdatedAt:  time.Now(),
			}, nil
		}
		return nil, fmt.Errorf("failed to get filter: %w", err)
	}

	return &dto.FilterResponse{
		ID:         f.ID,
		ServerType: string(f.ServerType),
		Config:     mapToFilterConfig(f.FilterConfig),
		IsPending:  f.IsPending,
		UpdatedAt:  f.UpdatedAt,
	}, nil
}

// UpdateFilter updates or creates user's filter for a server type
func (s *DiscoveryService) UpdateFilter(ctx context.Context, userID, serverType string, req *dto.UpdateFilterRequest) (*dto.FilterResponse, error) {
	// Check if filter exists
	f, err := s.entClient.Filter.
		Query().
		Where(filter.UserIDEQ(userID)).
		Where(filter.ServerTypeEQ(filter.ServerType(serverType))).
		Only(ctx)

	configMap := filterConfigToMap(req.Config)

	if ent.IsNotFound(err) {
		// Create new filter
		filterID := uuid.New().String()
		f, err = s.entClient.Filter.
			Create().
			SetID(filterID).
			SetUserID(userID).
			SetServerType(filter.ServerType(serverType)).
			SetFilterConfig(configMap).
			SetIsPending(true).
			Save(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to create filter: %w", err)
		}
	} else if err != nil {
		return nil, fmt.Errorf("failed to get filter: %w", err)
	} else {
		// Update existing filter
		f, err = f.Update().
			SetFilterConfig(configMap).
			SetIsPending(true).
			Save(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to update filter: %w", err)
		}
	}

	return &dto.FilterResponse{
		ID:         f.ID,
		ServerType: string(f.ServerType),
		Config:     req.Config,
		IsPending:  f.IsPending,
		UpdatedAt:  f.UpdatedAt,
	}, nil
}

// GetDiscoveryBatch gets or creates the current discovery batch for a server type
func (s *DiscoveryService) GetDiscoveryBatch(ctx context.Context, userID, serverType string) (*dto.DiscoveryBatchResponse, error) {
	// Try to find active batch
	batch, err := s.entClient.DiscoveryBatch.
		Query().
		Where(discoverybatch.UserIDEQ(userID)).
		Where(discoverybatch.ServerTypeEQ(discoverybatch.ServerType(serverType))).
		Where(discoverybatch.BatchStatusEQ(discoverybatch.BatchStatusActive)).
		Where(discoverybatch.DeletedAtIsNil()).
		WithCards(func(q *ent.DiscoveryCardQuery) {
			q.Where(discoverycard.DeletedAtIsNil()).
				Order(ent.Asc(discoverycard.FieldDisplayOrder))
		}).
		Only(ctx)

	if ent.IsNotFound(err) {
		// Generate new batch
		return s.generateNewBatch(ctx, userID, serverType)
	} else if err != nil {
		return nil, fmt.Errorf("failed to get batch: %w", err)
	}

	return s.batchToResponse(ctx, batch)
}

// RefreshDiscovery triggers a new discovery batch (respects global cooldown and tier limits)
func (s *DiscoveryService) RefreshDiscovery(ctx context.Context, userID, serverType string) (*dto.DiscoveryBatchResponse, error) {
	// Check refresh availability
	u, err := s.entClient.User.Get(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	tier := string(u.SubscriptionTier)
	tierConfig := getTierConfig(tier)

	// Check if at connection capacity (Hard Lock)
	if u.ActiveConnectionCount >= tierConfig.ConnectionSlots {
		return nil, fmt.Errorf("at connection capacity (%d/%d slots). Focus on your current connections or upgrade", u.ActiveConnectionCount, tierConfig.ConnectionSlots)
	}

	// Check refresh cooldown
	if u.RefreshAvailableAt != nil && time.Now().Before(*u.RefreshAvailableAt) {
		return nil, fmt.Errorf("refresh not available yet")
	}

	// Mark existing batch as consumed
	_, err = s.entClient.DiscoveryBatch.
		Update().
		Where(discoverybatch.UserIDEQ(userID)).
		Where(discoverybatch.ServerTypeEQ(discoverybatch.ServerType(serverType))).
		Where(discoverybatch.BatchStatusEQ(discoverybatch.BatchStatusActive)).
		SetBatchStatus(discoverybatch.BatchStatusConsumed).
		Save(ctx)
	if err != nil {
		// Ignore if no batch to update
	}

	// Set next refresh time based on tier
	nextRefresh := time.Now().Add(tierConfig.RefreshCooldown)
	_, err = u.Update().
		SetRefreshAvailableAt(nextRefresh).
		SetLastGlobalRefreshAt(time.Now()).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to update refresh time: %w", err)
	}

	// Generate new batch
	return s.generateNewBatch(ctx, userID, serverType)
}

// TierConfig contains tier-specific limits
type tierConfigInternal struct {
	ConnectionSlots int
	RefreshCooldown time.Duration
}

func getTierConfig(tier string) tierConfigInternal {
	configs := map[string]tierConfigInternal{
		"free": {1, 24 * time.Hour},
		"plus": {2, 12 * time.Hour},
		"pro":  {4, 6 * time.Hour},
	}
	if c, ok := configs[tier]; ok {
		return c
	}
	return configs["free"]
}

// GetRefreshStatus returns the refresh availability status
func (s *DiscoveryService) GetRefreshStatus(ctx context.Context, userID string) (*dto.RefreshStatusResponse, error) {
	u, err := s.entClient.User.Get(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	if u.RefreshAvailableAt == nil || time.Now().After(*u.RefreshAvailableAt) {
		return &dto.RefreshStatusResponse{
			CanRefresh: true,
		}, nil
	}

	seconds := int(time.Until(*u.RefreshAvailableAt).Seconds())
	return &dto.RefreshStatusResponse{
		CanRefresh:          false,
		RefreshAvailableAt:  u.RefreshAvailableAt,
		SecondsUntilRefresh: &seconds,
	}, nil
}

// generateNewBatch creates a new discovery batch with 5 candidate cards
func (s *DiscoveryService) generateNewBatch(ctx context.Context, userID, serverType string) (*dto.DiscoveryBatchResponse, error) {
	// Find candidate users (simple version - just get random users except self)
	// In production, this would use filters, exclude blocked users, previous matches, etc.
	candidates, err := s.entClient.User.
		Query().
		Where(user.IDNEQ(userID)).
		Where(user.DeletedAtIsNil()).
		Limit(5).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to find candidates: %w", err)
	}

	// Create batch
	batchID := uuid.New().String()
	batch, err := s.entClient.DiscoveryBatch.
		Create().
		SetID(batchID).
		SetUserID(userID).
		SetServerType(discoverybatch.ServerType(serverType)).
		SetBatchStatus(discoverybatch.BatchStatusActive).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create batch: %w", err)
	}

	// Create cards for each candidate
	for i, candidate := range candidates {
		cardID := uuid.New().String()
		_, err = s.entClient.DiscoveryCard.
			Create().
			SetID(cardID).
			SetBatchID(batchID).
			SetCandidateUserID(candidate.ID).
			SetDisplayOrder(i + 1).
			Save(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to create card: %w", err)
		}
	}

	// Reload batch with cards
	batch, err = s.entClient.DiscoveryBatch.
		Query().
		Where(discoverybatch.IDEQ(batchID)).
		WithCards(func(q *ent.DiscoveryCardQuery) {
			q.Order(ent.Asc(discoverycard.FieldDisplayOrder))
		}).
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to reload batch: %w", err)
	}

	return s.batchToResponse(ctx, batch)
}

func (s *DiscoveryService) batchToResponse(ctx context.Context, batch *ent.DiscoveryBatch) (*dto.DiscoveryBatchResponse, error) {
	cards := make([]dto.DiscoveryCardResponse, len(batch.Edges.Cards))

	for i, card := range batch.Edges.Cards {
		cardResp, err := s.cardToResponse(ctx, card)
		if err != nil {
			return nil, err
		}
		cards[i] = *cardResp
	}

	return &dto.DiscoveryBatchResponse{
		ID:          batch.ID,
		ServerType:  string(batch.ServerType),
		BatchStatus: string(batch.BatchStatus),
		Cards:       cards,
		ExpiresAt:   batch.ExpiresAt,
		CreatedAt:   batch.CreatedAt,
	}, nil
}

func (s *DiscoveryService) cardToResponse(ctx context.Context, card *ent.DiscoveryCard) (*dto.DiscoveryCardResponse, error) {
	// Get candidate user with profile
	candidate, err := s.entClient.User.
		Query().
		Where(user.IDEQ(card.CandidateUserID)).
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get candidate: %w", err)
	}

	// Get profile
	p, _ := s.entClient.Profile.
		Query().
		Where(profile.UserIDEQ(card.CandidateUserID)).
		Where(profile.DeletedAtIsNil()).
		Only(ctx)

	// Get photos
	photos, _ := s.entClient.Photo.
		Query().
		Where(photo.UserIDEQ(card.CandidateUserID)).
		Where(photo.DeletedAtIsNil()).
		Order(ent.Asc(photo.FieldDisplayOrder)).
		All(ctx)

	// Get hobbies
	hobbies, _ := s.entClient.Hobby.
		Query().
		Where(hobby.UserIDEQ(card.CandidateUserID)).
		Where(hobby.DeletedAtIsNil()).
		WithHobbyOption().
		Order(ent.Asc(hobby.FieldDisplayOrder)).
		All(ctx)

	// Build user info
	userInfo := dto.CardUserInfo{
		ID:        candidate.ID,
		FirstName: ptrToString(candidate.FirstName),
	}

	if p != nil {
		userInfo.City = ptrToString(p.City)
		userInfo.Bio = ptrToString(p.Bio)
		userInfo.IntentStatement = ptrToString(p.IntentStatement)
		if p.DateOfBirth != nil {
			userInfo.Age = calculateAge(*p.DateOfBirth)
		}
	}

	// Build photos
	cardPhotos := make([]dto.CardPhoto, len(photos))
	for i, ph := range photos {
		url, _ := s.storageClient.GetPresignedDownloadURL(ctx, ph.StorageKey)
		cardPhotos[i] = dto.CardPhoto{
			URL:          url,
			DisplayOrder: ph.DisplayOrder,
			IsFacePhoto:  ph.IsFacePhoto,
		}
	}

	// Build hobbies
	cardHobbies := make([]dto.CardHobby, len(hobbies))
	for i, h := range hobbies {
		opt := h.Edges.HobbyOption
		cardHobbies[i] = dto.CardHobby{
			Name:             opt.Name,
			Category:         opt.Category,
			IconName:         ptrToString(opt.IconName),
			MicroDescription: ptrToString(h.MicroDescription),
		}
	}

	return &dto.DiscoveryCardResponse{
		ID:           card.ID,
		DisplayOrder: card.DisplayOrder,
		User:         userInfo,
		Photos:       cardPhotos,
		Hobbies:      cardHobbies,
		CreatedAt:    card.CreatedAt,
	}, nil
}

// Helper functions
func ptrToString(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func calculateAge(dob time.Time) int {
	now := time.Now()
	age := now.Year() - dob.Year()
	if now.YearDay() < dob.YearDay() {
		age--
	}
	return age
}

func mapToFilterConfig(m map[string]interface{}) dto.FilterConfig {
	config := dto.FilterConfig{}
	if m == nil {
		return config
	}
	// Convert map to FilterConfig (simplified)
	if v, ok := m["ageMin"].(float64); ok {
		i := int(v)
		config.AgeMin = &i
	}
	if v, ok := m["ageMax"].(float64); ok {
		i := int(v)
		config.AgeMax = &i
	}
	// Add more conversions as needed
	return config
}

func filterConfigToMap(config dto.FilterConfig) map[string]interface{} {
	m := make(map[string]interface{})
	if config.AgeMin != nil {
		m["ageMin"] = *config.AgeMin
	}
	if config.AgeMax != nil {
		m["ageMax"] = *config.AgeMax
	}
	if config.Genders != nil {
		m["genders"] = config.Genders
	}
	if config.Cities != nil {
		m["cities"] = config.Cities
	}
	if config.Hobbies != nil {
		m["hobbies"] = config.Hobbies
	}
	if config.Distance != nil {
		m["distance"] = *config.Distance
	}
	return m
}
