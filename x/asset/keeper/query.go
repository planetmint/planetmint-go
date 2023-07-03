package keeper

import (
	"planetmint-go/x/asset/types"
)

var _ types.QueryServer = Keeper{}
