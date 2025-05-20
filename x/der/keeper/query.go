package keeper

import (
	"github.com/planetmint/planetmint-go/x/der/types"
)

var _ types.QueryServer = Keeper{}
