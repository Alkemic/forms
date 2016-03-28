package forms

import (
	"fmt"
	"strconv"
)

// List of attributes that should not be
var noUseAttrs []string

func init() {
	noUseAttrs = []string{"type", "name", "value"}
}

// Type is interface that tells us
type Type interface {
	// Tells if fields type should accept multiple values
	IsMultiValue() bool
	// Cleans data before it goes to user
	CleanData(values []string) interface{}
	// Render form
	Render(*Field, []Choice, []string) string
}

type Input struct{}

func (i *Input) IsMultiValue() bool {
	return false
}

func (i *Input) CleanData(values []string) interface{} {
	if len(values) > 0 {
		return values[0]
	}

	return ""
}

func (i *Input) Render(f *Field, cs []Choice, vs []string) string {
	return renderInput(f.Attributes, f.Name, "input", noUseAttrs, vs)
}

type Radio struct{}

func (i *Radio) IsMultiValue() bool {
	return true
}

func (i *Radio) CleanData(values []string) interface{} {
	return values
}

func (i *Radio) Render(f *Field, cs []Choice, vs []string) string {
	field := ""
	attrs := Attributes{}
	for _, c := range cs {
		for k, v := range f.Attributes {
			attrs[k] = v
		}
		attrs["id"] = fmt.Sprintf("c_%s_%s", f.Name, c.Value)
		field = field + fmt.Sprintf(
			"<label for=\"c_%s_%s\">%s %s</label>\n",
			f.Name, c.Value,
			renderInput(attrs, f.Name, "radio", noUseAttrs, []string{c.Value}),
			c.Label,
		)
	}
	return field
}

type Textarea struct {
	Input
}

func (t *Textarea) Render(f *Field, cs []Choice, vs []string) string {
	value := ""
	if len(vs) > 0 && vs[0] != "" {
		value = vs[0]
	}

	return fmt.Sprintf(
		"<textarea id=\"f_%s\" name=\"%s\"%s>%s</textarea>", f.Name, f.Name,
		prepareAttributes(f.Attributes, noUseAttrs), value,
	)
}

type InputNumber struct{}

func (t *InputNumber) IsMultiValue() bool {
	return false
}

func (t *InputNumber) CleanData(values []string) interface{} {
	if len(values) == 1 {
		ival, err := strconv.ParseInt(values[0], 10, 64)
		if err == nil {
			return ival
		}

		fval, err := strconv.ParseFloat(values[0], 64)
		if err == nil {
			return fval
		}
	}

	return nil
}

func (t *InputNumber) Render(f *Field, cs []Choice, vs []string) string {
	return renderInput(f.Attributes, f.Name, "number", noUseAttrs, vs)
}

type Checkbox struct {
	*Input
}

func (t *Checkbox) CleanData(values []string) interface{} {
	if len(values) == 1 && values[0] != "" {
		return true
	}

	return false
}

func (t *Checkbox) Render(f *Field, cs []Choice, vs []string) string {
	var attrs Attributes
	if f.Attributes == nil {
		attrs = Attributes{}
	} else {
		attrs = f.Attributes
	}

	if len(vs) > 0 && vs[0] != "" {
		attrs["checked"] = "checked"
	}

	return renderInput(attrs, f.Name, "checkbox", noUseAttrs, nil)
}

type InputEmail struct {
	*Input
}

func (t *InputEmail) Render(f *Field, cs []Choice, vs []string) string {
	return renderInput(f.Attributes, f.Name, "email", noUseAttrs, vs)
}

type InputPassword struct {
	*Input
}

func (t *InputPassword) Render(f *Field, cs []Choice, vs []string) string {
	return renderInput(f.Attributes, f.Name, "password", noUseAttrs, vs)
}
