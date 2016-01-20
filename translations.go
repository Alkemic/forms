package forms

var translations map[string]string

func init() {
	translations = map[string]string{
		"REQUIRED": "This field can't be empty",

		"INCORRECT_EMAIL": "Entered value is not correct email address",

		"INCORRECT_MULTI_VAL":  "You supplied more than one value for this field",
		"INCORRECT_MIN_LENGTH": "Entered value need to be at least %d chars long",
		"INCORRECT_MAX_LENGTH": "Entered value need to be at max %d chars long",

		"NO_MATCH_PATTERN": "Value doesn't match pattern \"%s\"",

		"VALUE_NOT_FOUND": "Value \"%s\" not found in slice",
	}
}
