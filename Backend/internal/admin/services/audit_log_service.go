// internal/admin/services/audit_log_service.go
package services

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/google/uuid"

	ent "github.com/UnoraApp/be/ent/generated"
	"github.com/UnoraApp/be/ent/generated/auditlog"
	"github.com/UnoraApp/be/internal/admin/dto"
)

// AuditLogService handles audit logging
type AuditLogService struct {
	entClient *ent.Client
}

// NewAuditLogService creates a new audit log service
func NewAuditLogService(entClient *ent.Client) *AuditLogService {
	return &AuditLogService{
		entClient: entClient,
	}
}

// LogAction logs an admin action
func (s *AuditLogService) LogAction(ctx context.Context, adminIdentifier, action, resourceType string, resourceID *string, requestBody interface{}, success bool, errorMsg string, ipAddress, userAgent string) error {
	id := uuid.New().String()

	var reqBodyStr *string
	if requestBody != nil {
		data, _ := json.Marshal(requestBody)
		str := string(data)
		reqBodyStr = &str
	}

	_, err := s.entClient.AuditLog.
		Create().
		SetID(id).
		SetAdminIdentifier(adminIdentifier).
		SetAction(action).
		SetResourceType(resourceType).
		SetNillableResourceID(resourceID).
		SetNillableRequestBody(reqBodyStr).
		SetSuccess(success).
		SetNillableErrorMessage(strPtr(errorMsg)).
		SetNillableIPAddress(strPtr(ipAddress)).
		SetNillableUserAgent(strPtr(userAgent)).
		Save(ctx)
	return err
}

// ListAuditLogs returns paginated list of audit logs
func (s *AuditLogService) ListAuditLogs(ctx context.Context, page, pageSize int, action, resourceType string) (*dto.AuditLogListResponse, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 50
	}
	offset := (page - 1) * pageSize

	query := s.entClient.AuditLog.Query()

	if action != "" {
		query = query.Where(auditlog.ActionEQ(action))
	}
	if resourceType != "" {
		query = query.Where(auditlog.ResourceTypeEQ(resourceType))
	}

	total, err := query.Clone().Count(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to count logs: %w", err)
	}

	logs, err := query.
		Order(ent.Desc(auditlog.FieldCreatedAt)).
		Offset(offset).
		Limit(pageSize).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get logs: %w", err)
	}

	result := make([]dto.AuditLogResponse, len(logs))
	for i, l := range logs {
		result[i] = dto.AuditLogResponse{
			ID:              l.ID,
			AdminIdentifier: l.AdminIdentifier,
			Action:          l.Action,
			ResourceType:    l.ResourceType,
			ResourceID:      ptrToString(l.ResourceID),
			Success:         l.Success,
			ErrorMessage:    ptrToString(l.ErrorMessage),
			IPAddress:       ptrToString(l.IPAddress),
			CreatedAt:       l.CreatedAt,
		}
	}

	totalPages := (total + pageSize - 1) / pageSize
	return &dto.AuditLogListResponse{
		Logs:       result,
		Total:      total,
		Page:       page,
		PageSize:   pageSize,
		TotalPages: totalPages,
	}, nil
}

// GetAuditLog returns a single audit log
func (s *AuditLogService) GetAuditLog(ctx context.Context, id string) (*dto.AuditLogResponse, error) {
	l, err := s.entClient.AuditLog.Get(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("log not found: %w", err)
	}

	return &dto.AuditLogResponse{
		ID:              l.ID,
		AdminIdentifier: l.AdminIdentifier,
		Action:          l.Action,
		ResourceType:    l.ResourceType,
		ResourceID:      ptrToString(l.ResourceID),
		Success:         l.Success,
		ErrorMessage:    ptrToString(l.ErrorMessage),
		IPAddress:       ptrToString(l.IPAddress),
		CreatedAt:       l.CreatedAt,
	}, nil
}
