// internal/admin/dto/admin_dto.go
package dto

import "time"

// ===== USER MANAGEMENT =====

// AdminUserResponse represents a user in admin view
// @Description User details for admin
type AdminUserResponse struct {
	ID               string     `json:"id" example:"550e8400-e29b-41d4-a716-446655440000"`
	Email            string     `json:"email" example:"user@example.com"`
	PhoneNumber      string     `json:"phoneNumber,omitempty" example:"+919876543210"`
	FirstName        string     `json:"firstName" example:"John"`
	LastName         string     `json:"lastName,omitempty" example:"Doe"`
	DateOfBirth      *time.Time `json:"dateOfBirth,omitempty"`
	Gender           string     `json:"gender,omitempty" example:"male"`
	AccountStatus    string     `json:"accountStatus" example:"active"`
	OnboardingStatus string     `json:"onboardingStatus" example:"completed"`
	CreditBalance    int        `json:"creditBalance" example:"150"`
	PhotoCount       int        `json:"photoCount" example:"3"`
	ConnectionCount  int        `json:"connectionCount" example:"5"`
	ReportsReceived  int        `json:"reportsReceived" example:"0"`
	CreatedAt        time.Time  `json:"createdAt" example:"2024-01-01T00:00:00Z"`
	LastActiveAt     *time.Time `json:"lastActiveAt,omitempty"`
	SuspendedAt      *time.Time `json:"suspendedAt,omitempty"`
}

// AdminUserListResponse paginated user list
// @Description Paginated list of users
type AdminUserListResponse struct {
	Users      []AdminUserResponse `json:"users"`
	Total      int                 `json:"total" example:"100"`
	Page       int                 `json:"page" example:"1"`
	PageSize   int                 `json:"pageSize" example:"20"`
	TotalPages int                 `json:"totalPages" example:"5"`
}

// SuspendUserRequest request to suspend a user
// @Description Suspend a user account
type SuspendUserRequest struct {
	Reason   string `json:"reason" validate:"required,max=500" example:"Violation of community guidelines"`
	Duration string `json:"duration" validate:"required" example:"7d"` // "permanent", "7d", "30d", etc.
}

// AdjustCreditsRequest request to adjust user credits
// @Description Adjust user credit balance
type AdjustCreditsRequest struct {
	Amount      int    `json:"amount" validate:"required" example:"50"`
	Reason      string `json:"reason" validate:"required,max=500" example:"Compensation for issue"`
	Description string `json:"description,omitempty" validate:"max=200" example:"Support ticket #123"`
}

// ===== REPORT MANAGEMENT =====

// AdminReportResponse represents a report in admin view
// @Description Report details for admin review
type AdminReportResponse struct {
	ID             string     `json:"id" example:"550e8400-e29b-41d4-a716-446655440000"`
	ReporterID     string     `json:"reporterId" example:"550e8400-e29b-41d4-a716-446655440000"`
	ReporterName   string     `json:"reporterName" example:"John"`
	ReportedID     string     `json:"reportedId" example:"550e8400-e29b-41d4-a716-446655440000"`
	ReportedName   string     `json:"reportedName" example:"Jane"`
	Reason         string     `json:"reason" example:"harassment"`
	Description    string     `json:"description,omitempty" example:"User sent threatening messages"`
	ReferenceType  string     `json:"referenceType,omitempty" example:"message"`
	ReferenceID    string     `json:"referenceId,omitempty" example:"msg_123"`
	Status         string     `json:"status" example:"pending"`
	AdminNotes     string     `json:"adminNotes,omitempty" example:"Reviewed, evidence found"`
	ReviewedAt     *time.Time `json:"reviewedAt,omitempty"`
	CreatedAt      time.Time  `json:"createdAt" example:"2024-01-01T00:00:00Z"`
}

// AdminReportListResponse paginated report list
// @Description Paginated list of reports
type AdminReportListResponse struct {
	Reports    []AdminReportResponse `json:"reports"`
	Total      int                   `json:"total" example:"25"`
	Page       int                   `json:"page" example:"1"`
	PageSize   int                   `json:"pageSize" example:"20"`
	TotalPages int                   `json:"totalPages" example:"2"`
}

// ReviewReportRequest request to review a report
// @Description Review and take action on a report
type ReviewReportRequest struct {
	Status     string `json:"status" validate:"required,oneof=reviewed action_taken dismissed" example:"action_taken"`
	AdminNotes string `json:"adminNotes,omitempty" validate:"max=1000" example:"Evidence confirmed, user suspended"`
	Action     string `json:"action,omitempty" validate:"omitempty,oneof=warn suspend_7d suspend_30d suspend_permanent delete" example:"suspend_7d"`
}

// ===== CONTENT MANAGEMENT =====

// CreateHobbyOptionRequest creates a new hobby option
// @Description Create a new hobby option
type CreateHobbyOptionRequest struct {
	Category    string `json:"category" validate:"required" example:"sports"`
	Name        string `json:"name" validate:"required,max=100" example:"Basketball"`
	IconName    string `json:"iconName,omitempty" validate:"max=50" example:"basketball"`
	Description string `json:"description,omitempty" validate:"max=200"`
}

// UpdateCreditPackageRequest updates a credit package
// @Description Update credit package details
type UpdateCreditPackageRequest struct {
	Name            string   `json:"name,omitempty" validate:"max=100" example:"Starter Pack"`
	Description     string   `json:"description,omitempty"`
	CreditAmount    *int     `json:"creditAmount,omitempty" example:"50"`
	BonusCredits    *int     `json:"bonusCredits,omitempty" example:"5"`
	PriceAmount     *int     `json:"priceAmount,omitempty" example:"9900"`
	DiscountPercent *float64 `json:"discountPercent,omitempty" example:"10"`
	BadgeText       *string  `json:"badgeText,omitempty" example:"Popular"`
	IsPopular       *bool    `json:"isPopular,omitempty" example:"true"`
	IsActive        *bool    `json:"isActive,omitempty" example:"true"`
	SortOrder       *int     `json:"sortOrder,omitempty" example:"1"`
}

// ===== ANALYTICS =====

// DashboardStatsResponse aggregated dashboard stats
// @Description Dashboard statistics for admin
type DashboardStatsResponse struct {
	UserStats       UserStatsResponse       `json:"userStats"`
	PaymentStats    PaymentStatsResponse    `json:"paymentStats"`
	ConnectionStats ConnectionStatsResponse `json:"connectionStats"`
	ReportStats     ReportStatsResponse     `json:"reportStats"`
}

// UserStatsResponse user statistics
// @Description User registration and activity stats
type UserStatsResponse struct {
	TotalUsers       int `json:"totalUsers" example:"10000"`
	ActiveToday      int `json:"activeToday" example:"500"`
	ActiveThisWeek   int `json:"activeThisWeek" example:"2000"`
	NewThisWeek      int `json:"newThisWeek" example:"150"`
	NewThisMonth     int `json:"newThisMonth" example:"500"`
	SuspendedUsers   int `json:"suspendedUsers" example:"25"`
	OnboardingPending int `json:"onboardingPending" example:"100"`
}

// PaymentStatsResponse payment statistics
// @Description Payment and revenue stats
type PaymentStatsResponse struct {
	TotalRevenue       int `json:"totalRevenue" example:"500000"` // In paise
	RevenueThisMonth   int `json:"revenueThisMonth" example:"50000"`
	RevenueToday       int `json:"revenueToday" example:"5000"`
	TotalTransactions  int `json:"totalTransactions" example:"1000"`
	TransactionsToday  int `json:"transactionsToday" example:"50"`
	CreditsInCirculation int `json:"creditsInCirculation" example:"100000"`
}

// ConnectionStatsResponse connection statistics
// @Description Connection and matching stats
type ConnectionStatsResponse struct {
	TotalConnections       int `json:"totalConnections" example:"5000"`
	ActiveConnections      int `json:"activeConnections" example:"3000"`
	ConnectionsThisWeek    int `json:"connectionsThisWeek" example:"200"`
	ActiveStreaks          int `json:"activeStreaks" example:"2500"`
	CompletedStreaks       int `json:"completedStreaks" example:"500"`
	AverageStreakDay       float64 `json:"averageStreakDay" example:"7.5"`
}

// ReportStatsResponse report statistics
// @Description Report and moderation stats
type ReportStatsResponse struct {
	TotalReports       int `json:"totalReports" example:"100"`
	PendingReports     int `json:"pendingReports" example:"15"`
	ReportsThisWeek    int `json:"reportsThisWeek" example:"20"`
	ActionsTaken       int `json:"actionsTaken" example:"50"`
	DismissedReports   int `json:"dismissedReports" example:"30"`
}
