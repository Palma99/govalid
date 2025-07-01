package main

import (
	"fmt"

	"github.com/Palma99/govalid"
	"github.com/Palma99/govalid/validators"
)

func personValidation() {
	fmt.Println("--- Person Validation ---")

	type Person struct {
		Name    string
		Surname string
		Age     int
	}

	person := Person{
		Name:    "aa",
		Surname: "bb",
		Age:     3,
	}

	nameValidator := govalid.Group("name", person.Name,
		validators.NonEmptyRule(),
		validators.MinLengthRule(3),
	)

	surnameValidator := govalid.Group("surname", person.Surname,
		validators.MaxLengthRule(10),
	)

	ageValidator := govalid.Group("age", person.Age,
		validators.MinRule(14),
	)

	personValidator := govalid.Compose(
		nameValidator,
		surnameValidator,
		ageValidator,
	)

	res := govalid.ValidateShortCircuit(
		personValidator,
	)

	fmt.Printf("Is valid: %t\n", res.IsValid())
	fmt.Printf("Is surname valid: %t\n", res.IsFieldValid("surname"))
	fmt.Printf("Errors: %v\n", res.Errors())
	fmt.Printf("Grouped errors by fields: %+v\n", res.GroupedErrorsByField())
}

func composeOfComposed() {
	fmt.Println("\n--- Compose of composed ---")

	compose1 := govalid.Compose(
		validators.NonEmpty("name", ""),
		validators.NonEmpty("surname", ""),
	)

	compose2 := govalid.Compose(
		compose1,
		validators.NonEmpty("age", ""),
		validators.NonEmpty("city", ""),
	)

	compose3 := govalid.Compose(
		compose2,
		govalid.Group("email", "",
			validators.NonEmptyRule(),
			validators.IsEmailRule(),
		),
	)

	res := govalid.Validate(
		compose3,
	)

	fmt.Printf("Is valid: %t\n", res.IsValid())
	fmt.Printf("Errors count: %v\n", res.ErrorCount())
	fmt.Printf("Grouped errors by fields: %+v\n", res.GroupedErrorsByField())
}

func customValidators() {
	fmt.Println("\n--- Custom validators ---")

	type Person struct {
		Name    string
		Surname string
		Age     int
	}

	personIsValid := validators.CustomValidator(
		func(value Person) *string {
			if value.Name == "John" && value.Surname == "Doe" && value.Age == 30 {
				return nil
			}
			message := "must be John Doe and 30 years old"
			return &message
		},
	)

	person := Person{
		Name:    "John",
		Surname: "Doe",
		Age:     31,
	}

	res := govalid.Validate(
		personIsValid("person", person),
	)

	fmt.Printf("Is valid: %t\n", res.IsValid())
	fmt.Printf("Errors count: %v\n", res.ErrorCount())
	fmt.Printf("Grouped errors by fields: %+v\n", res.GroupedErrorsByField())

	fmt.Println("\n--- Custom number between validator, with custom message ---")
	numberBetween := func(min, max int) validators.Validator {
		return validators.CustomValidator(
			func(value int) *string {
				if value >= min && value <= max {
					return nil
				}
				message := fmt.Sprintf("must be between %d and %d, received: %d", min, max, value)
				return &message
			},
		)
	}

	res = govalid.Validate(
		numberBetween(1, 10)("number1", -1),
		numberBetween(1, 10)("number2", 11, "not between 1 and 10"),
	)

	fmt.Println(res.Errors())

	fmt.Println("\n--- Custom rule ---")

	customRule := validators.CustomRule[int](numberBetween(4, 15))

	res = govalid.Validate(
		govalid.Group("age", 33,
			customRule(),
		),
	)

	fmt.Println(res.Errors())
}

func main() {
	personValidation()
	composeOfComposed()

	customValidators()
}
