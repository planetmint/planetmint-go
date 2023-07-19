package keeper_test

import (
	"context"
	"testing"

	keepertest "planetmint-go/testutil/keeper"
	"planetmint-go/x/machine/keeper"
	"planetmint-go/x/machine/types"

	"planetmint-go/testutil/sample"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.MachineKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}

func TestMsgServer(t *testing.T) {
	ms, ctx := setupMsgServer(t)
	require.NotNil(t, ms)
	require.NotNil(t, ctx)
}

func TestMsgServerAttestMachine(t *testing.T) {
	_, pk := sample.KeyPair()
	machine := sample.Machine(pk, pk)
	msg := types.NewMsgAttestMachine(pk, &machine)
	msgServer, ctx := setupMsgServer(t)
	res, err := msgServer.AttestMachine(ctx, msg)
	if assert.NoError(t, err) {
		assert.Equal(t, &types.MsgAttestMachineResponse{}, res)
	}
}
