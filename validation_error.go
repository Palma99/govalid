package govalid

import (
	"fmt"
)

type ValidationError struct {
	field   string
	message string
}

func (e ValidationError) Error() error {
	return fmt.Errorf("%s: %s", e.Field(), e.Message())
}

func NewValidationError(field, message string) *ValidationError {
	return &ValidationError{
		field:   field,
		message: message,
	}
}

func NewValidationErrorf(field, format string, args ...interface{}) *ValidationError {
	return &ValidationError{
		field:   field,
		message: fmt.Sprintf(format, args...),
	}
}

func (e ValidationError) Field() string {
	return e.field
}

func (e ValidationError) Message() string {
	return e.message
}
