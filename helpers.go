package forms

import (
	"fmt"
	"reflect"
)

func IsSlice(value interface{}) bool {
	return reflect.TypeOf(value).Kind() == reflect.Slice
}

func ValueInSlice(s string, vs []string) bool {
	for _, v := range vs {
		if s == v {
			return true
		}
	}

	return false
}

func prepareAttributes(attrs Attributes, noUse []string) string {
	attributes := ""

	for k, v := range attrs {
		if !ValueInSlice(k, noUse) {
			attributes = attributes + fmt.Sprintf(" %s=\"%s\"", k, v)
		}
	}

	return attributes
}

func renderInput(as Attributes, n, t string, noUse, vs []string) string {
	if as == nil {
		as = Attributes{}
	}

	if _, ok := as["id"]; !ok {
		as["id"] = fmt.Sprintf("f_%s", n)
	}

	attributes := prepareAttributes(as, noUse)

	if len(vs) > 0 && vs[0] != "" {
		attributes = attributes + fmt.Sprintf(" value=\"%s\"", vs[0])
	}

	return fmt.Sprintf("<input name=\"%s\" type=\"%s\"%s />", n, t, attributes)
}
