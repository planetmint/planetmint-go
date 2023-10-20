package keeper

import (
	"github.com/planetmint/planetmint-go/x/asset/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) StoreAsset(ctx sdk.Context, msg types.MsgNotarizeAsset) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AssetKey))
	store.Set(GetAssetCIDBytes(msg.GetCid()), []byte(msg.GetCreator()))
}

func (k Keeper) GetAsset(ctx sdk.Context, cid string) (msg types.MsgNotarizeAsset, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AssetKey))
	creatorBytes := store.Get(GetAssetCIDBytes(cid))
	if creatorBytes == nil {
		return msg, false
	}
	msg.Cid = cid
	msg.Creator = string(creatorBytes)
	return msg, true
}

func (k Keeper) GetCidsByAddress(ctx sdk.Context, address string) (cids []string, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AssetKey))

	reverseIterator := store.ReverseIterator(nil, nil)
	defer reverseIterator.Close()
	for ; reverseIterator.Valid(); reverseIterator.Next() {
		addressBytes := reverseIterator.Value()
		cidBytes := reverseIterator.Key()

		if string(addressBytes) == address {
			cids = append(cids, string(cidBytes))
		}
	}
	return cids, len(cids) > 0
}

func GetAssetCIDBytes(cid string) []byte {
	bz := []byte(cid)
	return bz
}
