package govalid_test

import (
	"testing"

	"github.com/Palma99/govalid"
	"github.com/Palma99/govalid/internal"
	"github.com/Palma99/govalid/validators"
	"github.com/stretchr/testify/assert"
)

func TestCompose(t *testing.T) {
	t.Run("should compose multiple validations", func(t *testing.T) {
		var v govalid.ValidationFunc = func() *internal.ValidationError {
			return nil
		}

		composed := govalid.ComposeShortCircuit(v, v, v)

		assert.Nil(t, composed())
	})

	t.Run("should return first error", func(t *testing.T) {
		var v1 govalid.ValidationFunc = func() *internal.ValidationError {
			return nil
		}
		var v2 govalid.ValidationFunc = func() *internal.ValidationError {
			return internal.NewValidationError("name", "test error1")
		}
		var v3 govalid.ValidationFunc = func() *internal.ValidationError {
			return internal.NewValidationError("name", "test error2")
		}

		composed := govalid.ComposeShortCircuit(v1, v2, v3)
		assert.NotNil(t, composed())
		assert.Equal(t, "name", composed().Field())
		assert.Equal(t, "test error1", composed().Message())
	})
}

func TestGroup(t *testing.T) {
	t.Run("should group multiple validations", func(t *testing.T) {
		grouped := govalid.GroupShortCircuit("name", "Mario",
			validators.NonEmptyRule(),
			validators.MaxLengthRule(10),
		)

		assert.Nil(t, grouped())
	})

	t.Run("should return first error if multiple validations", func(t *testing.T) {
		grouped := govalid.GroupShortCircuit("name", "",
			validators.NonEmptyRule(),
			validators.MinLengthRule(100),
		)

		assert.NotNil(t, grouped())
		assert.Equal(t, "name", grouped().Field())
		assert.Equal(t, "must not be empty", grouped().Message())
	})

	t.Run("should return second error if multiple validations", func(t *testing.T) {
		grouped := govalid.GroupShortCircuit("name", "Mario",
			validators.NonEmptyRule(),
			validators.MinLengthRule(100),
		)

		assert.NotNil(t, grouped())
		assert.Equal(t, "name", grouped().Field())
		assert.Equal(t, "must be at least 100 characters", grouped().Message())
	})
}
