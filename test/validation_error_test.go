package govalid_test

import (
	"errors"
	"testing"

	"github.com/Palma99/govalid/internal"
	"github.com/stretchr/testify/assert"
)

func TestNewValidationError(t *testing.T) {
	err := internal.NewValidationError("name", "required")

	assert.NotNil(t, err)
	assert.Equal(t, "name", err.Field())
	assert.Equal(t, "required", err.Message())
}

func TestNewValidationErrorf(t *testing.T) {
	err := internal.NewValidationErrorf("age", "must be at least %d", 18)

	assert.NotNil(t, err)
	assert.Equal(t, "age", err.Field())
	assert.Equal(t, "must be at least 18", err.Message())
}

func TestValidationError_ErrorMethod(t *testing.T) {
	err := internal.NewValidationError("email", "invalid format")

	e := err.Error()
	assert.NotNil(t, e)
	assert.IsType(t, errors.New(""), e)

	expectedMsg := "email: invalid format"
	assert.Equal(t, expectedMsg, e.Error())
}
