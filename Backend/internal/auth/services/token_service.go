// internal/auth/services/token_service.go
package services

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"time"

	"golang.org/x/crypto/chacha20poly1305"

	"github.com/UnoraApp/be/internal/auth/dto"
	"github.com/UnoraApp/be/internal/config"
)

// TokenService handles JWT token generation and validation
type TokenService struct {
	secretKey              []byte
	accessTokenExpiryHours int
	refreshTokenExpiryDays int
}

// NewTokenService creates a new token service
func NewTokenService(cfg *config.Config) *TokenService {
	// Derive a 32-byte key from the secret
	key := deriveKey(cfg.Auth.SecretKey)

	return &TokenService{
		secretKey:              key,
		accessTokenExpiryHours: cfg.Auth.JWTAccessTokenExpiryHours,
		refreshTokenExpiryDays: cfg.Auth.JWTRefreshTokenExpiryDays,
	}
}

// tokenData is the internal structure for token data
type tokenData struct {
	Payload   dto.TokenPayload `json:"payload"`
	ExpiresAt int64            `json:"expiresAt"`
	IssuedAt  int64            `json:"issuedAt"`
	Type      string           `json:"type"` // "access" or "refresh"
}

// GenerateTokenPair generates both access and refresh tokens
func (s *TokenService) GenerateTokenPair(payload dto.TokenPayload) (*dto.AuthResponse, error) {
	now := time.Now()

	// Generate access token
	accessExpiry := now.Add(time.Duration(s.accessTokenExpiryHours) * time.Hour)
	accessToken, err := s.generateToken(payload, accessExpiry, "access")
	if err != nil {
		return nil, fmt.Errorf("failed to generate access token: %w", err)
	}

	// Generate refresh token
	refreshExpiry := now.Add(time.Duration(s.refreshTokenExpiryDays) * 24 * time.Hour)
	refreshToken, err := s.generateToken(payload, refreshExpiry, "refresh")
	if err != nil {
		return nil, fmt.Errorf("failed to generate refresh token: %w", err)
	}

	return &dto.AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresIn:    int64(s.accessTokenExpiryHours * 3600), // in seconds
		TokenType:    "Bearer",
	}, nil
}

// RefreshAccessToken generates a new access token from a refresh token
func (s *TokenService) RefreshAccessToken(refreshToken string) (*dto.RefreshTokenResponse, error) {
	// Validate refresh token
	payload, err := s.ValidateToken(refreshToken, "refresh")
	if err != nil {
		return nil, fmt.Errorf("invalid refresh token: %w", err)
	}

	// Generate new access token
	now := time.Now()
	accessExpiry := now.Add(time.Duration(s.accessTokenExpiryHours) * time.Hour)
	accessToken, err := s.generateToken(*payload, accessExpiry, "access")
	if err != nil {
		return nil, fmt.Errorf("failed to generate access token: %w", err)
	}

	return &dto.RefreshTokenResponse{
		AccessToken: accessToken,
		ExpiresIn:   int64(s.accessTokenExpiryHours * 3600),
		TokenType:   "Bearer",
	}, nil
}

// ValidateToken validates a token and returns its payload
func (s *TokenService) ValidateToken(token string, expectedType string) (*dto.TokenPayload, error) {
	// Decode token
	tokenBytes, err := base64.URLEncoding.DecodeString(token)
	if err != nil {
		return nil, fmt.Errorf("invalid token encoding: %w", err)
	}

	// Decrypt token
	aead, err := chacha20poly1305.NewX(s.secretKey)
	if err != nil {
		return nil, fmt.Errorf("failed to create cipher: %w", err)
	}

	nonceSize := aead.NonceSize()
	if len(tokenBytes) < nonceSize {
		return nil, fmt.Errorf("token too short")
	}

	nonce, ciphertext := tokenBytes[:nonceSize], tokenBytes[nonceSize:]
	plaintext, err := aead.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to decrypt token: %w", err)
	}

	// Parse token data
	var data tokenData
	if err := json.Unmarshal(plaintext, &data); err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}

	// Validate token type
	if data.Type != expectedType {
		return nil, fmt.Errorf("invalid token type: expected %s, got %s", expectedType, data.Type)
	}

	// Check expiry
	if time.Now().Unix() > data.ExpiresAt {
		return nil, fmt.Errorf("token expired")
	}

	return &data.Payload, nil
}

// generateToken generates an encrypted token
func (s *TokenService) generateToken(payload dto.TokenPayload, expiry time.Time, tokenType string) (string, error) {
	data := tokenData{
		Payload:   payload,
		ExpiresAt: expiry.Unix(),
		IssuedAt:  time.Now().Unix(),
		Type:      tokenType,
	}

	plaintext, err := json.Marshal(data)
	if err != nil {
		return "", fmt.Errorf("failed to marshal token data: %w", err)
	}

	// Encrypt token
	aead, err := chacha20poly1305.NewX(s.secretKey)
	if err != nil {
		return "", fmt.Errorf("failed to create cipher: %w", err)
	}

	nonce := make([]byte, aead.NonceSize())
	if _, err := rand.Read(nonce); err != nil {
		return "", fmt.Errorf("failed to generate nonce: %w", err)
	}

	ciphertext := aead.Seal(nonce, nonce, plaintext, nil)

	return base64.URLEncoding.EncodeToString(ciphertext), nil
}

// deriveKey derives a 32-byte key from a secret string
func deriveKey(secret string) []byte {
	// Simple key derivation - in production you might want to use HKDF or similar
	key := make([]byte, 32)
	copy(key, []byte(secret))
	return key
}
