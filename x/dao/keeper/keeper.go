package keeper

import (
	"fmt"

	"cosmossdk.io/math"
	"github.com/cometbft/cometbft/libs/log"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	disttypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"

	"github.com/planetmint/planetmint-go/config"
	"github.com/planetmint/planetmint-go/util"
	"github.com/planetmint/planetmint-go/x/dao/types"
)

type (
	Keeper struct {
		cdc                   codec.BinaryCodec
		storeKey              storetypes.StoreKey
		memKey                storetypes.StoreKey
		challengeKey          storetypes.StoreKey
		mintRequestHashKey    storetypes.StoreKey
		mintRequestAddressKey storetypes.StoreKey
		paramstore            paramtypes.Subspace

		bankKeeper    types.BankKeeper
		accountKeeper types.AccountKeeper
		authority     string
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey storetypes.StoreKey,
	challengeKey storetypes.StoreKey,
	mintRequestHashKey storetypes.StoreKey,
	mintRequestAddressKey storetypes.StoreKey,
	ps paramtypes.Subspace,

	bankKeeper types.BankKeeper,
	accountKeeper types.AccountKeeper,
	authority string,
) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{
		cdc:                   cdc,
		storeKey:              storeKey,
		memKey:                memKey,
		challengeKey:          challengeKey,
		mintRequestHashKey:    mintRequestHashKey,
		mintRequestAddressKey: mintRequestAddressKey,
		paramstore:            ps,

		bankKeeper:    bankKeeper,
		accountKeeper: accountKeeper,
		authority:     authority,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) DistributeCollectedFees(ctx sdk.Context) {
	ctx = sdk.UnwrapSDKContext(ctx)
	conf := config.GetConfig()

	balances := make(map[string]math.Int)
	totalStake := math.ZeroInt()
	k.accountKeeper.IterateAccounts(ctx, func(acc authtypes.AccountI) bool {
		addr := acc.GetAddress()
		balance := k.bankKeeper.SpendableCoins(ctx, addr)
		found, stake := balance.Find(conf.StakeDenom)
		if found {
			totalStake = totalStake.Add(stake.Amount)
			balances[addr.String()] = stake.Amount
		}
		return false
	})

	distAddr := k.accountKeeper.GetModuleAddress(disttypes.ModuleName)
	distSpendableCoins := k.bankKeeper.SpendableCoins(ctx, distAddr)
	found, coinToDistribute := distSpendableCoins.Find(conf.FeeDenom)

	if found {
		err := k.processBalances(ctx, balances, totalStake, coinToDistribute)
		if err != nil {
			util.GetAppLogger().Error(ctx, "Error processing balances:", err)
		}
	}
}

// Check if the address is blocked
func (k Keeper) isAddressBlocked(accAddress sdk.AccAddress) bool {
	return k.bankKeeper.BlockedAddr(accAddress)
}

// Send coins from the module to the account
func (k Keeper) sendCoinsFromModuleToAccount(ctx sdk.Context, accAddress sdk.AccAddress, coinClaim sdk.Coin) error {
	return k.bankKeeper.SendCoinsFromModuleToAccount(ctx, disttypes.ModuleName, accAddress, sdk.NewCoins(coinClaim))
}

// Calculate the claim for an address
func calculateClaimForAddress(stake math.Int, totalStake math.Int, coinToDistribute sdk.Coin) sdk.Dec {
	decTotalAmountToDistribute := sdk.NewDecFromInt(coinToDistribute.Amount)
	decTotalStake := sdk.NewDecFromInt(totalStake)
	decStake := sdk.NewDecFromInt(stake)

	share := decStake.Quo(decTotalStake)
	return decTotalAmountToDistribute.Mul(share)
}

func (k Keeper) processBalances(ctx sdk.Context, balances map[string]math.Int, totalStake math.Int, coinToDistribute sdk.Coin) error {
	conf := config.GetConfig()
	for addr, stake := range balances {
		claim := calculateClaimForAddress(stake, totalStake, coinToDistribute)

		if claim.GTE(sdk.OneDec()) {
			intClaim := claim.TruncateInt()
			coinClaim := sdk.NewCoin(conf.FeeDenom, intClaim)

			accAddress, err := sdk.AccAddressFromBech32(addr)
			if err != nil {
				return err
			}

			if !k.isAddressBlocked(accAddress) {
				err = k.sendCoinsFromModuleToAccount(ctx, accAddress, coinClaim)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}
