package forms

import (
	"fmt"
	"net/url"
)

type CleanedData map[string]interface{}

type Form struct {
	Data   map[string]string
	Fields map[string]*Field

	Errors []string

	Attributes  map[string]string
	CleanedData CleanedData
}

func (f *Form) Clear() {
	f.Errors = []string{}
	for _, field := range f.Fields {
		field.Errors = []string{}
	}
}

func (f *Form) IsValid(data url.Values) (isValid bool) {
	f.Clear()
	cleanedData := CleanedData{}

	for name, field := range f.Fields {
		values, _ := data[name]

		isValid = field.IsValid(values)

		if isValid {
			cleanedData[name] = field.Type.CleanData(values)
		}
	}

	if isValid {
		f.CleanedData = cleanedData
	}

	return isValid
}

// SetFormData populates data from http.Request.Form values
func (f *Form) IsValidArray(data url.Values) {
	return
}

func Example1() {
	fields := map[string]*Field{
		"email": &Field{
			Type: &Input{},
			Validators: []Validator{
				&Email{},
			},
			Attributes: map[string]string{
				"required": "",
				"class":    "",
			},
		},
	}

	loginForm := Form{
		Fields: fields,
	}

	fmt.Println(loginForm)
}
