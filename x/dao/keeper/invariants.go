package keeper

import (
	"fmt"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/planetmint/planetmint-go/config"
	"github.com/planetmint/planetmint-go/x/dao/types"
)

// RegisterInvariants registers all dao invariants
func RegisterInvariants(ir sdk.InvariantRegistry, k *Keeper) {
	ir.RegisterRoute(types.ModuleName, "mint-requests", MintRequestInvariants(k))
}

// AllInvariants runs all invariants of the dao module.
func AllInvariants(k *Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		return MintRequestInvariants(k)(ctx)
	}
}

// MintRequestInvariants checks that the total amount of PLMNT is not exceeding the amount of all MintRequests and Gentxs
func MintRequestInvariants(k *Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		cfg := config.GetConfig()
		totalAmount := math.ZeroInt()
		genAmount := math.ZeroInt()
		mintAmount := math.ZeroInt()

		// get total amount from bank all balances
		k.bankKeeper.IterateAllBalances(ctx, func(_ sdk.AccAddress, coin sdk.Coin) bool {
			if coin.Denom == cfg.TokenDenom {
				totalAmount.Add(coin.Amount)
			}
			return false
		})

		// collect bank gen state for initial amounts
		bankGenState := k.bankKeeper.ExportGenesis(ctx)
		for _, balance := range bankGenState.Balances {
			genAmount.Add(balance.Coins.AmountOf(cfg.TokenDenom))
		}

		// collect all mint request amounts
		k.IterateAllMintRequests(ctx, func(mr types.MintRequest) bool {
			mintAmount.Add(math.NewIntFromUint64(mr.Amount))
			return false
		})

		// compare total amount == inital amounts + mint request amounts
		broken := genAmount.Add(mintAmount) != totalAmount

		return sdk.FormatInvariant(types.ModuleName, "mint request and genesis amount coins", fmt.Sprintf("asd", 1)), broken
	}
}
