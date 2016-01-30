package forms

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"html"
	"testing"
)

type ValidatorResults []struct {
	validator Validator
	test      []string
	result    bool
	message   []string
}

type ValidatorTestsSet struct {
	name    string
	results ValidatorResults
}

func prepareMsg(msg string, v1, v2 interface{}) string {
	return fmt.Sprintf(fmt.Sprintf(msg, v1), v2)
}

func makeSafe(msgs []string) []string {
	safeMsgs := []string{}
	for _, msg := range msgs {
		safeMsgs = append(safeMsgs, html.EscapeString(msg))
	}

	return safeMsgs
}

func executeValidatorTests(t *testing.T, results ValidatorTestsSet) {
	for _, result := range results.results {
		r, msgs := result.validator.IsValid(result.test)

		assert.Equal(t, msgs, makeSafe(result.message), "Incorrect message for \"%s\"", result.validator)
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
			{&Required{}, []string{"asd"}, true, []string{}},
			{&Required{}, []string{}, false, []string{translations["REQUIRED"]}},
		},
	}

	executeValidatorTests(t, results)
}

func TestRegexpValidator(t *testing.T) {
	var results = ValidatorTestsSet{
		name: "Regexp",
		results: ValidatorResults{
			{&Regexp{}, []string{"asd"}, false, []string{prepareMsg(translations["NO_MATCH_PATTERN"], "", "asd")}},
			{&Regexp{""}, []string{"asd"}, false, []string{prepareMsg(translations["NO_MATCH_PATTERN"], "", "asd")}},
			{&Regexp{""}, []string{""}, true, []string{}},
			{&Regexp{"^[a-d]*$"}, []string{"accdddaabbcc"}, true, []string{}},
			{&Regexp{"^[a-d]*$"}, []string{"accdddaabbcce"}, false, []string{prepareMsg(translations["NO_MATCH_PATTERN"], "^[a-d]*$", "accdddaabbcce")}},
			{&Regexp{"[0-9]*"}, []string{"accddd123aabbcce"}, true, []string{}},
			{&Regexp{"^[0-9]*$"}, []string{"accddd123aabbcce"}, false, []string{prepareMsg(translations["NO_MATCH_PATTERN"], "^[0-9]*$", "accddd123aabbcce")}},
		},
	}

	executeValidatorTests(t, results)
}

func TestEmailValidator(t *testing.T) {
	var results = ValidatorTestsSet{
		name: "Email",
		results: ValidatorResults{
			{&Email{}, []string{}, true, []string{}},
			{&Email{}, []string{"foo"}, false, []string{fmt.Sprintf(translations["INCORRECT_EMAIL"], "foo")}},
			{&Email{}, []string{"foo@ham"}, false, []string{fmt.Sprintf(translations["INCORRECT_EMAIL"], "foo@ham")}},
			{&Email{}, []string{"foo@ham.p"}, false, []string{fmt.Sprintf(translations["INCORRECT_EMAIL"], "foo@ham.p")}},
			{&Email{}, []string{"foo@ham.pl"}, true, []string{}},
			{&Email{}, []string{"foo+bar@ham"}, false, []string{fmt.Sprintf(translations["INCORRECT_EMAIL"], "foo+bar@ham")}},
			{&Email{}, []string{"foo+bar@ham.pl"}, true, []string{}},

			{&Email{}, []string{"foo@h_am.pl"}, false, []string{fmt.Sprintf(translations["INCORRECT_EMAIL"], "foo@h_am.pl")}},
			{&Email{}, []string{"foo+bar@h_am.pl"}, false, []string{fmt.Sprintf(translations["INCORRECT_EMAIL"], "foo+bar@h_am.pl")}},
		},
	}

	executeValidatorTests(t, results)
}

func TestMinLengthValidator(t *testing.T) {
	var results = ValidatorTestsSet{
		name: "MinLength",
		results: ValidatorResults{
			{&MinLength{Min: 2}, []string{}, true, []string{}},
			{&MinLength{Min: 2}, []string{""}, false, []string{prepareMsg(translations["INCORRECT_MIN_LENGTH"], 2, "")}},
			{&MinLength{Min: 2}, []string{"foo"}, true, []string{}},
			{&MinLength{Min: 2}, []string{"foo", "a"}, false, []string{prepareMsg(translations["INCORRECT_MIN_LENGTH"], 2, "a")}},
			{&MinLength{Min: 3}, []string{"foo"}, true, []string{}},
			{&MinLength{Min: 4}, []string{"foo"}, false, []string{prepareMsg(translations["INCORRECT_MIN_LENGTH"], 4, "foo")}},
		},
	}

	executeValidatorTests(t, results)
}

func TestMaxLengthValidator(t *testing.T) {
	var results = ValidatorTestsSet{
		name: "MaxLength",
		results: ValidatorResults{
			{&MaxLength{Max: 2}, []string{}, true, []string{}},
			{&MaxLength{Max: 2}, []string{""}, true, []string{}},
			{&MaxLength{Max: 5}, []string{"foo", "asdasdd"}, false, []string{prepareMsg(translations["INCORRECT_MAX_LENGTH"], 5, "asdasdd")}},
			{&MaxLength{Max: 2}, []string{"foo"}, false, []string{prepareMsg(translations["INCORRECT_MAX_LENGTH"], 2, "foo")}},
			{&MaxLength{Max: 3}, []string{"foo"}, true, []string{}},
			{&MaxLength{Max: 4}, []string{"foo"}, true, []string{}},
		},
	}

	executeValidatorTests(t, results)
}

func TestInSliceValidator(t *testing.T) {
	testSlice := []string{"foo", "bar", "spam", "ham", "eggs"}
	var results = ValidatorTestsSet{
		name: "InSlice",
		results: ValidatorResults{
			{&InSlice{Values: []string{""}}, []string{"foo"}, false,
				[]string{fmt.Sprintf(translations["VALUE_NOT_FOUND"], "foo")}},
			{&InSlice{Values: []string{""}}, []string{""}, true, []string{}},
			{&InSlice{Values: []string{}}, []string{""}, false,
				[]string{fmt.Sprintf(translations["VALUE_NOT_FOUND"], "")}},
			{&InSlice{Values: testSlice}, []string{"spam"}, true, []string{}},
			{&InSlice{Values: testSlice}, []string{"spa", "asd"}, false,
				[]string{fmt.Sprintf(translations["VALUE_NOT_FOUND"], "spa"),
					fmt.Sprintf(translations["VALUE_NOT_FOUND"], "asd")}},
		},
	}

	executeValidatorTests(t, results)
}
