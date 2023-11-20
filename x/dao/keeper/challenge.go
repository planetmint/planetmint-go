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

func (k Keeper) GetChallenge(ctx sdk.Context, height uint64) (val types.Challenge, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ChallengeKey))
	challenge := store.Get(getChallengeBytes(height))
	if challenge == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(challenge, &val)
	return val, true
}

func (k Keeper) GetChallengeRange(ctx sdk.Context, start uint64) (val []types.Challenge) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ChallengeKey))
	iterator := store.Iterator(getChallengeBytes(start), nil)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var challenge types.Challenge
		if err := challenge.Unmarshal(iterator.Value()); err != nil {
			// TODO: handle better than panicing
			panic(err)
		}
		val = append(val, challenge)
	}
	return val
}

func getChallengeBytes(height uint64) []byte {
	// Adding 1 because 0 will be interpreted as nil, which is an invalid key
	return big.NewInt(int64(height + 1)).Bytes()
}
