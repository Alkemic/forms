# forms [![Build Status](https://travis-ci.org/Alkemic/forms.svg?branch=master)](https://travis-ci.org/Alkemic/forms) [![Coverage Status](https://coveralls.io/repos/github/Alkemic/forms/badge.svg?branch=master)](https://coveralls.io/github/Alkemic/forms?branch=master) [![Go Report Card](https://goreportcard.com/badge/github.com/Alkemic/forms)](https://goreportcard.com/report/github.com/Alkemic/forms) [![GoDoc](https://godoc.org/github.com/asaskevich/govalidator?status.png)](https://godoc.org/github.com/Alkemic/forms)
I found no sufficient forms library in GO, so I decided to write one by my self. ;-)

The main goal is separate validation from field rendering, so you won't find here silly structures like:

```go
Field{Required: True}
```

Validate if user entered any data should be done in validation, and if you want ``<input />`` to have a ``required`` attribute, you need to set this in field's attributes.

## Usage

Here is a basic usage for forms. Keep in mind, that validation is separated from field,
and inputs and albels HTML attributes are also separated.

```go
import (
	"fmt"
	"net/http"

	"github.com/Alkemic/forms"
)

func someView(w http.ResponseWriter, r *http.Request) {
	form := forms.New(
		map[string]*forms.Field{
			"email": &forms.Field{
				Validators: []forms.Validator{
					&forms.Required{},
				},
			},
			"password": &forms.Field{
				Type: &forms.InputPassword{},
				Validators: []forms.Validator{
					&forms.Required{},
				},
			},
		},
		forms.Attributes{"id": "login-form"},
	)

	if form.IsValid(r.PostForm) {
		// if valid you can access cleaned data in attribute CleanedData
		for k, v := range form.CleanedData{
			fmt.Println("key:", k, "value:", v)
		}
	} else {
		// else not ;-)
	}
}
```

I've decided to don't write whole form rendering method, because, let's be honest,
it won't give level of control over form that we need and in the end you will
have to do it by yourself. Insted of there are methods that will help you with
displaying form.

```html
{{.Form.OpenTag}}
{{if .Form.Fields.email.HasErrors}}
    {{.Form.Fields.email.RenderErrors}}
{{end}}
{{.Form.Fields.email.RenderLabel}}
{{.Form.Fields.email.Render}}

{{if .Form.Fields.password.HasErrors}}
    {{.Form.Fields.password.RenderErrors}}
{{end}}
{{.Form.Fields.password.RenderLabel}}
{{.Form.Fields.password.Render}}
{{.Form.CloseTag}}
```

Eventually you can render errors by yourself

```html
{{if .Form.Fields.email.HasErrors}}
    <ul>
        {{range .Form.Fields.email.Errors}}
        <li class="error">{{.}}</li>
        {{end}}
    </ul>
{{end}}
```

## Installation
As usual, no magic here:
```bash
$ go get github.com/Alkemic/forms
```

## Field

Fields are representation of single field in form, it's a container for validators and attributes.

```go
Field{
	Type: &Input{},
	Validators: []Validator{
		&Required{},
	},
	Attributes{"id": "test"},
	Label: "Test label",
	LabelAttributes: Attributes{
		"required": "required",
		"id":   "test",
		"attr": "value",
	},
}
```

When fields label is rendered (``field.RenderLabel``) attribute ``for`` is automaticly added as well
as attribute ``id`` to field.

## Field types

Types are responsible for field behavior: rendering, cleaning data and giving information if
field accept multiple value. Generally you should not access fields type directly, as its used by field.

### Cleaning data

The incoming data need to be cleaned after succesful validation, and before we give them to user.
By clean, we mean that we convert them to format/
All of this transformation are done by method ``CleanData`` on ``Type``.
For example, when we crate ``Field`` with type ``NumberInput`` in ``form.CleanedData`` we
find a number (``int``), for ``MultiSelect`` we find a slice with all selected values.

## TODO
**Big fat note: this library is under development, and it's API may or may not change.**
Currently this library works, but I don't recomend this for prodution or even thinking about production usage. ;-)

* [x] Field rendering
* [x] Initial data support
* [ ] Internationalization
* [ ] Field types (inc. types introduced in HTML5)
  * [x] Input
  * [x] Textarea
  * [X] Radio
  * [ ] Select
  * [X] Email
  * [X] Number
  * [ ] Color
  * [ ] File
  * [ ] Hidden
  * [ ] Image
  * [x] Month
  * [x] Password
  * [ ] Range
  * [x] Telephone
  * [ ] Time
  * [x] URL
  * [x] Week
  * [x] Date
  * [x] Datetime
  * [x] Datetime-local
* [ ] Validators
  * [x] Regexp
  * [x] Required
  * [x] Email
  * [x] MinLength
  * [x] MaxLength
  * [x] InSlice
  * [ ] MinValue
  * [ ] MaxValue
  * [ ] URL
  * [ ] Date
  * [ ] Time
  * [ ] DateTime
