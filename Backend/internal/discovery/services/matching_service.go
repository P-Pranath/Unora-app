// internal/discovery/services/matching_service.go
package services

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	ent "github.com/UnoraApp/be/ent/generated"
	"github.com/UnoraApp/be/ent/generated/connection"
	"github.com/UnoraApp/be/ent/generated/interest"
	"github.com/UnoraApp/be/ent/generated/photo"
	"github.com/UnoraApp/be/ent/generated/profile"
	"github.com/UnoraApp/be/ent/generated/streak"
	"github.com/UnoraApp/be/internal/discovery/dto"
	"github.com/UnoraApp/be/pkg/storage"
)

// MatchingService handles interest expression and connection management
type MatchingService struct {
	entClient     *ent.Client
	storageClient storage.Client
}

// NewMatchingService creates a new matching service
func NewMatchingService(entClient *ent.Client, storageClient storage.Client) *MatchingService {
	return &MatchingService{
		entClient:     entClient,
		storageClient: storageClient,
	}
}

// ExpressInterest expresses interest in a user from a discovery card
func (s *MatchingService) ExpressInterest(ctx context.Context, senderUserID string, req *dto.ExpressInterestRequest) (*dto.InterestResponse, error) {
	// Get the discovery card to find receiver
	card, err := s.entClient.DiscoveryCard.Get(ctx, req.DiscoveryCardID)
	if err != nil {
		return nil, fmt.Errorf("discovery card not found: %w", err)
	}

	// Get the batch to find server type
	batch, err := s.entClient.DiscoveryBatch.Get(ctx, card.BatchID)
	if err != nil {
		return nil, fmt.Errorf("batch not found: %w", err)
	}

	receiverUserID := card.CandidateUserID
	serverType := string(batch.ServerType)

	// Check if already expressed interest
	exists, _ := s.entClient.Interest.
		Query().
		Where(interest.SenderUserIDEQ(senderUserID)).
		Where(interest.ReceiverUserIDEQ(receiverUserID)).
		Where(interest.ServerTypeEQ(interest.ServerType(serverType))).
		Where(interest.InterestStatusEQ(interest.InterestStatusPending)).
		Where(interest.DeletedAtIsNil()).
		Exist(ctx)
	if exists {
		return nil, fmt.Errorf("interest already expressed")
	}

	// Check for mutual interest (receiver already interested in sender)
	mutualInterest, err := s.entClient.Interest.
		Query().
		Where(interest.SenderUserIDEQ(receiverUserID)).
		Where(interest.ReceiverUserIDEQ(senderUserID)).
		Where(interest.ServerTypeEQ(interest.ServerType(serverType))).
		Where(interest.InterestStatusEQ(interest.InterestStatusPending)).
		Where(interest.DeletedAtIsNil()).
		Only(ctx)

	interestID := uuid.New().String()
	now := time.Now()

	if ent.IsNotFound(err) {
		// No mutual interest - create pending interest
		i, err := s.entClient.Interest.
			Create().
			SetID(interestID).
			SetSenderUserID(senderUserID).
			SetReceiverUserID(receiverUserID).
			SetServerType(interest.ServerType(serverType)).
			SetDiscoveryCardID(req.DiscoveryCardID).
			SetInterestStatus(interest.InterestStatusPending).
			Save(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to create interest: %w", err)
		}

		return s.interestToResponse(ctx, i, receiverUserID)
	} else if err != nil {
		return nil, fmt.Errorf("failed to check mutual interest: %w", err)
	}

	// Mutual interest found! Create connection
	_, err = s.createConnection(ctx, senderUserID, receiverUserID, serverType)
	if err != nil {
		return nil, fmt.Errorf("failed to create connection: %w", err)
	}

	// Update both interests to matched
	_, err = mutualInterest.Update().
		SetInterestStatus(interest.InterestStatusMatched).
		SetMatchedAt(now).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to update mutual interest: %w", err)
	}

	// Create the new interest as matched
	i, err := s.entClient.Interest.
		Create().
		SetID(interestID).
		SetSenderUserID(senderUserID).
		SetReceiverUserID(receiverUserID).
		SetServerType(interest.ServerType(serverType)).
		SetDiscoveryCardID(req.DiscoveryCardID).
		SetInterestStatus(interest.InterestStatusMatched).
		SetMatchedAt(now).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create matched interest: %w", err)
	}

	return s.interestToResponse(ctx, i, receiverUserID)
}

// GetSentInterests returns interests sent by the user
func (s *MatchingService) GetSentInterests(ctx context.Context, userID string) ([]*dto.InterestResponse, error) {
	interests, err := s.entClient.Interest.
		Query().
		Where(interest.SenderUserIDEQ(userID)).
		Where(interest.DeletedAtIsNil()).
		Order(ent.Desc(interest.FieldCreatedAt)).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get sent interests: %w", err)
	}

	result := make([]*dto.InterestResponse, len(interests))
	for i, intr := range interests {
		resp, err := s.interestToResponse(ctx, intr, intr.ReceiverUserID)
		if err != nil {
			continue
		}
		result[i] = resp
	}

	return result, nil
}

// GetReceivedInterests returns interests received by the user
func (s *MatchingService) GetReceivedInterests(ctx context.Context, userID string) ([]*dto.InterestResponse, error) {
	interests, err := s.entClient.Interest.
		Query().
		Where(interest.ReceiverUserIDEQ(userID)).
		Where(interest.InterestStatusEQ(interest.InterestStatusPending)).
		Where(interest.DeletedAtIsNil()).
		Order(ent.Desc(interest.FieldCreatedAt)).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get received interests: %w", err)
	}

	result := make([]*dto.InterestResponse, len(interests))
	for i, intr := range interests {
		resp, err := s.interestToResponse(ctx, intr, intr.SenderUserID)
		if err != nil {
			continue
		}
		result[i] = resp
	}

	return result, nil
}

// GetConnections returns all connections for a user
func (s *MatchingService) GetConnections(ctx context.Context, userID string) ([]*dto.ConnectionResponse, error) {
	// Find connections where user is either A or B
	connections, err := s.entClient.Connection.
		Query().
		Where(
			connection.Or(
				connection.UserAIDEQ(userID),
				connection.UserBIDEQ(userID),
			),
		).
		Where(connection.ConnectionStatusEQ(connection.ConnectionStatusActive)).
		Where(connection.DeletedAtIsNil()).
		WithStreak().
		Order(ent.Desc(connection.FieldCreatedAt)).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get connections: %w", err)
	}

	result := make([]*dto.ConnectionResponse, len(connections))
	for i, conn := range connections {
		resp, err := s.connectionToResponse(ctx, conn, userID)
		if err != nil {
			continue
		}
		result[i] = resp
	}

	return result, nil
}

// GetConnection returns a specific connection
func (s *MatchingService) GetConnection(ctx context.Context, userID, connectionID string) (*dto.ConnectionDetailResponse, error) {
	conn, err := s.entClient.Connection.
		Query().
		Where(connection.IDEQ(connectionID)).
		Where(
			connection.Or(
				connection.UserAIDEQ(userID),
				connection.UserBIDEQ(userID),
			),
		).
		Where(connection.DeletedAtIsNil()).
		WithStreak().
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, fmt.Errorf("connection not found")
		}
		return nil, fmt.Errorf("failed to get connection: %w", err)
	}

	basic, err := s.connectionToResponse(ctx, conn, userID)
	if err != nil {
		return nil, err
	}

	return &dto.ConnectionDetailResponse{
		ConnectionResponse: *basic,
	}, nil
}

// TerminateConnection terminates a connection
func (s *MatchingService) TerminateConnection(ctx context.Context, userID, connectionID string) error {
	_, err := s.entClient.Connection.
		Update().
		Where(connection.IDEQ(connectionID)).
		Where(
			connection.Or(
				connection.UserAIDEQ(userID),
				connection.UserBIDEQ(userID),
			),
		).
		SetConnectionStatus(connection.ConnectionStatusTerminated).
		SetTerminatedAt(time.Now()).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("failed to terminate connection: %w", err)
	}
	return nil
}

// createConnection creates a new connection between two users
func (s *MatchingService) createConnection(ctx context.Context, userA, userB, serverType string) (*ent.Connection, error) {
	// Ensure userA < userB for uniqueness
	if userA > userB {
		userA, userB = userB, userA
	}

	connID := uuid.New().String()
	conn, err := s.entClient.Connection.
		Create().
		SetID(connID).
		SetUserAID(userA).
		SetUserBID(userB).
		SetServerType(connection.ServerType(serverType)).
		SetConnectionStatus(connection.ConnectionStatusActive).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create connection: %w", err)
	}

	// Create associated streak
	streakID := uuid.New().String()
	_, err = s.entClient.Streak.
		Create().
		SetID(streakID).
		SetConnectionID(connID).
		SetStreakState(streak.StreakStateActive).
		SetCurrentDay(1).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create streak: %w", err)
	}

	// Increment active connection count for both users
	s.incrementConnectionCount(ctx, userA)
	s.incrementConnectionCount(ctx, userB)

	// Execute Total Wipe for both users if at capacity
	s.executeTotalWipeIfAtCapacity(ctx, userA)
	s.executeTotalWipeIfAtCapacity(ctx, userB)

	return conn, nil
}

// incrementConnectionCount increments a user's active connection count
func (s *MatchingService) incrementConnectionCount(ctx context.Context, userID string) {
	u, err := s.entClient.User.Get(ctx, userID)
	if err != nil {
		return
	}
	u.Update().SetActiveConnectionCount(u.ActiveConnectionCount + 1).Save(ctx)
}

// executeTotalWipeIfAtCapacity deletes all pending interests if user is at capacity
func (s *MatchingService) executeTotalWipeIfAtCapacity(ctx context.Context, userID string) {
	u, err := s.entClient.User.Get(ctx, userID)
	if err != nil {
		return
	}

	tier := string(u.SubscriptionTier)
	tierConfig := matchingGetTierConfig(tier)

	// Check if at capacity
	if u.ActiveConnectionCount >= tierConfig.ConnectionSlots {
		// Total Wipe: Delete all pending outgoing interests
		s.entClient.Interest.
			Update().
			Where(interest.SenderUserIDEQ(userID)).
			Where(interest.InterestStatusEQ(interest.InterestStatusPending)).
			SetDeletedAt(time.Now()).
			Save(ctx)
	}
}

// matchingGetTierConfig for matching service
func matchingGetTierConfig(tier string) struct {
	ConnectionSlots int
	RefreshCooldown time.Duration
} {
	configs := map[string]struct {
		ConnectionSlots int
		RefreshCooldown time.Duration
	}{
		"free": {1, 24 * time.Hour},
		"plus": {2, 12 * time.Hour},
		"pro":  {4, 6 * time.Hour},
	}
	if c, ok := configs[tier]; ok {
		return c
	}
	return configs["free"]
}

func (s *MatchingService) interestToResponse(ctx context.Context, i *ent.Interest, otherUserID string) (*dto.InterestResponse, error) {
	userInfo, err := s.getUserInfo(ctx, otherUserID)
	if err != nil {
		return nil, err
	}

	return &dto.InterestResponse{
		ID:         i.ID,
		ServerType: string(i.ServerType),
		Status:     string(i.InterestStatus),
		User:       *userInfo,
		CreatedAt:  i.CreatedAt,
		MatchedAt:  i.MatchedAt,
	}, nil
}

func (s *MatchingService) connectionToResponse(ctx context.Context, conn *ent.Connection, currentUserID string) (*dto.ConnectionResponse, error) {
	// Determine partner ID
	partnerID := conn.UserAID
	if conn.UserAID == currentUserID {
		partnerID = conn.UserBID
	}

	partner, err := s.getPartnerInfo(ctx, partnerID)
	if err != nil {
		return nil, err
	}

	resp := &dto.ConnectionResponse{
		ID:         conn.ID,
		ServerType: string(conn.ServerType),
		Status:     string(conn.ConnectionStatus),
		Partner:    *partner,
		CreatedAt:  conn.CreatedAt,
	}

	if conn.Edges.Streak != nil {
		st := conn.Edges.Streak
		resp.Streak = &dto.StreakSummary{
			ID:           st.ID,
			CurrentDay:   st.CurrentDay,
			State:        string(st.StreakState),
			NeedsCheckIn: true, // Simplified - would check actual check-ins
		}
	}

	return resp, nil
}

func (s *MatchingService) getUserInfo(ctx context.Context, userID string) (*dto.InterestUserInfo, error) {
	u, err := s.entClient.User.Get(ctx, userID)
	if err != nil {
		return nil, err
	}

	p, _ := s.entClient.Profile.
		Query().
		Where(profile.UserIDEQ(userID)).
		Where(profile.DeletedAtIsNil()).
		Only(ctx)

	ph, _ := s.entClient.Photo.
		Query().
		Where(photo.UserIDEQ(userID)).
		Where(photo.DeletedAtIsNil()).
		Order(ent.Asc(photo.FieldDisplayOrder)).
		First(ctx)

	info := &dto.InterestUserInfo{
		ID:        u.ID,
		FirstName: ptrToString(u.FirstName),
	}

	if p != nil {
		info.City = ptrToString(p.City)
		if p.DateOfBirth != nil {
			info.Age = calculateAge(*p.DateOfBirth)
		}
	}

	if ph != nil {
		info.PhotoURL, _ = s.storageClient.GetPresignedDownloadURL(ctx, ph.StorageKey)
	}

	return info, nil
}

func (s *MatchingService) getPartnerInfo(ctx context.Context, userID string) (*dto.ConnectionPartner, error) {
	u, err := s.entClient.User.Get(ctx, userID)
	if err != nil {
		return nil, err
	}

	p, _ := s.entClient.Profile.
		Query().
		Where(profile.UserIDEQ(userID)).
		Where(profile.DeletedAtIsNil()).
		Only(ctx)

	ph, _ := s.entClient.Photo.
		Query().
		Where(photo.UserIDEQ(userID)).
		Where(photo.DeletedAtIsNil()).
		Order(ent.Asc(photo.FieldDisplayOrder)).
		First(ctx)

	partner := &dto.ConnectionPartner{
		ID:        u.ID,
		FirstName: ptrToString(u.FirstName),
	}

	if p != nil {
		partner.City = ptrToString(p.City)
		partner.Bio = ptrToString(p.Bio)
		if p.DateOfBirth != nil {
			partner.Age = calculateAge(*p.DateOfBirth)
		}
	}

	if ph != nil {
		partner.PhotoURL, _ = s.storageClient.GetPresignedDownloadURL(ctx, ph.StorageKey)
	}

	return partner, nil
}
