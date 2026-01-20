# Complete Feature Development Guide

This guide covers the **entire workflow** from database schema to production-ready feature, using **Ent ORM**, **Google Wire** for DI, and established patterns.

---

## Table of Contents

1. [Overview](#overview)
2. [Step 1: Create Ent Schema](#step-1-create-ent-schema)
3. [Step 2: Generate Ent Code](#step-2-generate-ent-code)
4. [Step 3: Create Migration](#step-3-create-migration)
5. [Step 4: Create DTOs](#step-4-create-dtos)
6. [Step 5: Create Service Layer](#step-5-create-service-layer)
7. [Step 6: Create Handler Layer](#step-6-create-handler-layer)
8. [Step 7: Create Wire Provider Set](#step-7-create-wire-provider-set)
9. [Step 8: Register in DI Container](#step-8-register-in-di-container)
10. [Step 9: Create Routes](#step-9-create-routes)
11. [Step 10: Wire Up in Router](#step-10-wire-up-in-router)
12. [Step 11: Generate Wire Code](#step-11-generate-wire-code)
13. [Step 12: Test & Verify](#step-12-test--verify)
14. [Step 13: Add API Documentation (Swagger)](#step-13-add-api-documentation-swagger)
15. [Step 14: Generate Documentation & Types](#step-14-generate-documentation--types)
16. [Quick Reference Commands](#quick-reference-commands)
17. [Checklist for New Features](#checklist-for-new-features)

---

## Overview

### Development Flow

```
┌──────────────────────────────────────────────────────────────────┐
│                    FEATURE DEVELOPMENT FLOW                       │
├──────────────────────────────────────────────────────────────────┤
│                                                                   │
│  1. Schema (Ent)     →  Define your data model                   │
│         ↓                                                        │
│  2. Generate Ent     →  make ent-generate                        │
│         ↓                                                        │
│  3. Migration        →  make migration-create                    │
│         ↓                                                        │
│  4. DTOs             →  Request/Response structures              │
│         ↓                                                        │
│  5. Service          →  Business logic                           │
│         ↓                                                        │
│  6. Handler          →  HTTP endpoints                           │
│         ↓                                                        │
│  7. Wire Providers   →  Define DI provider set                   │
│         ↓                                                        │
│  8. DI Container     →  Register in internal/di/wire.go         │
│         ↓                                                        │
│  9. Routes           →  Register endpoints                       │
│         ↓                                                        │
│ 10. Wire Generate    →  make wire                                │
│         ↓                                                        │
│ 11. Test             →  make run && test endpoints               │
│                                                                   │
└──────────────────────────────────────────────────────────────────┘
```

### Project Structure
```
internal/
├── di/                      # Dependency Injection (Wire)
│   ├── wire.go             # Injector definitions (edit this)
│   └── wire_gen.go         # Generated code (don't edit)
├── auth/                    # Example module
│   ├── dto/
│   ├── handlers/
│   ├── middlewares/
│   ├── routes/
│   └── services/
└── {your_module}/           # New module
    ├── dto/
    ├── handlers/
    ├── services/
    └── routes/
```

---

## Step 1: Create Ent Schema

Create a database schema using Ent ORM.

```bash
make ent-init
# Enter schema name: User
```

This creates `ent/schema/user.go`. Edit it:

```go
// ent/schema/user.go
package schema

import (
    "time"

    "entgo.io/ent"
    "entgo.io/ent/dialect"
    "entgo.io/ent/schema/field"
    "entgo.io/ent/schema/index"
)

// User holds the schema definition for the User entity.
type User struct {
    ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
    return []ent.Field{
        field.String("id").
            Unique().
            Immutable(),
        field.String("email").
            Unique().
            NotEmpty(),
        field.String("name").
            NotEmpty(),
        field.String("first_name").
            Optional(),
        field.String("last_name").
            Optional(),
        field.String("picture").
            Optional(),
        field.String("provider").
            Default("google"),
        field.Bool("is_active").
            Default(true),
        field.Time("created_at").
            Default(time.Now).
            Immutable(),
        field.Time("updated_at").
            Default(time.Now).
            UpdateDefault(time.Now),
    }
}

// Indexes of the User.
func (User) Indexes() []ent.Index {
    return []ent.Index{
        index.Fields("email").Unique(),
        index.Fields("provider"),
    }
}

// Edges of the User (relationships).
func (User) Edges() []ent.Edge {
    return nil // Add relationships here later
}
```

### Common Field Types

| Type | Example |
|------|---------|
| String | `field.String("name")` |
| Int | `field.Int("age")` |
| Bool | `field.Bool("is_active")` |
| Time | `field.Time("created_at")` |
| UUID | `field.UUID("id", uuid.UUID{})` |
| Enum | `field.Enum("status").Values("active", "inactive")` |
| JSON | `field.JSON("metadata", map[string]any{})` |

### Common Field Modifiers

| Modifier | Purpose |
|----------|---------|
| `.Unique()` | Unique constraint |
| `.NotEmpty()` | Validation: not empty |
| `.Optional()` | Nullable field |
| `.Default(value)` | Default value |
| `.Immutable()` | Can't be updated after creation |

---

## Step 2: Generate Ent Code

Generate the Go code from your schema:

```bash
make ent-generate
```

This generates code in `ent/generated/` including:
- Client and transaction code
- Query builders
- Create/Update builders
- Model structs

---

## Step 3: Create Migration

Create a SQL migration for the schema changes:

```bash
make migration-create
# Enter name: create_users_table
```

Edit `migrations/versions/XXXXXX_create_users_table.sql`:

```sql
-- +goose Up
CREATE TABLE users (
    id VARCHAR(36) NOT NULL PRIMARY KEY,
    email VARCHAR(255) NOT NULL UNIQUE,
    name VARCHAR(255) NOT NULL,
    first_name VARCHAR(255),
    last_name VARCHAR(255),
    picture TEXT,
    provider VARCHAR(50) DEFAULT 'google',
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_users_email (email),
    INDEX idx_users_provider (provider)
);

-- +goose Down
DROP TABLE IF EXISTS users;
```

Run the migration:

```bash
make migration-up
```

---

## Step 4: Create DTOs

Create request/response structures in `internal/{module}/dto/`:

```go
// internal/user/dto/user_dto.go
package dto

// CreateUserRequest - request for creating a user
type CreateUserRequest struct {
    Email     string `json:"email" validate:"required,email"`
    Name      string `json:"name" validate:"required,min=2,max=100"`
    FirstName string `json:"firstName" validate:"max=50"`
    LastName  string `json:"lastName" validate:"max=50"`
}

// UpdateUserRequest - request for updating a user
type UpdateUserRequest struct {
    Name      *string `json:"name" validate:"omitempty,min=2,max=100"`
    FirstName *string `json:"firstName" validate:"omitempty,max=50"`
    LastName  *string `json:"lastName" validate:"omitempty,max=50"`
    Picture   *string `json:"picture" validate:"omitempty,url"`
}

// UserResponse - user data returned to client
type UserResponse struct {
    ID        string `json:"id"`
    Email     string `json:"email"`
    Name      string `json:"name"`
    FirstName string `json:"firstName,omitempty"`
    LastName  string `json:"lastName,omitempty"`
    Picture   string `json:"picture,omitempty"`
    Provider  string `json:"provider"`
    IsActive  bool   `json:"isActive"`
    CreatedAt string `json:"createdAt"`
    UpdatedAt string `json:"updatedAt"`
}

// ListUsersRequest - query params for listing users
type ListUsersRequest struct {
    Page   int    `form:"page" validate:"gte=1"`
    Limit  int    `form:"limit" validate:"gte=1,lte=100"`
    Search string `form:"search" validate:"max=100"`
}
```

### Validation Rules Reference

| Rule | Example | Description |
|------|---------|-------------|
| required | `validate:"required"` | Field must be present |
| email | `validate:"email"` | Valid email format |
| url | `validate:"url"` | Valid URL format |
| min/max | `validate:"min=2,max=100"` | String length bounds |
| gte/lte | `validate:"gte=1,lte=100"` | Numeric bounds |
| oneof | `validate:"oneof=admin user"` | Must be one of values |
| omitempty | `validate:"omitempty,email"` | Skip if empty |

---

## Step 5: Create Service Layer

Create business logic in `internal/{module}/services/`:

```go
// internal/user/services/user_service.go
package services

import (
    "context"
    "fmt"
    "time"

    "github.com/google/uuid"

    ent "github.com/UnoraApp/be/ent/generated"
    "github.com/UnoraApp/be/ent/generated/user"
    "github.com/UnoraApp/be/internal/user/dto"
    "github.com/UnoraApp/be/pkg/errors"
)

// UserService handles user business logic
type UserService struct {
    entClient *ent.Client
}

// NewUserService creates a new user service (constructor for Wire)
func NewUserService(entClient *ent.Client) *UserService {
    return &UserService{entClient: entClient}
}

// Create creates a new user
func (s *UserService) Create(ctx context.Context, req *dto.CreateUserRequest) (*dto.UserResponse, error) {
    // Check if user already exists
    exists, err := s.entClient.User.Query().
        Where(user.Email(req.Email)).
        Exist(ctx)
    if err != nil {
        return nil, errors.Internal("failed to check existing user", err)
    }
    if exists {
        return nil, errors.Conflict("user with this email already exists")
    }

    // Create user
    u, err := s.entClient.User.Create().
        SetID(uuid.New().String()).
        SetEmail(req.Email).
        SetName(req.Name).
        SetFirstName(req.FirstName).
        SetLastName(req.LastName).
        Save(ctx)
    if err != nil {
        return nil, errors.Internal("failed to create user", err)
    }

    return s.toResponse(u), nil
}

// GetByID retrieves a user by ID
func (s *UserService) GetByID(ctx context.Context, id string) (*dto.UserResponse, error) {
    u, err := s.entClient.User.Get(ctx, id)
    if err != nil {
        if ent.IsNotFound(err) {
            return nil, errors.NotFound("user")
        }
        return nil, errors.Internal("failed to get user", err)
    }

    return s.toResponse(u), nil
}

// List retrieves users with pagination
func (s *UserService) List(ctx context.Context, req *dto.ListUsersRequest) ([]*dto.UserResponse, int64, error) {
    query := s.entClient.User.Query().
        Where(user.IsActive(true))

    // Apply search filter
    if req.Search != "" {
        query = query.Where(
            user.Or(
                user.NameContains(req.Search),
                user.EmailContains(req.Search),
            ),
        )
    }

    // Get total count
    total, err := query.Count(ctx)
    if err != nil {
        return nil, 0, errors.Internal("failed to count users", err)
    }

    // Apply pagination
    offset := (req.Page - 1) * req.Limit
    users, err := query.
        Offset(offset).
        Limit(req.Limit).
        Order(ent.Desc(user.FieldCreatedAt)).
        All(ctx)
    if err != nil {
        return nil, 0, errors.Internal("failed to list users", err)
    }

    // Map to response
    responses := make([]*dto.UserResponse, len(users))
    for i, u := range users {
        responses[i] = s.toResponse(u)
    }

    return responses, int64(total), nil
}

// Update updates a user
func (s *UserService) Update(ctx context.Context, id string, req *dto.UpdateUserRequest) (*dto.UserResponse, error) {
    builder := s.entClient.User.UpdateOneID(id)

    if req.Name != nil {
        builder.SetName(*req.Name)
    }
    if req.FirstName != nil {
        builder.SetFirstName(*req.FirstName)
    }
    if req.LastName != nil {
        builder.SetLastName(*req.LastName)
    }
    if req.Picture != nil {
        builder.SetPicture(*req.Picture)
    }

    u, err := builder.Save(ctx)
    if err != nil {
        if ent.IsNotFound(err) {
            return nil, errors.NotFound("user")
        }
        return nil, errors.Internal("failed to update user", err)
    }

    return s.toResponse(u), nil
}

// Delete soft-deletes a user
func (s *UserService) Delete(ctx context.Context, id string) error {
    _, err := s.entClient.User.UpdateOneID(id).
        SetIsActive(false).
        Save(ctx)
    if err != nil {
        if ent.IsNotFound(err) {
            return errors.NotFound("user")
        }
        return errors.Internal("failed to delete user", err)
    }
    return nil
}

// toResponse converts Ent model to DTO
func (s *UserService) toResponse(u *ent.User) *dto.UserResponse {
    return &dto.UserResponse{
        ID:        u.ID,
        Email:     u.Email,
        Name:      u.Name,
        FirstName: u.FirstName,
        LastName:  u.LastName,
        Picture:   u.Picture,
        Provider:  u.Provider,
        IsActive:  u.IsActive,
        CreatedAt: u.CreatedAt.Format(time.RFC3339),
        UpdatedAt: u.UpdatedAt.Format(time.RFC3339),
    }
}
```

---

## Step 6: Create Handler Layer

Create HTTP handlers in `internal/{module}/handlers/`:

```go
// internal/user/handlers/user_handler.go
package handlers

import (
    "net/http"

    "github.com/gin-gonic/gin"

    "github.com/UnoraApp/be/internal/user/dto"
    "github.com/UnoraApp/be/internal/user/services"
    "github.com/UnoraApp/be/pkg/errors"
    "github.com/UnoraApp/be/pkg/response"
)

// UserHandler handles user HTTP requests
type UserHandler struct {
    userService *services.UserService
}

// NewUserHandler creates a handler (constructor for Wire)
func NewUserHandler(userService *services.UserService) *UserHandler {
    return &UserHandler{userService: userService}
}

// Create handles POST /users
func (h *UserHandler) Create(c *gin.Context) {
    var req dto.CreateUserRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        response.Error(c, http.StatusBadRequest, "INVALID_REQUEST", err.Error())
        return
    }

    result, err := h.userService.Create(c.Request.Context(), &req)
    if err != nil {
        h.handleError(c, err)
        return
    }

    response.JSON(c, http.StatusCreated, result)
}

// GetByID handles GET /users/:id
func (h *UserHandler) GetByID(c *gin.Context) {
    id := c.Param("id")

    result, err := h.userService.GetByID(c.Request.Context(), id)
    if err != nil {
        h.handleError(c, err)
        return
    }

    response.JSON(c, http.StatusOK, result)
}

// List handles GET /users
func (h *UserHandler) List(c *gin.Context) {
    var req dto.ListUsersRequest
    if err := c.ShouldBindQuery(&req); err != nil {
        response.Error(c, http.StatusBadRequest, "INVALID_REQUEST", err.Error())
        return
    }

    // Set defaults
    if req.Page == 0 {
        req.Page = 1
    }
    if req.Limit == 0 {
        req.Limit = 20
    }

    results, total, err := h.userService.List(c.Request.Context(), &req)
    if err != nil {
        h.handleError(c, err)
        return
    }

    response.Paginated(c, http.StatusOK, results, req.Page, req.Limit, total)
}

// Update handles PUT /users/:id
func (h *UserHandler) Update(c *gin.Context) {
    id := c.Param("id")

    var req dto.UpdateUserRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        response.Error(c, http.StatusBadRequest, "INVALID_REQUEST", err.Error())
        return
    }

    result, err := h.userService.Update(c.Request.Context(), id, &req)
    if err != nil {
        h.handleError(c, err)
        return
    }

    response.JSON(c, http.StatusOK, result)
}

// Delete handles DELETE /users/:id
func (h *UserHandler) Delete(c *gin.Context) {
    id := c.Param("id")

    if err := h.userService.Delete(c.Request.Context(), id); err != nil {
        h.handleError(c, err)
        return
    }

    response.JSON(c, http.StatusOK, gin.H{"message": "User deleted successfully"})
}

// handleError maps AppError to HTTP response
func (h *UserHandler) handleError(c *gin.Context, err error) {
    if appErr, ok := err.(*errors.AppError); ok {
        switch appErr.Code {
        case errors.ErrCodeNotFound:
            response.Error(c, http.StatusNotFound, string(appErr.Code), appErr.Message)
        case errors.ErrCodeBadRequest, errors.ErrCodeValidation:
            response.Error(c, http.StatusBadRequest, string(appErr.Code), appErr.Message)
        case errors.ErrCodeConflict:
            response.Error(c, http.StatusConflict, string(appErr.Code), appErr.Message)
        case errors.ErrCodeUnauthorized:
            response.Error(c, http.StatusUnauthorized, string(appErr.Code), appErr.Message)
        case errors.ErrCodeForbidden:
            response.Error(c, http.StatusForbidden, string(appErr.Code), appErr.Message)
        default:
            response.Error(c, http.StatusInternalServerError, "INTERNAL_ERROR", appErr.Message)
        }
        return
    }
    response.Error(c, http.StatusInternalServerError, "INTERNAL_ERROR", err.Error())
}
```

---

## Step 7: Create Wire Provider Set

Create providers for Wire in `internal/{module}/services/providers.go`:

```go
// internal/user/services/providers.go
package services

import (
    "github.com/google/wire"
)

// ProviderSet is the Wire provider set for user services
var ProviderSet = wire.NewSet(
    NewUserService,
    // Add more service constructors here as needed
)
```

And `internal/{module}/handlers/providers.go`:

```go
// internal/user/handlers/providers.go
package handlers

import (
    "github.com/google/wire"
)

// ProviderSet is the Wire provider set for user handlers
var ProviderSet = wire.NewSet(
    NewUserHandler,
    // Add more handler constructors here as needed
)
```

---

## Step 8: Register in DI Container

Add your module to `internal/di/wire.go`:

```go
//go:build wireinject
// +build wireinject

package di

import (
    "github.com/google/wire"
    "github.com/redis/go-redis/v9"

    ent "github.com/UnoraApp/be/ent/generated"
    authhandlers "github.com/UnoraApp/be/internal/auth/handlers"
    authservices "github.com/UnoraApp/be/internal/auth/services"
    userhandlers "github.com/UnoraApp/be/internal/user/handlers"  // ADD
    userservices "github.com/UnoraApp/be/internal/user/services"  // ADD
    "github.com/UnoraApp/be/internal/config"
    "github.com/UnoraApp/be/pkg/storage"
)

// Container holds all initialized dependencies
type Container struct {
    Config        *config.Config
    EntClient     *ent.Client
    RedisClient   *redis.Client
    StorageClient storage.Client

    // Auth
    AuthService *authservices.AuthService
    AuthHandler *authhandlers.AuthHandler

    // User - ADD YOUR MODULE HERE
    UserService *userservices.UserService
    UserHandler *userhandlers.UserHandler
}

// Module sets
var AuthSet = wire.NewSet(
    authservices.ProviderSet,
    authhandlers.ProviderSet,
)

// ADD YOUR MODULE SET
var UserSet = wire.NewSet(
    userservices.ProviderSet,
    userhandlers.ProviderSet,
)

// Full provider set
var ProviderSet = wire.NewSet(
    AuthSet,
    UserSet,  // ADD
)

// Container injector
func InitializeContainer(...) (*Container, error) {
    wire.Build(ProviderSet, wire.Struct(new(Container), "*"))
    return nil, nil
}

// ADD MODULE-SPECIFIC INJECTORS
func InitializeUserHandler(entClient *ent.Client) *userhandlers.UserHandler {
    wire.Build(UserSet)
    return nil
}

func InitializeUserService(entClient *ent.Client) *userservices.UserService {
    wire.Build(userservices.ProviderSet)
    return nil
}
```

---

## Step 9: Create Routes

Create route registration in `internal/{module}/routes/`:

```go
// internal/user/routes/user_routes.go
package routes

import (
    "github.com/gin-gonic/gin"
    "github.com/redis/go-redis/v9"

    ent "github.com/UnoraApp/be/ent/generated"
    authmiddleware "github.com/UnoraApp/be/internal/auth/middlewares"
    authservices "github.com/UnoraApp/be/internal/auth/services"
    "github.com/UnoraApp/be/internal/config"
    "github.com/UnoraApp/be/internal/di"
)

// RegisterUserRoutes registers all user routes
func RegisterUserRoutes(
    router *gin.RouterGroup,
    entClient *ent.Client,
    redisClient *redis.Client,
    cfg *config.Config,
) {
    // Use Wire-generated injector
    userHandler := di.InitializeUserHandler(entClient)

    // For auth middleware
    authService := di.InitializeAuthService(entClient, redisClient, cfg)

    // All user routes are protected
    users := router.Group("/users")
    users.Use(authmiddleware.AuthMiddleware(authService))
    {
        users.GET("/", userHandler.List)
        users.GET("/:id", userHandler.GetByID)
        users.POST("/", userHandler.Create)
        users.PUT("/:id", userHandler.Update)
        users.DELETE("/:id", userHandler.Delete)
    }
}
```

---

## Step 10: Wire Up in Router

Add your routes to `internal/router/api_v1.go`:

```go
import (
    // ... existing imports
    userroutes "github.com/UnoraApp/be/internal/user/routes"  // ADD
)

func RegisterRoutes(...) {
    // ... existing code

    // Auth routes
    authroutes.RegisterAuthRoutes(api, entClient, redisClient, cfg)

    // User routes - ADD
    userroutes.RegisterUserRoutes(api, entClient, redisClient, cfg)
}
```

---

## Step 11: Generate Wire Code

Regenerate Wire dependency injection code:

```bash
make wire
```

This updates `internal/di/wire_gen.go` with the new injector implementations.

---

## Step 12: Test & Verify

### Build and run:
```bash
make docker-infra  # Start MySQL, Redis, MinIO
make migration-up  # Run migrations
make run           # Start server
```

### Test endpoints:
```bash
# Create user
curl -X POST http://localhost:8080/api/v1/users \
  -H "Authorization: Bearer <token>" \
  -H "Content-Type: application/json" \
  -d '{"email":"test@example.com","name":"Test User"}'

# List users
curl http://localhost:8080/api/v1/users?page=1&limit=10 \
  -H "Authorization: Bearer <token>"

# Get user
curl http://localhost:8080/api/v1/users/<id> \
  -H "Authorization: Bearer <token>"
```

---

## Step 13: Add API Documentation (Swagger)

After creating your handlers, add Swagger annotations for automatic API documentation.

### Adding Swagger Annotations to Handlers

Add godoc comments above each handler method:

```go
// internal/user/handlers/user_handler.go

// Create godoc
// @Summary      Create a new user
// @Description  Create a new user with the provided information
// @Tags         users
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        request body dto.CreateUserRequest true "User creation data"
// @Success      201 {object} response.APIResponse{data=dto.UserResponse} "User created"
// @Failure      400 {object} response.APIResponse "Invalid request"
// @Failure      401 {object} response.APIResponse "Unauthorized"
// @Failure      409 {object} response.APIResponse "User already exists"
// @Router       /users [post]
func (h *UserHandler) Create(c *gin.Context) {
    // ... handler code
}

// GetByID godoc
// @Summary      Get user by ID
// @Description  Retrieve a single user by their ID
// @Tags         users
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id path string true "User ID"
// @Success      200 {object} response.APIResponse{data=dto.UserResponse} "User found"
// @Failure      401 {object} response.APIResponse "Unauthorized"
// @Failure      404 {object} response.APIResponse "User not found"
// @Router       /users/{id} [get]
func (h *UserHandler) GetByID(c *gin.Context) {
    // ... handler code
}

// List godoc
// @Summary      List users
// @Description  Get a paginated list of users
// @Tags         users
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        page query int false "Page number" default(1)
// @Param        limit query int false "Items per page" default(20)
// @Param        search query string false "Search by name or email"
// @Success      200 {object} response.PaginatedResponse{data=[]dto.UserResponse} "Users list"
// @Failure      401 {object} response.APIResponse "Unauthorized"
// @Router       /users [get]
func (h *UserHandler) List(c *gin.Context) {
    // ... handler code
}

// Update godoc
// @Summary      Update user
// @Description  Update an existing user's information
// @Tags         users
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id path string true "User ID"
// @Param        request body dto.UpdateUserRequest true "Update data"
// @Success      200 {object} response.APIResponse{data=dto.UserResponse} "User updated"
// @Failure      400 {object} response.APIResponse "Invalid request"
// @Failure      401 {object} response.APIResponse "Unauthorized"
// @Failure      404 {object} response.APIResponse "User not found"
// @Router       /users/{id} [put]
func (h *UserHandler) Update(c *gin.Context) {
    // ... handler code
}

// Delete godoc
// @Summary      Delete user
// @Description  Soft-delete a user by ID
// @Tags         users
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id path string true "User ID"
// @Success      200 {object} response.APIResponse{data=map[string]string} "Deletion successful"
// @Failure      401 {object} response.APIResponse "Unauthorized"
// @Failure      404 {object} response.APIResponse "User not found"
// @Router       /users/{id} [delete]
func (h *UserHandler) Delete(c *gin.Context) {
    // ... handler code
}
```

### Adding Swagger Annotations to DTOs

Add descriptions and examples to your DTOs:

```go
// internal/user/dto/user_dto.go
package dto

// CreateUserRequest - request for creating a user
// @Description User creation request body
type CreateUserRequest struct {
    // User's email address (must be unique)
    Email string `json:"email" validate:"required,email" example:"user@example.com"`
    // User's full name
    Name string `json:"name" validate:"required,min=2,max=100" example:"John Doe"`
    // User's first name
    FirstName string `json:"firstName" validate:"max=50" example:"John"`
    // User's last name
    LastName string `json:"lastName" validate:"max=50" example:"Doe"`
}

// UserResponse - user data returned to client
// @Description User profile information
type UserResponse struct {
    ID        string `json:"id" example:"550e8400-e29b-41d4-a716-446655440000"`
    Email     string `json:"email" example:"user@example.com"`
    Name      string `json:"name" example:"John Doe"`
    FirstName string `json:"firstName,omitempty" example:"John"`
    LastName  string `json:"lastName,omitempty" example:"Doe"`
    Picture   string `json:"picture,omitempty" example:"https://example.com/avatar.jpg"`
    Provider  string `json:"provider" example:"google"`
    IsActive  bool   `json:"isActive" example:"true"`
    CreatedAt string `json:"createdAt" example:"2024-01-01T00:00:00Z"`
    UpdatedAt string `json:"updatedAt" example:"2024-01-01T00:00:00Z"`
}
```

### Swagger Annotation Reference

| Annotation | Description | Example |
|------------|-------------|---------|
| `@Summary` | Short description | `@Summary Create a new user` |
| `@Description` | Detailed description | `@Description Creates a user with the given data` |
| `@Tags` | API group/category | `@Tags users` |
| `@Accept` | Request content type | `@Accept json` |
| `@Produce` | Response content type | `@Produce json` |
| `@Security` | Auth requirement | `@Security BearerAuth` |
| `@Param` | Parameter definition | See below |
| `@Success` | Success response | `@Success 200 {object} dto.Response` |
| `@Failure` | Error response | `@Failure 400 {object} response.APIResponse` |
| `@Router` | Route path and method | `@Router /users/{id} [get]` |

### @Param Syntax

```
@Param  name  location  type  required  "description"
```

| Location | Use Case |
|----------|----------|
| `path` | URL parameter like `/users/{id}` |
| `query` | Query string like `?page=1` |
| `body` | Request body (use with `true` struct type) |
| `header` | Request header |
| `formData` | Form data |

**Examples:**
```go
// Path parameter
// @Param id path string true "User ID"

// Query parameter
// @Param page query int false "Page number" default(1)

// Body parameter (use struct type)
// @Param request body dto.CreateUserRequest true "Request body"

// Optional with default
// @Param limit query int false "Items per page" default(20) minimum(1) maximum(100)
```

---

## Step 14: Generate Documentation & Types

### Generate Swagger Docs

After adding annotations, regenerate Swagger documentation:

```bash
make swagger
```

This creates:
- `docs/swagger/swagger.json` - OpenAPI spec (JSON)
- `docs/swagger/swagger.yaml` - OpenAPI spec (YAML)
- `docs/swagger/docs.go` - Go embed file

### View Swagger UI

Start the server and navigate to:
```
http://localhost:8080/swagger/index.html
```

> Note: Swagger UI is only available in non-production mode.

### Generate TypeScript Types

Generate types for your mobile/frontend app:

```bash
make types
```

This creates:
- `docs/contracts/types.ts` - TypeScript type definitions
- `docs/contracts/api-client.ts` - API client template

**Copy to your mobile app:**
```bash
cp docs/contracts/types.ts ../mobile-app/src/api/
```

### Complete Generation

Generate everything at once:

```bash
make generate-all
# = ent-generate + wire + swagger + types
```

---

## Quick Reference Commands

| Command | Description |
|---------|-------------|
| `make ent-init` | Create new Ent schema |
| `make ent-generate` | Generate Ent code |
| `make migration-create` | Create new migration |
| `make migration-up` | Run migrations |
| `make wire` | Generate Wire DI code |
| `make swagger` | Generate Swagger/OpenAPI docs |
| `make types` | Generate TypeScript types |
| `make generate` | Generate Ent + Wire |
| `make generate-all` | Generate all (Ent + Wire + Swagger + Types) |
| `make run` | Run server |
| `make docker-infra` | Start MySQL, Redis, MinIO |
| `make lint` | Run linter |
| `make test` | Run tests |

---

## Checklist for New Features

- [ ] Create Ent schema (`make ent-init`)
- [ ] Edit schema fields and indexes
- [ ] Generate Ent code (`make ent-generate`)
- [ ] Create migration (`make migration-create`)
- [ ] Run migration (`make migration-up`)
- [ ] Create DTOs in `internal/{module}/dto/` with Swagger examples
- [ ] Create service in `internal/{module}/services/`
- [ ] Create handler in `internal/{module}/handlers/` with Swagger annotations
- [ ] Create providers.go in services/ and handlers/
- [ ] Register in `internal/di/wire.go`
- [ ] Create routes in `internal/{module}/routes/`
- [ ] Register routes in `internal/router/api_v1.go`
- [ ] Generate Wire code (`make wire`)
- [ ] Generate Swagger docs (`make swagger`)
- [ ] Generate TypeScript types (`make types`)
- [ ] Test endpoints
- [ ] Copy types to mobile app
