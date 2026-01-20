// internal/discovery/config/tier_config.go
package config

import "time"

// TierConfig contains tier-specific limits and constraints
type TierConfig struct {
	ConnectionSlots   int
	RefreshCooldown   time.Duration
	NudgesPerDay      int
	FreeRecoveries    int
	EarnedReveals     int     // Reveals earned through milestones
	PurchasableReveals int    // Additional reveals that can be purchased
	RevealDays        []int   // Days when reveals unlock (internal)
}

// TierLimits maps subscription tiers to their configurations
var TierLimits = map[string]TierConfig{
	"free": {
		ConnectionSlots:    1,
		RefreshCooldown:    24 * time.Hour,
		NudgesPerDay:       1,
		FreeRecoveries:     0,
		EarnedReveals:      2,
		PurchasableReveals: 3,
		RevealDays:         []int{5, 12}, // Day 5, 12
	},
	"plus": {
		ConnectionSlots:    2,
		RefreshCooldown:    12 * time.Hour,
		NudgesPerDay:       3,
		FreeRecoveries:     1,
		EarnedReveals:      3,
		PurchasableReveals: 2,
		RevealDays:         []int{4, 8, 12}, // Day 4, 8, 12
	},
	"pro": {
		ConnectionSlots:    4,
		RefreshCooldown:    6 * time.Hour,
		NudgesPerDay:       4,
		FreeRecoveries:     2,
		EarnedReveals:      4,
		PurchasableReveals: 1,
		RevealDays:         []int{3, 6, 9, 12}, // Day 3, 6, 9, 12
	},
}

// GetTierConfig returns the configuration for a given tier
func GetTierConfig(tier string) TierConfig {
	if config, ok := TierLimits[tier]; ok {
		return config
	}
	return TierLimits["free"] // Default to free
}

// CanRefresh checks if user can refresh based on their tier and last refresh time
func CanRefresh(tier string, lastRefreshAt *time.Time) bool {
	if lastRefreshAt == nil {
		return true
	}
	config := GetTierConfig(tier)
	return time.Since(*lastRefreshAt) >= config.RefreshCooldown
}

// GetRefreshAvailableAt calculates when the next refresh will be available
func GetRefreshAvailableAt(tier string, lastRefreshAt time.Time) time.Time {
	config := GetTierConfig(tier)
	return lastRefreshAt.Add(config.RefreshCooldown)
}

// CanConnect checks if user has available connection slots
func CanConnect(tier string, activeConnections int) bool {
	config := GetTierConfig(tier)
	return activeConnections < config.ConnectionSlots
}

// HasFreeRecovery checks if user has free recoveries left
func HasFreeRecovery(tier string, recoveriesUsed int) bool {
	config := GetTierConfig(tier)
	return recoveriesUsed < config.FreeRecoveries
}

// CanSendNudge checks if user can send more nudges today
func CanSendNudge(tier string, nudgesSentToday int) bool {
	config := GetTierConfig(tier)
	return nudgesSentToday < config.NudgesPerDay
}
