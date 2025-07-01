package govalid

import "github.com/Palma99/govalid/internal"

type ValidationFunc func() *internal.ValidationError

type ValidationRule func(field string, value any) ValidationFunc

// ComposeShortCircuit combines all validation functions into one and return only the first error if any
//
// In this example only the name error will be returned, surname validation will not be evaluated
//
//	composed := govalid.ComposeShortCircuit(
//		govalid.NonEmpty("name", "")
//		govalid.NonEmpty("surname", ""),
//	)
//
//	res := govalid.Validate(composed)
func ComposeShortCircuit(validations ...any) ValidationFunc {
	return func() *internal.ValidationError {
		for _, v := range validations {
			switch validator := v.(type) {
			case ValidationFunc:
				if err := validator(); err != nil {
					return err
				}
			case []ValidationFunc:
				for _, validation := range validator {
					if err := validation(); err != nil {
						return err
					}
				}
			default:
				panic("ComposeShortCircuit: unsupported type")
			}
		}
		return nil
	}

}

// Compose all validation functions into one
// Validating an object created with Compose will return all errors
func Compose(validations ...any) []ValidationFunc {
	funcs := make([]ValidationFunc, 0, len(validations))
	for _, v := range validations {
		switch validator := v.(type) {
		case ValidationFunc:
			funcs = append(funcs, validator)
		case []ValidationFunc:
			funcs = append(funcs, validator...)
		}
	}
	return funcs
}

// Utility function to create a group of validation rules for a field
// Validating a group will return the first error
func GroupShortCircuit(fieldName string, value any, rules ...ValidationRule) ValidationFunc {
	funcs := Group(fieldName, value, rules...)
	return ComposeShortCircuit(funcs)
}

// Utility function to create a group of validation rules for a field
// Validating an object created with Group will return all errors
func Group(fieldName string, value any, rules ...ValidationRule) []ValidationFunc {
	funcs := make([]ValidationFunc, 0, len(rules))
	for _, rule := range rules {
		funcs = append(funcs, rule(fieldName, value))
	}
	return funcs
}
