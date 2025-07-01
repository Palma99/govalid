package govalid_test

import (
	"testing"

	"github.com/Palma99/govalid"
	"github.com/stretchr/testify/assert"
)

func TestValidationResultConstructorWithoutErrors(t *testing.T) {
	res := govalid.NewValidationResult()

	assert.False(t, res.HasErrors())
	assert.Len(t, res.Errors(), 0)
}

func TestValidationResultConstructorWithErrors(t *testing.T) {
	res := govalid.NewValidationResult(
		*govalid.NewValidationError("name", "test error"),
		*govalid.NewValidationError("surname", "test error"),
		*govalid.NewValidationError("age", "test error"),
	)

	assert.True(t, res.HasErrors())
	assert.Len(t, res.Errors(), 3)
}

func TestValidationResultAddError(t *testing.T) {
	res := govalid.NewValidationResult()

	res.AddError(*govalid.NewValidationError("name", "test error"))
	res.AddError(*govalid.NewValidationError("surname", "test error"))
	res.AddError(*govalid.NewValidationError("age", "test error"))

	assert.True(t, res.HasErrors())
	assert.Len(t, res.Errors(), 3)
}

func TestValidationResultNextError(t *testing.T) {
	type testCase struct {
		name          string
		expectedError *govalid.ValidationError
		testErrors    []*govalid.ValidationError
	}

	testCases := []testCase{
		{
			name:          "should return nil	if no errors",
			expectedError: nil,
			testErrors:    []*govalid.ValidationError{},
		},
		{
			name:          "should return the first error if called one time",
			expectedError: govalid.NewValidationError("name", "test error1"),
			testErrors: []*govalid.ValidationError{
				govalid.NewValidationError("name", "test error1"),
				govalid.NewValidationError("surname", "test error2"),
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := govalid.NewValidationResult()

			for _, err := range tc.testErrors {
				res.AddError(*err)
			}

			validationError := res.FirstError()
			assert.Equal(t, tc.expectedError, validationError)
		})
	}
}
