package util

import (
	"testing"

	"gotest.tools/assert"
)

func Test2FloatConvertion(t *testing.T) {
	t.Parallel()
	var expectedValue uint64 = 99869000000
	value := RDDLToken2Uint(998.69)
	assert.Equal(t, expectedValue, value)
}

func Test2UintConvertion(t *testing.T) {
	t.Parallel()
	expectedValue := 998.69
	value := RDDLToken2Float(99869000000)
	assert.Equal(t, expectedValue, value)
}

func TestStringToFloat(t *testing.T) {
	t.Parallel()
	expectedValue := 998.69
	value, err := RDDLTokenStringToFloat("998.69")
	assert.Equal(t, expectedValue, value)
	assert.Equal(t, nil, err)
}

func TestStringToUint(t *testing.T) {
	t.Parallel()
	var expectedValue uint64 = 99869000000
	value, err := RDDLTokenStringToUint("998.69")
	assert.Equal(t, expectedValue, value)
	assert.Equal(t, nil, err)
}

func TestAddPrecisionLongerThan8(t *testing.T) {
	t.Parallel()

	var input uint64 = 99869000000
	expectedValue := "998.69000000"
	rddlTokenString := UintValueToRDDLTokenString(input)
	assert.Equal(t, expectedValue, rddlTokenString)
}

func TestAddPrecisionEqual8(t *testing.T) {
	t.Parallel()

	var input uint64 = 69000000
	expectedValue := "0.69000000"
	rddlTokenString := UintValueToRDDLTokenString(input)
	assert.Equal(t, expectedValue, rddlTokenString)
}

func TestAddPrecisionShorterThan8(t *testing.T) {
	t.Parallel()

	var input uint64 = 9000000
	expectedValue := "0.09000000"
	rddlTokenString := UintValueToRDDLTokenString(input)
	assert.Equal(t, expectedValue, rddlTokenString)
}
