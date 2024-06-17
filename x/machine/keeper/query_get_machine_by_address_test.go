package keeper_test

import (
	"testing"

	keepertest "github.com/planetmint/planetmint-go/testutil/keeper"
	"github.com/planetmint/planetmint-go/x/machine/types"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestGetMachineByAddress(t *testing.T) {
	keeper, ctx := keepertest.MachineKeeper(t)
	msgs := createNMachine(keeper, ctx, 1)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetMachineByAddressRequest
		response *types.QueryGetMachineByAddressResponse
		err      error
	}{
		{
			desc:     "GetMachineByAddress",
			request:  &types.QueryGetMachineByAddressRequest{Address: msgs[0].Address},
			response: &types.QueryGetMachineByAddressResponse{Machine: &msgs[0]},
		}, {
			desc:    "MachineNotFound",
			request: &types.QueryGetMachineByAddressRequest{Address: "invalid address"},
			err:     status.Error(codes.NotFound, "machine not found"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.GetMachineByAddress(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.Equal(t, tc.response, response)
			}
		})
	}
}
