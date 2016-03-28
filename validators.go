package forms

import (
	"fmt"
	"html"
	"regexp"
)

const (
	emailPattern = `(?i)\b[A-Z0-9._%+-]+@[A-Z0-9.-]+\.[A-Z]{2,}\b`
)

func patternMatched(pattern, value string) bool {
	if pattern == "" && value != "" {
		return false
	}
	matched, _ := regexp.MatchString(pattern, value)
	return matched
}

type checkFunc func(string) bool

func validate(fn checkFunc, values []string, msg string) (bool, []string) {
	result := true
	msgs := []string{}
	for _, value := range values {
		if value != "" && !fn(value) {
			result = false
			msgs = append(msgs, html.EscapeString(fmt.Sprintf(msg, value)))
		}
	}

	return result, msgs
}

// Validator is interface for all validators
type Validator interface {
	IsValid(values []string) (bool, []string)
}

// Required is validator that require some data input
type Required struct{}

// IsValid checks is entered data are correct
func (r *Required) IsValid(values []string) (bool, []string) {
	if len(values) > 0 && len(values[0]) > 0 {
		return true, []string{}
	}

	return false, []string{html.EscapeString(translations["REQUIRED"])}
}

// Regexp validator checks if given value match pattern
//     validator := &Regexp{"\d{4}.\d{2}.\d{2} \d{2}:\d{2}:\d{2}"}
type Regexp struct {
	Pattern string
}

// IsValid checks is entered data are correct
func (r *Regexp) IsValid(values []string) (bool, []string) {
	return validate(func(value string) bool {
		return patternMatched(r.Pattern, value)
	}, values, fmt.Sprintf(translations["NO_MATCH_PATTERN"], r.Pattern))
}

// Email validator checks if given value is proper email
type Email struct{}

// IsValid checks is entered data are correct
func (v *Email) IsValid(values []string) (bool, []string) {
	return validate(func(value string) bool {
		return patternMatched(emailPattern, value)
	}, values, translations["INCORRECT_EMAIL"])
}

// MinLength validator checks if given values length is under value
//     validator := &MaxLength{32}
type MinLength struct {
	Min int
}

// IsValid checks is entered data are correct
func (v *MinLength) IsValid(values []string) (bool, []string) {
	return validate(func(value string) bool {
		return len(value) >= v.Min
	}, values, fmt.Sprintf(translations["INCORRECT_MIN_LENGTH"], v.Min))
}

// MaxLength validator checks if given values length doesn't exceed given value
//     validator := &MaxLength{32}
type MaxLength struct {
	Max int
}

// IsValid checks is entered data are correct
func (v *MaxLength) IsValid(values []string) (bool, []string) {
	return validate(func(value string) bool {
		return len(value) <= v.Max
	}, values, fmt.Sprintf(translations["INCORRECT_MAX_LENGTH"], v.Max))
}

// InSlice validator checks if given value is in slice
//     validator := &MaxLength{[]string{"ham", "spam", "eggs"}]}
type InSlice struct {
	Values []string
}

// IsValid checks is entered data are correct
func (v *InSlice) IsValid(values []string) (bool, []string) {
	return validate(func(value string) bool {
		return valueInSlice(value, v.Values)
	}, values, translations["VALUE_NOT_FOUND"])
}
