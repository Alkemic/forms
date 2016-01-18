package forms

type Field struct {
	Name string

	Value string
	Type  Type

	Validators []Validator

	Errors []string

	Attributes map[string]string
}

func (f *Field) IsValid(values []string) (result bool) {
	c := len(values)

	if f.Type == nil {
		f.Type = &Input{}
	}

	if !f.Type.IsMultiValue() && c > 1 {
		f.Errors = append(f.Errors, translations["INCORRECT_MULTI_VAL"])
	}

	// var result errors.Error
	result = true
	for _, validator := range f.Validators {
		result, msg := validator.IsValid(f.Value)
		if !result {
			f.Errors = append(f.Errors, msg)
			result = false
		}
	}

	return result
}

func (f *Field) Field() (field string) {
	return field
}

func (f *Field) Label() (l string) {
	return l
}
