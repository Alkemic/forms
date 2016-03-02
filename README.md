# forms
I found no sufficient forms library in GO, so I decided to write one by my self. ;-)

The main goal is separate validation from field rendering, so you won't find here silly structures like:

```go
Field{Required: True}
```

Validate if user entered any data should be done in validation, and if you want ``<input />`` to have a ``required`` attribute, you need to set this in field's attributes.

## Installation
As usual, no magic here:
```bash
$ go get github.com/Alkemic/forms
```

## Usage
Some day you will find here some documentation, but now here is this example and you can dig through tests files.
```go
import "github.com/Alkemic/forms"

form := forms.New(
	map[string]*Field{
		"field1": &Field{},
		"field2": &Field{},
	},
	Attributes{"id": "test"},
)

if form.IsValid(r.PostForm) {
    // if valid
} else {
    // else not ;-)
}
```

## TODO
**Big fat note: this library is under development, and it's API may or may not change.**
Currently this library works, but I don't recomend this for prodution or even thinking about production usage. ;-)

* [ ] Field rendering
* [ ] Initial data support
* [ ] Field types (inc. types introduced in HTML5)
 * [x] Input
 * [x] Textarea
 * [ ] Radio
 * [ ] Select
 * [ ] Email
 * [ ] Number
 * [ ] Color
 * [ ] File
 * [ ] Hidden
 * [ ] Image
 * [ ] Month
 * [ ] Password
 * [ ] Range
 * [ ] Telephone
 * [ ] Time
 * [ ] URL
 * [ ] Week
 * [ ] Date
 * [ ] Datetime
 * [ ] Datetime-local
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
