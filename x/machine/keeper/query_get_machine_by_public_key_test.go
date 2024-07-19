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

func TestGetMachineByPublicKey(t *testing.T) {
	keeper, ctx := keepertest.MachineKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNMachine(keeper, ctx, 1)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetMachineByPublicKeyRequest
		response *types.QueryGetMachineByPublicKeyResponse
		err      error
	}{
		{
			desc:     "GetByMachineId",
			request:  &types.QueryGetMachineByPublicKeyRequest{PublicKey: msgs[0].MachineId},
			response: &types.QueryGetMachineByPublicKeyResponse{Machine: &msgs[0]},
		},
		{
			desc:     "GetByIssuerPlanetmint",
			request:  &types.QueryGetMachineByPublicKeyRequest{PublicKey: msgs[0].IssuerPlanetmint},
			response: &types.QueryGetMachineByPublicKeyResponse{Machine: &msgs[0]},
		},
		{
			desc:     "GetByIssuerLiquid",
			request:  &types.QueryGetMachineByPublicKeyRequest{PublicKey: msgs[0].IssuerLiquid},
			response: &types.QueryGetMachineByPublicKeyResponse{Machine: &msgs[0]},
		}, {
			desc:    "MachineNotFound",
			request: &types.QueryGetMachineByPublicKeyRequest{PublicKey: "invalid key"},
			err:     status.Error(codes.NotFound, "machine not found by public key: invalid key"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.GetMachineByPublicKey(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.Equal(t, tc.response, response)
			}
		})
	}
}
