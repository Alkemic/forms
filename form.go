package forms

import (
	"fmt"
	"net/url"
	"reflect"
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

func (f *Form) IsValid(data url.Values) bool {
	f.Clear()
	isValid := true
	cleanedData := CleanedData{}

	for name, field := range f.Fields {
		values, _ := data[name]

		result := field.IsValid(values)

		if result {
			cleanedData[name] = field.Type.CleanData(values)
		} else {
			isValid = false
		}
	}

	if isValid {
		f.CleanedData = cleanedData
	}

	return isValid
}

// IsValidMap populates data from map
// IsValidMap accepts map of string/strings with keys as field names
func (f *Form) IsValidMap(values map[string]interface{}) bool {
	data := url.Values{}

	for k, v := range values {
		if IsSlice(v) {
			s := reflect.ValueOf(v)
			for i := 0; i < s.Len(); i++ {
				data.Add(k, s.Index(i).String())
			}
		} else {
			str, _ := v.(string)
			data.Set(k, str)
		}
	}

	return f.IsValid(data)
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
