package core

func Validate(validations ...ValidationFunc) ValidationResult {
	result := NewValidationResult()
	for _, validation := range validations {
		if err := validation(); err != nil {
			result.AddError(*err)
		}
	}

	return result
}
