// Package validation provides input validation utilities for the interview process.
package validation

import (
	"errors"
	"fmt"
	"net/url"
	"regexp"
	"strings"
)

// projectNameRegex validates project names (lowercase letters, numbers, hyphens, no consecutive hyphens).
var projectNameRegex = regexp.MustCompile(`^[a-z0-9]+(-[a-z0-9]+)*$`)

// emailRegex provides basic email validation.
var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)

// ValidateProjectName validates a project name according to the rules:
// - Must be 3-64 characters
// - Only lowercase letters, numbers, and hyphens
// - Cannot start or end with a hyphen
// - No consecutive hyphens
func ValidateProjectName(name string) error {
	if name == "" {
		return errors.New("project name is required")
	}

	if len(name) < 3 {
		return errors.New("project name must be at least 3 characters")
	}

	if len(name) > 64 {
		return errors.New("project name must be at most 64 characters")
	}

	if strings.HasPrefix(name, "-") || strings.HasSuffix(name, "-") {
		return errors.New("project name cannot start or end with a hyphen")
	}

	if !projectNameRegex.MatchString(name) {
		return errors.New("project name can only contain lowercase letters, numbers, and hyphens")
	}

	return nil
}

// ValidateNotEmpty validates that a string is not empty or whitespace-only.
func ValidateNotEmpty(value, fieldName string) error {
	if strings.TrimSpace(value) == "" {
		return fmt.Errorf("%s is required", fieldName)
	}
	return nil
}

// ValidateMinLength validates that a string has at least minLen characters.
func ValidateMinLength(value string, minLen int, fieldName string) error {
	if len(value) < minLen {
		return fmt.Errorf("%s must be at least %d characters", fieldName, minLen)
	}
	return nil
}

// ValidateMaxLength validates that a string has at most maxLen characters.
func ValidateMaxLength(value string, maxLen int, fieldName string) error {
	if len(value) > maxLen {
		return fmt.Errorf("%s must be at most %d characters", fieldName, maxLen)
	}
	return nil
}

// ValidateSelection validates that a value is one of the allowed options.
func ValidateSelection(value string, options []string, fieldName string) error {
	if value == "" {
		return fmt.Errorf("%s is required", fieldName)
	}

	for _, opt := range options {
		if value == opt {
			return nil
		}
	}

	return fmt.Errorf("%s must be one of: %s", fieldName, strings.Join(options, ", "))
}

// ValidateURL validates that a string is a valid HTTP or HTTPS URL.
func ValidateURL(value string) error {
	if value == "" {
		return errors.New("URL is required")
	}

	u, err := url.Parse(value)
	if err != nil {
		return fmt.Errorf("invalid URL: %w", err)
	}

	if u.Scheme != "http" && u.Scheme != "https" {
		return errors.New("URL must use http or https scheme")
	}

	if u.Host == "" {
		return errors.New("URL must have a host")
	}

	return nil
}

// ValidateEmail validates that a string is a valid email address.
func ValidateEmail(value string) error {
	if value == "" {
		return errors.New("email is required")
	}

	if !emailRegex.MatchString(value) {
		return errors.New("invalid email address")
	}

	return nil
}

// ValidationError represents a validation error with field context.
type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("%s: %s", e.Field, e.Message)
}

// NewValidationError creates a new validation error.
func NewValidationError(field, message string) *ValidationError {
	return &ValidationError{
		Field:   field,
		Message: message,
	}
}

// Validator provides a fluent interface for validating values.
type Validator struct {
	errors []error
}

// NewValidator creates a new Validator.
func NewValidator() *Validator {
	return &Validator{
		errors: make([]error, 0),
	}
}

// CheckProjectName validates a project name.
func (v *Validator) CheckProjectName(name string) *Validator {
	if err := ValidateProjectName(name); err != nil {
		v.errors = append(v.errors, err)
	}
	return v
}

// CheckNotEmpty validates that a field is not empty.
func (v *Validator) CheckNotEmpty(value, fieldName string) *Validator {
	if err := ValidateNotEmpty(value, fieldName); err != nil {
		v.errors = append(v.errors, err)
	}
	return v
}

// CheckMinLength validates minimum length.
func (v *Validator) CheckMinLength(value string, minLen int, fieldName string) *Validator {
	if err := ValidateMinLength(value, minLen, fieldName); err != nil {
		v.errors = append(v.errors, err)
	}
	return v
}

// CheckMaxLength validates maximum length.
func (v *Validator) CheckMaxLength(value string, maxLen int, fieldName string) *Validator {
	if err := ValidateMaxLength(value, maxLen, fieldName); err != nil {
		v.errors = append(v.errors, err)
	}
	return v
}

// IsValid returns true if no validation errors occurred.
func (v *Validator) IsValid() bool {
	return len(v.errors) == 0
}

// Errors returns all validation errors.
func (v *Validator) Errors() []error {
	return v.errors
}

// FirstError returns the first validation error, or nil if none.
func (v *Validator) FirstError() error {
	if len(v.errors) > 0 {
		return v.errors[0]
	}
	return nil
}
