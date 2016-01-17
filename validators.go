package forms

import (
	"fmt"
	"regexp"
)

func patternMatched(pattern, value string) bool {
	matched, _ := regexp.MatchString(pattern, value)
	return matched
}

type Validator interface {
	IsValid(value string) (bool, string)
}

type Required struct{}

func (r *Required) IsValid(value string) (bool, string) {
	if len(value) > 0 {
		return true, ""
	}

	return false, translations["REQUIRED"]
}

type Regexp struct {
	Pattern string
}

func (r *Regexp) IsValid(value string) (bool, string) {
	msg := ""
	if r.Pattern == "" {
		return false, ""
	}
	// matched, err := regexp.MatchString(r.Pattern, value)
	// fmt.Println(r.Pattern, value, matched, err)
	// return matched
	m := patternMatched(r.Pattern, value)
	if !m {
		msg = fmt.Sprintf("Value doesn't match pattern \"%s\"", r.Pattern)
	}

	return m, msg
}

type Email struct{}

func (v *Email) IsValid(value string) (bool, string) {
	m := patternMatched(`(?i)\b[A-Z0-9._%+-]+@[A-Z0-9.-]+\.[A-Z]{2,}\b`, value)
	if m {
		return m, ""
	}

	return m, "Entered value is not correct email address"
}

type MinLength struct {
	Min int
}

func (v *MinLength) IsValid(value string) (bool, string) {
	if len(value) >= v.Min {
		return true, ""
	}

	return false, fmt.Sprintf(translations["INCORRECT_MIN_LENGTH"], v.Min)
}

type MaxLength struct {
	Max int
}

func (v *MaxLength) IsValid(value string) (bool, string) {
	if len(value) <= v.Max {
		return true, ""
	}

	return false, fmt.Sprintf(translations["INCORRECT_MAX_LENGTH"], v.Max)
}

func notMain() {
	min := &MinLength{Min: 6}
	emailValidator := &Email{}

	fmt.Println(min.IsValid("luelue"))
	fmt.Println(min.IsValid("valuesss"))
	fmt.Println(emailValidator.IsValid("value"))
	fmt.Println(emailValidator.IsValid("alkemic7@gmail.com"))
}
