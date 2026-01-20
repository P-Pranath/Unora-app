// internal/safety/handlers/safety_handler.go
package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/UnoraApp/be/internal/safety/dto"
	"github.com/UnoraApp/be/internal/safety/services"
	"github.com/UnoraApp/be/pkg/apperror"
	"github.com/UnoraApp/be/pkg/response"
)

// SafetyHandler handles safety-related HTTP requests
type SafetyHandler struct {
	safetyService *services.SafetyService
}

// NewSafetyHandler creates a new safety handler
func NewSafetyHandler(safetyService *services.SafetyService) *SafetyHandler {
	return &SafetyHandler{
		safetyService: safetyService,
	}
}

// BlockUser godoc
// @Summary      Block user
// @Description  Block a user (also terminates connection and cancels interests)
// @Tags         safety
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        request body dto.BlockUserRequest true "User to block"
// @Success      201 {object} response.APIResponse{data=dto.BlockResponse} "User blocked"
// @Failure      400 {object} response.APIResponse "Invalid request"
// @Failure      401 {object} response.APIResponse "Not authenticated"
// @Router       /blocks [post]
func (h *SafetyHandler) BlockUser(c *gin.Context) {
	userID, _ := c.Get("userID")

	var req dto.BlockUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		apperror.HandleError(c, apperror.BadRequest(err.Error()))
		return
	}

	block, err := h.safetyService.BlockUser(c.Request.Context(), userID.(string), &req)
	if err != nil {
		apperror.HandleError(c, apperror.BadRequest(err.Error()))
		return
	}
	response.JSON(c, http.StatusCreated, block)
}

// UnblockUser godoc
// @Summary      Unblock user
// @Description  Unblock a previously blocked user
// @Tags         safety
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        request body dto.UnblockUserRequest true "User to unblock"
// @Success      200 {object} response.APIResponse{data=map[string]string} "User unblocked"
// @Failure      400 {object} response.APIResponse "User not blocked"
// @Failure      401 {object} response.APIResponse "Not authenticated"
// @Router       /blocks/unblock [post]
func (h *SafetyHandler) UnblockUser(c *gin.Context) {
	userID, _ := c.Get("userID")

	var req dto.UnblockUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		apperror.HandleError(c, apperror.BadRequest(err.Error()))
		return
	}

	if err := h.safetyService.UnblockUser(c.Request.Context(), userID.(string), &req); err != nil {
		apperror.HandleError(c, apperror.BadRequest(err.Error()))
		return
	}
	response.JSON(c, http.StatusOK, gin.H{"message": "User unblocked"})
}

// GetBlockedUsers godoc
// @Summary      Get blocked users
// @Description  Get list of users blocked by the authenticated user
// @Tags         safety
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200 {object} response.APIResponse{data=dto.BlockedUsersResponse} "Blocked users"
// @Failure      401 {object} response.APIResponse "Not authenticated"
// @Router       /blocks [get]
func (h *SafetyHandler) GetBlockedUsers(c *gin.Context) {
	userID, _ := c.Get("userID")

	blocks, err := h.safetyService.GetBlockedUsers(c.Request.Context(), userID.(string))
	if err != nil {
		apperror.HandleError(c, apperror.InternalError(err))
		return
	}
	response.JSON(c, http.StatusOK, blocks)
}

// GetBlockStatus godoc
// @Summary      Check block status
// @Description  Check if a user is blocked or blocking the authenticated user
// @Tags         safety
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        userId path string true "User ID to check"
// @Success      200 {object} response.APIResponse{data=dto.BlockStatusResponse} "Block status"
// @Failure      401 {object} response.APIResponse "Not authenticated"
// @Router       /blocks/status/{userId} [get]
func (h *SafetyHandler) GetBlockStatus(c *gin.Context) {
	userID, _ := c.Get("userID")
	otherUserID := c.Param("userId")

	status, err := h.safetyService.GetBlockStatus(c.Request.Context(), userID.(string), otherUserID)
	if err != nil {
		apperror.HandleError(c, apperror.InternalError(err))
		return
	}
	response.JSON(c, http.StatusOK, status)
}

// ReportUser godoc
// @Summary      Report user
// @Description  Report a user for policy violation
// @Tags         safety
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        request body dto.ReportUserRequest true "Report details"
// @Success      201 {object} response.APIResponse{data=dto.ReportResponse} "Report submitted"
// @Failure      400 {object} response.APIResponse "Invalid request"
// @Failure      401 {object} response.APIResponse "Not authenticated"
// @Router       /reports [post]
func (h *SafetyHandler) ReportUser(c *gin.Context) {
	userID, _ := c.Get("userID")

	var req dto.ReportUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		apperror.HandleError(c, apperror.BadRequest(err.Error()))
		return
	}

	report, err := h.safetyService.ReportUser(c.Request.Context(), userID.(string), &req)
	if err != nil {
		apperror.HandleError(c, apperror.BadRequest(err.Error()))
		return
	}
	response.JSON(c, http.StatusCreated, report)
}
