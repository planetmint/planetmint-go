package keeper

import (
	"github.com/planetmint/planetmint-go/x/machine/types"
)

var _ types.QueryServer = Keeper{}
