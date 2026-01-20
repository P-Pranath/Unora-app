// internal/profile/handlers/profile_handler.go
package handlers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/UnoraApp/be/internal/profile/dto"
	"github.com/UnoraApp/be/internal/profile/services"
	"github.com/UnoraApp/be/pkg/apperror"
	"github.com/UnoraApp/be/pkg/response"
)

// ProfileHandler handles profile-related HTTP requests
type ProfileHandler struct {
	profileService *services.ProfileService
	photoService   *services.PhotoService
}

// NewProfileHandler creates a new profile handler with dependencies injected
func NewProfileHandler(profileService *services.ProfileService, photoService *services.PhotoService) *ProfileHandler {
	return &ProfileHandler{
		profileService: profileService,
		photoService:   photoService,
	}
}

// CreateProfile godoc
// @Summary      Create user profile
// @Description  Create a new profile for the authenticated user
// @Tags         profile
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        request body dto.CreateProfileRequest true "Profile data"
// @Success      201 {object} response.APIResponse{data=dto.ProfileResponse} "Profile created"
// @Failure      400 {object} response.APIResponse "Invalid request body"
// @Failure      401 {object} response.APIResponse "Not authenticated"
// @Failure      409 {object} response.APIResponse "Profile already exists"
// @Router       /profile [post]
func (h *ProfileHandler) CreateProfile(c *gin.Context) {
	userID, _ := c.Get("userID")

	var req dto.CreateProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		apperror.HandleError(c, apperror.BadRequest("Invalid request body: "+err.Error()))
		return
	}

	profile, err := h.profileService.CreateProfile(c.Request.Context(), userID.(string), &req)
	if err != nil {
		if strings.Contains(err.Error(), "already exists") {
			apperror.HandleError(c, apperror.Conflict(err.Error()))
			return
		}
		apperror.HandleError(c, apperror.BadRequest(err.Error()))
		return
	}

	response.JSON(c, http.StatusCreated, profile)
}

// GetProfile godoc
// @Summary      Get user profile
// @Description  Get the authenticated user's profile
// @Tags         profile
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200 {object} response.APIResponse{data=dto.ProfileResponse} "Profile data"
// @Failure      401 {object} response.APIResponse "Not authenticated"
// @Failure      404 {object} response.APIResponse "Profile not found"
// @Router       /profile [get]
func (h *ProfileHandler) GetProfile(c *gin.Context) {
	userID, _ := c.Get("userID")

	profile, err := h.profileService.GetProfile(c.Request.Context(), userID.(string))
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			apperror.HandleError(c, apperror.NotFound("profile"))
			return
		}
		apperror.HandleError(c, apperror.InternalError(err))
		return
	}

	response.JSON(c, http.StatusOK, profile)
}

// UpdateProfile godoc
// @Summary      Update user profile
// @Description  Update the authenticated user's profile
// @Tags         profile
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        request body dto.UpdateProfileRequest true "Profile update data"
// @Success      200 {object} response.APIResponse{data=dto.ProfileResponse} "Profile updated"
// @Failure      400 {object} response.APIResponse "Invalid request body"
// @Failure      401 {object} response.APIResponse "Not authenticated"
// @Failure      404 {object} response.APIResponse "Profile not found"
// @Router       /profile [put]
func (h *ProfileHandler) UpdateProfile(c *gin.Context) {
	userID, _ := c.Get("userID")

	var req dto.UpdateProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		apperror.HandleError(c, apperror.BadRequest("Invalid request body: "+err.Error()))
		return
	}

	profile, err := h.profileService.UpdateProfile(c.Request.Context(), userID.(string), &req)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			apperror.HandleError(c, apperror.NotFound("profile"))
			return
		}
		apperror.HandleError(c, apperror.BadRequest(err.Error()))
		return
	}

	response.JSON(c, http.StatusOK, profile)
}

// DeleteProfile godoc
// @Summary      Delete user profile
// @Description  Soft delete the authenticated user's profile
// @Tags         profile
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200 {object} response.APIResponse{data=map[string]string} "Profile deleted"
// @Failure      401 {object} response.APIResponse "Not authenticated"
// @Router       /profile [delete]
func (h *ProfileHandler) DeleteProfile(c *gin.Context) {
	userID, _ := c.Get("userID")

	if err := h.profileService.DeleteProfile(c.Request.Context(), userID.(string)); err != nil {
		apperror.HandleError(c, apperror.InternalError(err))
		return
	}

	response.JSON(c, http.StatusOK, gin.H{"message": "Profile deleted successfully"})
}

// GetOnboardingStatus godoc
// @Summary      Get onboarding status
// @Description  Get the onboarding completion status for the authenticated user
// @Tags         profile
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200 {object} response.APIResponse{data=dto.OnboardingStatusResponse} "Onboarding status"
// @Failure      401 {object} response.APIResponse "Not authenticated"
// @Router       /profile/onboarding-status [get]
func (h *ProfileHandler) GetOnboardingStatus(c *gin.Context) {
	userID, _ := c.Get("userID")

	status, err := h.profileService.GetOnboardingStatus(c.Request.Context(), userID.(string))
	if err != nil {
		apperror.HandleError(c, apperror.InternalError(err))
		return
	}

	response.JSON(c, http.StatusOK, status)
}

// ========== Photo Handlers ==========

// GetUploadURL godoc
// @Summary      Get photo upload URL
// @Description  Get a presigned URL for uploading a photo
// @Tags         photos
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        request body dto.UploadPhotoURLRequest true "Upload request"
// @Success      200 {object} response.APIResponse{data=dto.UploadPhotoURLResponse} "Upload URL"
// @Failure      400 {object} response.APIResponse "Invalid request body"
// @Failure      401 {object} response.APIResponse "Not authenticated"
// @Router       /photos/upload-url [post]
func (h *ProfileHandler) GetUploadURL(c *gin.Context) {
	userID, _ := c.Get("userID")

	var req dto.UploadPhotoURLRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		apperror.HandleError(c, apperror.BadRequest("Invalid request body: "+err.Error()))
		return
	}

	result, err := h.photoService.GetUploadURL(c.Request.Context(), userID.(string), &req)
	if err != nil {
		apperror.HandleError(c, apperror.InternalError(err))
		return
	}

	response.JSON(c, http.StatusOK, result)
}

// RegisterPhoto godoc
// @Summary      Register uploaded photo
// @Description  Register a photo after successful upload to storage
// @Tags         photos
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        request body dto.RegisterPhotoRequest true "Photo registration data"
// @Success      201 {object} response.APIResponse{data=dto.PhotoResponse} "Photo registered"
// @Failure      400 {object} response.APIResponse "Invalid request body or max photos reached"
// @Failure      401 {object} response.APIResponse "Not authenticated"
// @Router       /photos [post]
func (h *ProfileHandler) RegisterPhoto(c *gin.Context) {
	userID, _ := c.Get("userID")

	var req dto.RegisterPhotoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		apperror.HandleError(c, apperror.BadRequest("Invalid request body: "+err.Error()))
		return
	}

	photo, err := h.photoService.RegisterPhoto(c.Request.Context(), userID.(string), &req)
	if err != nil {
		apperror.HandleError(c, apperror.BadRequest(err.Error()))
		return
	}

	response.JSON(c, http.StatusCreated, photo)
}

// GetUserPhotos godoc
// @Summary      Get user photos
// @Description  Get all photos for the authenticated user
// @Tags         photos
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200 {object} response.APIResponse{data=[]dto.PhotoResponse} "User photos"
// @Failure      401 {object} response.APIResponse "Not authenticated"
// @Router       /photos [get]
func (h *ProfileHandler) GetUserPhotos(c *gin.Context) {
	userID, _ := c.Get("userID")

	photos, err := h.photoService.GetUserPhotos(c.Request.Context(), userID.(string))
	if err != nil {
		apperror.HandleError(c, apperror.InternalError(err))
		return
	}

	response.JSON(c, http.StatusOK, photos)
}

// UpdatePhotoOrder godoc
// @Summary      Update photo order
// @Description  Update the display order of a photo
// @Tags         photos
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id path string true "Photo ID"
// @Param        request body dto.UpdatePhotoOrderRequest true "New order"
// @Success      200 {object} response.APIResponse{data=dto.PhotoResponse} "Photo updated"
// @Failure      400 {object} response.APIResponse "Invalid request body"
// @Failure      401 {object} response.APIResponse "Not authenticated"
// @Failure      404 {object} response.APIResponse "Photo not found"
// @Router       /photos/{id}/order [put]
func (h *ProfileHandler) UpdatePhotoOrder(c *gin.Context) {
	userID, _ := c.Get("userID")
	photoID := c.Param("id")

	var req dto.UpdatePhotoOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		apperror.HandleError(c, apperror.BadRequest("Invalid request body: "+err.Error()))
		return
	}

	photo, err := h.photoService.UpdatePhotoOrder(c.Request.Context(), userID.(string), photoID, &req)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			apperror.HandleError(c, apperror.NotFound("photo"))
			return
		}
		apperror.HandleError(c, apperror.BadRequest(err.Error()))
		return
	}

	response.JSON(c, http.StatusOK, photo)
}

// DeletePhoto godoc
// @Summary      Delete photo
// @Description  Delete a photo from user's profile
// @Tags         photos
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id path string true "Photo ID"
// @Success      200 {object} response.APIResponse{data=map[string]string} "Photo deleted"
// @Failure      401 {object} response.APIResponse "Not authenticated"
// @Failure      404 {object} response.APIResponse "Photo not found"
// @Router       /photos/{id} [delete]
func (h *ProfileHandler) DeletePhoto(c *gin.Context) {
	userID, _ := c.Get("userID")
	photoID := c.Param("id")

	if err := h.photoService.DeletePhoto(c.Request.Context(), userID.(string), photoID); err != nil {
		if strings.Contains(err.Error(), "not found") {
			apperror.HandleError(c, apperror.NotFound("photo"))
			return
		}
		apperror.HandleError(c, apperror.InternalError(err))
		return
	}

	response.JSON(c, http.StatusOK, gin.H{"message": "Photo deleted successfully"})
}

// ========== Hobby Handlers ==========

// GetHobbyOptions godoc
// @Summary      Get hobby options
// @Description  Get all available hobby options
// @Tags         hobbies
// @Accept       json
// @Produce      json
// @Param        category query string false "Filter by category"
// @Success      200 {object} response.APIResponse{data=[]dto.HobbyOptionResponse} "Hobby options"
// @Router       /hobby-options [get]
func (h *ProfileHandler) GetHobbyOptions(c *gin.Context) {
	category := c.Query("category")

	options, err := h.profileService.GetHobbyOptions(c.Request.Context(), category)
	if err != nil {
		apperror.HandleError(c, apperror.InternalError(err))
		return
	}

	response.JSON(c, http.StatusOK, options)
}

// AddHobby godoc
// @Summary      Add hobby
// @Description  Add a hobby to user's profile
// @Tags         hobbies
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        request body dto.AddHobbyRequest true "Hobby data"
// @Success      201 {object} response.APIResponse{data=dto.HobbyResponse} "Hobby added"
// @Failure      400 {object} response.APIResponse "Invalid request body or hobby already added"
// @Failure      401 {object} response.APIResponse "Not authenticated"
// @Router       /hobbies [post]
func (h *ProfileHandler) AddHobby(c *gin.Context) {
	userID, _ := c.Get("userID")

	var req dto.AddHobbyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		apperror.HandleError(c, apperror.BadRequest("Invalid request body: "+err.Error()))
		return
	}

	hobby, err := h.profileService.AddHobby(c.Request.Context(), userID.(string), &req)
	if err != nil {
		apperror.HandleError(c, apperror.BadRequest(err.Error()))
		return
	}

	response.JSON(c, http.StatusCreated, hobby)
}

// GetUserHobbies godoc
// @Summary      Get user hobbies
// @Description  Get all hobbies for the authenticated user
// @Tags         hobbies
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200 {object} response.APIResponse{data=[]dto.HobbyResponse} "User hobbies"
// @Failure      401 {object} response.APIResponse "Not authenticated"
// @Router       /hobbies [get]
func (h *ProfileHandler) GetUserHobbies(c *gin.Context) {
	userID, _ := c.Get("userID")

	hobbies, err := h.profileService.GetUserHobbies(c.Request.Context(), userID.(string))
	if err != nil {
		apperror.HandleError(c, apperror.InternalError(err))
		return
	}

	response.JSON(c, http.StatusOK, hobbies)
}

// DeleteHobby godoc
// @Summary      Delete hobby
// @Description  Remove a hobby from user's profile
// @Tags         hobbies
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id path string true "Hobby ID"
// @Success      200 {object} response.APIResponse{data=map[string]string} "Hobby deleted"
// @Failure      401 {object} response.APIResponse "Not authenticated"
// @Router       /hobbies/{id} [delete]
func (h *ProfileHandler) DeleteHobby(c *gin.Context) {
	userID, _ := c.Get("userID")
	hobbyID := c.Param("id")

	if err := h.profileService.DeleteHobby(c.Request.Context(), userID.(string), hobbyID); err != nil {
		apperror.HandleError(c, apperror.InternalError(err))
		return
	}

	response.JSON(c, http.StatusOK, gin.H{"message": "Hobby deleted successfully"})
}
