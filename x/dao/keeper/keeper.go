package keeper

import (
	"fmt"

	db "github.com/cometbft/cometbft-db"
	"github.com/cometbft/cometbft/libs/log"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"

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
		accountKeeperKey      storetypes.StoreKey
		paramstore            paramtypes.Subspace

		bankKeeper    types.BankKeeper
		accountKeeper types.AccountKeeper
		machineKeeper types.MachineKeeper
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
	accountKeeperKey storetypes.StoreKey,
	ps paramtypes.Subspace,

	bankKeeper types.BankKeeper,
	accountKeeper types.AccountKeeper,
	machineKeeper types.MachineKeeper,
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
		accountKeeperKey:      accountKeeperKey,
		paramstore:            ps,

		bankKeeper:    bankKeeper,
		accountKeeper: accountKeeper,
		machineKeeper: machineKeeper,
		authority:     authority,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) SelectPopParticipants(ctx sdk.Context) (challenger string, challengee string) {
	var startAccountNumber uint64
	lastPopHeight := ctx.BlockHeight() - k.GetParams(ctx).PopEpochs
	lastPop, found := k.LookupChallenge(ctx, lastPopHeight)
	if lastPopHeight > 0 && found && lastPop.Challengee != "" {
		lastAccountAddr := sdk.MustAccAddressFromBech32(lastPop.Challengee)
		lastAccount := k.accountKeeper.GetAccount(ctx, lastAccountAddr)
		startAccountNumber = lastAccount.GetAccountNumber() + 1
	}

	var participants []sdk.AccAddress
	k.iterateAccountsForMachines(ctx, startAccountNumber, &participants, true)
	if len(participants) != 2 {
		k.iterateAccountsForMachines(ctx, startAccountNumber, &participants, false)
	}

	// Not enough participants
	if len(participants) != 2 {
		return
	}

	challenger = participants[0].String()
	challengee = participants[1].String()

	return
}

func (k Keeper) iterateAccountsForMachines(ctx sdk.Context, start uint64, participants *[]sdk.AccAddress, iterateFromStart bool) {
	store := ctx.KVStore(k.accountKeeperKey)
	accountStore := prefix.NewStore(store, authtypes.AccountNumberStoreKeyPrefix)
	var iterator db.Iterator
	if iterateFromStart {
		iterator = accountStore.Iterator(sdk.Uint64ToBigEndian(start), nil)
	} else {
		iterator = accountStore.Iterator(nil, sdk.Uint64ToBigEndian(start))
	}
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		participant := sdk.AccAddress(iterator.Value())
		_, found := k.machineKeeper.GetMachineIndexByAddress(ctx, participant.String())
		if found {
			available, err := util.GetMqttStatusOfParticipant(participant.String(), k.GetParams(ctx).MqttResponseTimeout)
			if err == nil && available {
				*participants = append(*participants, participant)
			}
		}

		if len(*participants) == 2 {
			return
		}
	}
}
