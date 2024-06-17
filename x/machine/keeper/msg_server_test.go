package keeper_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/planetmint/planetmint-go/x/machine/types"

	"github.com/planetmint/planetmint-go/x/machine/keeper"

	keepertest "github.com/planetmint/planetmint-go/testutil/keeper"
	"github.com/planetmint/planetmint-go/testutil/moduleobject"
	"github.com/planetmint/planetmint-go/testutil/sample"
)

func setupMsgServer(t testing.TB) (keeper.Keeper, types.MsgServer, context.Context) {
	k, ctx := keepertest.MachineKeeper(t)
	return k, keeper.NewMsgServerImpl(k), ctx
}

func TestMsgServer(t *testing.T) {
	k, ms, ctx := setupMsgServer(t)
	require.NotNil(t, ms)
	require.NotNil(t, ctx)
	require.NotEmpty(t, k)
}

func TestMsgServerAttestMachine(t *testing.T) {
	t.Parallel()
	sk, pk := sample.KeyPair()
	ta := moduleobject.TrustAnchor(pk)
	taMsg := types.NewMsgRegisterTrustAnchor(pk, &ta)
	machine := moduleobject.Machine(pk, pk, sk, "")
	msg := types.NewMsgAttestMachine(pk, &machine)
	_, msgServer, ctx := setupMsgServer(t)
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
	_, msgServer, ctx := setupMsgServer(t)
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
	_, msgServer, ctx := setupMsgServer(t)
	res, err := msgServer.RegisterTrustAnchor(ctx, msg)
	if assert.NoError(t, err) {
		assert.Equal(t, &types.MsgRegisterTrustAnchorResponse{}, res)
	}
}

func TestMsgServerRegisterTrustAnchorUncompressedKey(t *testing.T) {
	t.Parallel()
	pk := "6003d0ab9af4ec112629195a7266a244aecf1ac7691da0084be3e7ceea2ee71571b0963fffd9c80a640317509a681ac66c2ed70ecc9f317a0d2b1a9bff94ff74"
	ta := moduleobject.TrustAnchor(pk)
	msg := types.NewMsgRegisterTrustAnchor(pk, &ta)
	_, msgServer, ctx := setupMsgServer(t)
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
	_, msgServer, ctx := setupMsgServer(t)
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
	_, msgServer, ctx := setupMsgServer(t)
	_, err := msgServer.RegisterTrustAnchor(ctx, msg)
	assert.EqualError(t, err, "invalid trust anchor pubkey")
}
