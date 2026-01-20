// internal/auth/services/auth_service.go
package services

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"

	entgen "github.com/UnoraApp/be/ent/generated"
	entuser "github.com/UnoraApp/be/ent/generated/user"
	"github.com/UnoraApp/be/internal/auth/dto"
	"github.com/UnoraApp/be/internal/config"
)

// AuthService handles authentication business logic
type AuthService struct {
	entClient     *entgen.Client
	redisClient   *redis.Client
	googleService *GoogleService
	tokenService  *TokenService
	cfg           *config.Config
}

// NewAuthService creates a new auth service with all dependencies injected
func NewAuthService(
	entClient *entgen.Client,
	redisClient *redis.Client,
	cfg *config.Config,
) *AuthService {
	return &AuthService{
		entClient:     entClient,
		redisClient:   redisClient,
		googleService: NewGoogleService(),
		tokenService:  NewTokenService(cfg),
		cfg:           cfg,
	}
}

// GoogleLogin handles Google OAuth login flow for mobile apps
func (s *AuthService) GoogleLogin(ctx context.Context, req *dto.GoogleOAuthRequest) (*dto.AuthResponse, error) {
	// 1. Verify Google ID token and extract user info
	googleUser, err := s.googleService.VerifyIDToken(ctx, req.IDToken)
	if err != nil {
		return nil, fmt.Errorf("failed to verify Google token: %w", err)
	}

	// 2. Find or create user in database
	user, err := s.findOrCreateUser(ctx, googleUser)
	if err != nil {
		return nil, fmt.Errorf("failed to find/create user: %w", err)
	}

	// 3. Generate token pair
	tokenPayload := dto.TokenPayload{
		UserID:   user.ID,
		Email:    user.Email,
		Provider: "google",
	}

	authResponse, err := s.tokenService.GenerateTokenPair(tokenPayload)
	if err != nil {
		return nil, fmt.Errorf("failed to generate tokens: %w", err)
	}

	// 4. Store refresh token in Redis for revocation capability
	if err := s.storeRefreshToken(ctx, user.ID, authResponse.RefreshToken); err != nil {
		return nil, fmt.Errorf("failed to store refresh token: %w", err)
	}

	// 5. Attach user info to response
	authResponse.User = dto.UserResponse{
		ID:        user.ID,
		Email:     user.Email,
		Name:      user.Name,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Picture:   user.Picture,
		Provider:  user.Provider,
		CreatedAt: user.CreatedAt,
	}

	return authResponse, nil
}

// RefreshToken handles token refresh
func (s *AuthService) RefreshToken(ctx context.Context, req *dto.RefreshTokenRequest) (*dto.RefreshTokenResponse, error) {
	// 1. Validate the refresh token exists in Redis (not revoked)
	payload, err := s.tokenService.ValidateToken(req.RefreshToken, "refresh")
	if err != nil {
		return nil, fmt.Errorf("invalid refresh token: %w", err)
	}

	// 2. Check if refresh token is in Redis (not revoked)
	key := fmt.Sprintf("refresh_token:%s", payload.UserID)
	storedToken, err := s.redisClient.Get(ctx, key).Result()
	if err == redis.Nil || storedToken != req.RefreshToken {
		return nil, fmt.Errorf("refresh token revoked or expired")
	}
	if err != nil {
		return nil, fmt.Errorf("failed to verify refresh token: %w", err)
	}

	// 3. Generate new access token
	response, err := s.tokenService.RefreshAccessToken(req.RefreshToken)
	if err != nil {
		return nil, fmt.Errorf("failed to refresh token: %w", err)
	}

	return response, nil
}

// Logout revokes all tokens for a user
func (s *AuthService) Logout(ctx context.Context, userID string) error {
	key := fmt.Sprintf("refresh_token:%s", userID)
	return s.redisClient.Del(ctx, key).Err()
}

// ValidateAccessToken validates an access token and returns the payload
func (s *AuthService) ValidateAccessToken(token string) (*dto.TokenPayload, error) {
	return s.tokenService.ValidateToken(token, "access")
}

// UserRecord is an internal representation of user data
type UserRecord struct {
	ID        string
	Email     string
	Name      string
	FirstName string
	LastName  string
	Picture   string
	Provider  string
	CreatedAt string
}

// findOrCreateUser finds an existing user by email or creates a new one using Ent ORM
func (s *AuthService) findOrCreateUser(ctx context.Context, googleUser *dto.GoogleUserInfo) (*UserRecord, error) {
	// Try to find existing user by email
	existingUser, err := s.entClient.User.
		Query().
		Where(entuser.EmailEQ(googleUser.Email)).
		Where(entuser.DeletedAtIsNil()).
		Only(ctx)
	
	if err == nil {
		// User found, return it
		return &UserRecord{
			ID:        existingUser.ID,
			Email:     ptrToString(existingUser.Email),
			Name:      ptrToString(existingUser.Name),
			FirstName: ptrToString(existingUser.FirstName),
			LastName:  ptrToString(existingUser.LastName),
			Picture:   ptrToString(existingUser.Picture),
			Provider:  ptrToString(existingUser.Provider),
			CreatedAt: existingUser.CreatedAt.Format(time.RFC3339),
		}, nil
	}

	// Check if error is "not found" - if so, create new user
	if !entgen.IsNotFound(err) {
		return nil, fmt.Errorf("failed to query user: %w", err)
	}

	// Create new user in database
	userID := uuid.New().String()
	
	newUser, err := s.entClient.User.
		Create().
		SetID(userID).
		SetEmail(googleUser.Email).
		SetName(googleUser.Name).
		SetFirstName(googleUser.FirstName).
		SetLastName(googleUser.LastName).
		SetPicture(googleUser.Picture).
		SetProvider("google").
		SetProviderUserID(googleUser.ID).
		Save(ctx)
	
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return &UserRecord{
		ID:        newUser.ID,
		Email:     ptrToString(newUser.Email),
		Name:      ptrToString(newUser.Name),
		FirstName: ptrToString(newUser.FirstName),
		LastName:  ptrToString(newUser.LastName),
		Picture:   ptrToString(newUser.Picture),
		Provider:  ptrToString(newUser.Provider),
		CreatedAt: newUser.CreatedAt.Format(time.RFC3339),
	}, nil
}

// ptrToString safely dereferences a string pointer, returning empty string if nil
func ptrToString(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

// storeRefreshToken stores refresh token in Redis for revocation capability
func (s *AuthService) storeRefreshToken(ctx context.Context, userID string, token string) error {
	key := fmt.Sprintf("refresh_token:%s", userID)
	expiry := time.Duration(s.cfg.Auth.JWTRefreshTokenExpiryDays) * 24 * time.Hour
	return s.redisClient.Set(ctx, key, token, expiry).Err()
}
