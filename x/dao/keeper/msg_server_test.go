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

func TestMsgServerReportPoPResult(t *testing.T) {
	initiator := sample.Secp256k1AccAddress()
	challenger := sample.Secp256k1AccAddress()
	challengee := sample.Secp256k1AccAddress()
	description := "sample text"

	testCases := []struct {
		name   string
		msg    types.MsgReportPopResult
		errMsg string
	}{
		{
			"report pop result",
			types.MsgReportPopResult{
				Creator: challenger.String(),
				Challenge: &types.Challenge{
					Initiator:   initiator.String(),
					Challenger:  challenger.String(),
					Challengee:  challengee.String(),
					Height:      1,
					Description: description,
					Success:     true,
				},
			},
			"",
		},
		{
			"success not set",
			types.MsgReportPopResult{
				Creator: challenger.String(),
				Challenge: &types.Challenge{
					Initiator:   initiator.String(),
					Challenger:  challenger.String(),
					Challengee:  challengee.String(),
					Height:      1,
					Description: description,
				},
			},
			"", // no error because Go defaults bool to false
		},
		{
			"initiator not set",
			types.MsgReportPopResult{
				Creator: challenger.String(),
				Challenge: &types.Challenge{
					Challenger:  challenger.String(),
					Challengee:  challengee.String(),
					Height:      1,
					Description: description,
					Success:     true,
				},
			},
			"Initiator is not set: invalid challenge",
		},
	}

	msgServer, ctx := setupMsgServer(t)

	for _, tc := range testCases {
		res, err := msgServer.ReportPopResult(ctx, &tc.msg)

		if tc.errMsg != "" {
			assert.EqualError(t, err, tc.errMsg)
		} else {
			assert.Equal(t, &types.MsgReportPopResultResponse{}, res)
		}
	}
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
