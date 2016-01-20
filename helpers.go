package forms

import (
	"reflect"
)

func IsSlice(value interface{}) bool {
	return reflect.TypeOf(value).Kind() == reflect.Slice
}
