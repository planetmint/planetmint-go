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

func TestMsgServerAttestMachineInvalidLiquidKey(t *testing.T) {
	_, pk := sample.KeyPair()
	machine := sample.Machine(pk, pk)
	machine.IssuerLiquid = "invalidkey"
	msg := types.NewMsgAttestMachine(pk, &machine)
	msgServer, ctx := setupMsgServer(t)
	_, err := msgServer.AttestMachine(ctx, msg)
	assert.EqualError(t, err, "invalid liquid key")
}

func TestMsgServerRegisterTrustAnchor(t *testing.T) {
	_, pk := sample.KeyPair()
	ta := sample.TrustAnchor()
	msg := types.NewMsgRegisterTrustAnchor(pk, &ta)
	msgServer, ctx := setupMsgServer(t)
	res, err := msgServer.RegisterTrustAnchor(ctx, msg)
	if assert.NoError(t, err) {
		assert.Equal(t, &types.MsgRegisterTrustAnchorResponse{}, res)
	}
}

func TestMsgServerRegisterTrustAnchorTwice(t *testing.T) {
	_, pk := sample.KeyPair()
	ta := sample.TrustAnchor()
	msg := types.NewMsgRegisterTrustAnchor(pk, &ta)
	msgServer, ctx := setupMsgServer(t)
	res, err := msgServer.RegisterTrustAnchor(ctx, msg)
	if assert.NoError(t, err) {
		assert.Equal(t, &types.MsgRegisterTrustAnchorResponse{}, res)
	}
	_, err = msgServer.RegisterTrustAnchor(ctx, msg)
	assert.EqualError(t, err, "trust anchor is already registered")
}

func TestMsgServerRegisterTrustAnchorInvalidPubkey(t *testing.T) {
	_, pk := sample.KeyPair()
	ta := types.TrustAnchor{
		Pubkey: "invalidpublickey",
	}
	msg := types.NewMsgRegisterTrustAnchor(pk, &ta)
	msgServer, ctx := setupMsgServer(t)
	_, err := msgServer.RegisterTrustAnchor(ctx, msg)
	assert.EqualError(t, err, "invalid trust anchor pubkey")
}
