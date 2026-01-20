// internal/streak/handlers/streak_handler.go
package handlers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/UnoraApp/be/internal/streak/dto"
	"github.com/UnoraApp/be/internal/streak/services"
	"github.com/UnoraApp/be/pkg/apperror"
	"github.com/UnoraApp/be/pkg/response"
)

// StreakHandler handles streak-related HTTP requests
type StreakHandler struct {
	streakService *services.StreakService
}

// NewStreakHandler creates a new streak handler
func NewStreakHandler(streakService *services.StreakService) *StreakHandler {
	return &StreakHandler{
		streakService: streakService,
	}
}

// GetStreak godoc
// @Summary      Get streak
// @Description  Get streak details for a connection
// @Tags         streaks
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        connectionId path string true "Connection ID"
// @Success      200 {object} response.APIResponse{data=dto.StreakResponse} "Streak details"
// @Failure      401 {object} response.APIResponse "Not authenticated"
// @Failure      404 {object} response.APIResponse "Streak not found"
// @Router       /connections/{connectionId}/streak [get]
func (h *StreakHandler) GetStreak(c *gin.Context) {
	userID, _ := c.Get("userID")
	connectionID := c.Param("connectionId")

	streak, err := h.streakService.GetStreakByConnection(c.Request.Context(), userID.(string), connectionID)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			apperror.HandleError(c, apperror.NotFound("streak"))
			return
		}
		apperror.HandleError(c, apperror.InternalError(err))
		return
	}
	response.JSON(c, http.StatusOK, streak)
}

// CheckIn godoc
// @Summary      Daily check-in
// @Description  Perform daily check-in for a streak
// @Tags         streaks
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        connectionId path string true "Connection ID"
// @Param        request body dto.CheckInRequest false "Check-in data"
// @Success      201 {object} response.APIResponse{data=dto.CheckInResponse} "Check-in created"
// @Failure      400 {object} response.APIResponse "Already checked in or streak not active"
// @Failure      401 {object} response.APIResponse "Not authenticated"
// @Router       /connections/{connectionId}/streak/check-in [post]
func (h *StreakHandler) CheckIn(c *gin.Context) {
	userID, _ := c.Get("userID")
	connectionID := c.Param("connectionId")

	var req dto.CheckInRequest
	c.ShouldBindJSON(&req) // Optional body

	checkIn, err := h.streakService.CheckIn(c.Request.Context(), userID.(string), connectionID, &req)
	if err != nil {
		apperror.HandleError(c, apperror.BadRequest(err.Error()))
		return
	}
	response.JSON(c, http.StatusCreated, checkIn)
}

// GetTodayStreaks godoc
// @Summary      Get today's streaks
// @Description  Get all streaks requiring check-in today
// @Tags         streaks
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200 {object} response.APIResponse{data=dto.TodayStreaksResponse} "Today's streaks"
// @Failure      401 {object} response.APIResponse "Not authenticated"
// @Router       /streaks/today [get]
func (h *StreakHandler) GetTodayStreaks(c *gin.Context) {
	userID, _ := c.Get("userID")

	streaks, err := h.streakService.GetTodayStreaks(c.Request.Context(), userID.(string))
	if err != nil {
		apperror.HandleError(c, apperror.InternalError(err))
		return
	}
	response.JSON(c, http.StatusOK, streaks)
}

// SendNudge godoc
// @Summary      Send nudge
// @Description  Send a nudge to streak partner
// @Tags         streaks
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        connectionId path string true "Connection ID"
// @Param        request body dto.SendNudgeRequest false "Nudge message"
// @Success      201 {object} response.APIResponse{data=dto.NudgeResponse} "Nudge sent"
// @Failure      400 {object} response.APIResponse "Already nudged today"
// @Failure      401 {object} response.APIResponse "Not authenticated"
// @Router       /connections/{connectionId}/nudge [post]
func (h *StreakHandler) SendNudge(c *gin.Context) {
	userID, _ := c.Get("userID")
	connectionID := c.Param("connectionId")

	var req dto.SendNudgeRequest
	c.ShouldBindJSON(&req) // Optional body

	nudge, err := h.streakService.SendNudge(c.Request.Context(), userID.(string), connectionID, &req)
	if err != nil {
		apperror.HandleError(c, apperror.BadRequest(err.Error()))
		return
	}
	response.JSON(c, http.StatusCreated, nudge)
}

// GetReceivedNudges godoc
// @Summary      Get received nudges
// @Description  Get nudges received by the authenticated user
// @Tags         streaks
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200 {object} response.APIResponse{data=[]dto.NudgeResponse} "Received nudges"
// @Failure      401 {object} response.APIResponse "Not authenticated"
// @Router       /nudges/received [get]
func (h *StreakHandler) GetReceivedNudges(c *gin.Context) {
	userID, _ := c.Get("userID")

	nudges, err := h.streakService.GetReceivedNudges(c.Request.Context(), userID.(string))
	if err != nil {
		apperror.HandleError(c, apperror.InternalError(err))
		return
	}
	response.JSON(c, http.StatusOK, nudges)
}

// GetRecoveryOptions godoc
// @Summary      Get recovery options
// @Description  Get available options for recovering a broken streak
// @Tags         streaks
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        streakId path string true "Streak ID"
// @Success      200 {object} response.APIResponse{data=dto.RecoveryOptionsResponse} "Recovery options"
// @Failure      401 {object} response.APIResponse "Not authenticated"
// @Failure      404 {object} response.APIResponse "Streak not found"
// @Router       /streaks/{streakId}/recovery-options [get]
func (h *StreakHandler) GetRecoveryOptions(c *gin.Context) {
	userID, _ := c.Get("userID")
	streakID := c.Param("streakId")

	options, err := h.streakService.GetRecoveryOptions(c.Request.Context(), userID.(string), streakID)
	if err != nil {
		apperror.HandleError(c, apperror.NotFound("streak"))
		return
	}
	response.JSON(c, http.StatusOK, options)
}

// RecoverStreak godoc
// @Summary      Recover streak
// @Description  Recover a broken streak using credits
// @Tags         streaks
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        streakId path string true "Streak ID"
// @Param        request body dto.RecoverStreakRequest true "Recovery request"
// @Success      200 {object} response.APIResponse{data=dto.RecoverStreakResponse} "Recovery result"
// @Failure      400 {object} response.APIResponse "Cannot recover or insufficient credits"
// @Failure      401 {object} response.APIResponse "Not authenticated"
// @Router       /streaks/{streakId}/recover [post]
func (h *StreakHandler) RecoverStreak(c *gin.Context) {
	userID, _ := c.Get("userID")
	streakID := c.Param("streakId")

	var req dto.RecoverStreakRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		apperror.HandleError(c, apperror.BadRequest(err.Error()))
		return
	}

	result, err := h.streakService.RecoverStreak(c.Request.Context(), userID.(string), streakID, &req)
	if err != nil {
		apperror.HandleError(c, apperror.BadRequest(err.Error()))
		return
	}
	response.JSON(c, http.StatusOK, result)
}
