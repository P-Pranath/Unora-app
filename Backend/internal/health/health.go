// internal/health/health.go
package health

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"

	ent "github.com/UnoraApp/be/ent/generated"
)

// Handler handles health check endpoints
type Handler struct {
	entClient   *ent.Client
	redisClient *redis.Client
	startTime   time.Time
}

// NewHandler creates a new health handler
func NewHandler(entClient *ent.Client, redisClient *redis.Client) *Handler {
	return &Handler{
		entClient:   entClient,
		redisClient: redisClient,
		startTime:   time.Now(),
	}
}

// HealthResponse represents the health check response
type HealthResponse struct {
	Status    string            `json:"status" example:"healthy"`
	Timestamp string            `json:"timestamp" example:"2026-01-14T12:00:00Z"`
	Uptime    string            `json:"uptime" example:"24h30m15s"`
	Checks    map[string]string `json:"checks,omitempty"`
}

// Healthz godoc
// @Summary      Health check
// @Description  Basic health check - returns OK if service is running
// @Tags         health
// @Produce      json
// @Success      200 {object} HealthResponse
// @Router       /healthz [get]
func (h *Handler) Healthz(c *gin.Context) {
	c.JSON(http.StatusOK, HealthResponse{
		Status:    "healthy",
		Timestamp: time.Now().UTC().Format(time.RFC3339),
		Uptime:    time.Since(h.startTime).Round(time.Second).String(),
	})
}

// Readyz godoc
// @Summary      Readiness check
// @Description  Checks if service is ready to accept traffic (DB, Redis connectivity)
// @Tags         health
// @Produce      json
// @Success      200 {object} HealthResponse
// @Failure      503 {object} HealthResponse
// @Router       /readyz [get]
func (h *Handler) Readyz(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	checks := make(map[string]string)
	healthy := true

	// Check database
	if h.entClient != nil {
		if err := h.checkDatabase(ctx); err != nil {
			checks["database"] = "unhealthy: " + err.Error()
			healthy = false
		} else {
			checks["database"] = "healthy"
		}
	}

	// Check Redis
	if h.redisClient != nil {
		if err := h.checkRedis(ctx); err != nil {
			checks["redis"] = "unhealthy: " + err.Error()
			healthy = false
		} else {
			checks["redis"] = "healthy"
		}
	}

	status := "ready"
	httpStatus := http.StatusOK
	if !healthy {
		status = "not_ready"
		httpStatus = http.StatusServiceUnavailable
	}

	c.JSON(httpStatus, HealthResponse{
		Status:    status,
		Timestamp: time.Now().UTC().Format(time.RFC3339),
		Uptime:    time.Since(h.startTime).Round(time.Second).String(),
		Checks:    checks,
	})
}

// Livez godoc
// @Summary      Liveness check
// @Description  Simple liveness probe for Kubernetes
// @Tags         health
// @Produce      json
// @Success      200 {object} map[string]string
// @Router       /livez [get]
func (h *Handler) Livez(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "alive"})
}

func (h *Handler) checkDatabase(ctx context.Context) error {
	// Simple query to check database connectivity
	_, err := h.entClient.User.Query().Limit(1).All(ctx)
	return err
}

func (h *Handler) checkRedis(ctx context.Context) error {
	return h.redisClient.Ping(ctx).Err()
}

// RegisterRoutes registers health endpoints
func RegisterRoutes(router *gin.Engine, entClient *ent.Client, redisClient *redis.Client) {
	handler := NewHandler(entClient, redisClient)
	
	router.GET("/healthz", handler.Healthz)
	router.GET("/readyz", handler.Readyz)
	router.GET("/livez", handler.Livez)
}
