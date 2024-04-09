package util

import (
	"strconv"
)

const (
	Factor = 100000000.0
)

func RDDLToken2Uint(amount float64) uint64 {
	return uint64(amount * Factor)
}

func RDDLToken2Float(amount uint64) float64 {
	return float64(amount) / Factor
}

func RDDLTokenStringToFloat(amount string) (amountFloat float64, err error) {
	amountFloat, err = strconv.ParseFloat(amount, 64)
	return amountFloat, err
}

func RDDLTokenStringToUint(amount string) (amountUint uint64, err error) {
	amountFloat, err := RDDLTokenStringToFloat(amount)
	if err == nil {
		amountUint = RDDLToken2Uint(amountFloat)
	}
	return amountUint, err
}

func addPrecision(valueString string) string {
	length := len(valueString)
	if length > 8 {
		return valueString[:length-8] + "." + valueString[length-8:]
	}

	resultString := "0."
	for i := 0; i < 8-length; i++ {
		resultString += "0"
	}
	return resultString + valueString
}

func UintValueToRDDLTokenString(value uint64) (rddlString string) {
	uint64String := strconv.FormatUint(value, 10)
	rddlString = addPrecision(uint64String)
	return
}
