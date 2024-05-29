package keeper

import (
	"github.com/planetmint/planetmint-go/util"
	"github.com/planetmint/planetmint-go/x/asset/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) setAddresAssetCount(ctx sdk.Context, address string, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AssetKey))
	store.Set(types.AddressCountKey(address), util.SerializeUint64(count))
}

func (k Keeper) GetAddressAssetCount(ctx sdk.Context, address string) (count uint64, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AssetKey))
	countBytes := store.Get(types.AddressCountKey(address))
	if countBytes == nil {
		return 0, false
	}
	return util.DeserializeUint64(countBytes), true
}

func (k Keeper) incrementAddressAssetCount(ctx sdk.Context, address string) uint64 {
	count, _ := k.GetAddressAssetCount(ctx, address)
	k.setAddresAssetCount(ctx, address, count+1)
	return count + 1
}

func (k Keeper) StoreAddressAsset(ctx sdk.Context, msg types.MsgNotarizeAsset) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.AddressPrefix(msg.GetCreator()))
	count := k.incrementAddressAssetCount(ctx, msg.GetCreator())
	store.Set(util.SerializeUint64(count), []byte(msg.GetCid()))
}

func (k Keeper) StoreAsset(ctx sdk.Context, msg types.MsgNotarizeAsset) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AssetKey))
	store.Set(util.SerializeString(msg.GetCid()), []byte(msg.GetCreator()))
	k.StoreAddressAsset(ctx, msg)
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

func (k Keeper) GetAssetByAddressAndID(ctx sdk.Context, address string, id uint64) (cid string, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.AddressPrefix(address))
	cidBytes := store.Get(util.SerializeUint64(id))
	if cidBytes == nil {
		return cid, false
	}
	return string(cidBytes), true
}

func (k Keeper) GetAssetsByAddress(ctx sdk.Context, address string, start []byte, end []byte) (cids []string, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.AddressPrefix(address))

	iterator := store.ReverseIterator(start, end)
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		cidBytes := iterator.Value()
		cids = append(cids, string(cidBytes))
	}
	return cids, len(cids) > 0
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
