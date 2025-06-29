package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidationResultConstructorWithoutErrors(t *testing.T) {
	res := NewValidationResult()

	assert.False(t, res.HasErrors())
	assert.Len(t, res.Errors(), 0)
}

func TestValidationResultConstructorWithErrors(t *testing.T) {
	res := NewValidationResult(
		*NewValidationError("name", "test error"),
		*NewValidationError("surname", "test error"),
		*NewValidationError("age", "test error"),
	)

	assert.True(t, res.HasErrors())
	assert.Len(t, res.Errors(), 3)
}

func TestValidationResultAddError(t *testing.T) {
	res := NewValidationResult()

	res.AddError(*NewValidationError("name", "test error"))
	res.AddError(*NewValidationError("surname", "test error"))
	res.AddError(*NewValidationError("age", "test error"))

	assert.True(t, res.HasErrors())
	assert.Len(t, res.Errors(), 3)
}

func TestValidationResultNextError(t *testing.T) {
	type testCase struct {
		name            string
		expectedHasMore bool
		expectedError   *ValidationError
		testErrors      []*ValidationError
	}

	testCases := []testCase{
		{
			name:            "should return nil	if no errors",
			expectedHasMore: false,
			expectedError:   nil,
			testErrors:      []*ValidationError{},
		},
		{
			name:          "should return the first error if called one time",
			expectedError: NewValidationError("name", "test error1"),
			testErrors: []*ValidationError{
				NewValidationError("name", "test error1"),
				NewValidationError("surname", "test error2"),
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := NewValidationResult()

			for _, err := range tc.testErrors {
				res.AddError(*err)
			}

			validationError := res.NextError()
			assert.Equal(t, tc.expectedError, validationError)
		})
	}
}

func TestValidationResultNextErrorShouldReturnAllErrors(t *testing.T) {
	res := NewValidationResult()

	res.AddError(*NewValidationError("name", "test error1"))
	res.AddError(*NewValidationError("surname", "test error2"))
	res.AddError(*NewValidationError("age", "test error3"))

	for i := range 3 {
		validationError := res.NextError()

		assert.NotNil(t, validationError)
		assert.Equal(t, res.Errors()[i].Field(), validationError.Field())
		assert.Equal(t, res.Errors()[i].Message(), validationError.Message())
	}
}

func TestValidationResultNextErrorShouldReturnNilIfAllErrorsAreConsumed(t *testing.T) {
	res := NewValidationResult(
		*NewValidationError("name", "test error1"),
		*NewValidationError("surname", "test error2"),
		*NewValidationError("age", "test error3"),
	)

	for range res.Errors() {
		err := res.NextError()
		assert.NotNil(t, err)
	}

	assert.Nil(t, res.NextError())
	assert.Nil(t, res.NextError())
}
