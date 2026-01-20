// internal/profile/services/profile_service.go
package services

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	ent "github.com/UnoraApp/be/ent/generated"
	"github.com/UnoraApp/be/ent/generated/hobby"
	"github.com/UnoraApp/be/ent/generated/hobbyoption"
	"github.com/UnoraApp/be/ent/generated/photo"
	"github.com/UnoraApp/be/ent/generated/profile"
	"github.com/UnoraApp/be/internal/profile/dto"
	"github.com/UnoraApp/be/pkg/storage"
)

// ProfileService handles profile-related business logic
type ProfileService struct {
	entClient     *ent.Client
	storageClient storage.Client
}

// NewProfileService creates a new profile service
func NewProfileService(entClient *ent.Client, storageClient storage.Client) *ProfileService {
	return &ProfileService{
		entClient:     entClient,
		storageClient: storageClient,
	}
}

// CreateProfile creates a new profile for a user
func (s *ProfileService) CreateProfile(ctx context.Context, userID string, req *dto.CreateProfileRequest) (*dto.ProfileResponse, error) {
	// Check if profile already exists
	exists, err := s.entClient.Profile.
		Query().
		Where(profile.UserIDEQ(userID)).
		Where(profile.DeletedAtIsNil()).
		Exist(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to check existing profile: %w", err)
	}
	if exists {
		return nil, fmt.Errorf("profile already exists for this user")
	}

	// Parse date of birth
	dob, err := time.Parse("2006-01-02", req.DateOfBirth)
	if err != nil {
		return nil, fmt.Errorf("invalid date of birth format, expected YYYY-MM-DD: %w", err)
	}

	// Create profile
	profileID := uuid.New().String()
	p, err := s.entClient.Profile.
		Create().
		SetID(profileID).
		SetUserID(userID).
		SetFirstName(req.FirstName).
		SetNillableLastName(strPtr(req.LastName)).
		SetDateOfBirth(dob).
		SetGender(req.Gender).
		SetCity(req.City).
		SetNillableBio(strPtr(req.Bio)).
		SetNillableIntentStatement(strPtr(req.IntentStatement)).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create profile: %w", err)
	}

	return profileToResponse(p), nil
}

// GetProfile gets the current user's profile
func (s *ProfileService) GetProfile(ctx context.Context, userID string) (*dto.ProfileResponse, error) {
	p, err := s.entClient.Profile.
		Query().
		Where(profile.UserIDEQ(userID)).
		Where(profile.DeletedAtIsNil()).
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, fmt.Errorf("profile not found")
		}
		return nil, fmt.Errorf("failed to get profile: %w", err)
	}

	return profileToResponse(p), nil
}

// UpdateProfile updates the user's profile
func (s *ProfileService) UpdateProfile(ctx context.Context, userID string, req *dto.UpdateProfileRequest) (*dto.ProfileResponse, error) {
	p, err := s.entClient.Profile.
		Query().
		Where(profile.UserIDEQ(userID)).
		Where(profile.DeletedAtIsNil()).
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, fmt.Errorf("profile not found")
		}
		return nil, fmt.Errorf("failed to get profile: %w", err)
	}

	// Build update
	update := p.Update()

	if req.FirstName != nil {
		update.SetFirstName(*req.FirstName)
	}
	if req.LastName != nil {
		update.SetLastName(*req.LastName)
	}
	if req.DateOfBirth != nil {
		dob, err := time.Parse("2006-01-02", *req.DateOfBirth)
		if err != nil {
			return nil, fmt.Errorf("invalid date of birth format: %w", err)
		}
		update.SetDateOfBirth(dob)
	}
	if req.Gender != nil {
		update.SetGender(*req.Gender)
	}
	if req.City != nil {
		update.SetCity(*req.City)
	}
	if req.Bio != nil {
		update.SetBio(*req.Bio)
	}
	if req.IntentStatement != nil {
		update.SetIntentStatement(*req.IntentStatement)
	}

	updatedProfile, err := update.Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to update profile: %w", err)
	}

	return profileToResponse(updatedProfile), nil
}

// DeleteProfile soft deletes the user's profile
func (s *ProfileService) DeleteProfile(ctx context.Context, userID string) error {
	_, err := s.entClient.Profile.
		Update().
		Where(profile.UserIDEQ(userID)).
		Where(profile.DeletedAtIsNil()).
		SetDeletedAt(time.Now()).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("failed to delete profile: %w", err)
	}
	return nil
}

// GetOnboardingStatus returns the onboarding completion status
func (s *ProfileService) GetOnboardingStatus(ctx context.Context, userID string) (*dto.OnboardingStatusResponse, error) {
	// Check profile
	profileExists, _ := s.entClient.Profile.
		Query().
		Where(profile.UserIDEQ(userID)).
		Where(profile.DeletedAtIsNil()).
		Exist(ctx)

	// Count photos
	photoCount, _ := s.entClient.Photo.
		Query().
		Where(photo.UserIDEQ(userID)).
		Where(photo.DeletedAtIsNil()).
		Count(ctx)

	// Count hobbies
	hobbyCount, _ := s.entClient.Hobby.
		Query().
		Where(hobby.UserIDEQ(userID)).
		Where(hobby.DeletedAtIsNil()).
		Count(ctx)

	const photosRequired = 3
	const hobbiesRequired = 2

	status := &dto.OnboardingStatusResponse{
		ProfileComplete: profileExists,
		PhotosComplete:  photoCount >= photosRequired,
		PhotoCount:      photoCount,
		PhotosRequired:  photosRequired,
		HobbiesComplete: hobbyCount >= hobbiesRequired,
		HobbyCount:      hobbyCount,
		HobbiesRequired: hobbiesRequired,
	}
	status.IsComplete = status.ProfileComplete && status.PhotosComplete && status.HobbiesComplete

	return status, nil
}

// GetHobbyOptions returns all active hobby options
func (s *ProfileService) GetHobbyOptions(ctx context.Context, category string) ([]*dto.HobbyOptionResponse, error) {
	query := s.entClient.HobbyOption.
		Query().
		Where(hobbyoption.IsActiveEQ(true)).
		Order(ent.Asc(hobbyoption.FieldSortOrder))

	if category != "" {
		query = query.Where(hobbyoption.CategoryEQ(category))
	}

	options, err := query.All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get hobby options: %w", err)
	}

	result := make([]*dto.HobbyOptionResponse, len(options))
	for i, opt := range options {
		result[i] = &dto.HobbyOptionResponse{
			ID:                opt.ID,
			Name:              opt.Name,
			Category:          opt.Category,
			IconName:          ptrToString(opt.IconName),
			MicroDescriptions: opt.MicroDescriptions,
		}
	}

	return result, nil
}

// AddHobby adds a hobby to user's profile
func (s *ProfileService) AddHobby(ctx context.Context, userID string, req *dto.AddHobbyRequest) (*dto.HobbyResponse, error) {
	// Verify hobby option exists
	opt, err := s.entClient.HobbyOption.Get(ctx, req.HobbyOptionID)
	if err != nil {
		return nil, fmt.Errorf("hobby option not found: %w", err)
	}

	// Check if already added
	exists, _ := s.entClient.Hobby.
		Query().
		Where(hobby.UserIDEQ(userID)).
		Where(hobby.HobbyOptionIDEQ(req.HobbyOptionID)).
		Where(hobby.DeletedAtIsNil()).
		Exist(ctx)
	if exists {
		return nil, fmt.Errorf("hobby already added to profile")
	}

	// Create hobby
	hobbyID := uuid.New().String()
	h, err := s.entClient.Hobby.
		Create().
		SetID(hobbyID).
		SetUserID(userID).
		SetHobbyOptionID(req.HobbyOptionID).
		SetNillableMicroDescription(strPtr(req.MicroDescription)).
		SetDisplayOrder(req.DisplayOrder).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to add hobby: %w", err)
	}

	return &dto.HobbyResponse{
		ID: h.ID,
		HobbyOption: dto.HobbyOptionResponse{
			ID:                opt.ID,
			Name:              opt.Name,
			Category:          opt.Category,
			IconName:          ptrToString(opt.IconName),
			MicroDescriptions: opt.MicroDescriptions,
		},
		MicroDescription: ptrToString(h.MicroDescription),
		DisplayOrder:     h.DisplayOrder,
		CreatedAt:        h.CreatedAt,
	}, nil
}

// GetUserHobbies returns user's hobbies
func (s *ProfileService) GetUserHobbies(ctx context.Context, userID string) ([]*dto.HobbyResponse, error) {
	hobbies, err := s.entClient.Hobby.
		Query().
		Where(hobby.UserIDEQ(userID)).
		Where(hobby.DeletedAtIsNil()).
		WithHobbyOption().
		Order(ent.Asc(hobby.FieldDisplayOrder)).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get hobbies: %w", err)
	}

	result := make([]*dto.HobbyResponse, len(hobbies))
	for i, h := range hobbies {
		opt := h.Edges.HobbyOption
		result[i] = &dto.HobbyResponse{
			ID: h.ID,
			HobbyOption: dto.HobbyOptionResponse{
				ID:                opt.ID,
				Name:              opt.Name,
				Category:          opt.Category,
				IconName:          ptrToString(opt.IconName),
				MicroDescriptions: opt.MicroDescriptions,
			},
			MicroDescription: ptrToString(h.MicroDescription),
			DisplayOrder:     h.DisplayOrder,
			CreatedAt:        h.CreatedAt,
		}
	}

	return result, nil
}

// DeleteHobby removes a hobby from user's profile
func (s *ProfileService) DeleteHobby(ctx context.Context, userID, hobbyID string) error {
	_, err := s.entClient.Hobby.
		Update().
		Where(hobby.IDEQ(hobbyID)).
		Where(hobby.UserIDEQ(userID)).
		Where(hobby.DeletedAtIsNil()).
		SetDeletedAt(time.Now()).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("failed to delete hobby: %w", err)
	}
	return nil
}

// Helper functions
func profileToResponse(p *ent.Profile) *dto.ProfileResponse {
	return &dto.ProfileResponse{
		ID:              p.ID,
		UserID:          p.UserID,
		FirstName:       ptrToString(p.FirstName),
		LastName:        ptrToString(p.LastName),
		DateOfBirth:     p.DateOfBirth,
		Gender:          ptrToString(p.Gender),
		City:            ptrToString(p.City),
		Bio:             ptrToString(p.Bio),
		IntentStatement: ptrToString(p.IntentStatement),
		CreatedAt:       p.CreatedAt,
		UpdatedAt:       p.UpdatedAt,
	}
}

func ptrToString(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func strPtr(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}
