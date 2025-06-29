package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func createValidatorSpy(counter *int, testErr *ValidationError) ValidationFunc {
	return func() *ValidationError {
		*counter += 1
		if testErr != nil {
			return testErr
		}
		return nil
	}
}

func TestValidateShouldCallAllValidators(t *testing.T) {
	callCount := 0

	Validate(
		createValidatorSpy(&callCount, nil),
		createValidatorSpy(&callCount, nil),
		createValidatorSpy(&callCount, nil),
	)

	assert.Equal(t, 3, callCount)
}

func TestValidateShouldReturnAllErrors(t *testing.T) {
	callCount := 0

	res := Validate(
		createValidatorSpy(&callCount, NewValidationError("field1", "test error1")),
		createValidatorSpy(&callCount, NewValidationError("field2", "test error2")),
		createValidatorSpy(&callCount, nil),
		createValidatorSpy(&callCount, NewValidationError("field3", "test error3")),
	)

	assert.Equal(t, 4, callCount)

	assert.True(t, res.HasErrors())

	errors := res.Errors()
	assert.Len(t, errors, 3)

	assert.Equal(t, "field1", errors[0].Field())
	assert.Equal(t, "test error1", errors[0].Message())
	assert.Equal(t, "field2", errors[1].Field())
	assert.Equal(t, "test error2", errors[1].Message())
	assert.Equal(t, "field3", errors[2].Field())
	assert.Equal(t, "test error3", errors[2].Message())
}

func TestValidateShouldReturnNoErrors(t *testing.T) {
	callCount := 0

	res := Validate(
		createValidatorSpy(&callCount, nil),
		createValidatorSpy(&callCount, nil),
	)

	assert.Equal(t, 2, callCount)
	assert.False(t, res.HasErrors())
	assert.Len(t, res.Errors(), 0)
}
