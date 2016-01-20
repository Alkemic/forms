package forms

type Field struct {
	Name string

	Value      interface{}
	Type       Type
	Attributes map[string]string

	Validators []Validator
	Errors     []string
}

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
		result, msg := validator.IsValid(values)
		if !result {
			f.Errors = append(f.Errors, msg)
			isValid = false
		}
	}

	return isValid
}

func (f *Field) Field() (field string) {
	return field
}

func (f *Field) Label() (l string) {
	return l
}
