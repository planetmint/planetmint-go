package keeper_test

import (
	"context"
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/planetmint/planetmint-go/testutil/keeper"
	"github.com/planetmint/planetmint-go/testutil/sample"
	"github.com/planetmint/planetmint-go/x/dao/keeper"
	"github.com/planetmint/planetmint-go/x/dao/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.DaoKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}

func TestMsgServer(t *testing.T) {
	ms, ctx := setupMsgServer(t)
	require.NotNil(t, ms)
	require.NotNil(t, ctx)
}

func TestMsgServerMintToken(t *testing.T) {
	minter := sample.AccAddress()
	beneficiary := sample.ConstBech32Addr
	mintRequest := sample.MintRequest(beneficiary, 1000, "hash")

	msg := types.NewMsgMintToken(minter, &mintRequest)
	msgServer, ctx := setupMsgServer(t)
	res, err := msgServer.MintToken(ctx, msg)
	if assert.NoError(t, err) {
		assert.Equal(t, &types.MsgMintTokenResponse{}, res)
	}

	// should throw error because hash has already been used
	_, err = msgServer.MintToken(ctx, msg)
	if assert.Error(t, err) {
		assert.EqualError(t, err, fmt.Sprintf("liquid tx hash %s has already been minted: already minted", "hash"))
	}
}

func TestMsgServerMintTokenInvalidAddress(t *testing.T) {
	minter := sample.AccAddress()
	beneficiary := "invalid address"
	mintRequest := sample.MintRequest(beneficiary, 1000, "hash")

	msg := types.NewMsgMintToken(minter, &mintRequest)
	msgServer, ctx := setupMsgServer(t)
	_, err := msgServer.MintToken(ctx, msg)
	if assert.Error(t, err) {
		assert.EqualError(t, err, fmt.Sprintf("for provided address %s: invalid address", beneficiary))
	}
}
