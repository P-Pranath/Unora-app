// internal/admin/dto/extended_admin_dto.go
package dto

import "time"

// ===== SERVER MANAGEMENT =====

// CreateServerRequest creates a new server
// @Description Create a new server type
type CreateServerRequest struct {
	ServerType  string `json:"serverType" validate:"required,oneof=partner friend growth" example:"partner"`
	DisplayName string `json:"displayName" validate:"required,max=50" example:"Partner"`
	IconName    string `json:"iconName" validate:"required,max=50" example:"heart"`
	SortOrder   int    `json:"sortOrder" example:"1"`
}

// UpdateServerRequest updates a server
// @Description Update server details
type UpdateServerRequest struct {
	DisplayName *string `json:"displayName,omitempty" validate:"max=50" example:"Partner"`
	IconName    *string `json:"iconName,omitempty" validate:"max=50" example:"heart"`
	SortOrder   *int    `json:"sortOrder,omitempty" example:"1"`
}

// ===== REVEAL MILESTONE MANAGEMENT =====

// CreateRevealMilestoneRequest creates a reveal milestone
// @Description Create a new reveal milestone
type CreateRevealMilestoneRequest struct {
	RevealNumber int    `json:"revealNumber" validate:"required,min=1" example:"1"`
	DayRequired  int    `json:"dayRequired" validate:"required,min=1" example:"5"`
	RevealType   string `json:"revealType" validate:"required" example:"personality"`
	Title        string `json:"title" validate:"required,max=100" example:"Personality Reveal"`
	Description  string `json:"description,omitempty" validate:"max=500"`
	IconName     string `json:"iconName,omitempty" validate:"max=50" example:"sparkles"`
	CreditCost   int    `json:"creditCost" example:"0"`
}

// UpdateRevealMilestoneRequest updates a reveal milestone
// @Description Update reveal milestone
type UpdateRevealMilestoneRequest struct {
	DayRequired *int    `json:"dayRequired,omitempty" example:"5"`
	Title       *string `json:"title,omitempty" example:"Personality Reveal"`
	Description *string `json:"description,omitempty"`
	IconName    *string `json:"iconName,omitempty" example:"sparkles"`
	CreditCost  *int    `json:"creditCost,omitempty" example:"0"`
	IsActive    *bool   `json:"isActive,omitempty" example:"true"`
}

// ===== CONNECTION MANAGEMENT =====

// AdminConnectionResponse detailed connection for admin
// @Description Connection details for admin
type AdminConnectionResponse struct {
	ID               string     `json:"id" example:"550e8400-e29b-41d4-a716-446655440000"`
	UserAID          string     `json:"userAId"`
	UserAName        string     `json:"userAName" example:"John"`
	UserBID          string     `json:"userBId"`
	UserBName        string     `json:"userBName" example:"Jane"`
	ServerType       string     `json:"serverType" example:"partner"`
	ConnectionStatus string     `json:"connectionStatus" example:"active"`
	StreakDay        int        `json:"streakDay" example:"7"`
	StreakState      string     `json:"streakState" example:"active"`
	CreatedAt        time.Time  `json:"createdAt"`
	TerminatedAt     *time.Time `json:"terminatedAt,omitempty"`
}

// AdminConnectionListResponse paginated connection list
// @Description Paginated list of connections
type AdminConnectionListResponse struct {
	Connections []AdminConnectionResponse `json:"connections"`
	Total       int                       `json:"total"`
	Page        int                       `json:"page"`
	PageSize    int                       `json:"pageSize"`
	TotalPages  int                       `json:"totalPages"`
}

// TerminateConnectionRequest terminates a connection
// @Description Terminate a connection
type TerminateConnectionRequest struct {
	Reason string `json:"reason" validate:"required,max=500" example:"Moderation action"`
}

// ===== STREAK MANAGEMENT =====

// AdminStreakResponse detailed streak for admin
// @Description Streak details for admin
type AdminStreakResponse struct {
	ID             string    `json:"id" example:"550e8400-e29b-41d4-a716-446655440000"`
	ConnectionID   string    `json:"connectionId"`
	UserAName      string    `json:"userAName" example:"John"`
	UserBName      string    `json:"userBName" example:"Jane"`
	CurrentDay     int       `json:"currentDay" example:"7"`
	StreakState    string    `json:"streakState" example:"active"`
	LastCheckInAt  *time.Time `json:"lastCheckInAt,omitempty"`
	CreatedAt      time.Time `json:"createdAt"`
}

// AdminStreakListResponse paginated streak list
// @Description Paginated list of streaks
type AdminStreakListResponse struct {
	Streaks    []AdminStreakResponse `json:"streaks"`
	Total      int                   `json:"total"`
	Page       int                   `json:"page"`
	PageSize   int                   `json:"pageSize"`
	TotalPages int                   `json:"totalPages"`
}

// AdjustStreakRequest adjusts streak day
// @Description Adjust streak day
type AdjustStreakRequest struct {
	NewDay int    `json:"newDay" validate:"required,min=0,max=15" example:"7"`
	Reason string `json:"reason" validate:"required,max=500" example:"Support adjustment"`
}

// ResetStreakRequest resets a streak
// @Description Reset a streak
type ResetStreakRequest struct {
	Reason string `json:"reason" validate:"required,max=500" example:"Testing purposes"`
}

// ===== PAYMENT MANAGEMENT =====

// AdminPaymentOrderResponse detailed payment order for admin
// @Description Payment order details for admin
type AdminPaymentOrderResponse struct {
	ID                string     `json:"id"`
	UserID            string     `json:"userId"`
	UserEmail         string     `json:"userEmail,omitempty"`
	PackageID         string     `json:"packageId,omitempty"`
	PackageName       string     `json:"packageName,omitempty"`
	RazorpayOrderID   string     `json:"razorpayOrderId"`
	RazorpayPaymentID string     `json:"razorpayPaymentId,omitempty"`
	Amount            int        `json:"amount" example:"9900"`
	Currency          string     `json:"currency" example:"INR"`
	CreditsToAdd      int        `json:"creditsToAdd" example:"55"`
	OrderStatus       string     `json:"orderStatus" example:"paid"`
	FailureReason     string     `json:"failureReason,omitempty"`
	CreatedAt         time.Time  `json:"createdAt"`
	PaidAt            *time.Time `json:"paidAt,omitempty"`
}

// AdminPaymentListResponse paginated payment order list
// @Description Paginated list of payment orders
type AdminPaymentListResponse struct {
	Orders     []AdminPaymentOrderResponse `json:"orders"`
	Total      int                         `json:"total"`
	Page       int                         `json:"page"`
	PageSize   int                         `json:"pageSize"`
	TotalPages int                         `json:"totalPages"`
}

// RefundPaymentRequest refunds a payment
// @Description Refund a payment
type RefundPaymentRequest struct {
	Reason string `json:"reason" validate:"required,max=500" example:"Customer request"`
}

// ===== PHOTO MODERATION =====

// AdminPhotoResponse photo for moderation
// @Description Photo details for moderation
type AdminPhotoResponse struct {
	ID           string    `json:"id"`
	UserID       string    `json:"userId"`
	UserName     string    `json:"userName"`
	PhotoURL     string    `json:"photoUrl"`
	SlotNumber   int       `json:"slotNumber" example:"1"`
	IsVerified   bool      `json:"isVerified" example:"false"`
	IsFlagged    bool      `json:"isFlagged" example:"false"`
	FlagReason   string    `json:"flagReason,omitempty"`
	UploadedAt   time.Time `json:"uploadedAt"`
}

// AdminPhotoListResponse paginated photo list
// @Description Paginated list of photos for moderation
type AdminPhotoListResponse struct {
	Photos     []AdminPhotoResponse `json:"photos"`
	Total      int                  `json:"total"`
	Page       int                  `json:"page"`
	PageSize   int                  `json:"pageSize"`
	TotalPages int                  `json:"totalPages"`
}

// RejectPhotoRequest rejects a photo
// @Description Reject a photo
type RejectPhotoRequest struct {
	Reason string `json:"reason" validate:"required,max=500" example:"Inappropriate content"`
}

// ===== AUDIT LOGS =====

// AuditLogResponse audit log entry
// @Description Audit log entry
type AuditLogResponse struct {
	ID              string    `json:"id"`
	AdminIdentifier string    `json:"adminIdentifier"`
	Action          string    `json:"action" example:"suspend_user"`
	ResourceType    string    `json:"resourceType" example:"user"`
	ResourceID      string    `json:"resourceId,omitempty"`
	Success         bool      `json:"success"`
	ErrorMessage    string    `json:"errorMessage,omitempty"`
	IPAddress       string    `json:"ipAddress,omitempty"`
	CreatedAt       time.Time `json:"createdAt"`
}

// AuditLogListResponse paginated audit log list
// @Description Paginated list of audit logs
type AuditLogListResponse struct {
	Logs       []AuditLogResponse `json:"logs"`
	Total      int                `json:"total"`
	Page       int                `json:"page"`
	PageSize   int                `json:"pageSize"`
	TotalPages int                `json:"totalPages"`
}
