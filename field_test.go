package forms

import (
	"github.com/stretchr/testify/assert"
	"testing"
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
		LabelAttributes: map[string]string{
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
		LabelAttributes: map[string]string{
			"for": "asd",
		},
	}
	label = f.RenderLabel()
	assert.Contains(t, label, "<label for=\"f_test\"", "")
	assert.Contains(t, label, ">Test label</label>", "")
	assert.NotContains(t, label, " for=\"asd\"", "")
}
