package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/planetmint/planetmint-go/x/dao/types"
)

func (k Keeper) StoreMintRequest(ctx sdk.Context, mintRequest types.MintRequest) {
	hashStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MintRequestHashKey))
	hashAppendValue := k.cdc.MustMarshal(&mintRequest)
	hashStore.Set(getMintRequestKeyBytes(mintRequest.LiquidTxHash), hashAppendValue)

	addressStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MintRequestAddressKey))
	mintRequests, _ := k.GetMintRequestsByAddress(ctx, mintRequest.Beneficiary)
	mintRequests.Requests = append(mintRequests.Requests, &mintRequest)
	addressAppendValue := k.cdc.MustMarshal(&mintRequests)
	addressStore.Set(getMintRequestKeyBytes(mintRequest.Beneficiary), addressAppendValue)
}

func (k Keeper) GetMintRequestsByAddress(ctx sdk.Context, address string) (val types.MintRequests, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MintRequestAddressKey))
	mintRequests := store.Get(getMintRequestKeyBytes(address))
	if mintRequests == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(mintRequests, &val)
	return val, true
}

func (k Keeper) GetMintRequestByHash(ctx sdk.Context, hash string) (val types.MintRequest, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MintRequestHashKey))
	mintRequest := store.Get(getMintRequestKeyBytes(hash))
	if mintRequest == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(mintRequest, &val)
	return val, true
}

func getMintRequestKeyBytes(key string) []byte {
	bz := []byte(key)
	return bz
}
