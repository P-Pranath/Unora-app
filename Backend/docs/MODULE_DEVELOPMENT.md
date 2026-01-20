# Module Development Guide

This guide explains how to develop new modules for this backend following established patterns and best practices.

## Quick Start

Use the scaffold script to create a new module:
```bash
make scaffold-module
# Enter module name, e.g., "product"
```

## Module Structure

Each module follows this structure:
```
internal/{module}/
├── dto/                  # Data Transfer Objects
│   └── {module}_dto.go
├── handlers/             # HTTP handlers
│   └── {module}_handler.go
├── services/             # Business logic
│   └── {module}_service.go
├── repositories/         # Data access (optional)
│   └── {module}_repository.go
├── middlewares/          # Module-specific middleware
│   └── {module}_middleware.go
└── routes/               # Route registration
    └── {module}_routes.go
```

---

## Step-by-Step: Creating a New Module

### Step 1: Define DTOs

Create `internal/{module}/dto/{module}_dto.go`:

```go
package dto

// CreateRequest - request for creating a resource
type CreateRequest struct {
    Name        string `json:"name" validate:"required,min=2,max=100"`
    Description string `json:"description" validate:"max=500"`
}

// UpdateRequest - request for updating a resource
type UpdateRequest struct {
    Name        *string `json:"name" validate:"omitempty,min=2,max=100"`
    Description *string `json:"description" validate:"omitempty,max=500"`
}

// Response - the response structure
type Response struct {
    ID          string `json:"id"`
    Name        string `json:"name"`
    Description string `json:"description"`
    CreatedAt   string `json:"createdAt"`
    UpdatedAt   string `json:"updatedAt"`
}
```

**Key Points:**
- Use JSON tags with camelCase
- Use `validate` tags for validation rules
- Use pointers for optional update fields
- Separate create/update/response DTOs

---

### Step 2: Create Service

Create `internal/{module}/services/{module}_service.go`:

```go
package services

import (
    "context"
    "fmt"

    ent "github.com/UnoraApp/be/ent/generated"
    "github.com/UnoraApp/be/internal/{module}/dto"
)

// Service handles business logic
type Service struct {
    entClient *ent.Client
}

// NewService creates a new service with dependencies
func NewService(entClient *ent.Client) *Service {
    return &Service{entClient: entClient}
}

// Create creates a new resource
func (s *Service) Create(ctx context.Context, req *dto.CreateRequest) (*dto.Response, error) {
    // Your business logic here
    return nil, nil
}

// GetByID retrieves a resource by ID
func (s *Service) GetByID(ctx context.Context, id string) (*dto.Response, error) {
    // Your business logic here
    return nil, nil
}

// List retrieves all resources with pagination
func (s *Service) List(ctx context.Context, page, limit int) ([]*dto.Response, int64, error) {
    // Your business logic here
    return nil, 0, nil
}
```

**Key Points:**
- Accept context as first parameter
- Return errors, let handlers decide HTTP response
- Use constructor function for DI
- Keep business logic in services, not handlers

---

### Step 3: Create Handler

Create `internal/{module}/handlers/{module}_handler.go`:

```go
package handlers

import (
    "net/http"

    "github.com/gin-gonic/gin"

    "github.com/UnoraApp/be/internal/{module}/dto"
    "github.com/UnoraApp/be/internal/{module}/services"
    "github.com/UnoraApp/be/pkg/response"
)

// Handler handles HTTP requests
type Handler struct {
    service *services.Service
}

// NewHandler creates a handler with dependencies
func NewHandler(service *services.Service) *Handler {
    return &Handler{service: service}
}

// Create handles POST requests
func (h *Handler) Create(c *gin.Context) {
    var req dto.CreateRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        response.Error(c, http.StatusBadRequest, "INVALID_REQUEST", err.Error())
        return
    }

    result, err := h.service.Create(c.Request.Context(), &req)
    if err != nil {
        response.Error(c, http.StatusInternalServerError, "CREATE_FAILED", err.Error())
        return
    }

    response.JSON(c, http.StatusCreated, result)
}

// GetByID handles GET /:id requests
func (h *Handler) GetByID(c *gin.Context) {
    id := c.Param("id")

    result, err := h.service.GetByID(c.Request.Context(), id)
    if err != nil {
        response.Error(c, http.StatusNotFound, "NOT_FOUND", "Resource not found")
        return
    }

    response.JSON(c, http.StatusOK, result)
}
```

**Key Points:**
- Use `pkg/response` for consistent responses
- Validate input with `ShouldBindJSON`
- Pass `c.Request.Context()` to services
- Keep handlers thin, delegate to services

---

### Step 4: Register Routes

Create `internal/{module}/routes/{module}_routes.go`:

```go
package routes

import (
    "github.com/gin-gonic/gin"
    "github.com/redis/go-redis/v9"

    ent "github.com/UnoraApp/be/ent/generated"
    "github.com/UnoraApp/be/internal/{module}/handlers"
    "github.com/UnoraApp/be/internal/{module}/services"
    authmiddleware "github.com/UnoraApp/be/internal/auth/middlewares"
    authservices "github.com/UnoraApp/be/internal/auth/services"
    "github.com/UnoraApp/be/internal/config"
)

// RegisterRoutes registers all routes for this module
func RegisterRoutes(
    router *gin.RouterGroup,
    entClient *ent.Client,
    redisClient *redis.Client,
    cfg *config.Config,
) {
    // Create dependencies (DI)
    service := services.NewService(entClient)
    handler := handlers.NewHandler(service)

    // For auth middleware
    authService := authservices.NewAuthService(entClient, redisClient, cfg)

    // Public routes
    public := router.Group("/{module}s")
    {
        public.GET("/", handler.List)
        public.GET("/:id", handler.GetByID)
    }

    // Protected routes
    protected := router.Group("/{module}s")
    protected.Use(authmiddleware.AuthMiddleware(authService))
    {
        protected.POST("/", handler.Create)
        protected.PUT("/:id", handler.Update)
        protected.DELETE("/:id", handler.Delete)
    }
}
```

---

### Step 5: Wire Up in Router

Edit `internal/router/api_v1.go`:

```go
import (
    // ... existing imports
    {module}routes "github.com/UnoraApp/be/internal/{module}/routes"
)

func RegisterRoutes(...) {
    // ... existing code

    // Register your module
    {module}routes.RegisterRoutes(api, entClient, redisClient, cfg)
}
```

---

## Dependency Injection Pattern

All modules follow constructor-based DI:

```
Routes → Handler → Service → (Repository) → Ent Client
         ↓
    Dependencies passed in, not created internally
```

**Benefits:**
- Easy testing with mocks
- Clear dependency graph
- Loose coupling

---

## Auth Integration

### Protecting Routes

```go
import authmiddleware "github.com/UnoraApp/be/internal/auth/middlewares"

// In your routes.go
protected := router.Group("/resource")
protected.Use(authmiddleware.AuthMiddleware(authService))
```

### Accessing User in Handler

```go
func (h *Handler) Create(c *gin.Context) {
    // Get authenticated user
    userID := c.GetString("userID")
    email := c.GetString("userEmail")

    // Or get full payload
    payload, _ := c.Get("userPayload")
    tokenPayload := payload.(*dto.TokenPayload)
}
```

---

## Response Format

All responses use `pkg/response`:

**Success:**
```json
{
  "success": true,
  "data": { ... }
}
```

**Error:**
```json
{
  "success": false,
  "error": {
    "code": "ERROR_CODE",
    "message": "Human readable message"
  }
}
```

**Paginated:**
```json
{
  "success": true,
  "data": [...],
  "pagination": {
    "page": 1,
    "limit": 10,
    "total": 100,
    "totalPages": 10
  }
}
```

---

## Validation

Use struct tags with `go-playground/validator`:

```go
type Request struct {
    Email    string `json:"email" validate:"required,email"`
    Name     string `json:"name" validate:"required,min=2,max=100"`
    Age      int    `json:"age" validate:"gte=0,lte=150"`
    Website  string `json:"website" validate:"omitempty,url"`
}
```

Common validators:
- `required` - field must be present
- `omitempty` - skip validation if empty
- `email` - valid email format
- `url` - valid URL
- `min=N, max=N` - string length
- `gte=N, lte=N` - numeric range
- `oneof=a b c` - must be one of values

---

## Error Handling

Use `pkg/errors` for typed errors:

```go
import "github.com/UnoraApp/be/pkg/errors"

// In service
if notFound {
    return nil, errors.NotFound("User")
}
if badInput {
    return nil, errors.BadRequest("Invalid email format")
}

// In handler
result, err := h.service.Create(ctx, req)
if err != nil {
    if appErr, ok := err.(*errors.AppError); ok {
        switch appErr.Code {
        case errors.ErrCodeNotFound:
            response.Error(c, http.StatusNotFound, string(appErr.Code), appErr.Message)
        case errors.ErrCodeBadRequest:
            response.Error(c, http.StatusBadRequest, string(appErr.Code), appErr.Message)
        default:
            response.Error(c, http.StatusInternalServerError, "INTERNAL_ERROR", appErr.Message)
        }
        return
    }
    response.Error(c, http.StatusInternalServerError, "INTERNAL_ERROR", err.Error())
    return
}
```

---

## Example: Auth Module Reference

The `internal/auth` module is a complete reference implementation:

| File | Purpose |
|------|---------|
| `dto/auth_dto.go` | Request/response structures |
| `services/google_service.go` | Google token verification |
| `services/token_service.go` | JWT generation/validation |
| `services/auth_service.go` | Main business logic with DI |
| `handlers/auth_handler.go` | HTTP endpoints |
| `middlewares/auth_middleware.go` | Token validation middleware |
| `routes/auth_routes.go` | Route registration |

**Endpoints:**
- `POST /api/v1/auth/google` - Login with Google ID token
- `POST /api/v1/auth/refresh` - Refresh access token
- `POST /api/v1/auth/logout` - Logout (protected)
- `GET /api/v1/auth/me` - Get current user (protected)

---

## Checklist for New Modules

- [ ] Create DTOs in `dto/`
- [ ] Create service with constructor DI in `services/`
- [ ] Create handler with service injection in `handlers/`
- [ ] Create routes registration in `routes/`
- [ ] Add middleware if needed in `middlewares/`
- [ ] Register routes in `internal/router/api_v1.go`
- [ ] Add tests (optional but recommended)
