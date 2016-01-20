package forms

import (
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
