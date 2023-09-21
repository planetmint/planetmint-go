package keeper

import (
	"github.com/planetmint/planetmint-go/x/asset/types"
)

var _ types.QueryServer = Keeper{}
