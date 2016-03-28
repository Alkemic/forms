package forms

import (
	"github.com/stretchr/testify/assert"
	"testing"
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
