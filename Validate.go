package govalid

import "fmt"

const (
	failFastMode    = true
	validateAllMode = false
)

func applyValidations(failFastMode bool, validations ...ValidationFunc) ValidationResult {
	result := NewValidationResult()
	for _, validation := range validations {
		if err := validation(); err != nil {
			result.addError(*err)
			if failFastMode {
				return result
			}
		}
	}

	return result
}

func validate(failFastMode bool, validator any) ValidationResult {
	result := NewValidationResult()

	switch v := validator.(type) {
	case ValidationFunc:
		if err := v(); err != nil {
			result.addError(*err)
		}
	case []ValidationFunc:
		vE := applyValidations(failFastMode, v...)
		result.errors = append(result.errors, vE.errors...)
	default:
		panic(fmt.Sprintf("Validate: unsupported type %T", v))
	}

	return result
}

// Runs all validations and returns errors
// It accepts ValidationFunc and/or []ValidationFunc and panics if other type is passed
//
// i.e. Validating a group and a composed
//
//	composed := govalid.Compose(
//		govalid.NonEmpty("name", person.Name),
//		govalid.NonEmpty("surname", person.Surname),
//	)
//
//	group := govalid.Group("email", person.Email,
//		govalid.NonEmptyRule(),
//		govalid.IsEmailRule()
//	)
//
//	govalid.Validate(composed, group)
func Validate(validations ...any) ValidationResult {
	result := NewValidationResult()

	for _, v := range validations {
		validationResult := validate(validateAllMode, v)
		result.errors = append(result.errors, validationResult.Errors()...)
	}

	return result
}

// Runs all validators and stops at the first error, if any
func ValidateShortCircuit(validations ...any) ValidationResult {
	result := NewValidationResult()

	for _, v := range validations {
		validationResult := validate(failFastMode, v)
		if validationResult.HasErrors() {
			return validationResult
		}
	}

	return result
}
