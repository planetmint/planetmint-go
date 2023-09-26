package keeper_test

import (
	"context"
	"encoding/hex"
	"testing"

	"github.com/planetmint/planetmint-go/config"
	keepertest "github.com/planetmint/planetmint-go/testutil/keeper"
	"github.com/planetmint/planetmint-go/testutil/sample"
	"github.com/planetmint/planetmint-go/x/asset/keeper"
	"github.com/planetmint/planetmint-go/x/asset/types"

	"github.com/btcsuite/btcd/btcutil/hdkeychain"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.AssetKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}

func TestMsgServer(t *testing.T) {
	ms, ctx := setupMsgServer(t)
	require.NotNil(t, ms)
	require.NotNil(t, ctx)
}

func TestMsgServerNotarizeAsset(t *testing.T) {
	ext_sk, _ := sample.ExtendedKeyPair(config.PlmntNetParams)
	xskKey, _ := hdkeychain.NewKeyFromString(ext_sk)
	privKey, _ := xskKey.ECPrivKey()
	byte_key := privKey.Serialize()
	sk := hex.EncodeToString(byte_key)
	cid, _ := sample.Asset(sk)

	msg := types.NewMsgNotarizeAsset(sk, cid)
	msgServer, ctx := setupMsgServer(t)
	res, err := msgServer.NotarizeAsset(ctx, msg)
	if assert.NoError(t, err) {
		assert.Equal(t, &types.MsgNotarizeAssetResponse{}, res)
	}
}
