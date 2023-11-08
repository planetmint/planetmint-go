package util

import (
	"testing"

	"gotest.tools/assert"
)

func Test2FloatConvertion(t *testing.T) {
	var expectedValue uint64 = 99869000000
	value := RDDLToken2Uint(998.69)
	assert.Equal(t, expectedValue, value)
}

func Test2UintConvertion(t *testing.T) {
	var expectedValue float64 = 998.69
	value := RDDLToken2Float(99869000000)
	assert.Equal(t, expectedValue, value)
}

func TestStringToFloat(t *testing.T) {
	var expectedValue float64 = 998.69
	value, err := RDDLTokenStringToFloat("998.69")
	assert.Equal(t, expectedValue, value)
	assert.Equal(t, nil, err)
}

func TestStringToUint(t *testing.T) {
	var expectedValue uint64 = 99869000000
	value, err := RDDLTokenStringToUint("998.69")
	assert.Equal(t, expectedValue, value)
	assert.Equal(t, nil, err)
}
