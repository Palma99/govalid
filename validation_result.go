package govalid

type ValidationResult struct {
	errors []ValidationError
}

func NewValidationResult(
	errors ...ValidationError,
) ValidationResult {
	return ValidationResult{
		errors: errors,
	}
}

func (r *ValidationResult) AddError(err ValidationError) {
	r.errors = append(r.errors, err)
}

func (r ValidationResult) HasErrors() bool {
	return len(r.errors) > 0
}

func (r ValidationResult) Errors() []ValidationError {
	return r.errors
}

func (r *ValidationResult) FirstError() *ValidationError {
	if !r.HasErrors() {
		return nil
	}

	return &r.errors[0]
}
