package forms

import (
	"fmt"
	"net/url"
	"reflect"
)

// Attributes is structure that contains forms or fields attributes
type Attributes map[string]interface{}

// Data is structure in which we store cleaned and initial data
type Data map[string]interface{}

// Form is structure that hold all fields, data, and
type Form struct {
	// Keeps all the fields
	Fields map[string]*Field

	// Form attributes
	Attributes Attributes

	// Data that are used in validation
	IncomingData url.Values
	// After validation is done we put data here
	CleanedData Data
	// Initial data are used before form validation
	InitialData Data
}

// Clear clears error and data on fields in form
func (f *Form) Clear() {
	f.CleanedData = nil
	for _, field := range f.Fields {
		field.Errors = []string{}
	}
}

// SetInitial sets initial data on form
func (f *Form) SetInitial(data Data) {
	f.InitialData = data
	for name, field := range f.Fields {
		if values, ok := data[name]; ok {
			field.InitialValue = values
		}
	}
}

// IsValid validate all fields and if all is correct assign cleaned data from
// every field to forms CleanedData attribute
func (f *Form) IsValid(data url.Values) bool {
	f.Clear()
	f.IncomingData = data
	isValid := true
	cleanedData := Data{}

	for name, field := range f.Fields {
		values, _ := data[name]
		field.Value = values

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

// IsValidMap populates data from map.
// It accepts map of string/strings with keys as field names.
func (f *Form) IsValidMap(values map[string]interface{}) bool {
	data := url.Values{}

	for k, v := range values {
		if isSlice(v) {
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

// OpenTag render opening tag of the form with given attributes
func (f *Form) OpenTag() string {
	return fmt.Sprintf("<form%s>", prepareAttributes(f.Attributes, nil))
}

// CloseTag render closing tag for form
func (f *Form) CloseTag() string {
	return "</form>"
}

// New is shorthand, and preferred way, to create new form.
// Main difference is that, this approach add field name, basing on key in map,
// to a field instance
// Example
//     form := forms.New(
//         map[string]*forms.Field{
//             "field1": &forms.Field{},
//             "field2": &forms.Field{},
//         },
//         forms.Attributes{"id": "test"},
//     )
func New(fields map[string]*Field, attrs Attributes) *Form {
	for fieldName, field := range fields {
		field.Name = fieldName
	}

	return &Form{
		Fields:     fields,
		Attributes: attrs,
	}
}
