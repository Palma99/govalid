package govalid_test

import (
	"errors"
	"testing"

	"github.com/Palma99/govalid"
	"github.com/stretchr/testify/assert"
)

func TestNewValidationError(t *testing.T) {
	err := govalid.NewValidationError("name", "required")

	assert.NotNil(t, err)
	assert.Equal(t, "name", err.Field())
	assert.Equal(t, "required", err.Message())
}

func TestNewValidationErrorf(t *testing.T) {
	err := govalid.NewValidationErrorf("age", "must be at least %d", 18)

	assert.NotNil(t, err)
	assert.Equal(t, "age", err.Field())
	assert.Equal(t, "must be at least 18", err.Message())
}

func TestValidationError_ErrorMethod(t *testing.T) {
	err := govalid.NewValidationError("email", "invalid format")

	e := err.Error()
	assert.NotNil(t, e)
	assert.IsType(t, errors.New(""), e)

	expectedMsg := "email: invalid format"
	assert.Equal(t, expectedMsg, e.Error())
}
