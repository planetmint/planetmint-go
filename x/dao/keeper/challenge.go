package keeper

import (
	db "github.com/cometbft/cometbft-db"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/planetmint/planetmint-go/util"
	"github.com/planetmint/planetmint-go/x/dao/types"
)

func (k Keeper) StoreChallenge(ctx sdk.Context, challenge types.Challenge) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ChallengeKey))
	appendValue := k.cdc.MustMarshal(&challenge)
	store.Set(util.SerializeInt64(challenge.Height), appendValue)
}

func (k Keeper) LookupChallenge(ctx sdk.Context, height int64) (val types.Challenge, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ChallengeKey))
	challenge := store.Get(util.SerializeInt64(height))
	if challenge == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(challenge, &val)
	return val, true
}

func (k Keeper) GetChallengeRangeToEnd(ctx sdk.Context, start int64) (val []types.Challenge, err error) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ChallengeKey))
	// adding 1 to end because end is exclusive on store.Iterator
	iterator := store.Iterator(util.SerializeInt64(start), nil)
	defer iterator.Close()

	return k.getChallengeRangeFromStore(ctx, iterator)
}

func (k Keeper) GetChallengeRange(ctx sdk.Context, start int64, end int64) (val []types.Challenge, err error) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ChallengeKey))
	// adding 1 to end because end is exclusive on store.Iterator
	iterator := store.Iterator(util.SerializeInt64(start), util.SerializeInt64(end+1))
	defer iterator.Close()

	return k.getChallengeRangeFromStore(ctx, iterator)
}

func (k Keeper) getChallengeRangeFromStore(ctx sdk.Context, iterator db.Iterator) (val []types.Challenge, err error) {
	for ; iterator.Valid(); iterator.Next() {
		var challenge types.Challenge
		if err := challenge.Unmarshal(iterator.Value()); err != nil {
			util.GetAppLogger().Error(ctx, err, "unable to unmarshal challenge")
			return nil, err // or continue TODO make decision
		}
		val = append(val, challenge)
	}
	return val, nil
}

func (k Keeper) GetChallenges(ctx sdk.Context) (challenges []types.Challenge, err error) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, nil)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var event types.Challenge
		if err = event.Unmarshal(iterator.Value()); err != nil {
			util.GetAppLogger().Error(ctx, err, "unable to unmarshal challenge")
			return nil, err // or continue TODO make decision
		}
		challenges = append(challenges, event)
	}
	return
}

func (k Keeper) StoreChallangeInitiatorReward(ctx sdk.Context, height int64, amount uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PoPInitiatorReward))
	appendValue := util.SerializeUint64(amount)
	store.Set(util.SerializeInt64(height), appendValue)
}

func (k Keeper) getChallengeInitiatorReward(ctx sdk.Context, height int64) (amount uint64, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PoPInitiatorReward))
	amountBytes := store.Get(util.SerializeInt64(height))
	if amountBytes == nil {
		return 0, false
	}
	amount = util.DeserializeUint64(amountBytes)
	return amount, true
}
