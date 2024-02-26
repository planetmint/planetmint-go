package util

import (
	"testing"

	"gotest.tools/assert"
)

const (
	rddlTokenAmount string = "998.85844748"
)

func Test2FloatConvertion(t *testing.T) {
	t.Parallel()
	var expectedValue uint64 = 99885844748
	value := RDDLToken2Uint(998.85844748)
	assert.Equal(t, expectedValue, value)
}

func Test2UintConvertion(t *testing.T) {
	t.Parallel()
	expectedValue := 998.85844748
	value := RDDLToken2Float(99885844748)
	assert.Equal(t, expectedValue, value)
}

func TestStringToFloat(t *testing.T) {
	t.Parallel()
	expectedValue := 998.85844748
	value, err := RDDLTokenStringToFloat(rddlTokenAmount)
	assert.Equal(t, expectedValue, value)
	assert.Equal(t, nil, err)
}

func TestStringToUint(t *testing.T) {
	t.Parallel()
	var expectedValue uint64 = 99885844748
	value, err := RDDLTokenStringToUint(rddlTokenAmount)
	assert.Equal(t, expectedValue, value)
	assert.Equal(t, nil, err)
}

func TestAddPrecisionLongerThan8(t *testing.T) {
	t.Parallel()

	var input uint64 = 99885844748
	expectedValue := rddlTokenAmount
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
