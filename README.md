![Go Test](https://github.com/Palma99/govalid/actions/workflows/test.yml/badge.svg)

# govalid

**govalid** is a lightweight, extensible validation library for Go that helps you validate fields and data structures in a declarative, composable way. It supports both full validation and short-circuit (fail-fast) strategies.

## Features

- Simple, readable API
- Compose complex validation logic with ease
- Fail-fast or full error collection modes
- Built-in validators for common use cases
- Group-based field validation
- Custom validation rules support

## Installation

```bash
go get github.com/Palma99/govalid
```

## Example
```go
import "github.com/Palma99/govalid"

person := struct {
	Name    string
	Surname string
	Email   string
}{
	Name:    "",
	Surname: "",
	Email:   "invalid-email",
}

// Compose validations to gather all errors
composed := govalid.Compose(
	govalid.NonEmpty("name", person.Name),
	govalid.NonEmpty("surname", person.Surname),
)

// Group validations for a field
emailGroup := govalid.Group("email", person.Email,
	govalid.NonEmptyRule(),
	govalid.IsEmailRule(),
)

result := govalid.Validate(composed, emailGroup)

if result.HasErrors() {
	fmt.Println(result.Errors())
}
```


## API
### Validation Modes
`Validate(...)`

Runs all validations and returns a list of all errors.

```go
result := govalid.Validate(validation1, validation2)
```

`ValidateShortCircuit(...)`

Stops at the first error encountered.

```go
result := govalid.ValidateShortCircuit(validation1, validation2)
```

### Composition Helpers

`Compose(...)`

Combines multiple validation functions. All will be evaluated.

```go
composed := govalid.Compose(
	govalid.NonEmpty("username", input.Username),
	govalid.NonEmpty("password", input.Password),
)
```

`ComposeShortCircuit(...)`

Combines multiple validation functions. All will be evaluated.

```go
composed := govalid.ComposeShortCircuit(
	govalid.NonEmpty("username", ""),
	govalid.NonEmpty("password", ""),
)
```

### Group validation

`Group(field, value, rules...)`

Applies multiple `ValidationRules` to a field. Returns all errors.

```go
group := govalid.Group("email", user.Email,
	govalid.NonEmptyRule(),
	govalid.IsEmailRule(),
)
```

`GroupShortCircuit(field, value, rules...)`

Stops at the first rule that fails.

```go
group := govalid.GroupShortCircuit("email", user.Email,
	govalid.NonEmptyRule(),
	govalid.IsEmailRule(),
)
```

## Built-in Validators

### Validators

`NonEmpty(field, value)`

Validates that a field is not empty. Works with string, []any, and map[any]any.

`Min[T internal.Number](fieldName string, value any, min T, args ...string)`

Check if a number is greater than a defined value.

`Max[T internal.Number](fieldName string, value any, max T, args ...string)`

Check if a number is greater than a defined value.

`MinLength(fieldName string, value any, min int, args ...string)` and `MaxLength(fieldName string, value any, max int, args ...string)`

Works with strings, arrays and maps, and perform a check on the length of the value 

`MatchesRegex(fieldName, value, pattern string, args ...string)`

Check if value matches the defined pattern

`IsEmail(fieldName, value string, args ...string)`

Apply a simple regex for validating an email

### Rules

Convenient set of rules to use with `Group()` 

- `NonEmptyRule(...customMessage)`
- `MinLengthRule(min, ...customMessage)`
- `MaxLengthRule(max, ...customMessage)`
- `MinRule(min, ...customMessage)`
- `MaxRule(max, ...customMessage)`
- `MatchesRegexRule(pattern, ...customMessage)`
- `IsEmailRule(...customMessage)`


## Custom Validator

You can define your own validation logic using `CustomValidator`

```go
personIsValid := validators.CustomValidator(
  func(value Person) *string {
    if value.Name == "John" && value.Surname == "Doe" && value.Age == 30 {
      return nil
    }
    message := "must be John Doe and 30 years old"
    return &message
  },
)

res := govalid.Validate(
  personIsValid("person", person),
)
```

For Group usage, the `CustomRule` function is available

```go
customValidator := validators.CustomValidator(...)

customRule := validators.CustomRule[int](customValidator)

govalid.Group("field", value, customRule)
```