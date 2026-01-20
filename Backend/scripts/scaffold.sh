#!/bin/bash

# Module scaffolding script
# Creates a new module with handlers, services, repositories, routes, and DTOs

set -e

# Colors
RED='\033[0;31m'
GREEN='\033[0;32m'
BLUE='\033[0;34m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

echo -e "${BLUE}=== Module Scaffolding ===${NC}"
echo ""

read -p "Enter module name (e.g., 'user', 'product'): " MODULE

if [ -z "$MODULE" ]; then
    echo -e "${RED}Error: Module name is required${NC}"
    exit 1
fi

# Convert to lowercase
MODULE=$(echo "$MODULE" | tr '[:upper:]' '[:lower:]')
# Capitalize first letter for struct names
MODULE_UPPER="$(tr '[:lower:]' '[:upper:]' <<< ${MODULE:0:1})${MODULE:1}"

MODULE_PATH="internal/$MODULE"

if [ -d "$MODULE_PATH" ]; then
    echo -e "${RED}Error: Module '$MODULE' already exists${NC}"
    exit 1
fi

echo -e "${YELLOW}Creating module structure for '$MODULE'...${NC}"

# Create directories
mkdir -p "$MODULE_PATH/handlers"
mkdir -p "$MODULE_PATH/services"
mkdir -p "$MODULE_PATH/repositories"
mkdir -p "$MODULE_PATH/routes"
mkdir -p "$MODULE_PATH/dto"

# Create handler file
cat > "$MODULE_PATH/handlers/${MODULE}_handler.go" << EOF
// internal/$MODULE/handlers/${MODULE}_handler.go
package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ZethicTech/be/internal/$MODULE/services"
	"github.com/ZethicTech/be/pkg/response"
)

// ${MODULE_UPPER}Handler handles HTTP requests for $MODULE
type ${MODULE_UPPER}Handler struct {
	service *services.${MODULE_UPPER}Service
}

// New${MODULE_UPPER}Handler creates a new handler
func New${MODULE_UPPER}Handler(service *services.${MODULE_UPPER}Service) *${MODULE_UPPER}Handler {
	return &${MODULE_UPPER}Handler{service: service}
}

// List handles GET requests to list all items
func (h *${MODULE_UPPER}Handler) List(c *gin.Context) {
	items, err := h.service.List(c.Request.Context())
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "INTERNAL_ERROR", err.Error())
		return
	}
	response.JSON(c, http.StatusOK, items)
}

// GetByID handles GET requests to get an item by ID
func (h *${MODULE_UPPER}Handler) GetByID(c *gin.Context) {
	id := c.Param("id")

	item, err := h.service.GetByID(c.Request.Context(), id)
	if err != nil {
		response.Error(c, http.StatusNotFound, "NOT_FOUND", "Item not found")
		return
	}
	response.JSON(c, http.StatusOK, item)
}

// Create handles POST requests to create a new item
func (h *${MODULE_UPPER}Handler) Create(c *gin.Context) {
	// Implement create logic
	response.JSON(c, http.StatusCreated, gin.H{"message": "Created"})
}

// Update handles PUT requests to update an item
func (h *${MODULE_UPPER}Handler) Update(c *gin.Context) {
	// Implement update logic
	response.JSON(c, http.StatusOK, gin.H{"message": "Updated"})
}

// Delete handles DELETE requests to delete an item
func (h *${MODULE_UPPER}Handler) Delete(c *gin.Context) {
	// Implement delete logic
	response.JSON(c, http.StatusOK, gin.H{"message": "Deleted"})
}
EOF

# Create service file
cat > "$MODULE_PATH/services/${MODULE}_service.go" << EOF
// internal/$MODULE/services/${MODULE}_service.go
package services

import (
	"context"

	ent "github.com/ZethicTech/be/ent/generated"
)

// ${MODULE_UPPER}Service handles business logic for $MODULE
type ${MODULE_UPPER}Service struct {
	client *ent.Client
}

// New${MODULE_UPPER}Service creates a new service
func New${MODULE_UPPER}Service(client *ent.Client) *${MODULE_UPPER}Service {
	return &${MODULE_UPPER}Service{client: client}
}

// List returns all items
func (s *${MODULE_UPPER}Service) List(ctx context.Context) ([]interface{}, error) {
	// Implement list logic
	return []interface{}{}, nil
}

// GetByID returns an item by ID
func (s *${MODULE_UPPER}Service) GetByID(ctx context.Context, id string) (interface{}, error) {
	// Implement get by ID logic
	return nil, nil
}

// Create creates a new item
func (s *${MODULE_UPPER}Service) Create(ctx context.Context, data interface{}) (interface{}, error) {
	// Implement create logic
	return nil, nil
}

// Update updates an item
func (s *${MODULE_UPPER}Service) Update(ctx context.Context, id string, data interface{}) (interface{}, error) {
	// Implement update logic
	return nil, nil
}

// Delete deletes an item
func (s *${MODULE_UPPER}Service) Delete(ctx context.Context, id string) error {
	// Implement delete logic
	return nil
}
EOF

# Create routes file
cat > "$MODULE_PATH/routes/${MODULE}_routes.go" << EOF
// internal/$MODULE/routes/${MODULE}_routes.go
package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"

	ent "github.com/ZethicTech/be/ent/generated"
	"github.com/ZethicTech/be/internal/$MODULE/handlers"
	"github.com/ZethicTech/be/internal/$MODULE/services"
)

// Register${MODULE_UPPER}Routes registers all routes for $MODULE
func Register${MODULE_UPPER}Routes(router *gin.RouterGroup, entClient *ent.Client, redisClient *redis.Client) {
	// Create service and handler
	service := services.New${MODULE_UPPER}Service(entClient)
	handler := handlers.New${MODULE_UPPER}Handler(service)

	// Register routes
	group := router.Group("/${MODULE}s")
	{
		group.GET("/", handler.List)
		group.GET("/:id", handler.GetByID)
		group.POST("/", handler.Create)
		group.PUT("/:id", handler.Update)
		group.DELETE("/:id", handler.Delete)
	}
}
EOF

# Create DTO files
cat > "$MODULE_PATH/dto/${MODULE}_dto.go" << EOF
// internal/$MODULE/dto/${MODULE}_dto.go
package dto

// ${MODULE_UPPER}CreateDTO is the request body for creating a $MODULE
type ${MODULE_UPPER}CreateDTO struct {
	Name string \`json:"name" validate:"required"\`
}

// ${MODULE_UPPER}UpdateDTO is the request body for updating a $MODULE
type ${MODULE_UPPER}UpdateDTO struct {
	Name string \`json:"name"\`
}

// ${MODULE_UPPER}ResponseDTO is the response body for a $MODULE
type ${MODULE_UPPER}ResponseDTO struct {
	ID   string \`json:"id"\`
	Name string \`json:"name"\`
}
EOF

echo -e "${GREEN}Module '$MODULE' created successfully!${NC}"
echo ""
echo "Created files:"
echo "  - $MODULE_PATH/handlers/${MODULE}_handler.go"
echo "  - $MODULE_PATH/services/${MODULE}_service.go"
echo "  - $MODULE_PATH/routes/${MODULE}_routes.go"
echo "  - $MODULE_PATH/dto/${MODULE}_dto.go"
echo ""
echo -e "${YELLOW}Don't forget to register the routes in internal/router/api_v1.go:${NC}"
echo "  ${MODULE}routes.Register${MODULE_UPPER}Routes(api, entClient, redisClient)"
