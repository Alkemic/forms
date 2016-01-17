package forms

import (
	"strconv"
)

type Type interface {
	IsMultiValue() bool
	CleanData(values []string) interface{}
	GetType() string
}

type Input struct{}

func (i *Input) IsMultiValue() bool {
	return false
}

func (i *Input) GetType() string {
	return "input"
}

func (i *Input) CleanData(values []string) interface{} {
	if len(values) > 0 {
		return values[0]
	}

	return ""
}

type Radio struct{}

func (i *Radio) IsMultiValue() bool {
	return true
}

func (i *Radio) CleanData(values []string) interface{} {
	return values
}

func (r *Radio) GetType() string {
	return "radio"
}

type Textarea struct {
	Input
}

type InputNumber struct {
	Input
}

func (i *InputNumber) CleanData(values []string) interface{} {
	if len(values) == 1 {
		fval, err := strconv.ParseFloat(values[0], 64)
		if err != nil {
			return fval
		}

		ival, err := strconv.ParseInt(values[0], 10, 64)
		if err != nil {
			return ival
		}
	}
	return nil
}
