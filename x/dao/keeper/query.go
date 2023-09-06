package keeper

import (
	"planetmint-go/x/dao/types"
)

var _ types.QueryServer = Keeper{}
