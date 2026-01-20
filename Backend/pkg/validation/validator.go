// pkg/validation/validator.go
package validation

import (
	"github.com/go-playground/validator/v10"
)

// Validator wraps the validator instance
type Validator struct {
	validate *validator.Validate
}

// New returns a configured Validator
func New() *Validator {
	v := validator.New()

	// Register custom validations here if needed
	// v.RegisterValidation("custom", customFunc)

	return &Validator{validate: v}
}

// Validate validates any struct using the validator instance
func (v *Validator) Validate(i interface{}) error {
	return v.validate.Struct(i)
}
