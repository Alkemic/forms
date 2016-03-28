package forms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrepareAttributes(t *testing.T) {
	prepared := prepareAttributes(Attributes{
		"v":         "asd",
		"id":        "test",
		"attr":      "value",
		"name":      "value",
		"forbidden": "value",
	}, []string{"name", "forbidden"})
	assert.Contains(t, prepared, " v=\"asd\"", "")
	assert.Contains(t, prepared, " id=\"test\"", "")
	assert.Contains(t, prepared, " attr=\"value\"", "")
	assert.NotContains(t, prepared, " name=\"value\"", "")
	assert.NotContains(t, prepared, " forbidden=\"value\"", "")
}

func TestAnyToString(t *testing.T) {
	var s string
	var b bool

	s, b = anyToString(8)
	assert.True(t, b)
	assert.Equal(t, s, "8")

	s, b = anyToString(8.8)
	assert.True(t, b)
	assert.Equal(t, s, "8.8")

	s, b = anyToString(float32(8.8))
	assert.True(t, b)
	assert.Equal(t, s, "8.8")

	s, b = anyToString(uint32(123123123))
	assert.True(t, b)
	assert.Equal(t, s, "123123123")

	s, b = anyToString(true)
	assert.True(t, b)
	assert.Equal(t, s, "1")

	s, b = anyToString(false)
	assert.True(t, b)
	assert.Equal(t, s, "0")

	s, b = anyToString(Textarea{})
	assert.False(t, b)
	assert.Equal(t, s, "forms.Textarea")
}
