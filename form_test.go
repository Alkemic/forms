package forms

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
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

func TestFormNewFunc(t *testing.T) {
	f := New(
		map[string]*Field{
			"field1": &Field{},
			"field2": &Field{},
		},
		Attributes{"id": "test"},
	)

	assert.Equal(t, f.Fields["field1"].Name, "field1", "Field name should propagate")
	assert.Equal(t, f.Fields["field2"].Name, "field2", "Field name should propagate")
	assert.Equal(t, f.Attributes, Attributes{"id": "test"}, "Attributes should be set")
}

func TestFormOpenTag(t *testing.T) {
	openTag := New(nil, nil).OpenTag()

	assert.Len(t, openTag, 6)
	assert.Equal(t, openTag, `<form>`)

	openTag = New(
		nil,
		Attributes{"id": "test", "class": "register-form rwd-form"},
	).OpenTag()

	assert.Len(t, openTag, 47)
	assert.Contains(t, openTag, `<form `)
	assert.Contains(t, openTag, ` id="test"`)
	assert.Contains(t, openTag, ` class="register-form rwd-form"`)
	assert.Contains(t, openTag, `>`)
}

func TestFormCloseTag(t *testing.T) {
	f := New(nil, Attributes{"id": "test"})

	assert.Equal(t, f.CloseTag(), "</form>")
}
