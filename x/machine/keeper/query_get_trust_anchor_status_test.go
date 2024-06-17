package keeper_test

import (
	"testing"

	keepertest "github.com/planetmint/planetmint-go/testutil/keeper"
	"github.com/planetmint/planetmint-go/x/machine/types"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestGetTrustAnchorQuery(t *testing.T) {
	keeper, ctx := keepertest.MachineKeeper(t)
	msgs := createNTrustAnchor(t, keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetTrustAnchorStatusRequest
		response *types.QueryGetTrustAnchorStatusResponse
		err      error
	}{
		{
			desc:     "GetByMachineId",
			request:  &types.QueryGetTrustAnchorStatusRequest{MachineId: msgs[0].Pubkey},
			response: &types.QueryGetTrustAnchorStatusResponse{MachineId: msgs[0].Pubkey, IsActivated: false},
		},
		{
			desc:     "Not Activated",
			request:  &types.QueryGetTrustAnchorStatusRequest{MachineId: msgs[1].Pubkey},
			response: &types.QueryGetTrustAnchorStatusResponse{MachineId: msgs[1].Pubkey, IsActivated: true},
		},
		{
			desc:    "NotFound",
			request: &types.QueryGetTrustAnchorStatusRequest{MachineId: "invalid MachineID"},
			err:     status.Error(codes.NotFound, "trust anchor not found"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.GetTrustAnchorStatus(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.Equal(t, tc.response, response)
			}
		})
	}
}
