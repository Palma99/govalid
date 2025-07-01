package validators_test

import (
	"testing"

	"github.com/Palma99/govalid/validators"
	"github.com/stretchr/testify/assert"
)

func TestNonEmptyRule(t *testing.T) {
	rule := validators.NonEmptyRule()
	validator := rule("username", "")
	err := validator()
	assert.NotNil(t, err)
	assert.Equal(t, "username", err.Field())

	validator = rule("username", "john")
	assert.Nil(t, validator())
}

func TestMaxLengthRule(t *testing.T) {
	rule := validators.MaxLengthRule(5)
	validator := rule("code", "123456")
	err := validator()
	assert.NotNil(t, err)
	assert.Equal(t, "code", err.Field())

	validator = rule("code", "123")
	assert.Nil(t, validator())
}

func TestMinLengthRule(t *testing.T) {
	rule := validators.MinLengthRule(3)
	validator := rule("password", "ab")
	err := validator()
	assert.NotNil(t, err)
	assert.Equal(t, "password", err.Field())

	validator = rule("password", "abc")
	assert.Nil(t, validator())
}

func TestMatchesRegexRule(t *testing.T) {
	pattern := `^[a-z]+$`
	rule := validators.MatchesRegexRule(pattern)

	validator := rule("slug", "abc123")
	err := validator()
	assert.NotNil(t, err)
	assert.Equal(t, "slug", err.Field())

	validator = rule("slug", "abc")
	assert.Nil(t, validator())
}

func TestIsEmailRule(t *testing.T) {
	rule := validators.IsEmailRule()

	validator := rule("email", "invalid-email")
	err := validator()
	assert.NotNil(t, err)
	assert.Equal(t, "email", err.Field())

	validator = rule("email", "user@example.com")
	assert.Nil(t, validator())
}

func TestMaxRule(t *testing.T) {
	rule := validators.MaxRule(5)
	validator := rule("code", 6)
	err := validator()
	assert.NotNil(t, err)
	assert.Equal(t, "code", err.Field())

	validator = rule("code", 3)
	assert.Nil(t, validator())
}

func TestMinRule(t *testing.T) {
	rule := validators.MinRule(3)
	validator := rule("code", 2)
	err := validator()
	assert.NotNil(t, err)
	assert.Equal(t, "code", err.Field())

	validator = rule("code", 5)
	assert.Nil(t, validator())
}
