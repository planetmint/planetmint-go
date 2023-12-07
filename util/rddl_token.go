package util

import (
	"math"
	"strconv"

	"github.com/planetmint/planetmint-go/config"
)

const (
	factor       = 100000000.0
	PopsPerCycle = 1051200.0
)

func GetPopNumber(blockHeight int64) float64 {
	return float64(blockHeight) / float64(config.GetConfig().PopEpochs)
}

func GetPopReward(blockHeight int64) (total uint64, challenger uint64, challengee uint64) {
	PopNumber := GetPopNumber(blockHeight)
	exactCycleID := PopNumber / PopsPerCycle

	switch cycleID := math.Floor(exactCycleID); cycleID {
	case 0:
		return 7990867578, 1997716894, 5993150684
	case 1:
		return 3995433789, 998858447, 2996575342
	case 2:
		return 1997716894, 499429223, 1498287671
	case 3:
		return 998858446, 249714611, 749143835
	case 4:
		return 499429222, 124857305, 374571917
	default:
		return 0, 0, 0
	}
}

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
