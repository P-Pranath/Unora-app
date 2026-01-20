// internal/auth/services/google_service.go
package services

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/UnoraApp/be/internal/auth/dto"
)

// GoogleService handles Google OAuth operations
type GoogleService struct {
	httpClient *http.Client
}

// NewGoogleService creates a new Google service
func NewGoogleService() *GoogleService {
	return &GoogleService{
		httpClient: &http.Client{Timeout: 10 * time.Second},
	}
}

// googleTokenInfo is the response from Google's token verification endpoint
type googleTokenInfo struct {
	Iss           string `json:"iss"`
	Azp           string `json:"azp"`
	Aud           string `json:"aud"`
	Sub           string `json:"sub"`
	Email         string `json:"email"`
	EmailVerified string `json:"email_verified"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Picture       string `json:"picture"`
	Locale        string `json:"locale"`
	Iat           string `json:"iat"`
	Exp           string `json:"exp"`
	Alg           string `json:"alg"`
	Kid           string `json:"kid"`
	Typ           string `json:"typ"`
}

// VerifyIDToken verifies a Google ID token and returns user info
func (s *GoogleService) VerifyIDToken(ctx context.Context, idToken string) (*dto.GoogleUserInfo, error) {
	// Google's token info endpoint for verifying ID tokens
	url := fmt.Sprintf("https://oauth2.googleapis.com/tokeninfo?id_token=%s", idToken)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := s.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to verify token: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("invalid token: status %d", resp.StatusCode)
	}

	var tokenInfo googleTokenInfo
	if err := json.NewDecoder(resp.Body).Decode(&tokenInfo); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	// Validate issuer
	if tokenInfo.Iss != "accounts.google.com" && tokenInfo.Iss != "https://accounts.google.com" {
		return nil, fmt.Errorf("invalid token issuer: %s", tokenInfo.Iss)
	}

	return &dto.GoogleUserInfo{
		ID:            tokenInfo.Sub,
		Email:         tokenInfo.Email,
		EmailVerified: tokenInfo.EmailVerified == "true",
		Name:          tokenInfo.Name,
		FirstName:     tokenInfo.GivenName,
		LastName:      tokenInfo.FamilyName,
		Picture:       tokenInfo.Picture,
	}, nil
}
