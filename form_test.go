package forms

import (
	"net/url"
	"testing"
)

func TestBaseForm(t *testing.T) {
	postData := url.Values{}
	postData.Set("field1", "Foo")
	postData.Set("field2", "Bar")
	postData.Set("fieldX", "Ham")

	f := Form{
		Fields: map[string]*Field{
			"field1": &Field{
				Type: &Input{},
				// Validators: []Validator{},
			},
			"field2": &Field{
				Type:       &Input{},
				Validators: []Validator{},
			},
		},
	}

	if !f.IsValid(postData) {
		t.Error("Expected true, got false")
	}
}
