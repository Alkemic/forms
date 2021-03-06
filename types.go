package forms

import (
	"fmt"
	"html/template"
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
	Render(*Field, []Choice, []string) template.HTML
}

// Input is basic input type
type Input struct{}

// IsMultiValue returns if basic input allow multiple values
func (i *Input) IsMultiValue() bool {
	return false
}

// CleanData returns cleaned values for basic input
func (i *Input) CleanData(values []string) interface{} {
	if len(values) > 0 {
		return values[0]
	}

	return ""
}

// Render returns string with rendered basic input
func (i *Input) Render(f *Field, cs []Choice, vs []string) template.HTML {
	return renderInput(f.Attributes, f.Name, "input", noUseAttrs, vs)
}

// Radio is radio input type
type Radio struct{}

// IsMultiValue returns if radio input allow multiple values
func (i *Radio) IsMultiValue() bool {
	return true
}

// CleanData returns cleaned values for radio
func (i *Radio) CleanData(values []string) interface{} {
	return values
}

// Render returns string with rendered radio input
func (i *Radio) Render(f *Field, cs []Choice, vs []string) template.HTML {
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
	return template.HTML(field)
}

// Textarea is textarea type
type Textarea struct {
	Input
}

// Render returns string with rendered textarea
func (t *Textarea) Render(f *Field, cs []Choice, vs []string) template.HTML {
	value := ""
	if len(vs) > 0 && vs[0] != "" {
		value = vs[0]
	}

	return template.HTML(fmt.Sprintf(
		"<textarea id=\"f_%s\" name=\"%s\"%s>%s</textarea>", f.Name, f.Name,
		prepareAttributes(f.Attributes, noUseAttrs), value,
	))
}

// InputNumber is number input type
type InputNumber struct{}

// IsMultiValue returns if numeric input allow multiple values
func (t *InputNumber) IsMultiValue() bool {
	return false
}

// CleanData returns cleaned values for number input
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

// Render returns string with rendered number input
func (t *InputNumber) Render(f *Field, cs []Choice, vs []string) template.HTML {
	return renderInput(f.Attributes, f.Name, "number", noUseAttrs, vs)
}

// Checkbox is checkbox input type
type Checkbox struct {
	*Input
}

// CleanData returns cleaned values for checkbox
func (t *Checkbox) CleanData(values []string) interface{} {
	if len(values) == 1 && values[0] != "" {
		return true
	}

	return false
}

// Render returns string with rendered checkbox input
func (t *Checkbox) Render(f *Field, cs []Choice, vs []string) template.HTML {
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

// InputEmail is email input type
type InputEmail struct {
	*Input
}

// Render returns string with rendered email input
func (t *InputEmail) Render(f *Field, cs []Choice, vs []string) template.HTML {
	return renderInput(f.Attributes, f.Name, "email", noUseAttrs, vs)
}

// InputPassword is password input type
type InputPassword struct {
	*Input
}

// Render returns string with rendered password input
func (t *InputPassword) Render(f *Field, cs []Choice, vs []string) template.HTML {
	return renderInput(f.Attributes, f.Name, "password", noUseAttrs, vs)
}

// InputDate is date input type
type InputDate struct {
	*Input
}

// Render returns string with rendered date input
func (t *InputDate) Render(f *Field, cs []Choice, vs []string) template.HTML {
	return renderInput(f.Attributes, f.Name, "date", noUseAttrs, vs)
}

// InputTime is date input type
type InputTime struct {
	*Input
}

// Render returns string with rendered time input
func (t *InputTime) Render(f *Field, cs []Choice, vs []string) template.HTML {
	return renderInput(f.Attributes, f.Name, "time", noUseAttrs, vs)
}

// InputDateTime is date datetime (uses datetime-local) input type
type InputDateTime struct {
	*Input
}

// Render returns string with rendered datetime input
func (t *InputDateTime) Render(f *Field, cs []Choice, vs []string) template.HTML {
	return renderInput(f.Attributes, f.Name, "datetime-local", noUseAttrs, vs)
}

// InputMonth is month input type
type InputMonth struct {
	*Input
}

// Render returns string with rendered month input
func (t *InputMonth) Render(f *Field, cs []Choice, vs []string) template.HTML {
	return renderInput(f.Attributes, f.Name, "month", noUseAttrs, vs)
}

// InputWeek is week input type
type InputWeek struct {
	*Input
}

// Render returns string with rendered week input
func (t *InputWeek) Render(f *Field, cs []Choice, vs []string) template.HTML {
	return renderInput(f.Attributes, f.Name, "week", noUseAttrs, vs)
}

// InputURL is url input type
type InputURL struct {
	*Input
}

// Render returns string with rendered url input
func (t *InputURL) Render(f *Field, cs []Choice, vs []string) template.HTML {
	return renderInput(f.Attributes, f.Name, "url", noUseAttrs, vs)
}

// InputTel is tel input type
type InputTel struct {
	*Input
}

// Render returns string with rendered tel input
func (t *InputTel) Render(f *Field, cs []Choice, vs []string) template.HTML {
	return renderInput(f.Attributes, f.Name, "tel", noUseAttrs, vs)
}

// InputSearch is tel search type
type InputSearch struct {
	*Input
}

// Render returns string with rendered search input
func (t *InputSearch) Render(f *Field, cs []Choice, vs []string) template.HTML {
	return renderInput(f.Attributes, f.Name, "search", noUseAttrs, vs)
}
