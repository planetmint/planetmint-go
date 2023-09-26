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
	creator_bytes := store.Get(GetAssetCIDBytes(cid))
	if creator_bytes == nil {
		return msg, false
	}
	msg.Cid = cid
	msg.Creator = string(creator_bytes)
	return msg, true
}

func GetAssetCIDBytes(cid string) []byte {
	bz := []byte(cid)
	return bz
}
