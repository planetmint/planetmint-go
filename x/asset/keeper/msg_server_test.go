package keeper_test

import (
	"context"
	"testing"

	keepertest "planetmint-go/testutil/keeper"
	"planetmint-go/testutil/sample"
	"planetmint-go/x/asset/keeper"
	"planetmint-go/x/asset/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.AssetKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}

func TestMsgServerNotarizeAsset(t *testing.T) {
	_, pk := sample.KeyPair()
	msg := types.NewMsgNotarizeAsset(pk, "cid", "sign", "pubkey")
	msgServer, ctx := setupMsgServer(t)
	res, err := msgServer.NotarizeAsset(ctx, msg)
	if assert.NoError(t, err) {
		assert.Equal(t, &types.MsgNotarizeAssetResponse{}, res)
	}
}

func TestMsgServerNotarizeAssetMachineNotFound(t *testing.T) {
	sk, _ := sample.KeyPair()
	msg := types.NewMsgNotarizeAsset(sk, "cid", "sign", "pubkey")
	msgServer, ctx := setupMsgServer(t)
	_, err := msgServer.NotarizeAsset(ctx, msg)
	assert.EqualError(t, err, "machine not found")
}

func TestMsgServerNotarizeAssetInvalidAsset(t *testing.T) {
	_, pk := sample.KeyPair()
	msg := types.NewMsgNotarizeAsset(pk, "cid", "sign", "pubkey")
	msgServer, ctx := setupMsgServer(t)
	_, err := msgServer.NotarizeAsset(ctx, msg)
	assert.EqualError(t, err, "invalid signature")
}
