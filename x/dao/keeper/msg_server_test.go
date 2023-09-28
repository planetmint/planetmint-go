package keeper_test

import (
	"context"
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
