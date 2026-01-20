// internal/admin/handlers/admin_handler.go
package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/UnoraApp/be/internal/admin/dto"
	"github.com/UnoraApp/be/internal/admin/services"
	"github.com/UnoraApp/be/pkg/response"
)

// AdminHandler handles admin HTTP requests
type AdminHandler struct {
	userMgmtService    *services.UserManagementService
	reportMgmtService  *services.ReportManagementService
	analyticsService   *services.AnalyticsService
	contentMgmtService *services.ContentManagementService
}

// NewAdminHandler creates a new admin handler
func NewAdminHandler(
	userMgmtService *services.UserManagementService,
	reportMgmtService *services.ReportManagementService,
	analyticsService *services.AnalyticsService,
	contentMgmtService *services.ContentManagementService,
) *AdminHandler {
	return &AdminHandler{
		userMgmtService:    userMgmtService,
		reportMgmtService:  reportMgmtService,
		analyticsService:   analyticsService,
		contentMgmtService: contentMgmtService,
	}
}

// ========== DASHBOARD ==========

// GetDashboardStats godoc
// @Summary      Get dashboard stats
// @Description  Get aggregated dashboard statistics
// @Tags         admin
// @Accept       json
// @Produce      json
// @Security     AdminAPIKey
// @Success      200 {object} response.APIResponse{data=dto.DashboardStatsResponse} "Dashboard stats"
// @Router       /admin/dashboard [get]
func (h *AdminHandler) GetDashboardStats(c *gin.Context) {
	stats, err := h.analyticsService.GetDashboardStats(c.Request.Context())
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "GET_STATS_FAILED", err.Error())
		return
	}
	response.JSON(c, http.StatusOK, stats)
}

// ========== USER MANAGEMENT ==========

// ListUsers godoc
// @Summary      List users
// @Description  Get paginated list of users with optional filters
// @Tags         admin
// @Accept       json
// @Produce      json
// @Security     AdminAPIKey
// @Param        page query int false "Page number" default(1)
// @Param        pageSize query int false "Page size" default(20)
// @Param        status query string false "Account status filter"
// @Param        search query string false "Search by email or name"
// @Success      200 {object} response.APIResponse{data=dto.AdminUserListResponse} "Users"
// @Router       /admin/users [get]
func (h *AdminHandler) ListUsers(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "20"))
	status := c.Query("status")
	search := c.Query("search")

	users, err := h.userMgmtService.ListUsers(c.Request.Context(), page, pageSize, status, search)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "LIST_USERS_FAILED", err.Error())
		return
	}
	response.JSON(c, http.StatusOK, users)
}

// GetUser godoc
// @Summary      Get user details
// @Description  Get detailed user information
// @Tags         admin
// @Accept       json
// @Produce      json
// @Security     AdminAPIKey
// @Param        userId path string true "User ID"
// @Success      200 {object} response.APIResponse{data=dto.AdminUserResponse} "User details"
// @Router       /admin/users/{userId} [get]
func (h *AdminHandler) GetUser(c *gin.Context) {
	userID := c.Param("userId")

	user, err := h.userMgmtService.GetUser(c.Request.Context(), userID)
	if err != nil {
		response.Error(c, http.StatusNotFound, "USER_NOT_FOUND", err.Error())
		return
	}
	response.JSON(c, http.StatusOK, user)
}

// SuspendUser godoc
// @Summary      Suspend user
// @Description  Suspend a user account
// @Tags         admin
// @Accept       json
// @Produce      json
// @Security     AdminAPIKey
// @Param        userId path string true "User ID"
// @Param        request body dto.SuspendUserRequest true "Suspension details"
// @Success      200 {object} response.APIResponse "User suspended"
// @Router       /admin/users/{userId}/suspend [post]
func (h *AdminHandler) SuspendUser(c *gin.Context) {
	userID := c.Param("userId")

	var req dto.SuspendUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "INVALID_REQUEST", err.Error())
		return
	}

	if err := h.userMgmtService.SuspendUser(c.Request.Context(), userID, &req); err != nil {
		response.Error(c, http.StatusBadRequest, "SUSPEND_FAILED", err.Error())
		return
	}
	response.JSON(c, http.StatusOK, gin.H{"message": "User suspended"})
}

// UnsuspendUser godoc
// @Summary      Unsuspend user
// @Description  Reactivate a suspended user account
// @Tags         admin
// @Accept       json
// @Produce      json
// @Security     AdminAPIKey
// @Param        userId path string true "User ID"
// @Success      200 {object} response.APIResponse "User unsuspended"
// @Router       /admin/users/{userId}/unsuspend [post]
func (h *AdminHandler) UnsuspendUser(c *gin.Context) {
	userID := c.Param("userId")

	if err := h.userMgmtService.UnsuspendUser(c.Request.Context(), userID); err != nil {
		response.Error(c, http.StatusBadRequest, "UNSUSPEND_FAILED", err.Error())
		return
	}
	response.JSON(c, http.StatusOK, gin.H{"message": "User unsuspended"})
}

// DeleteUser godoc
// @Summary      Delete user
// @Description  Soft delete a user account
// @Tags         admin
// @Accept       json
// @Produce      json
// @Security     AdminAPIKey
// @Param        userId path string true "User ID"
// @Success      200 {object} response.APIResponse "User deleted"
// @Router       /admin/users/{userId} [delete]
func (h *AdminHandler) DeleteUser(c *gin.Context) {
	userID := c.Param("userId")

	if err := h.userMgmtService.DeleteUser(c.Request.Context(), userID); err != nil {
		response.Error(c, http.StatusBadRequest, "DELETE_FAILED", err.Error())
		return
	}
	response.JSON(c, http.StatusOK, gin.H{"message": "User deleted"})
}

// AdjustCredits godoc
// @Summary      Adjust user credits
// @Description  Add or deduct credits from user balance
// @Tags         admin
// @Accept       json
// @Produce      json
// @Security     AdminAPIKey
// @Param        userId path string true "User ID"
// @Param        request body dto.AdjustCreditsRequest true "Credit adjustment"
// @Success      200 {object} response.APIResponse "Credits adjusted"
// @Router       /admin/users/{userId}/credits [post]
func (h *AdminHandler) AdjustCredits(c *gin.Context) {
	userID := c.Param("userId")

	var req dto.AdjustCreditsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "INVALID_REQUEST", err.Error())
		return
	}

	if err := h.userMgmtService.AdjustCredits(c.Request.Context(), userID, &req); err != nil {
		response.Error(c, http.StatusBadRequest, "ADJUST_FAILED", err.Error())
		return
	}
	response.JSON(c, http.StatusOK, gin.H{"message": "Credits adjusted"})
}

// ========== REPORT MANAGEMENT ==========

// ListReports godoc
// @Summary      List reports
// @Description  Get paginated list of user reports
// @Tags         admin
// @Accept       json
// @Produce      json
// @Security     AdminAPIKey
// @Param        page query int false "Page number" default(1)
// @Param        pageSize query int false "Page size" default(20)
// @Param        status query string false "Report status filter"
// @Success      200 {object} response.APIResponse{data=dto.AdminReportListResponse} "Reports"
// @Router       /admin/reports [get]
func (h *AdminHandler) ListReports(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "20"))
	status := c.Query("status")

	reports, err := h.reportMgmtService.ListReports(c.Request.Context(), page, pageSize, status)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "LIST_REPORTS_FAILED", err.Error())
		return
	}
	response.JSON(c, http.StatusOK, reports)
}

// GetReport godoc
// @Summary      Get report details
// @Description  Get detailed report information
// @Tags         admin
// @Accept       json
// @Produce      json
// @Security     AdminAPIKey
// @Param        reportId path string true "Report ID"
// @Success      200 {object} response.APIResponse{data=dto.AdminReportResponse} "Report details"
// @Router       /admin/reports/{reportId} [get]
func (h *AdminHandler) GetReport(c *gin.Context) {
	reportID := c.Param("reportId")

	report, err := h.reportMgmtService.GetReport(c.Request.Context(), reportID)
	if err != nil {
		response.Error(c, http.StatusNotFound, "REPORT_NOT_FOUND", err.Error())
		return
	}
	response.JSON(c, http.StatusOK, report)
}

// ReviewReport godoc
// @Summary      Review report
// @Description  Review a report and optionally take action
// @Tags         admin
// @Accept       json
// @Produce      json
// @Security     AdminAPIKey
// @Param        reportId path string true "Report ID"
// @Param        request body dto.ReviewReportRequest true "Review details"
// @Success      200 {object} response.APIResponse "Report reviewed"
// @Router       /admin/reports/{reportId}/review [post]
func (h *AdminHandler) ReviewReport(c *gin.Context) {
	reportID := c.Param("reportId")

	var req dto.ReviewReportRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "INVALID_REQUEST", err.Error())
		return
	}

	if err := h.reportMgmtService.ReviewReport(c.Request.Context(), reportID, &req); err != nil {
		response.Error(c, http.StatusBadRequest, "REVIEW_FAILED", err.Error())
		return
	}
	response.JSON(c, http.StatusOK, gin.H{"message": "Report reviewed"})
}

// ========== CONTENT MANAGEMENT ==========

// ListHobbyOptions godoc
// @Summary      List hobby options
// @Description  Get all hobby options
// @Tags         admin
// @Accept       json
// @Produce      json
// @Security     AdminAPIKey
// @Param        includeInactive query bool false "Include inactive" default(false)
// @Success      200 {object} response.APIResponse "Hobby options"
// @Router       /admin/hobby-options [get]
func (h *AdminHandler) ListHobbyOptions(c *gin.Context) {
	includeInactive := c.Query("includeInactive") == "true"

	options, err := h.contentMgmtService.ListHobbyOptions(c.Request.Context(), includeInactive)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "LIST_FAILED", err.Error())
		return
	}
	response.JSON(c, http.StatusOK, options)
}

// CreateHobbyOption godoc
// @Summary      Create hobby option
// @Description  Create a new hobby option
// @Tags         admin
// @Accept       json
// @Produce      json
// @Security     AdminAPIKey
// @Param        request body dto.CreateHobbyOptionRequest true "Hobby option details"
// @Success      201 {object} response.APIResponse "Hobby option created"
// @Router       /admin/hobby-options [post]
func (h *AdminHandler) CreateHobbyOption(c *gin.Context) {
	var req dto.CreateHobbyOptionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "INVALID_REQUEST", err.Error())
		return
	}

	option, err := h.contentMgmtService.CreateHobbyOption(c.Request.Context(), &req)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "CREATE_FAILED", err.Error())
		return
	}
	response.JSON(c, http.StatusCreated, option)
}

// DeleteHobbyOption godoc
// @Summary      Delete hobby option
// @Description  Soft delete a hobby option
// @Tags         admin
// @Accept       json
// @Produce      json
// @Security     AdminAPIKey
// @Param        optionId path string true "Hobby Option ID"
// @Success      200 {object} response.APIResponse "Hobby option deleted"
// @Router       /admin/hobby-options/{optionId} [delete]
func (h *AdminHandler) DeleteHobbyOption(c *gin.Context) {
	optionID := c.Param("optionId")

	if err := h.contentMgmtService.DeleteHobbyOption(c.Request.Context(), optionID); err != nil {
		response.Error(c, http.StatusBadRequest, "DELETE_FAILED", err.Error())
		return
	}
	response.JSON(c, http.StatusOK, gin.H{"message": "Hobby option deleted"})
}

// ListCreditPackages godoc
// @Summary      List credit packages
// @Description  Get all credit packages
// @Tags         admin
// @Accept       json
// @Produce      json
// @Security     AdminAPIKey
// @Param        includeInactive query bool false "Include inactive" default(false)
// @Success      200 {object} response.APIResponse "Credit packages"
// @Router       /admin/credit-packages [get]
func (h *AdminHandler) ListCreditPackages(c *gin.Context) {
	includeInactive := c.Query("includeInactive") == "true"

	packages, err := h.contentMgmtService.ListCreditPackages(c.Request.Context(), includeInactive)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "LIST_FAILED", err.Error())
		return
	}
	response.JSON(c, http.StatusOK, packages)
}

// UpdateCreditPackage godoc
// @Summary      Update credit package
// @Description  Update a credit package
// @Tags         admin
// @Accept       json
// @Produce      json
// @Security     AdminAPIKey
// @Param        packageId path string true "Package ID"
// @Param        request body dto.UpdateCreditPackageRequest true "Package updates"
// @Success      200 {object} response.APIResponse "Package updated"
// @Router       /admin/credit-packages/{packageId} [patch]
func (h *AdminHandler) UpdateCreditPackage(c *gin.Context) {
	packageID := c.Param("packageId")

	var req dto.UpdateCreditPackageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "INVALID_REQUEST", err.Error())
		return
	}

	if err := h.contentMgmtService.UpdateCreditPackage(c.Request.Context(), packageID, &req); err != nil {
		response.Error(c, http.StatusBadRequest, "UPDATE_FAILED", err.Error())
		return
	}
	response.JSON(c, http.StatusOK, gin.H{"message": "Package updated"})
}
