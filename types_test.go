package forms

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type TypeTestsResults []struct {
	values []string
	result interface{}
}

type TypeTestsSet struct {
	name       string
	inputType  string
	multiValue bool
	results    TypeTestsResults
}

func executeTypeTests(t *testing.T, _t Type, results TypeTestsSet) {
	if results.multiValue {
		assert.True(
			t, _t.IsMultiValue(),
			fmt.Sprintf("%s should be multi value", results.name))
	} else {
		assert.False(
			t, _t.IsMultiValue(),
			fmt.Sprintf("%s shouldn't be multi value", results.name))
	}

	assert.Equal(
		t, _t.GetType(), results.inputType,
		fmt.Sprintf("Should return \"%s\"", results.inputType))

	for _, result := range results.results {
		assert.Equal(
			t, _t.CleanData(result.values), result.result, "Should equal")
	}
}

func TestTypeInput(t *testing.T) {
	var resultsSet = TypeTestsSet{
		name:       "Input",
		inputType:  "input",
		multiValue: false,

		results: TypeTestsResults{
			{[]string{"accdddaabbcce"}, "accdddaabbcce"},
			{[]string{"a", "b"}, "a"}, // we assume that all
			{[]string{""}, ""},
		},
	}

	executeTypeTests(t, &Input{}, resultsSet)
}

func TestTypeRadio(t *testing.T) {
	var resultsSet = TypeTestsSet{
		name:       "Radio",
		inputType:  "radio",
		multiValue: true,

		results: TypeTestsResults{
			{[]string{"accdddaabbcce"}, []string{"accdddaabbcce"}},
			{[]string{"a", "b"}, []string{"a", "b"}},
			{[]string{""}, []string{""}},
			{nil, []string(nil)},
		},
	}

	executeTypeTests(t, &Radio{}, resultsSet)
}

func TestTypeTextarea(t *testing.T) {
	var resultsSet = TypeTestsSet{
		name:       "Textarea",
		inputType:  "input",
		multiValue: false,

		results: TypeTestsResults{
			{[]string{"accdddaabbcce"}, "accdddaabbcce"},
			{[]string{"a", "b"}, "a"},
			{[]string{""}, ""},
			{nil, ""},
		},
	}

	executeTypeTests(t, &Textarea{}, resultsSet)
}
