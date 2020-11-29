package forms

import (
	"fmt"
	"html/template"
	"log"
)

// Choice is used to store choices in field
// I choose to use type here because we need ordering of choices
type Choice struct {
	Value string
	Label string
}

// Field represent single field in form and its validatorss
type Field struct {
	Name string

	Label           string
	LabelAttributes Attributes

	Choices      []Choice
	Value        []string
	InitialValue interface{}
	Type         Type
	Attributes   Attributes

	Validators []Validator
	Errors     []string
}

// IsValid do data validation
func (f *Field) IsValid(values []string) (isValid bool) {
	c := len(values)

	if f.Type == nil {
		f.Type = &Input{}
	}

	if !f.Type.IsMultiValue() && c > 1 {
		f.Errors = append(f.Errors, translations["INCORRECT_MULTI_VAL"])
		return false
	}

	isValid = true
	for _, validator := range f.Validators {
		result, msgs := validator.IsValid(values)
		if !result {
			f.Errors = append(f.Errors, msgs...)
			isValid = false
		}
	}

	return isValid
}

// Render field (in matter of fact, only passing through to render method on type)
func (f *Field) Render() template.HTML {
	if f.Type == nil {
		f.Type = &Input{}
	}

	var values []string
	if f.Value == nil && f.InitialValue != nil {
		if isSlice(f.InitialValue) {
			for _, value := range f.InitialValue.([]interface{}) {
				if stringValue, ok := anyToString(value); ok {
					values = append(values, stringValue)
				} else {
					log.Println(value, "is incorrect type for InitialValue")
				}
			}
		} else {
			if stringValue, ok := anyToString(f.InitialValue); ok {
				values = append(values, stringValue)
			} else {
				log.Println(f.InitialValue, "is incorrect type for InitialValue")
			}
		}
	} else {
		values = f.Value
	}

	return f.Type.Render(f, f.Choices, values)
}

// RenderLabel render label for field
func (f *Field) RenderLabel() template.HTML {
	attributes := prepareAttributes(f.LabelAttributes, []string{"for"})

	return template.HTML(fmt.Sprintf("<label for=\"f_%s\"%s>%s</label>", f.Name, attributes, f.Label))
}

// HasErrors returns information if there are validation errors in this field
func (f *Field) HasErrors() bool {
	return len(f.Errors) > 0
}

// RenderErrors render all errors as list (<ul>) with class "errors"
func (f *Field) RenderErrors() template.HTML {
	if !f.HasErrors() {
		return ""
	}

	rendered := ""
	for _, err := range f.Errors {
		rendered += fmt.Sprintf("<li>%s</li>\n", err)
	}

	return template.HTML(fmt.Sprintf("<ul class=\"errors\">\n%s</ul>", rendered))
}
