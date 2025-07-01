package validators

import (
	"github.com/Palma99/govalid"
)

type Rule func(customMessage ...string) govalid.ValidationRule

// CustomRule is a function that returns a ValidationRule
// that uses a custom validator
func CustomRule[T any](validator Validator) Rule {
	return func(customMessage ...string) govalid.ValidationRule {
		return func(field string, value any) govalid.ValidationFunc {
			return validator(field, value, customMessage...)
		}
	}
}

func NonEmptyRule(customMessage ...string) govalid.ValidationRule {
	return func(field string, value any) govalid.ValidationFunc {
		return NonEmpty(field, value, customMessage...)
	}
}

func MaxLengthRule(max int, customMessage ...string) govalid.ValidationRule {
	return func(field string, value any) govalid.ValidationFunc {
		return MaxLength(field, value.(string), max, customMessage...)
	}
}

func MinLengthRule(min int, customMessage ...string) govalid.ValidationRule {
	return func(field string, value any) govalid.ValidationFunc {
		return MinLength(field, value.(string), min, customMessage...)
	}
}

func MaxRule(max int, customMessage ...string) govalid.ValidationRule {
	return func(field string, value any) govalid.ValidationFunc {
		return Max(field, value, max, customMessage...)
	}
}

func MinRule(min int, customMessage ...string) govalid.ValidationRule {
	return func(field string, value any) govalid.ValidationFunc {
		return Min(field, value, min, customMessage...)
	}
}

func MatchesRegexRule(pattern string, customMessage ...string) govalid.ValidationRule {
	return func(field string, value any) govalid.ValidationFunc {
		return MatchesRegex(field, value.(string), pattern, customMessage...)
	}
}

func IsEmailRule(customMessage ...string) govalid.ValidationRule {
	return func(field string, value any) govalid.ValidationFunc {
		return IsEmail(field, value.(string), customMessage...)
	}
}
