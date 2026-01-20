// internal/reveal/handlers/reveal_handler.go
package handlers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/UnoraApp/be/internal/reveal/dto"
	"github.com/UnoraApp/be/internal/reveal/services"
	"github.com/UnoraApp/be/pkg/apperror"
	"github.com/UnoraApp/be/pkg/response"
)

// RevealHandler handles reveal-related HTTP requests
type RevealHandler struct {
	revealService *services.RevealService
}

// NewRevealHandler creates a new reveal handler
func NewRevealHandler(revealService *services.RevealService) *RevealHandler {
	return &RevealHandler{
		revealService: revealService,
	}
}

// GetMilestones godoc
// @Summary      Get reveal milestones
// @Description  Get all reveal milestone configurations
// @Tags         reveals
// @Accept       json
// @Produce      json
// @Success      200 {object} response.APIResponse{data=[]dto.RevealMilestoneResponse} "Milestones"
// @Router       /reveal-milestones [get]
func (h *RevealHandler) GetMilestones(c *gin.Context) {
	milestones, err := h.revealService.GetMilestones(c.Request.Context())
	if err != nil {
		apperror.HandleError(c, apperror.InternalError(err))
		return
	}
	response.JSON(c, http.StatusOK, milestones)
}

// GetConnectionReveals godoc
// @Summary      Get connection reveals
// @Description  Get all reveals for a specific connection
// @Tags         reveals
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        connectionId path string true "Connection ID"
// @Success      200 {object} response.APIResponse{data=dto.ConnectionRevealsResponse} "Connection reveals"
// @Failure      401 {object} response.APIResponse "Not authenticated"
// @Failure      404 {object} response.APIResponse "Connection not found"
// @Router       /connections/{connectionId}/reveals [get]
func (h *RevealHandler) GetConnectionReveals(c *gin.Context) {
	userID, _ := c.Get("userID")
	connectionID := c.Param("connectionId")

	reveals, err := h.revealService.GetConnectionReveals(c.Request.Context(), userID.(string), connectionID)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			apperror.HandleError(c, apperror.NotFound("connection"))
			return
		}
		apperror.HandleError(c, apperror.InternalError(err))
		return
	}
	response.JSON(c, http.StatusOK, reveals)
}

// UnlockReveal godoc
// @Summary      Unlock reveal
// @Description  Unlock a reveal milestone for a connection
// @Tags         reveals
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        connectionId path string true "Connection ID"
// @Param        milestoneId path string true "Milestone ID"
// @Param        request body dto.UnlockRevealRequest true "Unlock options"
// @Success      200 {object} response.APIResponse{data=dto.UnlockRevealResponse} "Unlock result"
// @Failure      400 {object} response.APIResponse "Cannot unlock or insufficient credits"
// @Failure      401 {object} response.APIResponse "Not authenticated"
// @Router       /connections/{connectionId}/reveals/{milestoneId}/unlock [post]
func (h *RevealHandler) UnlockReveal(c *gin.Context) {
	userID, _ := c.Get("userID")
	connectionID := c.Param("connectionId")
	milestoneID := c.Param("milestoneId")

	var req dto.UnlockRevealRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		req = dto.UnlockRevealRequest{UseCredits: false}
	}

	result, err := h.revealService.UnlockReveal(c.Request.Context(), userID.(string), connectionID, milestoneID, &req)
	if err != nil {
		apperror.HandleError(c, apperror.BadRequest(err.Error()))
		return
	}
	response.JSON(c, http.StatusOK, result)
}

// MarkRevealViewed godoc
// @Summary      Mark reveal viewed
// @Description  Mark a reveal as viewed
// @Tags         reveals
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        revealId path string true "Reveal ID"
// @Success      200 {object} response.APIResponse{data=map[string]string} "Success"
// @Failure      401 {object} response.APIResponse "Not authenticated"
// @Failure      404 {object} response.APIResponse "Reveal not found"
// @Router       /reveals/{revealId}/viewed [post]
func (h *RevealHandler) MarkRevealViewed(c *gin.Context) {
	userID, _ := c.Get("userID")
	revealID := c.Param("revealId")

	if err := h.revealService.MarkRevealViewed(c.Request.Context(), userID.(string), revealID); err != nil {
		if strings.Contains(err.Error(), "not found") || strings.Contains(err.Error(), "unauthorized") {
			apperror.HandleError(c, apperror.NotFound("reveal"))
			return
		}
		apperror.HandleError(c, apperror.BadRequest(err.Error()))
		return
	}
	response.JSON(c, http.StatusOK, gin.H{"message": "Reveal marked as viewed"})
}
