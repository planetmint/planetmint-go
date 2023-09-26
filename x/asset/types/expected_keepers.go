package types

import (
	machinetypes "github.com/planetmint/planetmint-go/x/machine/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
)

type MachineKeeper interface {
	// Methods imported from machine should be defined here
	GetMachine(ctx sdk.Context, index machinetypes.MachineIndex) (val machinetypes.Machine, found bool)
	GetMachineIndexByPubKey(ctx sdk.Context, pubKey string) (val machinetypes.MachineIndex, found bool)
}

// AccountKeeper defines the expected account keeper used for simulations (noalias)
type AccountKeeper interface {
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) types.AccountI
	// Methods imported from account should be defined here
}

// BankKeeper defines the expected interface needed to retrieve account balances.
type BankKeeper interface {
	SpendableCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	// Methods imported from bank should be defined here
}
