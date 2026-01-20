// pkg/apperror/handler.go
package apperror

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ErrorResponse is the JSON response for errors
type ErrorResponse struct {
	Success bool                   `json:"success" example:"false"`
	Error   ErrorDetail            `json:"error"`
}

// ErrorDetail contains error information
type ErrorDetail struct {
	Code    string                 `json:"code" example:"VALIDATION_ERROR"`
	Message string                 `json:"message" example:"Invalid input"`
	Details map[string]interface{} `json:"details,omitempty"`
}

// HandleError sends an error response and logs it
func HandleError(c *gin.Context, err error) {
	var appErr *AppError
	if errors.As(err, &appErr) {
		c.JSON(appErr.HTTPStatus, ErrorResponse{
			Success: false,
			Error: ErrorDetail{
				Code:    appErr.Code,
				Message: appErr.Message,
				Details: appErr.Details,
			},
		})
		return
	}

	// Unknown error - treat as internal error
	c.JSON(http.StatusInternalServerError, ErrorResponse{
		Success: false,
		Error: ErrorDetail{
			Code:    ErrCodeInternal,
			Message: "An unexpected error occurred",
		},
	})
}

// AbortWithError aborts the request with an error
func AbortWithError(c *gin.Context, err error) {
	HandleError(c, err)
	c.Abort()
}
