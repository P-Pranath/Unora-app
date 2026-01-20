// internal/admin/services/report_management_service.go
package services

import (
	"context"
	"fmt"
	"time"

	ent "github.com/UnoraApp/be/ent/generated"
	"github.com/UnoraApp/be/ent/generated/userreport"
	"github.com/UnoraApp/be/internal/admin/dto"
)

// ReportManagementService handles admin report management
type ReportManagementService struct {
	entClient         *ent.Client
	userMgmtService   *UserManagementService
}

// NewReportManagementService creates a new report management service
func NewReportManagementService(entClient *ent.Client, userMgmtService *UserManagementService) *ReportManagementService {
	return &ReportManagementService{
		entClient:       entClient,
		userMgmtService: userMgmtService,
	}
}

// ListReports returns paginated list of reports
func (s *ReportManagementService) ListReports(ctx context.Context, page, pageSize int, status string) (*dto.AdminReportListResponse, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 20
	}
	offset := (page - 1) * pageSize

	query := s.entClient.UserReport.Query()

	// Apply status filter
	if status != "" {
		query = query.Where(userreport.ReportStatusEQ(userreport.ReportStatus(status)))
	}

	// Get total count
	total, err := query.Clone().Count(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to count reports: %w", err)
	}

	// Get reports
	reports, err := query.
		Order(ent.Desc(userreport.FieldCreatedAt)).
		Offset(offset).
		Limit(pageSize).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get reports: %w", err)
	}

	result := make([]dto.AdminReportResponse, len(reports))
	for i, r := range reports {
		result[i] = s.reportToAdminResponse(ctx, r)
	}

	totalPages := (total + pageSize - 1) / pageSize

	return &dto.AdminReportListResponse{
		Reports:    result,
		Total:      total,
		Page:       page,
		PageSize:   pageSize,
		TotalPages: totalPages,
	}, nil
}

// GetReport returns detailed report info
func (s *ReportManagementService) GetReport(ctx context.Context, reportID string) (*dto.AdminReportResponse, error) {
	r, err := s.entClient.UserReport.Get(ctx, reportID)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, fmt.Errorf("report not found")
		}
		return nil, fmt.Errorf("failed to get report: %w", err)
	}

	result := s.reportToAdminResponse(ctx, r)
	return &result, nil
}

// ReviewReport reviews a report and optionally takes action
func (s *ReportManagementService) ReviewReport(ctx context.Context, reportID string, req *dto.ReviewReportRequest) error {
	r, err := s.entClient.UserReport.Get(ctx, reportID)
	if err != nil {
		return fmt.Errorf("report not found: %w", err)
	}

	now := time.Now()

	// Update report status
	_, err = r.Update().
		SetReportStatus(userreport.ReportStatus(req.Status)).
		SetNillableAdminNotes(strPtr(req.AdminNotes)).
		SetReviewedAt(now).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("failed to update report: %w", err)
	}

	// Take action if specified
	if req.Action != "" {
		switch req.Action {
		case "warn":
			// Just mark as warned (could send notification)
		case "suspend_7d":
			s.userMgmtService.SuspendUser(ctx, r.ReportedUserID, &dto.SuspendUserRequest{
				Reason:   "Violation of community guidelines",
				Duration: "7d",
			})
		case "suspend_30d":
			s.userMgmtService.SuspendUser(ctx, r.ReportedUserID, &dto.SuspendUserRequest{
				Reason:   "Violation of community guidelines",
				Duration: "30d",
			})
		case "suspend_permanent":
			s.userMgmtService.SuspendUser(ctx, r.ReportedUserID, &dto.SuspendUserRequest{
				Reason:   "Serious violation of community guidelines",
				Duration: "permanent",
			})
		case "delete":
			s.userMgmtService.DeleteUser(ctx, r.ReportedUserID)
		}
	}

	return nil
}

func (s *ReportManagementService) reportToAdminResponse(ctx context.Context, r *ent.UserReport) dto.AdminReportResponse {
	// Get reporter name
	reporter, _ := s.entClient.User.Get(ctx, r.ReporterUserID)
	reporterName := ""
	if reporter != nil && reporter.FirstName != nil {
		reporterName = *reporter.FirstName
	}

	// Get reported name
	reported, _ := s.entClient.User.Get(ctx, r.ReportedUserID)
	reportedName := ""
	if reported != nil && reported.FirstName != nil {
		reportedName = *reported.FirstName
	}

	return dto.AdminReportResponse{
		ID:            r.ID,
		ReporterID:    r.ReporterUserID,
		ReporterName:  reporterName,
		ReportedID:    r.ReportedUserID,
		ReportedName:  reportedName,
		Reason:        string(r.ReportReason),
		Description:   ptrToString(r.Description),
		ReferenceType: ptrToString(r.ReferenceType),
		ReferenceID:   ptrToString(r.ReferenceID),
		Status:        string(r.ReportStatus),
		AdminNotes:    ptrToString(r.AdminNotes),
		ReviewedAt:    r.ReviewedAt,
		CreatedAt:     r.CreatedAt,
	}
}
