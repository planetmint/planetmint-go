package keeper

import (
	"github.com/planetmint/planetmint-go/util"
	"github.com/planetmint/planetmint-go/x/asset/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) StoreAsset(ctx sdk.Context, msg types.MsgNotarizeAsset) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AssetKey))
	store.Set(util.SerializeString(msg.GetCid()), []byte(msg.GetCreator()))
}

func (k Keeper) GetAsset(ctx sdk.Context, cid string) (msg types.MsgNotarizeAsset, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AssetKey))
	creatorBytes := store.Get(util.SerializeString(cid))
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
			break
		}
	}
	return cids, len(cids) > 0
}
