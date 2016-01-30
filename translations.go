package forms

var translations map[string]string

func init() {
	translations = map[string]string{
		"REQUIRED": "This field can't be empty",

		"INCORRECT_EMAIL": "\"%s\" is not correct email address",

		"INCORRECT_MULTI_VAL":  "You supplied more than one value for this field",
		"INCORRECT_MIN_LENGTH": "Value \"%%s\" need to be at least %d chars long",
		"INCORRECT_MAX_LENGTH": "Value \"%%s\" need to be at max %d chars long",

		"NO_MATCH_PATTERN": "Value \"%%s\" doesn't match pattern \"%s\"",

		"VALUE_NOT_FOUND": "Value \"%s\" not found in slice",
	}
}
