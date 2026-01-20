// internal/safety/dto/safety_dto.go
package dto

import "time"

// BlockUserRequest is the request to block a user
// @Description Block a user
type BlockUserRequest struct {
	UserID string `json:"userId" validate:"required" example:"550e8400-e29b-41d4-a716-446655440000"`
	Reason string `json:"reason,omitempty" validate:"max=500" example:"Inappropriate messages"`
}

// BlockResponse represents a user block
// @Description User block record
type BlockResponse struct {
	ID            string     `json:"id" example:"550e8400-e29b-41d4-a716-446655440000"`
	BlockedUserID string     `json:"blockedUserId" example:"550e8400-e29b-41d4-a716-446655440000"`
	BlockedName   string     `json:"blockedName" example:"John"`
	CreatedAt     time.Time  `json:"createdAt" example:"2024-01-01T00:00:00Z"`
	UnblockedAt   *time.Time `json:"unblockedAt,omitempty"`
}

// UnblockUserRequest is the request to unblock a user
// @Description Unblock a user
type UnblockUserRequest struct {
	UserID string `json:"userId" validate:"required" example:"550e8400-e29b-41d4-a716-446655440000"`
}

// ReportUserRequest is the request to report a user
// @Description Report a user for policy violation
type ReportUserRequest struct {
	UserID        string `json:"userId" validate:"required" example:"550e8400-e29b-41d4-a716-446655440000"`
	Reason        string `json:"reason" validate:"required,oneof=inappropriate_content harassment spam fake_profile underage offensive_behavior other" example:"harassment"`
	Description   string `json:"description,omitempty" validate:"max=1000" example:"User sent threatening messages"`
	ReferenceType string `json:"referenceType,omitempty" validate:"max=50" example:"message"`
	ReferenceID   string `json:"referenceId,omitempty" validate:"max=36" example:"msg_123"`
}

// ReportResponse represents a user report
// @Description User report record
type ReportResponse struct {
	ID             string    `json:"id" example:"550e8400-e29b-41d4-a716-446655440000"`
	ReportedUserID string    `json:"reportedUserId" example:"550e8400-e29b-41d4-a716-446655440000"`
	Reason         string    `json:"reason" example:"harassment"`
	Description    string    `json:"description,omitempty" example:"User sent threatening messages"`
	Status         string    `json:"status" example:"pending"`
	CreatedAt      time.Time `json:"createdAt" example:"2024-01-01T00:00:00Z"`
}

// BlockedUsersResponse returns list of blocked users
// @Description List of blocked users
type BlockedUsersResponse struct {
	Blocks []BlockResponse `json:"blocks"`
	Total  int             `json:"total" example:"3"`
}

// BlockStatusResponse returns whether a user is blocked
// @Description Block status check result
type BlockStatusResponse struct {
	IsBlocked   bool `json:"isBlocked" example:"true"`
	IsBlockedBy bool `json:"isBlockedBy" example:"false"`
}
