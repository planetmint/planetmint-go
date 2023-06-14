package keeper

import (
	"planetmint-go/x/machine/types"
)

var _ types.QueryServer = Keeper{}
