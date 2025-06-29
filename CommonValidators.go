package core

import (
	"fmt"
	"regexp"
)

func NonEmpty(fieldName, value string, args ...interface{}) ValidationFunc {
	message := "must not be empty"
	if len(args) > 0 {
		message = fmt.Sprintf(args[0].(string), args[1:]...)
	}

	return func() *ValidationError {
		if value == "" {
			return NewValidationError(
				fieldName,
				message,
			)
		}
		return nil
	}
}

func MaxLength(fieldName, value string, max int, args ...interface{}) ValidationFunc {
	return func() *ValidationError {
		if len(value) > max {
			return NewValidationErrorf(
				fieldName,
				"must be at most %d characters",
				max,
			)
		}
		return nil
	}
}

func MatchesRegex(fieldName, value, pattern string) ValidationFunc {
	return func() *ValidationError {
		matched, err := regexp.MatchString(pattern, value)
		if err != nil || !matched {
			return NewValidationError(
				fieldName,
				fmt.Sprintf("must match pattern %s", pattern),
			)
		}
		return nil
	}
}

func IsEmail(fieldName, value string) ValidationFunc {
	const emailPattern = `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`
	return MatchesRegex(fieldName, value, emailPattern)
}
