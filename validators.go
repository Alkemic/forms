package forms

import (
	"fmt"
	"html"
	"regexp"
)

const (
	EMAIL_PATTERN = `(?i)\b[A-Z0-9._%+-]+@[A-Z0-9.-]+\.[A-Z]{2,}\b`
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
		if !fn(value) {
			result = false
			msgs = append(msgs, html.EscapeString(fmt.Sprintf(msg, value)))
		}
	}

	return result, msgs
}

type Validator interface {
	IsValid(values []string) (bool, []string)
}

type Required struct{}

func (r *Required) IsValid(values []string) (bool, []string) {
	if len(values) > 0 && len(values[0]) > 0 {
		return true, []string{}
	}

	return false, []string{html.EscapeString(translations["REQUIRED"])}
}

type Regexp struct {
	Pattern string
}

func (r *Regexp) IsValid(values []string) (bool, []string) {
	return validate(func(value string) bool {
		return patternMatched(r.Pattern, value)
	}, values, fmt.Sprintf(translations["NO_MATCH_PATTERN"], r.Pattern))
}

type Email struct{}

func (v *Email) IsValid(values []string) (bool, []string) {
	return validate(func(value string) bool {
		return patternMatched(EMAIL_PATTERN, value)
	}, values, translations["INCORRECT_EMAIL"])
}

type MinLength struct {
	Min int
}

func (v *MinLength) IsValid(values []string) (bool, []string) {
	return validate(func(value string) bool {
		return len(value) >= v.Min
	}, values, fmt.Sprintf(translations["INCORRECT_MIN_LENGTH"], v.Min))
}

type MaxLength struct {
	Max int
}

func (v *MaxLength) IsValid(values []string) (bool, []string) {
	return validate(func(value string) bool {
		return len(value) <= v.Max
	}, values, fmt.Sprintf(translations["INCORRECT_MAX_LENGTH"], v.Max))
}

type InSlice struct {
	Values []string
}

func (v *InSlice) IsValid(values []string) (bool, []string) {
	return validate(func(value string) bool {
		return ValueInSlice(value, v.Values)
	}, values, translations["VALUE_NOT_FOUND"])
}
