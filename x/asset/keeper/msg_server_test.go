package keeper_test

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"testing"

	keepertest "planetmint-go/testutil/keeper"
	"planetmint-go/testutil/sample"
	"planetmint-go/x/asset/keeper"
	"planetmint-go/x/asset/types"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
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
	sk, pk := sample.KeyPair()
	cid := "cid"

	skBytes, err := hex.DecodeString(sk)
	if err != nil {
		assert.Equal(t, true, false)
	}
	privKey := &secp256k1.PrivKey{Key: skBytes}

	cidBytes, _ := hex.DecodeString(cid)
	hash := sha256.Sum256(cidBytes)

	sign, err := privKey.Sign(hash[:])
	if err != nil {
		assert.Equal(t, true, false)
	}

	signatureHex := hex.EncodeToString(sign)

	msg := types.NewMsgNotarizeAsset(pk, cid, signatureHex, pk)
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

func TestMsgServerNotarizeAssetInvalidAsset(t *testing.T) {
	_, pk := sample.KeyPair()
	msg := types.NewMsgNotarizeAsset(pk, "cid", "sign", pk)
	msgServer, ctx := setupMsgServer(t)
	_, err := msgServer.NotarizeAsset(ctx, msg)
	assert.EqualError(t, err, "invalid signature")
}