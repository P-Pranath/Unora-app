// internal/admin/handlers/extended_admin_handler.go
package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/UnoraApp/be/internal/admin/dto"
	"github.com/UnoraApp/be/internal/admin/services"
	"github.com/UnoraApp/be/pkg/response"
)

// ExtendedAdminHandler handles extended admin HTTP requests
type ExtendedAdminHandler struct {
	serverService         *services.ServerManagementService
	connectionService     *services.ConnectionManagementService
	streakService         *services.StreakManagementService
	paymentService        *services.PaymentManagementService
	photoService          *services.PhotoModerationService
	auditService          *services.AuditLogService
	revealMilestoneService *services.RevealMilestoneService
}

// NewExtendedAdminHandler creates a new extended admin handler
func NewExtendedAdminHandler(
	serverService *services.ServerManagementService,
	connectionService *services.ConnectionManagementService,
	streakService *services.StreakManagementService,
	paymentService *services.PaymentManagementService,
	photoService *services.PhotoModerationService,
	auditService *services.AuditLogService,
	revealMilestoneService *services.RevealMilestoneService,
) *ExtendedAdminHandler {
	return &ExtendedAdminHandler{
		serverService:         serverService,
		connectionService:     connectionService,
		streakService:         streakService,
		paymentService:        paymentService,
		photoService:          photoService,
		auditService:          auditService,
		revealMilestoneService: revealMilestoneService,
	}
}

// ========== SERVER MANAGEMENT ==========

// ListServers godoc
// @Summary      List servers
// @Description  Get all servers
// @Tags         admin
// @Security     AdminAPIKey
// @Success      200 {object} response.APIResponse "Servers"
// @Router       /admin/servers [get]
func (h *ExtendedAdminHandler) ListServers(c *gin.Context) {
	servers, err := h.serverService.ListServers(c.Request.Context())
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "LIST_FAILED", err.Error())
		return
	}
	response.JSON(c, http.StatusOK, servers)
}

// CreateServer godoc
// @Summary      Create server
// @Description  Create a new server
// @Tags         admin
// @Security     AdminAPIKey
// @Param        request body dto.CreateServerRequest true "Server details"
// @Success      201 {object} response.APIResponse "Server created"
// @Router       /admin/servers [post]
func (h *ExtendedAdminHandler) CreateServer(c *gin.Context) {
	var req dto.CreateServerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "INVALID_REQUEST", err.Error())
		return
	}
	srv, err := h.serverService.CreateServer(c.Request.Context(), &req)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "CREATE_FAILED", err.Error())
		return
	}
	response.JSON(c, http.StatusCreated, srv)
}

// UpdateServer godoc
// @Summary      Update server
// @Description  Update a server
// @Tags         admin
// @Security     AdminAPIKey
// @Param        serverId path string true "Server ID"
// @Param        request body dto.UpdateServerRequest true "Server updates"
// @Success      200 {object} response.APIResponse "Server updated"
// @Router       /admin/servers/{serverId} [patch]
func (h *ExtendedAdminHandler) UpdateServer(c *gin.Context) {
	serverID := c.Param("serverId")
	var req dto.UpdateServerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "INVALID_REQUEST", err.Error())
		return
	}
	if err := h.serverService.UpdateServer(c.Request.Context(), serverID, &req); err != nil {
		response.Error(c, http.StatusBadRequest, "UPDATE_FAILED", err.Error())
		return
	}
	response.JSON(c, http.StatusOK, gin.H{"message": "Server updated"})
}

// DeleteServer godoc
// @Summary      Delete server
// @Description  Delete a server
// @Tags         admin
// @Security     AdminAPIKey
// @Param        serverId path string true "Server ID"
// @Success      200 {object} response.APIResponse "Server deleted"
// @Router       /admin/servers/{serverId} [delete]
func (h *ExtendedAdminHandler) DeleteServer(c *gin.Context) {
	serverID := c.Param("serverId")
	if err := h.serverService.DeleteServer(c.Request.Context(), serverID); err != nil {
		response.Error(c, http.StatusBadRequest, "DELETE_FAILED", err.Error())
		return
	}
	response.JSON(c, http.StatusOK, gin.H{"message": "Server deleted"})
}

// ========== CONNECTION MANAGEMENT ==========

// ListConnections godoc
// @Summary      List connections
// @Description  Get paginated list of connections
// @Tags         admin
// @Security     AdminAPIKey
// @Param        page query int false "Page number" default(1)
// @Param        pageSize query int false "Page size" default(20)
// @Param        status query string false "Connection status filter"
// @Success      200 {object} response.APIResponse{data=dto.AdminConnectionListResponse} "Connections"
// @Router       /admin/connections [get]
func (h *ExtendedAdminHandler) ListConnections(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "20"))
	status := c.Query("status")
	connections, err := h.connectionService.ListConnections(c.Request.Context(), page, pageSize, status, "")
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "LIST_FAILED", err.Error())
		return
	}
	response.JSON(c, http.StatusOK, connections)
}

// GetConnection godoc
// @Summary      Get connection
// @Description  Get connection details
// @Tags         admin
// @Security     AdminAPIKey
// @Param        connectionId path string true "Connection ID"
// @Success      200 {object} response.APIResponse{data=dto.AdminConnectionResponse} "Connection"
// @Router       /admin/connections/{connectionId} [get]
func (h *ExtendedAdminHandler) GetConnection(c *gin.Context) {
	connectionID := c.Param("connectionId")
	conn, err := h.connectionService.GetConnection(c.Request.Context(), connectionID)
	if err != nil {
		response.Error(c, http.StatusNotFound, "NOT_FOUND", err.Error())
		return
	}
	response.JSON(c, http.StatusOK, conn)
}

// TerminateConnection godoc
// @Summary      Terminate connection
// @Description  Terminate a connection
// @Tags         admin
// @Security     AdminAPIKey
// @Param        connectionId path string true "Connection ID"
// @Param        request body dto.TerminateConnectionRequest true "Termination reason"
// @Success      200 {object} response.APIResponse "Connection terminated"
// @Router       /admin/connections/{connectionId}/terminate [post]
func (h *ExtendedAdminHandler) TerminateConnection(c *gin.Context) {
	connectionID := c.Param("connectionId")
	var req dto.TerminateConnectionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "INVALID_REQUEST", err.Error())
		return
	}
	if err := h.connectionService.TerminateConnection(c.Request.Context(), connectionID, &req); err != nil {
		response.Error(c, http.StatusBadRequest, "TERMINATE_FAILED", err.Error())
		return
	}
	response.JSON(c, http.StatusOK, gin.H{"message": "Connection terminated"})
}

// ========== STREAK MANAGEMENT ==========

// ListStreaks godoc
// @Summary      List streaks
// @Description  Get paginated list of streaks
// @Tags         admin
// @Security     AdminAPIKey
// @Param        page query int false "Page number" default(1)
// @Param        pageSize query int false "Page size" default(20)
// @Param        state query string false "Streak state filter"
// @Success      200 {object} response.APIResponse{data=dto.AdminStreakListResponse} "Streaks"
// @Router       /admin/streaks [get]
func (h *ExtendedAdminHandler) ListStreaks(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "20"))
	state := c.Query("state")
	streaks, err := h.streakService.ListStreaks(c.Request.Context(), page, pageSize, state)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "LIST_FAILED", err.Error())
		return
	}
	response.JSON(c, http.StatusOK, streaks)
}

// GetStreak godoc
// @Summary      Get streak
// @Description  Get streak details
// @Tags         admin
// @Security     AdminAPIKey
// @Param        streakId path string true "Streak ID"
// @Success      200 {object} response.APIResponse{data=dto.AdminStreakResponse} "Streak"
// @Router       /admin/streaks/{streakId} [get]
func (h *ExtendedAdminHandler) GetStreak(c *gin.Context) {
	streakID := c.Param("streakId")
	streak, err := h.streakService.GetStreak(c.Request.Context(), streakID)
	if err != nil {
		response.Error(c, http.StatusNotFound, "NOT_FOUND", err.Error())
		return
	}
	response.JSON(c, http.StatusOK, streak)
}

// AdjustStreak godoc
// @Summary      Adjust streak
// @Description  Adjust streak day
// @Tags         admin
// @Security     AdminAPIKey
// @Param        streakId path string true "Streak ID"
// @Param        request body dto.AdjustStreakRequest true "Adjustment details"
// @Success      200 {object} response.APIResponse "Streak adjusted"
// @Router       /admin/streaks/{streakId}/adjust [post]
func (h *ExtendedAdminHandler) AdjustStreak(c *gin.Context) {
	streakID := c.Param("streakId")
	var req dto.AdjustStreakRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "INVALID_REQUEST", err.Error())
		return
	}
	if err := h.streakService.AdjustStreak(c.Request.Context(), streakID, &req); err != nil {
		response.Error(c, http.StatusBadRequest, "ADJUST_FAILED", err.Error())
		return
	}
	response.JSON(c, http.StatusOK, gin.H{"message": "Streak adjusted"})
}

// ResetStreak godoc
// @Summary      Reset streak
// @Description  Reset streak to day 0
// @Tags         admin
// @Security     AdminAPIKey
// @Param        streakId path string true "Streak ID"
// @Param        request body dto.ResetStreakRequest true "Reset reason"
// @Success      200 {object} response.APIResponse "Streak reset"
// @Router       /admin/streaks/{streakId}/reset [post]
func (h *ExtendedAdminHandler) ResetStreak(c *gin.Context) {
	streakID := c.Param("streakId")
	var req dto.ResetStreakRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "INVALID_REQUEST", err.Error())
		return
	}
	if err := h.streakService.ResetStreak(c.Request.Context(), streakID, &req); err != nil {
		response.Error(c, http.StatusBadRequest, "RESET_FAILED", err.Error())
		return
	}
	response.JSON(c, http.StatusOK, gin.H{"message": "Streak reset"})
}

// ========== PAYMENT MANAGEMENT ==========

// ListPaymentOrders godoc
// @Summary      List payment orders
// @Description  Get paginated list of payment orders
// @Tags         admin
// @Security     AdminAPIKey
// @Param        page query int false "Page number" default(1)
// @Param        pageSize query int false "Page size" default(20)
// @Param        status query string false "Order status filter"
// @Success      200 {object} response.APIResponse{data=dto.AdminPaymentListResponse} "Payment orders"
// @Router       /admin/payments [get]
func (h *ExtendedAdminHandler) ListPaymentOrders(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "20"))
	status := c.Query("status")
	orders, err := h.paymentService.ListPaymentOrders(c.Request.Context(), page, pageSize, status)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "LIST_FAILED", err.Error())
		return
	}
	response.JSON(c, http.StatusOK, orders)
}

// RefundPayment godoc
// @Summary      Refund payment
// @Description  Mark payment as refunded and deduct credits
// @Tags         admin
// @Security     AdminAPIKey
// @Param        orderId path string true "Order ID"
// @Param        request body dto.RefundPaymentRequest true "Refund reason"
// @Success      200 {object} response.APIResponse "Payment refunded"
// @Router       /admin/payments/{orderId}/refund [post]
func (h *ExtendedAdminHandler) RefundPayment(c *gin.Context) {
	orderID := c.Param("orderId")
	var req dto.RefundPaymentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "INVALID_REQUEST", err.Error())
		return
	}
	if err := h.paymentService.RefundPayment(c.Request.Context(), orderID, &req); err != nil {
		response.Error(c, http.StatusBadRequest, "REFUND_FAILED", err.Error())
		return
	}
	response.JSON(c, http.StatusOK, gin.H{"message": "Payment refunded"})
}

// ========== PHOTO MODERATION ==========

// ListPendingPhotos godoc
// @Summary      List pending photos
// @Description  Get photos awaiting moderation
// @Tags         admin
// @Security     AdminAPIKey
// @Param        page query int false "Page number" default(1)
// @Param        pageSize query int false "Page size" default(20)
// @Success      200 {object} response.APIResponse{data=dto.AdminPhotoListResponse} "Pending photos"
// @Router       /admin/photos/pending [get]
func (h *ExtendedAdminHandler) ListPendingPhotos(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "20"))
	photos, err := h.photoService.ListPendingPhotos(c.Request.Context(), page, pageSize)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "LIST_FAILED", err.Error())
		return
	}
	response.JSON(c, http.StatusOK, photos)
}

// ListFlaggedPhotos godoc
// @Summary      List flagged photos
// @Description  Get flagged photos
// @Tags         admin
// @Security     AdminAPIKey
// @Param        page query int false "Page number" default(1)
// @Param        pageSize query int false "Page size" default(20)
// @Success      200 {object} response.APIResponse{data=dto.AdminPhotoListResponse} "Flagged photos"
// @Router       /admin/photos/flagged [get]
func (h *ExtendedAdminHandler) ListFlaggedPhotos(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "20"))
	photos, err := h.photoService.ListFlaggedPhotos(c.Request.Context(), page, pageSize)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "LIST_FAILED", err.Error())
		return
	}
	response.JSON(c, http.StatusOK, photos)
}

// ApprovePhoto godoc
// @Summary      Approve photo
// @Description  Approve a photo
// @Tags         admin
// @Security     AdminAPIKey
// @Param        photoId path string true "Photo ID"
// @Success      200 {object} response.APIResponse "Photo approved"
// @Router       /admin/photos/{photoId}/approve [post]
func (h *ExtendedAdminHandler) ApprovePhoto(c *gin.Context) {
	photoID := c.Param("photoId")
	if err := h.photoService.ApprovePhoto(c.Request.Context(), photoID); err != nil {
		response.Error(c, http.StatusBadRequest, "APPROVE_FAILED", err.Error())
		return
	}
	response.JSON(c, http.StatusOK, gin.H{"message": "Photo approved"})
}

// RejectPhoto godoc
// @Summary      Reject photo
// @Description  Reject and delete a photo
// @Tags         admin
// @Security     AdminAPIKey
// @Param        photoId path string true "Photo ID"
// @Param        request body dto.RejectPhotoRequest true "Rejection reason"
// @Success      200 {object} response.APIResponse "Photo rejected"
// @Router       /admin/photos/{photoId}/reject [post]
func (h *ExtendedAdminHandler) RejectPhoto(c *gin.Context) {
	photoID := c.Param("photoId")
	var req dto.RejectPhotoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "INVALID_REQUEST", err.Error())
		return
	}
	if err := h.photoService.RejectPhoto(c.Request.Context(), photoID, &req); err != nil {
		response.Error(c, http.StatusBadRequest, "REJECT_FAILED", err.Error())
		return
	}
	response.JSON(c, http.StatusOK, gin.H{"message": "Photo rejected"})
}

// ========== REVEAL MILESTONE MANAGEMENT ==========

// ListRevealMilestones godoc
// @Summary      List reveal milestones
// @Description  Get all reveal milestones
// @Tags         admin
// @Security     AdminAPIKey
// @Param        includeInactive query bool false "Include inactive" default(false)
// @Success      200 {object} response.APIResponse "Reveal milestones"
// @Router       /admin/reveal-milestones [get]
func (h *ExtendedAdminHandler) ListRevealMilestones(c *gin.Context) {
	includeInactive := c.Query("includeInactive") == "true"
	milestones, err := h.revealMilestoneService.ListRevealMilestones(c.Request.Context(), includeInactive)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "LIST_FAILED", err.Error())
		return
	}
	response.JSON(c, http.StatusOK, milestones)
}

// CreateRevealMilestone godoc
// @Summary      Create reveal milestone
// @Description  Create a new reveal milestone
// @Tags         admin
// @Security     AdminAPIKey
// @Param        request body dto.CreateRevealMilestoneRequest true "Milestone details"
// @Success      201 {object} response.APIResponse "Milestone created"
// @Router       /admin/reveal-milestones [post]
func (h *ExtendedAdminHandler) CreateRevealMilestone(c *gin.Context) {
	var req dto.CreateRevealMilestoneRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "INVALID_REQUEST", err.Error())
		return
	}
	milestone, err := h.revealMilestoneService.CreateRevealMilestone(c.Request.Context(), &req)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "CREATE_FAILED", err.Error())
		return
	}
	response.JSON(c, http.StatusCreated, milestone)
}

// UpdateRevealMilestone godoc
// @Summary      Update reveal milestone
// @Description  Update a reveal milestone
// @Tags         admin
// @Security     AdminAPIKey
// @Param        milestoneId path string true "Milestone ID"
// @Param        request body dto.UpdateRevealMilestoneRequest true "Milestone updates"
// @Success      200 {object} response.APIResponse "Milestone updated"
// @Router       /admin/reveal-milestones/{milestoneId} [patch]
func (h *ExtendedAdminHandler) UpdateRevealMilestone(c *gin.Context) {
	milestoneID := c.Param("milestoneId")
	var req dto.UpdateRevealMilestoneRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "INVALID_REQUEST", err.Error())
		return
	}
	if err := h.revealMilestoneService.UpdateRevealMilestone(c.Request.Context(), milestoneID, &req); err != nil {
		response.Error(c, http.StatusBadRequest, "UPDATE_FAILED", err.Error())
		return
	}
	response.JSON(c, http.StatusOK, gin.H{"message": "Milestone updated"})
}

// DeleteRevealMilestone godoc
// @Summary      Delete reveal milestone
// @Description  Soft delete a reveal milestone
// @Tags         admin
// @Security     AdminAPIKey
// @Param        milestoneId path string true "Milestone ID"
// @Success      200 {object} response.APIResponse "Milestone deleted"
// @Router       /admin/reveal-milestones/{milestoneId} [delete]
func (h *ExtendedAdminHandler) DeleteRevealMilestone(c *gin.Context) {
	milestoneID := c.Param("milestoneId")
	if err := h.revealMilestoneService.DeleteRevealMilestone(c.Request.Context(), milestoneID); err != nil {
		response.Error(c, http.StatusBadRequest, "DELETE_FAILED", err.Error())
		return
	}
	response.JSON(c, http.StatusOK, gin.H{"message": "Milestone deleted"})
}

// ========== AUDIT LOGS ==========

// ListAuditLogs godoc
// @Summary      List audit logs
// @Description  Get paginated list of audit logs
// @Tags         admin
// @Security     AdminAPIKey
// @Param        page query int false "Page number" default(1)
// @Param        pageSize query int false "Page size" default(50)
// @Param        action query string false "Action filter"
// @Param        resourceType query string false "Resource type filter"
// @Success      200 {object} response.APIResponse{data=dto.AuditLogListResponse} "Audit logs"
// @Router       /admin/audit-logs [get]
func (h *ExtendedAdminHandler) ListAuditLogs(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "50"))
	action := c.Query("action")
	resourceType := c.Query("resourceType")
	logs, err := h.auditService.ListAuditLogs(c.Request.Context(), page, pageSize, action, resourceType)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "LIST_FAILED", err.Error())
		return
	}
	response.JSON(c, http.StatusOK, logs)
}

// GetAuditLog godoc
// @Summary      Get audit log
// @Description  Get audit log details
// @Tags         admin
// @Security     AdminAPIKey
// @Param        logId path string true "Log ID"
// @Success      200 {object} response.APIResponse{data=dto.AuditLogResponse} "Audit log"
// @Router       /admin/audit-logs/{logId} [get]
func (h *ExtendedAdminHandler) GetAuditLog(c *gin.Context) {
	logID := c.Param("logId")
	log, err := h.auditService.GetAuditLog(c.Request.Context(), logID)
	if err != nil {
		response.Error(c, http.StatusNotFound, "NOT_FOUND", err.Error())
		return
	}
	response.JSON(c, http.StatusOK, log)
}
