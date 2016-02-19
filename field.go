package forms

import (
	"fmt"
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

	Choices    []Choice
	Value      []string
	Type       Type
	Attributes Attributes

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
func (f *Field) Render() string {
	if f.Type == nil {
		f.Type = &Input{}
	}

	return f.Type.Render(f, f.Choices, f.Value)
}

// RenderLabel render label for field
func (f *Field) RenderLabel() string {
	attributes := prepareAttributes(f.LabelAttributes, []string{"for"})

	return fmt.Sprintf("<label for=\"f_%s\"%s>%s</label>", f.Name, attributes, f.Label)
}

// HasErrors returns information if there are validation errors in this field
func (f *Field) HasErrors() bool {
	return len(f.Errors) > 0
}

// RenderErrors render all errors as list (<ul>) with class "errors"
func (f *Field) RenderErrors() string {
	if !f.HasErrors() {
		return ""
	}

	rendered := ""
	for _, err := range f.Errors {
		rendered += fmt.Sprintf("<li>%s</li>\n", err)
	}

	return fmt.Sprintf("<ul class=\"errors\">\n%s</li>", rendered)
}
