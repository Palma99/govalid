package govalid

import "fmt"

// Apply all validations and returns ValidationResult
func Validate(validations ...ValidationFunc) ValidationResult {
	result := NewValidationResult()
	for _, validation := range validations {
		if err := validation(); err != nil {
			result.AddError(*err)
		}
	}

	return result
}

// Validate complex validation groups and composes
// It accepts ValidationFunc and/or []ValidationFunc and panics if other type is passed
//
// i.e. Validating a group and a composed
//
//	composed := govalid.ComposeAll(
//		govalid.NonEmpty("name", person.Name),
//		govalid.NonEmpty("surname", person.Surname),
//	)
//
//	group := govalid.GroupAll("email", person.Email,
//		govalid.NonEmptyRule(),
//		govalid.IsEmailRule()
//	)
//
//	govalid.ValidateAll(composed, group)
func ValidateAll(validations ...any) ValidationResult {
	result := NewValidationResult()

	for _, v := range validations {
		switch vv := v.(type) {
		case ValidationFunc:
			if err := vv(); err != nil {
				result.AddError(*err)
			}
		case []ValidationFunc:
			vE := Validate(vv...)
			result.errors = append(result.errors, vE.errors...)
		default:
			panic(fmt.Sprintf("Validate: unsupported type %T", v))
		}
	}

	return result
}
