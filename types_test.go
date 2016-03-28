package forms

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var renderdParts, expectedParts []string

type TypeTestsResults []struct {
	values []string
	result interface{}
}

type TypeTestsSet struct {
	name       string
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

	for _, result := range results.results {
		assert.Equal(
			t, _t.CleanData(result.values), result.result, "Should equal")
	}
}

func TestTypeInput(t *testing.T) {
	var resultsSet = TypeTestsSet{
		name:       "Input",
		multiValue: false,

		results: TypeTestsResults{
			{[]string{"accdddaabbcce"}, "accdddaabbcce"},
			{[]string{"a", "b"}, "a"}, // we assume that all
			{[]string{""}, ""},
		},
	}

	_t := &Input{}
	executeTypeTests(t, _t, resultsSet)

	f := &Field{Name: "test1", Type: _t}
	assert.Equal(t, _t.Render(f, nil, []string{}), "<input name=\"test1\" type=\"input\" id=\"f_test1\" />")
	assert.Equal(t, _t.Render(f, nil, []string{""}), "<input name=\"test1\" type=\"input\" id=\"f_test1\" />")
	assert.Equal(t, _t.Render(f, nil, []string{"accdddaabbcce"}), "<input name=\"test1\" type=\"input\" id=\"f_test1\" value=\"accdddaabbcce\" />")
}

func TestTypeRadio(t *testing.T) {
	var resultsSet = TypeTestsSet{
		name:       "Radio",
		multiValue: true,

		results: TypeTestsResults{
			{[]string{"accdddaabbcce"}, []string{"accdddaabbcce"}},
			{[]string{"a", "b"}, []string{"a", "b"}},
			{[]string{""}, []string{""}},
			{nil, []string(nil)},
		},
	}

	choices := []Choice{
		{"pizza", "Pizza"},
		{"pasta", "Makaron"},
		{"risotto", "Risotto"},
	}

	_t := &Radio{}
	f := &Field{Name: "test1", Type: _t}
	executeTypeTests(t, _t, resultsSet)

	r := strings.Split(_t.Render(f, choices, []string{}), "\n")

	assert.Len(t, r, 4)

	assert.Len(t, r[0], 109)
	assert.True(t, strings.HasPrefix(r[0], "<label for=\"c_test1_pizza\"><input "))
	assert.Contains(t, r[0], " name=\"test1\" ")
	assert.Contains(t, r[0], " type=\"radio\" ")
	assert.Contains(t, r[0], " id=\"c_test1_pizza\" ")
	assert.Contains(t, r[0], " value=\"pizza\" ")
	assert.True(t, strings.HasSuffix(r[0], " /> Pizza</label>"))
	assert.NotContains(t, r[0], " checked=\"checked\" ")

	assert.Len(t, r[1], 111)
	assert.True(t, strings.HasPrefix(r[1], "<label for=\"c_test1_pasta\"><input "))
	assert.Contains(t, r[1], " name=\"test1\" ")
	assert.Contains(t, r[1], " type=\"radio\" ")
	assert.Contains(t, r[1], " id=\"c_test1_pasta\" ")
	assert.Contains(t, r[1], " value=\"pasta\" ")
	assert.True(t, strings.HasSuffix(r[1], " /> Makaron</label>"))
	assert.NotContains(t, r[1], " checked=\"checked\" ")

	assert.Len(t, r[2], 117)
	assert.True(t, strings.HasPrefix(r[2], "<label for=\"c_test1_risotto\"><input "))
	assert.Contains(t, r[2], " name=\"test1\" ")
	assert.Contains(t, r[2], " type=\"radio\" ")
	assert.Contains(t, r[2], " id=\"c_test1_risotto\" ")
	assert.Contains(t, r[2], " value=\"risotto\" ")
	assert.True(t, strings.HasSuffix(r[2], " /> Risotto</label>"))
	assert.NotContains(t, r[2], " checked=\"checked\" ")

	assert.Equal(t, r[3], "")

	f.Attributes = Attributes{"test": "ok"}
	r = strings.Split(_t.Render(f, choices, []string{}), "\n")
	assert.Len(t, r, 4)

	assert.Len(t, r[0], 119)
	assert.True(t, strings.HasPrefix(r[0], "<label for=\"c_test1_pizza\"><input "))
	assert.Contains(t, r[0], " name=\"test1\" ")
	assert.Contains(t, r[0], " type=\"radio\" ")
	assert.Contains(t, r[0], " test=\"ok\" ")
	assert.Contains(t, r[0], " id=\"c_test1_pizza\" ")
	assert.Contains(t, r[0], " value=\"pizza\" ")
	assert.True(t, strings.HasSuffix(r[0], " /> Pizza</label>"))
	assert.NotContains(t, r[0], " checked=\"checked\" ")

	assert.Len(t, r[1], 121)
	assert.True(t, strings.HasPrefix(r[1], "<label for=\"c_test1_pasta\"><input "))
	assert.Contains(t, r[1], " name=\"test1\" ")
	assert.Contains(t, r[1], " type=\"radio\" ")
	assert.Contains(t, r[1], " test=\"ok\" ")
	assert.Contains(t, r[1], " id=\"c_test1_pasta\" ")
	assert.Contains(t, r[1], " value=\"pasta\" ")
	assert.True(t, strings.HasSuffix(r[1], " /> Makaron</label>"))
	assert.NotContains(t, r[1], " checked=\"checked\" ")

	assert.Len(t, r[2], 127)
	assert.True(t, strings.HasPrefix(r[2], "<label for=\"c_test1_risotto\"><input "))
	assert.Contains(t, r[2], " name=\"test1\" ")
	assert.Contains(t, r[2], " type=\"radio\" ")
	assert.Contains(t, r[2], " test=\"ok\" ")
	assert.Contains(t, r[2], " id=\"c_test1_risotto\" ")
	assert.Contains(t, r[2], " value=\"risotto\" ")
	assert.True(t, strings.HasSuffix(r[2], " /> Risotto</label>"))
	assert.NotContains(t, r[2], " checked=\"checked\" ")

	assert.Equal(t, r[3], "")

	// Empty choices
	assert.Equal(t, _t.Render(f, nil, []string{}), "")
}

func TestTypeTextarea(t *testing.T) {
	var resultsSet = TypeTestsSet{
		name:       "Textarea",
		multiValue: false,

		results: TypeTestsResults{
			{[]string{"accdddaabbcce"}, "accdddaabbcce"},
			{[]string{"a", "b"}, "a"},
			{[]string{""}, ""},
			{nil, ""},
		},
	}

	_t := &Textarea{}
	executeTypeTests(t, _t, resultsSet)

	f := &Field{Name: "test1", Type: _t}
	assert.Equal(t, _t.Render(f, nil, []string{}), "<textarea id=\"f_test1\" name=\"test1\"></textarea>")
	assert.Equal(t, _t.Render(f, nil, []string{""}), "<textarea id=\"f_test1\" name=\"test1\"></textarea>")
	assert.Equal(t, _t.Render(f, nil, []string{"accdddaabbcce"}), "<textarea id=\"f_test1\" name=\"test1\">accdddaabbcce</textarea>")
}

func TestTypeInputNumber(t *testing.T) {
	var resultsSet = TypeTestsSet{
		name:       "InputNumber",
		multiValue: false,

		results: TypeTestsResults{
			{[]string{"12"}, int64(12)},
			{[]string{"1ab"}, nil},
			{[]string{"-1"}, int64(-1)},
			{nil, nil},
			{[]string{"-12.3"}, float64(-12.3)},
			{[]string{"-12.3s"}, nil},
			{[]string{"99999999999999.1234123"}, float64(99999999999999.1234123)},
		},
	}

	_t := &InputNumber{}
	executeTypeTests(t, _t, resultsSet)

	f := &Field{Name: "test1", Type: _t}
	assert.Equal(t, _t.Render(f, nil, []string{}), "<input name=\"test1\" type=\"number\" id=\"f_test1\" />")
	assert.Equal(t, _t.Render(f, nil, []string{""}), "<input name=\"test1\" type=\"number\" id=\"f_test1\" />")
	assert.Equal(t, _t.Render(f, nil, []string{"11"}), "<input name=\"test1\" type=\"number\" id=\"f_test1\" value=\"11\" />")
}

func TestTypeChecbox(t *testing.T) {
	var resultsSet = TypeTestsSet{
		name:       "Checkbox",
		multiValue: false,

		results: TypeTestsResults{
			{[]string{"12"}, true},
			{[]string{"1ab"}, true},
			{[]string{"-1"}, true},
			{[]string{""}, false},
			{nil, false},
			{[]string{"-12.3"}, true},
			{[]string{"-12.3s"}, true},
			{[]string{"99999999999999.1234123"}, true},
		},
	}

	_t := &Checkbox{}
	executeTypeTests(t, _t, resultsSet)

	f := &Field{Name: "test1", Type: _t}

	rendered := _t.Render(f, nil, nil)
	assert.Len(t, rendered, 51)
	assert.True(t, strings.HasPrefix(rendered, "<input "))
	assert.Contains(t, rendered, " name=\"test1\" ")
	assert.Contains(t, rendered, " id=\"f_test1\" ")
	assert.True(t, strings.HasSuffix(rendered, " />"))
	assert.NotContains(t, rendered, " checked=\"checked\" ")

	rendered = _t.Render(f, nil, []string{""})
	assert.Len(t, rendered, 51)
	assert.True(t, strings.HasPrefix(rendered, "<input "))
	assert.Contains(t, rendered, " name=\"test1\" ")
	assert.Contains(t, rendered, " id=\"f_test1\" ")
	assert.True(t, strings.HasSuffix(rendered, " />"))
	assert.NotContains(t, rendered, " checked=\"checked\" ")

	rendered = _t.Render(f, nil, []string{"true"})
	assert.Len(t, rendered, 69)
	assert.True(t, strings.HasPrefix(rendered, "<input "))
	assert.Contains(t, rendered, " name=\"test1\" ")
	assert.Contains(t, rendered, " checked=\"checked\" ")
	assert.Contains(t, rendered, " id=\"f_test1\" ")
	assert.True(t, strings.HasSuffix(rendered, " />"))

	f.Attributes = Attributes{"test": "ok"}

	rendered = _t.Render(f, nil, nil)
	assert.Len(t, rendered, 61)
	assert.True(t, strings.HasPrefix(rendered, "<input "))
	assert.Contains(t, rendered, " name=\"test1\" ")
	assert.Contains(t, rendered, " type=\"checkbox\" ")
	assert.Contains(t, rendered, " test=\"ok\" ")
	assert.Contains(t, rendered, " id=\"f_test1\" ")
	assert.True(t, strings.HasSuffix(rendered, " />"))
	assert.NotContains(t, rendered, " checked=\"checked\" ")

	rendered = _t.Render(f, nil, []string{""})
	assert.Len(t, rendered, 61)
	assert.True(t, strings.HasPrefix(rendered, "<input "))
	assert.Contains(t, rendered, " name=\"test1\" ")
	assert.Contains(t, rendered, " type=\"checkbox\" ")
	assert.Contains(t, rendered, " test=\"ok\" ")
	assert.Contains(t, rendered, " id=\"f_test1\" ")
	assert.True(t, strings.HasSuffix(rendered, " />"))
	assert.NotContains(t, rendered, " checked=\"checked\" ")

	rendered = _t.Render(f, nil, []string{"true"})
	assert.Len(t, rendered, 79)
	assert.True(t, strings.HasPrefix(rendered, "<input "))
	assert.Contains(t, rendered, " name=\"test1\" ")
	assert.Contains(t, rendered, " type=\"checkbox\" ")
	assert.Contains(t, rendered, " test=\"ok\" ")
	assert.Contains(t, rendered, " checked=\"checked\" ")
	assert.Contains(t, rendered, " id=\"f_test1\" ")
	assert.True(t, strings.HasSuffix(rendered, " />"))
}

func TestTypeInputEmail(t *testing.T) {
	var resultsSet = TypeTestsSet{
		name:       "InputEmail",
		multiValue: false,

		results: TypeTestsResults{
			{[]string{"12"}, "12"},
			{[]string{"1ab"}, "1ab"},
			{[]string{"-1"}, "-1"},
			{nil, ""},
			{[]string{"-12.3"}, "-12.3"},
			{[]string{"-12.3s"}, "-12.3s"},
			{[]string{"asdasda"}, "asdasda"},
		},
	}

	_t := &InputEmail{}
	executeTypeTests(t, _t, resultsSet)

	f := &Field{Name: "test1", Type: _t}
	assert.Equal(t, _t.Render(f, nil, []string{}), "<input name=\"test1\" type=\"email\" id=\"f_test1\" />")
	assert.Equal(t, _t.Render(f, nil, []string{""}), "<input name=\"test1\" type=\"email\" id=\"f_test1\" />")
	assert.Equal(t, _t.Render(f, nil, []string{"test_value"}), "<input name=\"test1\" type=\"email\" id=\"f_test1\" value=\"test_value\" />")
}

func TestTypeInputPassword(t *testing.T) {
	var resultsSet = TypeTestsSet{
		name:       "InputPassword",
		multiValue: false,

		results: TypeTestsResults{
			{[]string{"12"}, "12"},
			{[]string{"1ab"}, "1ab"},
			{[]string{"-1"}, "-1"},
			{nil, ""},
			{[]string{"-12.3"}, "-12.3"},
			{[]string{"-12.3s"}, "-12.3s"},
			{[]string{"asdasda"}, "asdasda"},
		},
	}

	_t := &InputPassword{}
	executeTypeTests(t, _t, resultsSet)

	f := &Field{Name: "pwd", Type: _t}
	assert.Equal(t, _t.Render(f, nil, []string{}), "<input name=\"pwd\" type=\"password\" id=\"f_pwd\" />")
}
