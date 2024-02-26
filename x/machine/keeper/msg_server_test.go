package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/planetmint/planetmint-go/testutil/keeper"
	"github.com/planetmint/planetmint-go/testutil/moduleobject"
	"github.com/planetmint/planetmint-go/x/machine/keeper"
	"github.com/planetmint/planetmint-go/x/machine/types"

	"github.com/planetmint/planetmint-go/testutil/sample"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.MachineKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}

func TestMsgServer(t *testing.T) {
	t.Parallel()
	ms, ctx := setupMsgServer(t)
	require.NotNil(t, ms)
	require.NotNil(t, ctx)
}

func TestMsgServerAttestMachine(t *testing.T) {
	t.Parallel()
	sk, pk := sample.KeyPair()
	ta := moduleobject.TrustAnchor(pk)
	taMsg := types.NewMsgRegisterTrustAnchor(pk, &ta)
	machine := moduleobject.Machine(pk, pk, sk, "")
	msg := types.NewMsgAttestMachine(pk, &machine)
	msgServer, ctx := setupMsgServer(t)
	_, err := msgServer.RegisterTrustAnchor(ctx, taMsg)
	assert.NoError(t, err)
	res, err := msgServer.AttestMachine(ctx, msg)
	if assert.NoError(t, err) {
		assert.Equal(t, &types.MsgAttestMachineResponse{}, res)
	}
}

func TestMsgServerAttestMachineInvalidLiquidKey(t *testing.T) {
	t.Parallel()
	sk, pk := sample.KeyPair()
	ta := moduleobject.TrustAnchor(pk)
	taMsg := types.NewMsgRegisterTrustAnchor(pk, &ta)
	machine := moduleobject.Machine(pk, pk, sk, "")
	machine.IssuerLiquid = "invalidkey"
	msg := types.NewMsgAttestMachine(pk, &machine)
	msgServer, ctx := setupMsgServer(t)
	_, err := msgServer.RegisterTrustAnchor(ctx, taMsg)
	assert.NoError(t, err)
	_, err = msgServer.AttestMachine(ctx, msg)
	assert.EqualError(t, err, "liquid: invalid key")
}

func TestMsgServerRegisterTrustAnchor(t *testing.T) {
	t.Parallel()
	_, pk := sample.KeyPair()
	ta := moduleobject.TrustAnchor(pk)
	msg := types.NewMsgRegisterTrustAnchor(pk, &ta)
	msgServer, ctx := setupMsgServer(t)
	res, err := msgServer.RegisterTrustAnchor(ctx, msg)
	if assert.NoError(t, err) {
		assert.Equal(t, &types.MsgRegisterTrustAnchorResponse{}, res)
	}
}

func TestMsgServerRegisterTrustAnchorTwice(t *testing.T) {
	t.Parallel()
	_, pk := sample.KeyPair()
	ta := moduleobject.TrustAnchor(pk)
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
	t.Parallel()
	_, pk := sample.KeyPair()
	ta := types.TrustAnchor{
		Pubkey: "invalidpublickey",
	}
	msg := types.NewMsgRegisterTrustAnchor(pk, &ta)
	msgServer, ctx := setupMsgServer(t)
	_, err := msgServer.RegisterTrustAnchor(ctx, msg)
	assert.EqualError(t, err, "invalid trust anchor pubkey")
}
