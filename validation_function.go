package govalid

type ValidationFunc func() *ValidationError

type ValidationRule func(field string, value any) ValidationFunc

// Compose all validation functions into one and return only the first error if any
//
// In this example only the name error will be returned, surname validation will not be evaluated
//
//	composed := govalid.Compose(
//		govalid.NonEmpty("name", "")
//		govalid.NonEmpty("surname", ""),
//	)
//
//	res := govalid.Validate(composed)
func Compose(validations ...ValidationFunc) ValidationFunc {
	return func() *ValidationError {
		for _, validation := range validations {
			if err := validation(); err != nil {
				return err
			}
		}
		return nil
	}
}

func ComposeAll(validations ...ValidationFunc) []ValidationFunc {
	return validations
}

// Utility function to create a group of validation rules for a field
// Validating a group will return the first error
func Group(fieldName string, value any, rules ...ValidationRule) ValidationFunc {
	funcs := GroupAll(fieldName, value, rules...)
	return Compose(funcs...)
}

// Utility function to create a group of validation rules for a field
// Validating an object created with GroupAll will return all errors
func GroupAll(fieldName string, value any, rules ...ValidationRule) []ValidationFunc {
	funcs := make([]ValidationFunc, 0, len(rules))
	for _, rule := range rules {
		funcs = append(funcs, rule(fieldName, value))
	}
	return funcs
}
