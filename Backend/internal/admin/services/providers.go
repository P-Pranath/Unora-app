// internal/admin/services/providers.go
package services

import "github.com/google/wire"

// ProviderSet is the wire provider set for admin services
var ProviderSet = wire.NewSet(
	NewUserManagementService,
	NewReportManagementService,
	NewAnalyticsService,
	NewContentManagementService,
)
