package forms

import (
	"fmt"
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

type Validator interface {
	IsValid(values []string) (bool, []string)
}

type Required struct{}

func (r *Required) IsValid(values []string) (bool, []string) {
	if len(values) > 0 && len(values[0]) > 0 {
		return true, []string{}
	}

	return false, []string{translations["REQUIRED"]}
}

type Regexp struct {
	Pattern string
}

func (r *Regexp) IsValid(values []string) (bool, []string) {
	result := true
	msgs := []string{}
	for _, value := range values {
		m := patternMatched(r.Pattern, value)
		if !m {
			result = false
			msgs = append(msgs, fmt.Sprintf(translations["NO_MATCH_PATTERN"], r.Pattern))
		}
	}

	return result, msgs
}

type Email struct{}

func (v *Email) IsValid(values []string) (bool, []string) {
	result := true
	msgs := []string{}
	for _, value := range values {
		m := patternMatched(EMAIL_PATTERN, value)
		if !m {
			result = false
			msgs = append(msgs, translations["INCORRECT_EMAIL"])
		}
	}

	return result, msgs
}

type MinLength struct {
	Min int
}

func (v *MinLength) IsValid(values []string) (bool, []string) {
	result := true
	msgs := []string{}
	for _, value := range values {
		m := len(value) >= v.Min
		if !m {
			result = false
			msgs = append(msgs, fmt.Sprintf(translations["INCORRECT_MIN_LENGTH"], v.Min))
		}
	}

	return result, msgs
}

type MaxLength struct {
	Max int
}

func (v *MaxLength) IsValid(values []string) (bool, []string) {
	result := true
	msgs := []string{}
	for _, value := range values {
		if !(len(value) <= v.Max) {
			result = false
			msgs = append(msgs, fmt.Sprintf(translations["INCORRECT_MAX_LENGTH"], v.Max))
		}
	}

	return result, msgs
}

type InSlice struct {
	Values []string
}

func (v *InSlice) IsValid(values []string) (bool, []string) {
	result := true
	msgs := []string{}
	for _, value := range values {
		if !ValueInSlice(value, v.Values) {
			result = false
			msgs = append(msgs, fmt.Sprintf(translations["VALUE_NOT_FOUND"], value))
		}
	}

	return result, msgs

	// for _, value := range values {
	// 	if !ValueInSlice(value, v.Values) {
	// 		return false, []string{fmt.Sprintf(translations["VALUE_NOT_FOUND"], value)}
	// 	}
	// }

	return true, []string{}
}
