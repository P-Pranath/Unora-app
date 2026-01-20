// internal/admin/routes/admin_routes.go
package routes

import (
	"github.com/gin-gonic/gin"

	ent "github.com/UnoraApp/be/ent/generated"
	"github.com/UnoraApp/be/internal/admin/handlers"
	"github.com/UnoraApp/be/internal/admin/middlewares"
	"github.com/UnoraApp/be/internal/admin/services"
)

// RegisterAdminRoutes registers all admin routes
func RegisterAdminRoutes(
	router *gin.RouterGroup,
	entClient *ent.Client,
) {
	// Create services
	userMgmtService := services.NewUserManagementService(entClient)
	reportMgmtService := services.NewReportManagementService(entClient, userMgmtService)
	analyticsService := services.NewAnalyticsService(entClient)
	contentMgmtService := services.NewContentManagementService(entClient)
	serverService := services.NewServerManagementService(entClient)
	connectionService := services.NewConnectionManagementService(entClient)
	streakService := services.NewStreakManagementService(entClient)
	paymentService := services.NewPaymentManagementService(entClient)
	photoService := services.NewPhotoModerationService(entClient)
	auditService := services.NewAuditLogService(entClient)
	revealMilestoneService := services.NewRevealMilestoneService(entClient)

	// Create handlers
	handler := handlers.NewAdminHandler(
		userMgmtService,
		reportMgmtService,
		analyticsService,
		contentMgmtService,
	)
	extHandler := handlers.NewExtendedAdminHandler(
		serverService,
		connectionService,
		streakService,
		paymentService,
		photoService,
		auditService,
		revealMilestoneService,
	)

	// Admin routes with API key authentication
	admin := router.Group("/admin")
	admin.Use(middlewares.AdminAPIKeyMiddleware())
	{
		// Dashboard
		admin.GET("/dashboard", handler.GetDashboardStats)

		// User management
		admin.GET("/users", handler.ListUsers)
		admin.GET("/users/:userId", handler.GetUser)
		admin.POST("/users/:userId/suspend", handler.SuspendUser)
		admin.POST("/users/:userId/unsuspend", handler.UnsuspendUser)
		admin.DELETE("/users/:userId", handler.DeleteUser)
		admin.POST("/users/:userId/credits", handler.AdjustCredits)

		// Report management
		admin.GET("/reports", handler.ListReports)
		admin.GET("/reports/:reportId", handler.GetReport)
		admin.POST("/reports/:reportId/review", handler.ReviewReport)

		// Content management
		admin.GET("/hobby-options", handler.ListHobbyOptions)
		admin.POST("/hobby-options", handler.CreateHobbyOption)
		admin.DELETE("/hobby-options/:optionId", handler.DeleteHobbyOption)
		admin.GET("/credit-packages", handler.ListCreditPackages)
		admin.PATCH("/credit-packages/:packageId", handler.UpdateCreditPackage)

		// ========== EXTENDED ADMIN ROUTES ==========

		// Server management
		admin.GET("/servers", extHandler.ListServers)
		admin.POST("/servers", extHandler.CreateServer)
		admin.PATCH("/servers/:serverId", extHandler.UpdateServer)
		admin.DELETE("/servers/:serverId", extHandler.DeleteServer)

		// Connection management
		admin.GET("/connections", extHandler.ListConnections)
		admin.GET("/connections/:connectionId", extHandler.GetConnection)
		admin.POST("/connections/:connectionId/terminate", extHandler.TerminateConnection)

		// Streak management
		admin.GET("/streaks", extHandler.ListStreaks)
		admin.GET("/streaks/:streakId", extHandler.GetStreak)
		admin.POST("/streaks/:streakId/adjust", extHandler.AdjustStreak)
		admin.POST("/streaks/:streakId/reset", extHandler.ResetStreak)

		// Payment management
		admin.GET("/payments", extHandler.ListPaymentOrders)
		admin.POST("/payments/:orderId/refund", extHandler.RefundPayment)

		// Photo moderation
		admin.GET("/photos/pending", extHandler.ListPendingPhotos)
		admin.GET("/photos/flagged", extHandler.ListFlaggedPhotos)
		admin.POST("/photos/:photoId/approve", extHandler.ApprovePhoto)
		admin.POST("/photos/:photoId/reject", extHandler.RejectPhoto)

		// Reveal milestone management
		admin.GET("/reveal-milestones", extHandler.ListRevealMilestones)
		admin.POST("/reveal-milestones", extHandler.CreateRevealMilestone)
		admin.PATCH("/reveal-milestones/:milestoneId", extHandler.UpdateRevealMilestone)
		admin.DELETE("/reveal-milestones/:milestoneId", extHandler.DeleteRevealMilestone)

		// Audit logs
		admin.GET("/audit-logs", extHandler.ListAuditLogs)
		admin.GET("/audit-logs/:logId", extHandler.GetAuditLog)
	}
}
