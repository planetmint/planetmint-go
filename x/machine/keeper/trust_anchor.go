package keeper

import (
	"errors"

	"github.com/planetmint/planetmint-go/util"
	"github.com/planetmint/planetmint-go/x/machine/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) StoreTrustAnchor(ctx sdk.Context, ta types.TrustAnchor, activated bool) error {
	store := prefix.NewStore(ctx.KVStore(k.taStoreKey), types.KeyPrefix(types.TrustAnchorKey))
	// if activated is set to true then store 1 else 0
	var appendValue []byte
	if activated {
		appendValue = []byte{1}
	} else {
		appendValue = []byte{0}
	}
	pubKeyBytes, err := util.SerializeHexString(ta.Pubkey)
	if err != nil {
		return errors.New("the given public key could not be decoded (malformed string)")
	}
	store.Set(pubKeyBytes, appendValue)
	return nil
}

func (k Keeper) GetTrustAnchor(ctx sdk.Context, pubKey string) (val types.TrustAnchor, activated bool, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.taStoreKey), types.KeyPrefix(types.TrustAnchorKey))
	pubKeyBytes, err := util.SerializeHexString(pubKey)
	if err != nil {
		return val, false, false
	}
	trustAnchorActivated := store.Get(pubKeyBytes)

	if trustAnchorActivated == nil {
		return val, false, false
	}

	// if stored byte is 1 then return activated equals true
	val.Pubkey = pubKey
	if trustAnchorActivated[0] == 1 {
		return val, true, true
	}
	return val, false, true
}
