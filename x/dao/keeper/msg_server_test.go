package keeper_test

import (
	"context"
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/planetmint/planetmint-go/testutil/keeper"
	"github.com/planetmint/planetmint-go/testutil/moduleobject"
	"github.com/planetmint/planetmint-go/testutil/sample"
	"github.com/planetmint/planetmint-go/x/dao/keeper"
	"github.com/planetmint/planetmint-go/x/dao/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context, *keeper.Keeper) {
	k, ctx := keepertest.DaoKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx), k
}

func TestMsgServer(t *testing.T) {
	t.Parallel()
	ms, ctx, _ := setupMsgServer(t)
	require.NotNil(t, ms)
	require.NotNil(t, ctx)
}

func TestMsgServerReportPoPResult(t *testing.T) {
	initiator := sample.Secp256k1AccAddress().String()
	challenger := sample.Secp256k1AccAddress()
	challengee := sample.Secp256k1AccAddress()
	errInvalidPopData := "Invalid pop data"

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
					Initiator:  initiator,
					Challenger: challenger.String(),
					Challengee: challengee.String(),
					Height:     1,
					Success:    true,
					Finished:   true,
				},
			},
			"",
		},
		{
			"success not set",
			types.MsgReportPopResult{
				Creator: challenger.String(),
				Challenge: &types.Challenge{
					Initiator:  initiator,
					Challenger: challenger.String(),
					Challengee: challengee.String(),
					Height:     2,
					Finished:   true,
				},
			},
			"", // no error because Go defaults bool to false
		},
		{
			"initiator not set",
			types.MsgReportPopResult{
				Creator: challenger.String(),
				Challenge: &types.Challenge{
					Challenger: challenger.String(),
					Challengee: challengee.String(),
					Height:     3,
					Success:    true,
					Finished:   true,
				},
			},
			"Initiator is not set: invalid challenge",
		},
		{
			errInvalidPopData,
			types.MsgReportPopResult{
				Creator: challenger.String(),
				Challenge: &types.Challenge{
					Initiator:  initiator,
					Challenger: challenger.String(),
					Challengee: challengee.String(),
					Height:     4,
					Success:    true,
					Finished:   true,
				},
			},
			"PoP report data does not match challenge: invalid challenge",
		},
		{
			errInvalidPopData,
			types.MsgReportPopResult{
				Creator: challenger.String(),
				Challenge: &types.Challenge{
					Initiator:  initiator,
					Challenger: challenger.String(),
					Challengee: challengee.String(),
					Height:     5,
					Success:    false,
					Finished:   false,
				},
			},
			"PoP reporter is not the challenger: invalid PoP reporter",
		},
		{
			errInvalidPopData,
			types.MsgReportPopResult{
				Creator: challenger.String(),
				Challenge: &types.Challenge{
					Initiator:  initiator,
					Challenger: challenger.String(),
					Challengee: challengee.String(),
					Height:     6,
					Success:    true,
					Finished:   true,
				},
			},
			"PoP report data does not match challenge: invalid challenge",
		},
		{
			"Non-Existing PoP",
			types.MsgReportPopResult{
				Creator: challenger.String(),
				Challenge: &types.Challenge{
					Initiator:  initiator,
					Challenger: challenger.String(),
					Challengee: challengee.String(),
					Height:     7,
					Success:    true,
					Finished:   true,
				},
			},
			"no challenge found for PoP report: invalid challenge",
		},
	}

	msgServer, ctx, k := setupMsgServer(t)
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	// set up the challenges, do not store the last challenge (special test case)
	for i := 0; i < len(testCases)-1; i++ {
		msg := testCases[i].msg
		challenge := msg.GetChallenge()
		k.StoreChallenge(sdkCtx, *challenge)
	}
	// adjust challenge 4 to satisfy the test case
	testCases[3].msg.Challenge.Challengee = testCases[3].msg.Challenge.Challenger
	testCases[4].msg.Challenge.Challenger = testCases[4].msg.Challenge.Challengee
	testCases[5].msg.Challenge.Initiator = challenger.String()

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			res, err := msgServer.ReportPopResult(ctx, &tc.msg)

			if tc.errMsg != "" {
				assert.EqualError(t, err, tc.errMsg)
			} else {
				assert.Equal(t, &types.MsgReportPopResultResponse{}, res)
			}
		})
	}
}
func TestMsgServerMintToken(t *testing.T) {
	t.Parallel()
	minter := sample.AccAddress()
	beneficiary := sample.ConstBech32Addr
	mintRequest := moduleobject.MintRequest(beneficiary, 1000, "hash")

	msg := types.NewMsgMintToken(minter, &mintRequest)
	msgServer, ctx, _ := setupMsgServer(t)
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
	t.Parallel()
	minter := sample.AccAddress()
	beneficiary := "invalid address"
	mintRequest := moduleobject.MintRequest(beneficiary, 1000, "hash")

	msg := types.NewMsgMintToken(minter, &mintRequest)
	msgServer, ctx, _ := setupMsgServer(t)
	_, err := msgServer.MintToken(ctx, msg)
	if assert.Error(t, err) {
		assert.EqualError(t, err, fmt.Sprintf("for provided address %s: invalid address", beneficiary))
	}
}
