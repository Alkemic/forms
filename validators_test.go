package forms

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type ValidatorResults []struct {
	validator Validator
	test      string
	result    bool
}

type ValidatorTestsSet struct {
	name    string
	results ValidatorResults
}

func executeValidatorTests(t *testing.T, results ValidatorTestsSet) {
	for _, result := range results.results {
		r, _ := result.validator.IsValid([]string{result.test})
		if result.result {
			assert.True(
				t, r, fmt.Sprintf(
					"%s validator for %s should pass",
					results.name, result.validator))
		} else {
			assert.False(
				t, r, fmt.Sprintf(
					"%s validator for %s shouldn't pass",
					results.name, result.validator))
		}
	}
}

func TestRequiredValidator(t *testing.T) {
	var results = ValidatorTestsSet{
		name: "Required",
		results: ValidatorResults{
			{&Required{}, "asd", true},
			{&Required{}, "", false},
		},
	}

	executeValidatorTests(t, results)
}

func TestRegexpValidator(t *testing.T) {
	var results = ValidatorTestsSet{
		name: "Regexp",
		results: ValidatorResults{
			{&Regexp{}, "asd", false},
			{&Regexp{""}, "asd", false},
			{&Regexp{"^[a-d]*$"}, "accdddaabbcc", true},
			{&Regexp{"^[a-d]*$"}, "accdddaabbcce", false},
			{&Regexp{"[0-9]*"}, "accddd123aabbcce", true},
			{&Regexp{"^[0-9]*$"}, "accddd123aabbcce", false},
		},
	}

	executeValidatorTests(t, results)
}

func TestEmailValidator(t *testing.T) {
	var results = ValidatorTestsSet{
		name: "Email",
		results: ValidatorResults{
			{&Email{}, "foo", false},
			{&Email{}, "foo@ham", false},
			{&Email{}, "foo@ham.p", false},
			{&Email{}, "foo@ham.pl", true},
			{&Email{}, "foo+bar@ham", false},
			{&Email{}, "foo+bar@ham.pl", true},

			{&Email{}, "foo@h_am.pl", false},
			{&Email{}, "foo+bar@h_am.pl", false},
		},
	}

	executeValidatorTests(t, results)
}

func TestMinLengthValidator(t *testing.T) {
	var results = ValidatorTestsSet{
		name: "MinLength",
		results: ValidatorResults{
			{&MinLength{Min: 2}, "foo", true},
			{&MinLength{Min: 3}, "foo", true},
			{&MinLength{Min: 4}, "foo", false},
		},
	}

	executeValidatorTests(t, results)
}

func TestMaxLengthValidator(t *testing.T) {
	var results = ValidatorTestsSet{
		name: "MaxLength",
		results: ValidatorResults{
			{&MaxLength{Max: 2}, "foo", false},
			{&MaxLength{Max: 3}, "foo", true},
			{&MaxLength{Max: 4}, "foo", true},
		},
	}

	executeValidatorTests(t, results)
}
