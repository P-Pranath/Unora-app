// internal/discovery/handlers/discovery_handler.go
package handlers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/UnoraApp/be/internal/discovery/dto"
	"github.com/UnoraApp/be/internal/discovery/services"
	"github.com/UnoraApp/be/pkg/apperror"
	"github.com/UnoraApp/be/pkg/response"
)

// DiscoveryHandler handles discovery and matching HTTP requests
type DiscoveryHandler struct {
	discoveryService *services.DiscoveryService
	matchingService  *services.MatchingService
}

// NewDiscoveryHandler creates a new discovery handler
func NewDiscoveryHandler(discoveryService *services.DiscoveryService, matchingService *services.MatchingService) *DiscoveryHandler {
	return &DiscoveryHandler{
		discoveryService: discoveryService,
		matchingService:  matchingService,
	}
}

// ========== Server Endpoints ==========

// GetServers godoc
// @Summary      Get servers
// @Description  Get all available server types (partner/friend/growth)
// @Tags         servers
// @Accept       json
// @Produce      json
// @Success      200 {object} response.APIResponse{data=[]dto.ServerResponse} "Server list"
// @Router       /servers [get]
func (h *DiscoveryHandler) GetServers(c *gin.Context) {
	servers, err := h.discoveryService.GetServers(c.Request.Context())
	if err != nil {
		apperror.HandleError(c, apperror.InternalError(err))
		return
	}
	response.JSON(c, http.StatusOK, servers)
}

// ========== Filter Endpoints ==========

// GetFilter godoc
// @Summary      Get filter
// @Description  Get user's filter settings for a server type
// @Tags         filters
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        serverType path string true "Server type (partner/friend/growth)"
// @Success      200 {object} response.APIResponse{data=dto.FilterResponse} "Filter settings"
// @Failure      401 {object} response.APIResponse "Not authenticated"
// @Router       /servers/{serverType}/filters [get]
func (h *DiscoveryHandler) GetFilter(c *gin.Context) {
	userID, _ := c.Get("userID")
	serverType := c.Param("serverType")

	filter, err := h.discoveryService.GetFilter(c.Request.Context(), userID.(string), serverType)
	if err != nil {
		apperror.HandleError(c, apperror.InternalError(err))
		return
	}
	response.JSON(c, http.StatusOK, filter)
}

// UpdateFilter godoc
// @Summary      Update filter
// @Description  Update user's filter settings for a server type
// @Tags         filters
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        serverType path string true "Server type (partner/friend/growth)"
// @Param        request body dto.UpdateFilterRequest true "Filter config"
// @Success      200 {object} response.APIResponse{data=dto.FilterResponse} "Updated filter"
// @Failure      400 {object} response.APIResponse "Invalid request"
// @Failure      401 {object} response.APIResponse "Not authenticated"
// @Router       /servers/{serverType}/filters [put]
func (h *DiscoveryHandler) UpdateFilter(c *gin.Context) {
	userID, _ := c.Get("userID")
	serverType := c.Param("serverType")

	var req dto.UpdateFilterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		apperror.HandleError(c, apperror.BadRequest(err.Error()))
		return
	}

	filter, err := h.discoveryService.UpdateFilter(c.Request.Context(), userID.(string), serverType, &req)
	if err != nil {
		apperror.HandleError(c, apperror.InternalError(err))
		return
	}
	response.JSON(c, http.StatusOK, filter)
}

// ========== Discovery Endpoints ==========

// GetDiscovery godoc
// @Summary      Get discovery cards
// @Description  Get current batch of discovery cards for a server type
// @Tags         discovery
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        serverType path string true "Server type (partner/friend/growth)"
// @Success      200 {object} response.APIResponse{data=dto.DiscoveryBatchResponse} "Discovery batch"
// @Failure      401 {object} response.APIResponse "Not authenticated"
// @Router       /servers/{serverType}/discover [get]
func (h *DiscoveryHandler) GetDiscovery(c *gin.Context) {
	userID, _ := c.Get("userID")
	serverType := c.Param("serverType")

	batch, err := h.discoveryService.GetDiscoveryBatch(c.Request.Context(), userID.(string), serverType)
	if err != nil {
		apperror.HandleError(c, apperror.InternalError(err))
		return
	}
	response.JSON(c, http.StatusOK, batch)
}

// RefreshDiscovery godoc
// @Summary      Refresh discovery
// @Description  Trigger a new discovery batch (respects global 24h cooldown)
// @Tags         discovery
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        serverType path string true "Server type (partner/friend/growth)"
// @Success      200 {object} response.APIResponse{data=dto.DiscoveryBatchResponse} "New batch"
// @Failure      400 {object} response.APIResponse "Refresh not available yet"
// @Failure      401 {object} response.APIResponse "Not authenticated"
// @Router       /servers/{serverType}/discover/refresh [post]
func (h *DiscoveryHandler) RefreshDiscovery(c *gin.Context) {
	userID, _ := c.Get("userID")
	serverType := c.Param("serverType")

	batch, err := h.discoveryService.RefreshDiscovery(c.Request.Context(), userID.(string), serverType)
	if err != nil {
		if strings.Contains(err.Error(), "not available") || strings.Contains(err.Error(), "capacity") {
			apperror.HandleError(c, apperror.RateLimited(err.Error()))
			return
		}
		apperror.HandleError(c, apperror.InternalError(err))
		return
	}
	response.JSON(c, http.StatusOK, batch)
}

// GetRefreshStatus godoc
// @Summary      Get refresh status
// @Description  Check if discovery refresh is available
// @Tags         discovery
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200 {object} response.APIResponse{data=dto.RefreshStatusResponse} "Refresh status"
// @Failure      401 {object} response.APIResponse "Not authenticated"
// @Router       /discover/refresh-status [get]
func (h *DiscoveryHandler) GetRefreshStatus(c *gin.Context) {
	userID, _ := c.Get("userID")

	status, err := h.discoveryService.GetRefreshStatus(c.Request.Context(), userID.(string))
	if err != nil {
		apperror.HandleError(c, apperror.InternalError(err))
		return
	}
	response.JSON(c, http.StatusOK, status)
}

// ========== Interest Endpoints ==========

// ExpressInterest godoc
// @Summary      Express interest
// @Description  Express interest in a user from a discovery card
// @Tags         interests
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        request body dto.ExpressInterestRequest true "Interest data"
// @Success      201 {object} response.APIResponse{data=dto.InterestResponse} "Interest created"
// @Failure      400 {object} response.APIResponse "Invalid request or already interested"
// @Failure      401 {object} response.APIResponse "Not authenticated"
// @Router       /interests [post]
func (h *DiscoveryHandler) ExpressInterest(c *gin.Context) {
	userID, _ := c.Get("userID")

	var req dto.ExpressInterestRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		apperror.HandleError(c, apperror.BadRequest(err.Error()))
		return
	}

	interest, err := h.matchingService.ExpressInterest(c.Request.Context(), userID.(string), &req)
	if err != nil {
		apperror.HandleError(c, apperror.BadRequest(err.Error()))
		return
	}
	response.JSON(c, http.StatusCreated, interest)
}

// GetSentInterests godoc
// @Summary      Get sent interests
// @Description  Get interests sent by the authenticated user
// @Tags         interests
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200 {object} response.APIResponse{data=[]dto.InterestResponse} "Sent interests"
// @Failure      401 {object} response.APIResponse "Not authenticated"
// @Router       /interests/sent [get]
func (h *DiscoveryHandler) GetSentInterests(c *gin.Context) {
	userID, _ := c.Get("userID")

	interests, err := h.matchingService.GetSentInterests(c.Request.Context(), userID.(string))
	if err != nil {
		apperror.HandleError(c, apperror.InternalError(err))
		return
	}
	response.JSON(c, http.StatusOK, interests)
}

// GetReceivedInterests godoc
// @Summary      Get received interests
// @Description  Get pending interests received by the authenticated user
// @Tags         interests
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200 {object} response.APIResponse{data=[]dto.InterestResponse} "Received interests"
// @Failure      401 {object} response.APIResponse "Not authenticated"
// @Router       /interests/received [get]
func (h *DiscoveryHandler) GetReceivedInterests(c *gin.Context) {
	userID, _ := c.Get("userID")

	interests, err := h.matchingService.GetReceivedInterests(c.Request.Context(), userID.(string))
	if err != nil {
		apperror.HandleError(c, apperror.InternalError(err))
		return
	}
	response.JSON(c, http.StatusOK, interests)
}

// ========== Connection Endpoints ==========

// GetConnections godoc
// @Summary      Get connections
// @Description  Get all active connections for the authenticated user
// @Tags         connections
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200 {object} response.APIResponse{data=[]dto.ConnectionResponse} "Connections"
// @Failure      401 {object} response.APIResponse "Not authenticated"
// @Router       /connections [get]
func (h *DiscoveryHandler) GetConnections(c *gin.Context) {
	userID, _ := c.Get("userID")

	connections, err := h.matchingService.GetConnections(c.Request.Context(), userID.(string))
	if err != nil {
		apperror.HandleError(c, apperror.InternalError(err))
		return
	}
	response.JSON(c, http.StatusOK, connections)
}

// GetConnection godoc
// @Summary      Get connection details
// @Description  Get detailed information about a specific connection
// @Tags         connections
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id path string true "Connection ID"
// @Success      200 {object} response.APIResponse{data=dto.ConnectionDetailResponse} "Connection details"
// @Failure      401 {object} response.APIResponse "Not authenticated"
// @Failure      404 {object} response.APIResponse "Connection not found"
// @Router       /connections/{id} [get]
func (h *DiscoveryHandler) GetConnection(c *gin.Context) {
	userID, _ := c.Get("userID")
	connectionID := c.Param("id")

	connection, err := h.matchingService.GetConnection(c.Request.Context(), userID.(string), connectionID)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			apperror.HandleError(c, apperror.NotFound("connection"))
			return
		}
		apperror.HandleError(c, apperror.InternalError(err))
		return
	}
	response.JSON(c, http.StatusOK, connection)
}

// TerminateConnection godoc
// @Summary      Terminate connection
// @Description  Terminate an active connection
// @Tags         connections
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id path string true "Connection ID"
// @Success      200 {object} response.APIResponse{data=map[string]string} "Connection terminated"
// @Failure      401 {object} response.APIResponse "Not authenticated"
// @Router       /connections/{id} [delete]
func (h *DiscoveryHandler) TerminateConnection(c *gin.Context) {
	userID, _ := c.Get("userID")
	connectionID := c.Param("id")

	if err := h.matchingService.TerminateConnection(c.Request.Context(), userID.(string), connectionID); err != nil {
		apperror.HandleError(c, apperror.InternalError(err))
		return
	}
	response.JSON(c, http.StatusOK, gin.H{"message": "Connection terminated successfully"})
}
