package validators

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/Palma99/govalid"
	"github.com/Palma99/govalid/internal"
	"github.com/Palma99/govalid/internal/utils"
)

func NonEmpty(fieldName string, value any, args ...string) govalid.ValidationFunc {
	return func() *govalid.ValidationError {
		validationError := govalid.NewValidationError(
			fieldName,
			utils.GetOptionalStringOrDefault("must not be empty", args...),
		)

		switch v := value.(type) {
		case string:
			if strings.TrimSpace(v) == "" {
				return validationError
			}
		case []any:
			if len(v) == 0 {
				return validationError
			}
		case map[any]any:
			if len(v) == 0 {
				return validationError
			}
		}
		return nil
	}
}

func Min[T internal.Number](fieldName string, value any, min T, args ...string) govalid.ValidationFunc {
	return func() *govalid.ValidationError {
		switch v := value.(type) {
		case T:
			if v < min {
				return govalid.NewValidationErrorf(
					fieldName,
					utils.GetOptionalStringOrDefault(
						fmt.Sprintf("must be at least %v", min),
						args...,
					),
				)
			}
		}
		return nil
	}
}

func Max[T internal.Number](fieldName string, value any, max T, args ...string) govalid.ValidationFunc {
	return func() *govalid.ValidationError {
		switch v := value.(type) {
		case T:
			if v > max {
				return govalid.NewValidationErrorf(
					fieldName,
					utils.GetOptionalStringOrDefault(
						fmt.Sprintf("must be at most %v", max),
						args...,
					),
				)
			}
		}
		return nil
	}
}

func MinLength(fieldName string, value any, min int, args ...string) govalid.ValidationFunc {
	return func() *govalid.ValidationError {
		length, err := utils.GetLength(value)
		if err != nil {
			return govalid.NewValidationError(fieldName, err.Error())
		}

		if length < min {
			msg := utils.GetOptionalStringOrDefault(
				fmt.Sprintf("must be at least %d characters", min),
				args...,
			)
			return govalid.NewValidationError(fieldName, msg)
		}
		return nil
	}
}

func MaxLength(fieldName string, value any, max int, args ...string) govalid.ValidationFunc {
	return func() *govalid.ValidationError {
		length, err := utils.GetLength(value)
		if err != nil {
			return govalid.NewValidationError(fieldName, err.Error())
		}

		if length > max {
			msg := utils.GetOptionalStringOrDefault(
				fmt.Sprintf("must be at most %d characters", max),
				args...,
			)
			return govalid.NewValidationError(fieldName, msg)
		}
		return nil
	}
}

func MatchesRegex(fieldName, value, pattern string, args ...string) govalid.ValidationFunc {
	return func() *govalid.ValidationError {
		matched, err := regexp.MatchString(pattern, value)
		if err != nil || !matched {
			return govalid.NewValidationError(
				fieldName,
				utils.GetOptionalStringOrDefault(
					fmt.Sprintf("must match pattern %s", pattern),
					args...,
				),
			)
		}
		return nil
	}
}

func IsEmail(fieldName, value string, args ...string) govalid.ValidationFunc {
	const emailPattern = `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`
	return MatchesRegex(fieldName, value, emailPattern, args...)
}
