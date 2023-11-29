package keeper

import (
	"math/big"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/planetmint/planetmint-go/x/dao/types"
)

func (k Keeper) StoreChallenge(ctx sdk.Context, challenge types.Challenge) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ChallengeKey))
	appendValue := k.cdc.MustMarshal(&challenge)
	store.Set(getChallengeBytes(challenge.Height), appendValue)
}

func (k Keeper) LookupChallenge(ctx sdk.Context, height int64) (val types.Challenge, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ChallengeKey))
	challenge := store.Get(getChallengeBytes(height))
	if challenge == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(challenge, &val)
	return val, true
}

func (k Keeper) GetChallengeRange(ctx sdk.Context, start int64, end int64) (val []types.Challenge, err error) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ChallengeKey))
	// adding 1 to end because end is exclusive on store.Iterator
	iterator := store.Iterator(getChallengeBytes(start), getChallengeBytes(end+1))
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var challenge types.Challenge
		if err := challenge.Unmarshal(iterator.Value()); err != nil {
			return nil, err
		}
		val = append(val, challenge)
	}
	return val, nil
}

func getChallengeBytes(height int64) []byte {
	// Adding 1 because 0 will be interpreted as nil, which is an invalid key
	return big.NewInt(height + 1).Bytes()
}
