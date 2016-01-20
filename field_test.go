package forms

import (
	"github.com/stretchr/testify/assert"
	// "net/url"
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
