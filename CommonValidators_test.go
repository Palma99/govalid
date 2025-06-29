package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNonEmptyValidator(t *testing.T) {
	t.Run("should return error when empty", func(t *testing.T) {
		err := NonEmpty("name", "")()
		assert.NotNil(t, err)
		assert.Equal(t, "name", err.Field())
		assert.Equal(t, "must not be empty", err.Message())
	})

	t.Run("should return nil when non-empty", func(t *testing.T) {
		err := NonEmpty("name", "Mario")()
		assert.Nil(t, err)
	})

	t.Run("should support custom error message", func(t *testing.T) {
		err := NonEmpty("name", "", "campo obbligatorio")()
		assert.NotNil(t, err)
		assert.Equal(t, "campo obbligatorio", err.Message())
	})
}

func TestMaxLengthValidator(t *testing.T) {
	t.Run("should return error when value exceeds max length", func(t *testing.T) {
		err := MaxLength("bio", "too long string", 5)()
		assert.NotNil(t, err)
		assert.Equal(t, "bio", err.Field())
		assert.Equal(t, "must be at most 5 characters", err.Message())
	})

	t.Run("should return nil when value is within limit", func(t *testing.T) {
		err := MaxLength("bio", "short", 10)()
		assert.Nil(t, err)
	})
}

func TestMatchesRegexValidator(t *testing.T) {
	t.Run("should return error if string doesn't match regex", func(t *testing.T) {
		err := MatchesRegex("code", "1234", `^[A-Z]{4}$`)()
		assert.NotNil(t, err)
		assert.Equal(t, "code", err.Field())
		assert.Contains(t, err.Message(), "must match pattern")
	})

	t.Run("should return nil if string matches regex", func(t *testing.T) {
		err := MatchesRegex("code", "ABCD", `^[A-Z]{4}$`)()
		assert.Nil(t, err)
	})
}

func TestIsEmailValidator(t *testing.T) {
	t.Run("should return error for invalid email", func(t *testing.T) {
		err := IsEmail("email", "not-an-email")()
		assert.NotNil(t, err)
		assert.Equal(t, "email", err.Field())
		assert.Contains(t, err.Message(), "must match pattern")
	})

	t.Run("should return nil for valid email", func(t *testing.T) {
		err := IsEmail("email", "test@example.com")()
		assert.Nil(t, err)
	})
}
