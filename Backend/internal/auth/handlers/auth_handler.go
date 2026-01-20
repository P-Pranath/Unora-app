// internal/auth/handlers/auth_handler.go
package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/UnoraApp/be/internal/auth/dto"
	"github.com/UnoraApp/be/internal/auth/services"
	"github.com/UnoraApp/be/pkg/apperror"
	"github.com/UnoraApp/be/pkg/response"
)

// AuthHandler handles authentication HTTP requests
type AuthHandler struct {
	authService *services.AuthService
}

// NewAuthHandler creates a new auth handler with dependencies injected
func NewAuthHandler(authService *services.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

// GoogleLogin godoc
// @Summary      Login with Google OAuth
// @Description  Authenticate user using Google ID token from mobile app. Returns access token, refresh token, and user info.
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request body dto.GoogleOAuthRequest true "Google ID Token from mobile app"
// @Success      200 {object} response.APIResponse{data=dto.AuthResponse} "Login successful"
// @Failure      400 {object} response.APIResponse "Invalid request body"
// @Failure      401 {object} response.APIResponse "Authentication failed - invalid token"
// @Router       /auth/google [post]
func (h *AuthHandler) GoogleLogin(c *gin.Context) {
	var req dto.GoogleOAuthRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		apperror.HandleError(c, apperror.BadRequest("Invalid request body: "+err.Error()))
		return
	}

	authResponse, err := h.authService.GoogleLogin(c.Request.Context(), &req)
	if err != nil {
		apperror.HandleError(c, apperror.Unauthorized("Authentication failed: "+err.Error()))
		return
	}

	response.JSON(c, http.StatusOK, authResponse)
}

// RefreshToken godoc
// @Summary      Refresh access token
// @Description  Get a new access token using a valid refresh token. Use this when access token expires.
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request body dto.RefreshTokenRequest true "Refresh token"
// @Success      200 {object} response.APIResponse{data=dto.RefreshTokenResponse} "Token refreshed"
// @Failure      400 {object} response.APIResponse "Invalid request body"
// @Failure      401 {object} response.APIResponse "Invalid or expired refresh token"
// @Router       /auth/refresh [post]
func (h *AuthHandler) RefreshToken(c *gin.Context) {
	var req dto.RefreshTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		apperror.HandleError(c, apperror.BadRequest("Invalid request body: "+err.Error()))
		return
	}

	tokenResponse, err := h.authService.RefreshToken(c.Request.Context(), &req)
	if err != nil {
		apperror.HandleError(c, apperror.Unauthorized("Token refresh failed: "+err.Error()))
		return
	}

	response.JSON(c, http.StatusOK, tokenResponse)
}

// Logout godoc
// @Summary      Logout user
// @Description  Revoke refresh token for the authenticated user. Requires valid access token.
// @Tags         auth
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200 {object} response.APIResponse{data=map[string]string} "Logout successful"
// @Failure      401 {object} response.APIResponse "Not authenticated"
// @Router       /auth/logout [post]
func (h *AuthHandler) Logout(c *gin.Context) {
	// Get user ID from context (set by auth middleware)
	userID, exists := c.Get("userID")
	if !exists {
		apperror.HandleError(c, apperror.Unauthorized(""))
		return
	}

	if err := h.authService.Logout(c.Request.Context(), userID.(string)); err != nil {
		apperror.HandleError(c, apperror.InternalError(err))
		return
	}

	response.JSON(c, http.StatusOK, gin.H{"message": "Successfully logged out"})
}

// GetMe godoc
// @Summary      Get current user information
// @Description  Returns the profile information of the currently authenticated user.
// @Tags         auth
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200 {object} response.APIResponse{data=dto.TokenPayload} "User information"
// @Failure      401 {object} response.APIResponse "Not authenticated"
// @Router       /auth/me [get]
func (h *AuthHandler) GetMe(c *gin.Context) {
	payload, exists := c.Get("userPayload")
	if !exists {
		apperror.HandleError(c, apperror.Unauthorized(""))
		return
	}

	response.JSON(c, http.StatusOK, payload)
}
