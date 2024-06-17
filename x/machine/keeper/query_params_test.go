package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/planetmint/planetmint-go/x/machine/types"

	keepertest "github.com/planetmint/planetmint-go/testutil/keeper"
)

func TestParamsQuery(t *testing.T) {
	keeper, ctx := keepertest.MachineKeeper(t)
	params := types.DefaultParams()
	require.NoError(t, keeper.SetParams(ctx, params))

	response, err := keeper.Params(ctx, &types.QueryParamsRequest{})
	require.NoError(t, err)
	require.Equal(t, &types.QueryParamsResponse{Params: params}, response)
}
