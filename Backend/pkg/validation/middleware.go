// pkg/validation/middleware.go
package validation

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/UnoraApp/be/pkg/response"
)

// ValidateRequest is a generic middleware for request validation
func ValidateRequest[T any](v *Validator) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req T

		// Read the raw body
		body, err := c.GetRawData()
		if err != nil {
			response.Error(c, http.StatusBadRequest, "INVALID_REQUEST", "Failed to read request body")
			c.Abort()
			return
		}

		// Parse JSON to check field names
		var rawJSON map[string]interface{}
		if err := json.Unmarshal(body, &rawJSON); err != nil {
			response.Error(c, http.StatusBadRequest, "INVALID_REQUEST", fmt.Sprintf("Invalid JSON: %v", err))
			c.Abort()
			return
		}

		// Get expected field names from struct tags
		expectedFields := getJSONFieldNames[T]()

		// Check if all provided fields match expected camelCase names
		for fieldName := range rawJSON {
			if !contains(expectedFields, fieldName) {
				response.Error(c, http.StatusBadRequest, "INVALID_REQUEST", fmt.Sprintf("Invalid field name: '%s'", fieldName))
				c.Abort()
				return
			}
		}

		// Use a case-sensitive JSON decoder
		decoder := json.NewDecoder(bytes.NewReader(body))
		decoder.DisallowUnknownFields()

		if err := decoder.Decode(&req); err != nil {
			response.Error(c, http.StatusBadRequest, "INVALID_REQUEST", fmt.Sprintf("JSON parsing error: %v", err))
			c.Abort()
			return
		}

		// Validate struct
		if err := v.Validate(&req); err != nil {
			response.Error(c, http.StatusBadRequest, "VALIDATION_ERROR", err.Error())
			c.Abort()
			return
		}

		// Save validated request in context
		c.Set("validatedBody", req)
		c.Next()
	}
}

// getJSONFieldNames extracts JSON field names from struct tags
func getJSONFieldNames[T any]() []string {
	var req T
	t := reflect.TypeOf(req)

	var fields []string
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		jsonTag := field.Tag.Get("json")
		if jsonTag != "" && jsonTag != "-" {
			fieldName := strings.Split(jsonTag, ",")[0]
			fields = append(fields, fieldName)
		}
	}
	return fields
}

// contains checks if a slice contains a string
func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
