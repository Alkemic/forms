package forms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFieldNoValidation(t *testing.T) {
	f := Field{}

	r := f.IsValid([]string{""})
	assert.True(t, r, "Field without validators should pass validation")
}

func TestFieldBasicValidation(t *testing.T) {
	f := Field{
		Validators: []Validator{
			&Required{},
		},
	}

	r := f.IsValid([]string{""})
	assert.False(t, r, "Field validation should fail")
}

func TestFieldDefaultType(t *testing.T) {
	f := Field{}
	_ = f.IsValid([]string{""})
	assert.Equal(t, f.Type, &Input{}, "Field should have defult type Input")

	f = Field{}
	_ = f.Render()
	assert.Equal(t, f.Type, &Input{}, "Field should have defult type Input")
}

func TestFieldMultiValue(t *testing.T) {
	f := Field{}
	r := f.IsValid([]string{"a", "b"})
	assert.False(t, r, "Validation should fail when supplying multivalue for non multi value field")
	assert.Equal(t, f.Errors, []string{translations["INCORRECT_MULTI_VAL"]}, "Field should have defult type Input")
}

func TestFieldRenderLabel(t *testing.T) {
	f := Field{
		Name:  "test",
		Label: "Test label",
		LabelAttributes: Attributes{
			"v":    "asd",
			"id":   "test",
			"attr": "value",
		},
	}
	label := f.RenderLabel()
	assert.Contains(t, label, "<label for=\"f_test\"", "")
	assert.Contains(t, label, " v=\"asd\"", "")
	assert.Contains(t, label, " id=\"test\"", "")
	assert.Contains(t, label, " attr=\"value\"", "")
	assert.Contains(t, label, ">Test label</label>", "")

	f = Field{
		Name:  "test",
		Label: "Test label",
		LabelAttributes: Attributes{
			"for": "asd",
		},
	}
	label = f.RenderLabel()
	assert.Contains(t, label, "<label for=\"f_test\"", "")
	assert.Contains(t, label, ">Test label</label>", "")
	assert.NotContains(t, label, " for=\"asd\"", "")
}

func TestFieldRender(t *testing.T) {
	var _t Type
	var f Field
	_t = &Input{}
	f = Field{Type: _t, Name: "test1"}
	assert.Equal(t, f.Render(), _t.Render(&f, nil, []string{}))

	_t = &Textarea{}
	f = Field{Type: _t, Name: "test1"}
	assert.Equal(t, f.Render(), _t.Render(&f, nil, []string{}))

	_t = &Radio{}
	f = Field{Type: _t, Name: "test1"}
	assert.Equal(t, f.Render(), _t.Render(&f, nil, []string{}))
}

func TestFieldHandlingErrors(t *testing.T) {
	var f Field
	f = Field{}
	assert.False(t, f.HasErrors())
	assert.Equal(t, f.RenderErrors(), "")
	f = Field{}
	f.Errors = []string{"Error"}
	assert.True(t, f.HasErrors())
	assert.Equal(t, f.RenderErrors(), "<ul class=\"errors\">\n<li>Error</li>\n</li>")
}

func TestFieldInitialValueRender(t *testing.T) {
	var f Field
	var _t Type

	_t = &Input{}
	f = Field{Type: _t, Name: "test1"}
	f.InitialValue = []interface{}{"a1", "b1"}
	assert.Contains(t, f.Render(), " value=\"a1\" ")

	_t = &Input{}
	f = Field{Type: _t, Name: "test1"}
	f.InitialValue = []interface{}{"a1", "b1"}
	f.Value = []string{"c1"}
	assert.Contains(t, f.Render(), " value=\"c1\" ")

	_t = &Input{}
	f = Field{Type: _t, Name: "test2"}
	f.InitialValue = "test"
	assert.Contains(t, f.Render(), " value=\"test\" ")

	_t = &Input{}
	f = Field{Type: _t, Name: "test2"}
	f.InitialValue = "test"
	f.Value = []string{"incoming2"}
	assert.Contains(t, f.Render(), " value=\"incoming2\" ")
}

func TestFieldRenderWithInitial(t *testing.T) {
	var f Field

	f = Field{Name: "test1", InitialValue: "123"}
	assert.Equal(t, f.Render(), "<input name=\"test1\" type=\"input\" id=\"f_test1\" value=\"123\" />")

	f = Field{Name: "test1", InitialValue: []interface{}{"123", "345"}}
	assert.Equal(t, f.Render(), "<input name=\"test1\" type=\"input\" id=\"f_test1\" value=\"123\" />")

	f = Field{Name: "test1", InitialValue: "123", Value: []string{"incoming"}}
	assert.Equal(t, f.Render(), "<input name=\"test1\" type=\"input\" id=\"f_test1\" value=\"incoming\" />")
}

func TestFieldRenderWithInitialAndErrors(t *testing.T) {
	var f Field

	f = Field{Name: "test1", InitialValue: Input{}}
	assert.Equal(t, f.Render(), "<input name=\"test1\" type=\"input\" id=\"f_test1\" />")

	f = Field{Name: "test1", InitialValue: []interface{}{Required{}, Input{}}}
	assert.Equal(t, f.Render(), "<input name=\"test1\" type=\"input\" id=\"f_test1\" />")
}
