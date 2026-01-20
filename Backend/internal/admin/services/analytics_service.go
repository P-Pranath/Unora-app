// internal/admin/services/analytics_service.go
package services

import (
	"context"
	"fmt"
	"time"

	ent "github.com/UnoraApp/be/ent/generated"
	"github.com/UnoraApp/be/ent/generated/connection"
	"github.com/UnoraApp/be/ent/generated/paymentorder"
	"github.com/UnoraApp/be/ent/generated/streak"
	"github.com/UnoraApp/be/ent/generated/user"
	"github.com/UnoraApp/be/ent/generated/userreport"
	"github.com/UnoraApp/be/internal/admin/dto"
)

// AnalyticsService handles admin analytics
type AnalyticsService struct {
	entClient *ent.Client
}

// NewAnalyticsService creates a new analytics service
func NewAnalyticsService(entClient *ent.Client) *AnalyticsService {
	return &AnalyticsService{
		entClient: entClient,
	}
}

// GetDashboardStats returns aggregated dashboard statistics
func (s *AnalyticsService) GetDashboardStats(ctx context.Context) (*dto.DashboardStatsResponse, error) {
	userStats, err := s.getUserStats(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get user stats: %w", err)
	}

	paymentStats, err := s.getPaymentStats(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get payment stats: %w", err)
	}

	connectionStats, err := s.getConnectionStats(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get connection stats: %w", err)
	}

	reportStats, err := s.getReportStats(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get report stats: %w", err)
	}

	return &dto.DashboardStatsResponse{
		UserStats:       *userStats,
		PaymentStats:    *paymentStats,
		ConnectionStats: *connectionStats,
		ReportStats:     *reportStats,
	}, nil
}

func (s *AnalyticsService) getUserStats(ctx context.Context) (*dto.UserStatsResponse, error) {
	now := time.Now()
	todayStart := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	weekStart := todayStart.AddDate(0, 0, -7)
	monthStart := todayStart.AddDate(0, -1, 0)

	totalUsers, _ := s.entClient.User.Query().Count(ctx)

	activeToday, _ := s.entClient.User.Query().
		Where(user.LastActiveAtGTE(todayStart)).
		Count(ctx)

	activeThisWeek, _ := s.entClient.User.Query().
		Where(user.LastActiveAtGTE(weekStart)).
		Count(ctx)

	newThisWeek, _ := s.entClient.User.Query().
		Where(user.CreatedAtGTE(weekStart)).
		Count(ctx)

	newThisMonth, _ := s.entClient.User.Query().
		Where(user.CreatedAtGTE(monthStart)).
		Count(ctx)

	suspendedUsers, _ := s.entClient.User.Query().
		Where(user.AccountStatusEQ(user.AccountStatusSuspended)).
		Count(ctx)

	onboardingPending, _ := s.entClient.User.Query().
		Where(user.OnboardingStatusNEQ(user.OnboardingStatusCompleted)).
		Count(ctx)

	return &dto.UserStatsResponse{
		TotalUsers:        totalUsers,
		ActiveToday:       activeToday,
		ActiveThisWeek:    activeThisWeek,
		NewThisWeek:       newThisWeek,
		NewThisMonth:      newThisMonth,
		SuspendedUsers:    suspendedUsers,
		OnboardingPending: onboardingPending,
	}, nil
}

func (s *AnalyticsService) getPaymentStats(ctx context.Context) (*dto.PaymentStatsResponse, error) {
	now := time.Now()
	todayStart := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	monthStart := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())

	// Total revenue from paid orders
	paidOrders, _ := s.entClient.PaymentOrder.Query().
		Where(paymentorder.OrderStatusEQ(paymentorder.OrderStatusPaid)).
		All(ctx)

	totalRevenue := 0
	revenueThisMonth := 0
	revenueToday := 0
	transactionsToday := 0

	for _, o := range paidOrders {
		totalRevenue += o.Amount
		if o.PaidAt != nil {
			if o.PaidAt.After(monthStart) {
				revenueThisMonth += o.Amount
			}
			if o.PaidAt.After(todayStart) {
				revenueToday += o.Amount
				transactionsToday++
			}
		}
	}

	totalTransactions := len(paidOrders)

	// Credits in circulation
	users, _ := s.entClient.User.Query().All(ctx)
	creditsInCirculation := 0
	for _, u := range users {
		creditsInCirculation += u.CreditBalance
	}

	return &dto.PaymentStatsResponse{
		TotalRevenue:         totalRevenue,
		RevenueThisMonth:     revenueThisMonth,
		RevenueToday:         revenueToday,
		TotalTransactions:    totalTransactions,
		TransactionsToday:    transactionsToday,
		CreditsInCirculation: creditsInCirculation,
	}, nil
}

func (s *AnalyticsService) getConnectionStats(ctx context.Context) (*dto.ConnectionStatsResponse, error) {
	now := time.Now()
	weekStart := now.AddDate(0, 0, -7)

	totalConnections, _ := s.entClient.Connection.Query().Count(ctx)

	activeConnections, _ := s.entClient.Connection.Query().
		Where(connection.ConnectionStatusEQ(connection.ConnectionStatusActive)).
		Count(ctx)

	connectionsThisWeek, _ := s.entClient.Connection.Query().
		Where(connection.CreatedAtGTE(weekStart)).
		Count(ctx)

	activeStreaks, _ := s.entClient.Streak.Query().
		Where(streak.StreakStateIn(streak.StreakStateActive, streak.StreakStateAtRisk)).
		Count(ctx)

	completedStreaks, _ := s.entClient.Streak.Query().
		Where(streak.StreakStateEQ(streak.StreakStateCompleted)).
		Count(ctx)

	// Calculate average streak day
	streaks, _ := s.entClient.Streak.Query().
		Where(streak.StreakStateIn(streak.StreakStateActive, streak.StreakStateAtRisk)).
		All(ctx)

	avgStreakDay := 0.0
	if len(streaks) > 0 {
		totalDays := 0
		for _, s := range streaks {
			totalDays += s.CurrentDay
		}
		avgStreakDay = float64(totalDays) / float64(len(streaks))
	}

	return &dto.ConnectionStatsResponse{
		TotalConnections:    totalConnections,
		ActiveConnections:   activeConnections,
		ConnectionsThisWeek: connectionsThisWeek,
		ActiveStreaks:       activeStreaks,
		CompletedStreaks:    completedStreaks,
		AverageStreakDay:    avgStreakDay,
	}, nil
}

func (s *AnalyticsService) getReportStats(ctx context.Context) (*dto.ReportStatsResponse, error) {
	now := time.Now()
	weekStart := now.AddDate(0, 0, -7)

	totalReports, _ := s.entClient.UserReport.Query().Count(ctx)

	pendingReports, _ := s.entClient.UserReport.Query().
		Where(userreport.ReportStatusEQ(userreport.ReportStatusPending)).
		Count(ctx)

	reportsThisWeek, _ := s.entClient.UserReport.Query().
		Where(userreport.CreatedAtGTE(weekStart)).
		Count(ctx)

	actionsTaken, _ := s.entClient.UserReport.Query().
		Where(userreport.ReportStatusEQ(userreport.ReportStatusActionTaken)).
		Count(ctx)

	dismissedReports, _ := s.entClient.UserReport.Query().
		Where(userreport.ReportStatusEQ(userreport.ReportStatusDismissed)).
		Count(ctx)

	return &dto.ReportStatsResponse{
		TotalReports:     totalReports,
		PendingReports:   pendingReports,
		ReportsThisWeek:  reportsThisWeek,
		ActionsTaken:     actionsTaken,
		DismissedReports: dismissedReports,
	}, nil
}
