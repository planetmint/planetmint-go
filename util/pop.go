package util

import (
	"math"

	"github.com/planetmint/planetmint-go/config"
)

const (
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
