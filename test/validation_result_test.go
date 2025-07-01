package govalid_test

import (
	"testing"

	"github.com/Palma99/govalid"
	"github.com/Palma99/govalid/internal"
	"github.com/stretchr/testify/assert"
)

func TestValidationResultConstructorWithoutErrors(t *testing.T) {
	res := govalid.NewValidationResult()

	assert.False(t, res.HasErrors())
	assert.Len(t, res.Errors(), 0)
}

func TestValidationResultConstructorWithErrors(t *testing.T) {
	res := govalid.NewValidationResult(
		*internal.NewValidationError("name", "test error"),
		*internal.NewValidationError("surname", "test error"),
		*internal.NewValidationError("age", "test error"),
	)

	assert.True(t, res.HasErrors())
	assert.Len(t, res.Errors(), 3)
}

func TestValidationResultAddError(t *testing.T) {
	res := govalid.NewValidationResult(
		*internal.NewValidationError("name", "test error"),
		*internal.NewValidationError("surname", "test error"),
		*internal.NewValidationError("age", "test error"),
	)

	assert.True(t, res.HasErrors())
	assert.Len(t, res.Errors(), 3)
}

func TestValidationResultNextError(t *testing.T) {
	type testCase struct {
		name          string
		expectedError *internal.ValidationError
		testErrors    []*internal.ValidationError
	}

	testCases := []testCase{
		{
			name:          "should return nil	if no errors",
			expectedError: nil,
			testErrors:    []*internal.ValidationError{},
		},
		{
			name:          "should return the first error if called one time",
			expectedError: internal.NewValidationError("name", "test error1"),
			testErrors: []*internal.ValidationError{
				internal.NewValidationError("name", "test error1"),
				internal.NewValidationError("surname", "test error2"),
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			errors := []internal.ValidationError{}
			for _, err := range tc.testErrors {
				errors = append(errors, *err)
			}

			res := govalid.NewValidationResult(errors...)

			validationError := res.FirstError()
			assert.Equal(t, tc.expectedError, validationError)
		})
	}
}

func TestFieldErrors(t *testing.T) {
	t.Run("should return errors for field", func(t *testing.T) {

		res := govalid.NewValidationResult(
			*internal.NewValidationError("name", "test error name"),
			*internal.NewValidationError("surname", "test error surname"),
			*internal.NewValidationError("age", "test error age"),
		)

		assert.Len(t, res.FieldErrors("name"), 1)
		assert.Equal(t, "test error name", res.FieldErrors("name")[0].Message())

		assert.Len(t, res.FieldErrors("surname"), 1)
		assert.Equal(t, "test error surname", res.FieldErrors("surname")[0].Message())

		assert.Len(t, res.FieldErrors("age"), 1)
		assert.Equal(t, "test error age", res.FieldErrors("age")[0].Message())
	})

	t.Run("should return all errors for field", func(t *testing.T) {

		res := govalid.NewValidationResult(
			*internal.NewValidationError("name", "test error name"),
			*internal.NewValidationError("name", "test error name2"),
			*internal.NewValidationError("name", "test error name3"),
		)

		assert.Len(t, res.FieldErrors("name"), 3)
		assert.Equal(t, "test error name", res.FieldErrors("name")[0].Message())
		assert.Equal(t, "test error name2", res.FieldErrors("name")[1].Message())
		assert.Equal(t, "test error name3", res.FieldErrors("name")[2].Message())
	})

	t.Run("should return empty array if field has no errors", func(t *testing.T) {
		res := govalid.NewValidationResult(
			*internal.NewValidationError("name", "test error name"),
		)

		assert.Len(t, res.FieldErrors("surname"), 0)
	})
}

func TestIsFieldValid(t *testing.T) {
	t.Run("should return true if field has no errors", func(t *testing.T) {
		res := govalid.NewValidationResult(
			*internal.NewValidationError("name", "test error name"),
		)

		assert.False(t, res.IsFieldValid("name"))
		assert.True(t, res.IsFieldValid("surname"))
	})
}
