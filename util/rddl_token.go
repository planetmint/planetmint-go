package util

import "strconv"

var factor float64 = 100000000.0

func RDDLToken2Uint(amount float64) uint64 {
	return uint64(amount * factor)
}

func RDDLToken2Float(amount uint64) float64 {
	return float64(amount) / factor
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
