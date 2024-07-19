package keeper_test

import (
	"testing"

	keepertest "github.com/planetmint/planetmint-go/testutil/keeper"
	"github.com/planetmint/planetmint-go/x/machine/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestGetTrustAnchorQuery(t *testing.T) {
	keeper, ctx := keepertest.MachineKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNTrustAnchor(t, keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetTrustAnchorStatusRequest
		response *types.QueryGetTrustAnchorStatusResponse
		err      error
	}{
		{
			desc:     "GetByMachineId",
			request:  &types.QueryGetTrustAnchorStatusRequest{Machineid: msgs[0].Pubkey},
			response: &types.QueryGetTrustAnchorStatusResponse{Machineid: msgs[0].Pubkey, Isactivated: false},
		},
		{
			desc:     "Not Activated",
			request:  &types.QueryGetTrustAnchorStatusRequest{Machineid: msgs[1].Pubkey},
			response: &types.QueryGetTrustAnchorStatusResponse{Machineid: msgs[1].Pubkey, Isactivated: true},
		},
		{
			desc:    "NotFound",
			request: &types.QueryGetTrustAnchorStatusRequest{Machineid: "invalid MachineID"},
			err:     status.Error(codes.NotFound, "trust anchor not found by machine id: invalid MachineID"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.GetTrustAnchorStatus(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.Equal(t, tc.response, response)
			}
		})
	}
}
