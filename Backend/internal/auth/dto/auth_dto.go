// internal/auth/dto/auth_dto.go
package dto

// GoogleOAuthRequest is the request body for Google OAuth login
// @Description Google OAuth login request - send ID token from mobile app
type GoogleOAuthRequest struct {
	// Google ID token obtained from Google Sign-In SDK on mobile
	IDToken string `json:"idToken" validate:"required" example:"eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9..."`
}

// GoogleUserInfo represents user info extracted from Google token
// @Description User information extracted from verified Google ID token
type GoogleUserInfo struct {
	ID            string `json:"id" example:"123456789"`
	Email         string `json:"email" example:"user@gmail.com"`
	EmailVerified bool   `json:"emailVerified" example:"true"`
	Name          string `json:"name" example:"John Doe"`
	FirstName     string `json:"firstName" example:"John"`
	LastName      string `json:"lastName" example:"Doe"`
	Picture       string `json:"picture" example:"https://lh3.googleusercontent.com/..."`
}

// AuthResponse is the response body for successful authentication
// @Description Successful authentication response with tokens and user info
type AuthResponse struct {
	// JWT access token - use in Authorization header for protected routes
	AccessToken string `json:"accessToken" example:"eyJhbGciOiJIUzI1NiIs..."`
	// Refresh token - use to get new access token when it expires
	RefreshToken string `json:"refreshToken" example:"eyJhbGciOiJIUzI1NiIs..."`
	// Access token expiry time in seconds
	ExpiresIn int64 `json:"expiresIn" example:"86400"`
	// Token type - always "Bearer"
	TokenType string `json:"tokenType" example:"Bearer"`
	// Authenticated user information
	User UserResponse `json:"user"`
}

// UserResponse is the user data returned after auth
// @Description User profile information
type UserResponse struct {
	ID        string `json:"id" example:"550e8400-e29b-41d4-a716-446655440000"`
	Email     string `json:"email" example:"user@gmail.com"`
	Name      string `json:"name" example:"John Doe"`
	FirstName string `json:"firstName" example:"John"`
	LastName  string `json:"lastName" example:"Doe"`
	Picture   string `json:"picture" example:"https://lh3.googleusercontent.com/..."`
	Provider  string `json:"provider" example:"google"`
	CreatedAt string `json:"createdAt" example:"2024-01-01T00:00:00Z"`
}

// RefreshTokenRequest is the request body for token refresh
// @Description Token refresh request
type RefreshTokenRequest struct {
	// Refresh token received from login
	RefreshToken string `json:"refreshToken" validate:"required" example:"eyJhbGciOiJIUzI1NiIs..."`
}

// RefreshTokenResponse is the response body for token refresh
// @Description Token refresh response with new access token
type RefreshTokenResponse struct {
	// New JWT access token
	AccessToken string `json:"accessToken" example:"eyJhbGciOiJIUzI1NiIs..."`
	// Token expiry time in seconds
	ExpiresIn int64 `json:"expiresIn" example:"86400"`
	// Token type - always "Bearer"
	TokenType string `json:"tokenType" example:"Bearer"`
}

// TokenPayload is the JWT payload structure
// @Description JWT token payload containing user identification
type TokenPayload struct {
	UserID   string `json:"userId" example:"550e8400-e29b-41d4-a716-446655440000"`
	Email    string `json:"email" example:"user@gmail.com"`
	Provider string `json:"provider" example:"google"`
}
