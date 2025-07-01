package govalid

import "github.com/Palma99/govalid/internal"

type ValidationResult struct {
	errors []internal.ValidationError
}

func NewValidationResult(
	errors ...internal.ValidationError,
) ValidationResult {
	return ValidationResult{
		errors: errors,
	}
}

// Returns a map of errors grouped by field
func (r ValidationResult) GroupedErrorsByField() map[string][]internal.ValidationError {
	groupedErrors := make(map[string][]internal.ValidationError)

	for _, err := range r.errors {
		groupedErrors[err.Field()] = append(groupedErrors[err.Field()], err)
	}

	return groupedErrors
}

// Returns true if the result contains no errors
func (r ValidationResult) IsValid() bool {
	return !r.HasErrors()
}

// Returns true if the field has no errors
func (r ValidationResult) IsFieldValid(field string) bool {
	return len(r.FieldErrors(field)) == 0
}

// Returns all errors for a given field
func (r ValidationResult) FieldErrors(field string) []internal.ValidationError {
	var errors []internal.ValidationError
	for _, err := range r.errors {
		if err.Field() == field {
			errors = append(errors, err)
		}
	}

	return errors
}

// Returns the number of errors
func (r ValidationResult) ErrorCount() int {
	return len(r.Errors())
}

// Returns true if the result contains at least one error
func (r ValidationResult) HasErrors() bool {
	return len(r.errors) > 0
}

// Returns all the collected errors
func (r ValidationResult) Errors() []internal.ValidationError {
	return r.errors
}

// Returns the first error, or nil
func (r *ValidationResult) FirstError() *internal.ValidationError {
	if !r.HasErrors() {
		return nil
	}

	return &r.errors[0]
}

func (r *ValidationResult) addError(err internal.ValidationError) {
	r.errors = append(r.errors, err)
}
