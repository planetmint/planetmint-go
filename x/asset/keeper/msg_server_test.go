package keeper_test

import (
	"context"
	"testing"

	keepertest "planetmint-go/testutil/keeper"
	sample "planetmint-go/testutil/sample"
	"planetmint-go/x/asset/keeper"
	"planetmint-go/x/asset/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.AssetKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}

func prepareMachine(t testing.TB, ctx sdk.Context) {
	// store machine to test
	machine := sample.Machine()
	k, _ := keepertest.MachineKeeper(t)
	k.StoreMachine(ctx, machine)
}

func TestMsgServerNotarizeAsset(t *testing.T) {
	// machine pubkey for now

	msg := types.NewMsgNotarizeAsset("pubkey", "cid", "sign", "pubkey")

	msgServer, ctx := setupMsgServer(t)
	prepareMachine(t, sdk.UnwrapSDKContext(ctx))
	res, err := msgServer.NotarizeAsset(ctx, msg)
	if assert.NoError(t, err) {
		assert.Equal(t, types.MsgNotarizeAssetResponse{}, res)
	}
}
