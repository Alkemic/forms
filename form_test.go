package forms

import (
	"github.com/stretchr/testify/assert"
	"net/url"
	"testing"
)

// TestFormIsValid tests data that come from request (url.Values)
func TestFormIsValid(t *testing.T) {
	postData := url.Values{}
	postData.Set("field1", "Foo")
	postData.Set("field2", "Bar")
	postData.Set("fieldX", "Ham")

	f := Form{
		Fields: map[string]*Field{
			"field1": &Field{
				Type: &Input{},
			},
			"field2": &Field{
				Type: &Input{},
			},
		},
	}

	assert.Equal(
		t, f.CleanedData, *new(CleanedData),
		"CleanedData should be empty at beggining")

	assert.True(t, f.IsValid(postData), "Form should pass")
	assert.Equal(
		t, f.CleanedData, CleanedData{"field1": "Foo", "field2": "Bar"},
		"Forms CleanedData field should contain cleaned data")

	assert.True(t, f.IsValid(url.Values{}), "Form should pass")
	assert.Equal(
		t, f.CleanedData, CleanedData{"field1": "", "field2": ""},
		"Form should pass")
}

// TestFromIsValidMap tests data that come from map
func TestFromIsValidMap(t *testing.T) {
	f := Form{
		Fields: map[string]*Field{
			"field1": &Field{
				Type: &Input{},
			},
			"field2": &Field{
				Type: &Input{},
			},
		},
	}

	values := map[string]interface{}{
		"field1": "Spam",
		"field2": []string{"Ham"},
	}

	assert.True(t, f.IsValidMap(values), "Form should pass")
	assert.Equal(
		t, f.CleanedData, CleanedData{"field1": "Spam", "field2": "Ham"},
		"Form should pass")

	values = map[string]interface{}{
		"field1": "Spam",
	}

	assert.True(t, f.IsValidMap(values), "Form should pass")
	assert.Equal(
		t, f.CleanedData, CleanedData{"field1": "Spam", "field2": ""},
		"Form should pass")
}

func TestDefaultFieldType(t *testing.T) {
	f := Form{
		Fields: map[string]*Field{
			"field1": &Field{},
			"field2": &Field{},
		},
	}

	values := map[string]interface{}{
		"field1": "Spam",
		"field2": []string{"Ham"},
	}

	assert.True(t, f.IsValidMap(values), "Form should pass")
}

func TestBasicValidation(t *testing.T) {
	f := Form{
		Fields: map[string]*Field{
			"field1": &Field{
				Validators: []Validator{
					&Required{},
				},
			},
			"field2": &Field{},
		},
	}

	values := map[string]interface{}{
		"field2": []string{"Ham"},
	}

	assert.False(t, f.IsValidMap(values), "Form shouldn't pass")
}
