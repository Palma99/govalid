package govalid_test

import (
	"testing"

	"github.com/Palma99/govalid"
	"github.com/Palma99/govalid/internal"
	"github.com/Palma99/govalid/validators"
	"github.com/stretchr/testify/assert"
)

func createValidatorSpy(counter *int, testErr *internal.ValidationError) govalid.ValidationFunc {
	return func() *internal.ValidationError {
		*counter += 1
		if testErr != nil {
			return testErr
		}
		return nil
	}
}

func TestValidateShouldCallAllValidators(t *testing.T) {
	callCount := 0

	govalid.Validate(
		createValidatorSpy(&callCount, nil),
		createValidatorSpy(&callCount, nil),
		createValidatorSpy(&callCount, nil),
	)

	assert.Equal(t, 3, callCount)
}

func TestValidateShouldReturnAllErrors(t *testing.T) {
	callCount := 0

	res := govalid.Validate(
		createValidatorSpy(&callCount, internal.NewValidationError("field1", "test error1")),
		createValidatorSpy(&callCount, internal.NewValidationError("field2", "test error2")),
		createValidatorSpy(&callCount, nil),
		createValidatorSpy(&callCount, internal.NewValidationError("field3", "test error3")),
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

func TestValidateShortCircuitShouldReturnFirstError(t *testing.T) {
	callCount := 0
	res := govalid.ValidateShortCircuit(
		createValidatorSpy(&callCount, internal.NewValidationError("field1", "test error1")),
		createValidatorSpy(&callCount, internal.NewValidationError("field2", "test error2")),
		createValidatorSpy(&callCount, nil),
		createValidatorSpy(&callCount, internal.NewValidationError("field3", "test error3")),
	)

	assert.True(t, res.HasErrors())
	assert.Equal(t, 1, callCount)

	assert.Equal(t, "field1", res.FirstError().Field())
	assert.Equal(t, "test error1", res.FirstError().Message())
}

func TestValidateShouldWorkWithCustomValidator(t *testing.T) {
	t.Run("should work with custom validator", func(t *testing.T) {
		res := govalid.Validate(
			(func(field, value string) govalid.ValidationFunc {
				return func() *internal.ValidationError {
					if field == value {
						return nil
					}
					return internal.NewValidationError(field, "must be equal to "+value)
				}
			})("field1", "test value"),

			(func(field, value string) govalid.ValidationFunc {
				return func() *internal.ValidationError {
					if value == "value" {
						return nil
					}
					return internal.NewValidationError(field, "must be equal to value")
				}
			})("field2", "value"),
		)

		assert.True(t, res.HasErrors())
		assert.Len(t, res.Errors(), 1)
		assert.Equal(t, "field1", res.FirstError().Field())
	})
}

func TestValidateShouldReturnNoErrors(t *testing.T) {
	callCount := 0

	res := govalid.Validate(
		createValidatorSpy(&callCount, nil),
		createValidatorSpy(&callCount, nil),
	)

	assert.Equal(t, 2, callCount)
	assert.False(t, res.HasErrors())
	assert.Len(t, res.Errors(), 0)
}

func TestValidateShouldWorkWithCompose(t *testing.T) {
	t.Run("should work with compose", func(t *testing.T) {
		composed1 := govalid.ComposeShortCircuit(
			validators.NonEmpty("name", ""),
			validators.NonEmpty("surname", ""),
		)

		composed2 := govalid.ComposeShortCircuit(
			validators.NonEmpty("age", ""),
			validators.NonEmpty("city", ""),
		)

		res := govalid.Validate(
			composed1,
			composed2,
		)

		assert.True(t, res.HasErrors())
		assert.Len(t, res.Errors(), 2)
		assert.Equal(t, "name", res.FirstError().Field())
		assert.Equal(t, "age", res.Errors()[1].Field())
	})
}

func TestValidateShouldWorkWithComposeAndGroup(t *testing.T) {
	t.Run("should work with compose and group", func(t *testing.T) {
		person := struct {
			Name string
			Age  int
		}{
			Name: "Mario",
			Age:  18,
		}

		nameValidator := govalid.Group("name", person.Name,
			validators.NonEmptyRule(),
			validators.MinLengthRule(3),
		)

		ageValidator := govalid.Group("age", person.Age,
			validators.MinRule(14),
		)

		personValidator := govalid.Compose(
			nameValidator,
			ageValidator,
		)

		res := govalid.Validate(
			personValidator,
		)

		assert.False(t, res.HasErrors())
	})

	t.Run("should work with compose and group with custom error message", func(t *testing.T) {
		person := struct {
			Name string
			Age  int
		}{
			Name: "Mario",
			Age:  12,
		}

		nameValidator := govalid.GroupShortCircuit("name", person.Name,
			validators.NonEmptyRule(),
			validators.MinLengthRule(3),
		)

		ageValidator := govalid.GroupShortCircuit("age", person.Age,
			validators.MinRule(14, "you are too young!"),
		)

		personValidator := govalid.ComposeShortCircuit(
			nameValidator,
			ageValidator,
		)

		res := govalid.Validate(
			personValidator,
		)

		assert.True(t, res.HasErrors())
		assert.Len(t, res.Errors(), 1)
		assert.Equal(t, "age", res.FirstError().Field())
		assert.Equal(t, "you are too young!", res.FirstError().Message())
	})
}

func TestValidateWithComposeAll(t *testing.T) {
	t.Run("validate with ComposeAll should return all errors", func(t *testing.T) {

		compose1 := govalid.Compose(
			validators.NonEmpty("name", ""),
			validators.NonEmpty("surname", ""),
		)

		compose2 := govalid.Compose(
			validators.NonEmpty("age", ""),
			validators.NonEmpty("city", ""),
		)

		emailGroup := govalid.Group("email", "",
			validators.NonEmptyRule(),
			validators.IsEmailRule(),
		)

		res := govalid.Validate(
			compose1, compose2, emailGroup,
		)

		assert.True(t, res.HasErrors())
		assert.Len(t, res.Errors(), 6)
	})

	t.Run("validate with ComposeAll Compose, Group, GroupAll should return all errors", func(t *testing.T) {

		compose := govalid.Compose(
			validators.NonEmpty("name", ""),
			validators.NonEmpty("surname", ""),
		)

		composeAll := govalid.ComposeShortCircuit(
			validators.NonEmpty("age", ""),
			validators.NonEmpty("city", ""),
			validators.NonEmpty("address", ""),
		)

		group := govalid.Group("email", "",
			validators.NonEmptyRule(),
			validators.IsEmailRule(),
		)

		groupAll := govalid.GroupShortCircuit("phone", "",
			validators.NonEmptyRule(),
		)

		res := govalid.Validate(
			compose, composeAll, group, groupAll,
		)

		assert.True(t, res.HasErrors())
		assert.Len(t, res.Errors(), 6)
	})

	t.Run("should throw error if validator is not of type govalid.Validator", func(t *testing.T) {
		assert.Panics(t, func() {
			govalid.Validate(1)
		})
	})
}

func TestValidateAndFailFast(t *testing.T) {
	t.Run("should work with compose and group", func(t *testing.T) {
		person := struct {
			Name string
			Age  int
		}{
			Name: "Mario",
			Age:  18,
		}

		nameValidator := govalid.GroupShortCircuit("name", person.Name,
			validators.NonEmptyRule(),
			validators.MinLengthRule(3),
		)

		ageValidator := govalid.GroupShortCircuit("age", person.Age,
			validators.MinRule(14),
		)

		personValidator := govalid.ComposeShortCircuit(
			nameValidator,
			ageValidator,
		)

		res := govalid.ValidateShortCircuit(
			personValidator,
		)

		assert.False(t, res.HasErrors())
	})

	t.Run("should return only the first error", func(t *testing.T) {
		composed1 := govalid.ComposeShortCircuit(
			validators.NonEmpty("name", ""),
			validators.NonEmpty("surname", ""),
		)

		composed2 := govalid.ComposeShortCircuit(
			validators.NonEmpty("age", ""),
			validators.NonEmpty("city", ""),
		)

		res := govalid.ValidateShortCircuit(
			composed1, composed2,
		)

		assert.True(t, res.HasErrors())
		assert.Len(t, res.Errors(), 1)
		assert.Equal(t, "name", res.FirstError().Field())
	})

	t.Run("Should panic if validator is not of type govalid.Validator", func(t *testing.T) {
		assert.Panics(t, func() {
			govalid.ValidateShortCircuit(1)
		})
	})

}
