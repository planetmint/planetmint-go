package keeper_test

import (
	"context"
	"testing"

	keepertest "planetmint-go/testutil/keeper"
	"planetmint-go/testutil/sample"
	"planetmint-go/x/machine/keeper"
	"planetmint-go/x/machine/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.MachineKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}

func TestMsgServerAttestMachine(t *testing.T) {
	_, pk := sample.KeyPair()
	machine := sample.Machine(pk, pk, pk)
	msg := types.NewMsgAttestMachine(pk, &machine)
	msgServer, ctx := setupMsgServer(t)
	res, err := msgServer.AttestMachine(ctx, msg)
	if assert.NoError(t, err) {
		assert.Equal(t, &types.MsgAttestMachineResponse{}, res)
	}
}
