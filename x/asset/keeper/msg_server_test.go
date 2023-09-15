package keeper_test

import (
	"context"
	"encoding/hex"
	"testing"

	"planetmint-go/config"
	keepertest "planetmint-go/testutil/keeper"
	"planetmint-go/testutil/sample"
	"planetmint-go/x/asset/keeper"
	"planetmint-go/x/asset/types"

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
	ext_sk, ppk := sample.ExtendedKeyPair(config.PlmntNetParams)
	xskKey, _ := hdkeychain.NewKeyFromString(ext_sk)
	privKey, _ := xskKey.ECPrivKey()
	byte_key := privKey.Serialize()
	sk := hex.EncodeToString(byte_key)
	cid, signatureHex := sample.Asset(sk)

	msg := types.NewMsgNotarizeAsset(sk, cid, signatureHex, ppk)
	msgServer, ctx := setupMsgServer(t)
	res, err := msgServer.NotarizeAsset(ctx, msg)
	if assert.NoError(t, err) {
		assert.Equal(t, &types.MsgNotarizeAssetResponse{}, res)
	}
}

func TestMsgServerNotarizeAssetMachineNotFound(t *testing.T) {
	sk, _ := sample.KeyPair()
	msg := types.NewMsgNotarizeAsset(sk, "cid", "sign", sk)
	msgServer, ctx := setupMsgServer(t)
	_, err := msgServer.NotarizeAsset(ctx, msg)
	assert.EqualError(t, err, "machine not found")
}

func TestMsgServerNotarizeAssetInvalidAssetSignatureType(t *testing.T) {
	_, pk := sample.ExtendedKeyPair(config.PlmntNetParams)
	hex_string := hex.EncodeToString([]byte("cid"))
	msg := types.NewMsgNotarizeAsset(pk, hex_string, "sign", pk)
	msgServer, ctx := setupMsgServer(t)
	_, err := msgServer.NotarizeAsset(ctx, msg)
	assert.EqualError(t, err, "invalid signature hex string")
}

func TestMsgServerNotarizeAssetInvalidAssetSignature(t *testing.T) {
	_, pk := sample.ExtendedKeyPair(config.PlmntNetParams)
	hex_string_cid := hex.EncodeToString([]byte("cid"))
	hex_string_sid := hex.EncodeToString([]byte("sign"))
	msg := types.NewMsgNotarizeAsset(pk, hex_string_cid, hex_string_sid, pk)
	msgServer, ctx := setupMsgServer(t)
	_, err := msgServer.NotarizeAsset(ctx, msg)
	assert.EqualError(t, err, "invalid signature")
}

func TestMsgServerNotarizeAssetInvalidXPubKey(t *testing.T) {
	_, pk := sample.KeyPair()
	msg := types.NewMsgNotarizeAsset(pk, "cid", "sign", pk)
	msgServer, ctx := setupMsgServer(t)
	_, err := msgServer.NotarizeAsset(ctx, msg)
	assert.EqualError(t, err, "could not convert xpub key to hex pub key")
}
