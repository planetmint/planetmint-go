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

	"planetmint-go/config"
	"planetmint-go/x/dao/types"
)

type (
	Keeper struct {
		cdc        codec.BinaryCodec
		storeKey   storetypes.StoreKey
		memKey     storetypes.StoreKey
		paramstore paramtypes.Subspace

		bankKeeper    types.BankKeeper
		accountKeeper types.AccountKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey storetypes.StoreKey,
	ps paramtypes.Subspace,

	bankKeeper types.BankKeeper,
	accountKeeper types.AccountKeeper,
) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{
		cdc:        cdc,
		storeKey:   storeKey,
		memKey:     memKey,
		paramstore: ps,

		bankKeeper:    bankKeeper,
		accountKeeper: accountKeeper,
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
		decTotalAmountToDistribute := sdk.NewDecFromInt(coinToDistribute.Amount)
		decTotalStake := sdk.NewDecFromInt(totalStake)
		for addr, stake := range balances {
			decStake := sdk.NewDecFromInt(stake)
			share := decStake.Quo(decTotalStake)
			claim := decTotalAmountToDistribute.Mul(share)
			if claim.GTE(sdk.OneDec()) {
				intClaim := claim.TruncateInt()
				coinClaim := sdk.NewCoin(conf.FeeDenom, intClaim)
				accAddress, err := sdk.AccAddressFromBech32(addr)
				if err != nil {
					panic(err)
				}
				if !k.bankKeeper.BlockedAddr(accAddress) {
					err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, disttypes.ModuleName, accAddress, sdk.NewCoins(coinClaim))
					if err != nil {
						panic(err)
					}
				}
			}
		}
	}
}
