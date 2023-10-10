package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

// AccountKeeper defines the expected account keeper used for simulations (noalias)
type AccountKeeper interface {
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) types.AccountI
	GetModuleAddress(module string) sdk.AccAddress
	IterateAccounts(sdk.Context, func(types.AccountI) bool)
	// Methods imported from account should be defined here
}

// BankKeeper defines the expected interface needed to retrieve account balances.
type BankKeeper interface {
	SpendableCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	SendCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
	BlockedAddr(addr sdk.AccAddress) bool
	MintCoins(ctx sdk.Context, moduleName string, amt sdk.Coins) error
	ExportGenesis(ctx sdk.Context) *banktypes.GenesisState
	IterateAllBalances(ctx sdk.Context, cb func(address sdk.AccAddress, coin sdk.Coin) (stop bool))
	// Methods imported from bank should be defined here
}
