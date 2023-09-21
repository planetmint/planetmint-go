package keeper

import (
	"github.com/planetmint/planetmint-go/x/dao/types"
)

var _ types.QueryServer = Keeper{}
