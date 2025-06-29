package govalid

type ValidationResult struct {
	errors    []ValidationError
	iterIndex int
}

func NewValidationResult(
	errors ...ValidationError,
) ValidationResult {
	return ValidationResult{
		errors:    errors,
		iterIndex: 0,
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

func (r *ValidationResult) NextError() *ValidationError {
	if r.iterIndex >= len(r.errors) {
		return nil
	}
	err := &r.errors[r.iterIndex]
	r.iterIndex++
	return err
}
