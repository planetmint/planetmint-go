package cli_test

import (
	"fmt"
	"strconv"
	"testing"

	tmcli "github.com/cometbft/cometbft/libs/cli"
	"github.com/cosmos/cosmos-sdk/client/flags"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/planetmint/planetmint-go/testutil/network"
	"github.com/planetmint/planetmint-go/testutil/nullify"
	"github.com/planetmint/planetmint-go/x/dao/client/cli"
	"github.com/planetmint/planetmint-go/x/dao/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func networkWithRedeemClaimObjects(t *testing.T, n int) (*network.Network, []types.RedeemClaim) {
	t.Helper()
	cfg := network.DefaultConfig()
	state := types.GenesisState{}
	for i := 0; i < n; i++ {
		redeemClaim := types.RedeemClaim{
			Beneficiary:  strconv.Itoa(i),
			LiquidTxHash: strconv.Itoa(i),
		}
		nullify.Fill(&redeemClaim)
		state.RedeemClaimList = append(state.RedeemClaimList, redeemClaim)
	}
	buf, err := cfg.Codec.MarshalJSON(&state)
	require.NoError(t, err)
	cfg.GenesisState[types.ModuleName] = buf
	return network.New(t, cfg), state.RedeemClaimList
}

func TestShowRedeemClaim(t *testing.T) {
	t.Parallel()
	net, objs := networkWithRedeemClaimObjects(t, 2)

	ctx := net.Validators[0].ClientCtx
	common := []string{
		fmt.Sprintf("--%s=json", tmcli.OutputFlag),
	}
	tests := []struct {
		desc           string
		idBeneficiary  string
		idLiquidTxHash string

		args []string
		err  error
		obj  types.RedeemClaim
	}{
		{
			desc:           "found",
			idBeneficiary:  objs[0].Beneficiary,
			idLiquidTxHash: objs[0].LiquidTxHash,

			args: common,
			obj:  objs[0],
		},
		{
			desc:           "not found",
			idBeneficiary:  strconv.Itoa(100000),
			idLiquidTxHash: strconv.Itoa(100000),

			args: common,
			err:  status.Error(codes.NotFound, "not found"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			t.Parallel()
			args := []string{
				tc.idBeneficiary,
				tc.idLiquidTxHash,
			}
			args = append(args, tc.args...)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdShowRedeemClaim(), args)
			if tc.err != nil {
				stat, ok := status.FromError(tc.err)
				require.True(t, ok)
				require.ErrorIs(t, stat.Err(), tc.err)
			} else {
				require.NoError(t, err)
				var resp types.QueryGetRedeemClaimResponse
				require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
				require.NotNil(t, resp.RedeemClaim)
				require.Equal(t,
					nullify.Fill(&tc.obj),
					nullify.Fill(&resp.RedeemClaim),
				)
			}
		})
	}
}

func TestListRedeemClaim(t *testing.T) {
	t.Parallel()
	net, objs := networkWithRedeemClaimObjects(t, 5)

	ctx := net.Validators[0].ClientCtx
	request := func(next []byte, offset, limit uint64, total bool) []string {
		args := []string{
			fmt.Sprintf("--%s=json", tmcli.OutputFlag),
		}
		if next == nil {
			args = append(args, fmt.Sprintf("--%s=%d", flags.FlagOffset, offset))
		} else {
			args = append(args, fmt.Sprintf("--%s=%s", flags.FlagPageKey, next))
		}
		args = append(args, fmt.Sprintf("--%s=%d", flags.FlagLimit, limit))
		if total {
			args = append(args, fmt.Sprintf("--%s", flags.FlagCountTotal))
		}
		return args
	}
	t.Run("ByOffset", func(t *testing.T) {
		t.Parallel()
		step := 2
		for i := 0; i < len(objs); i += step {
			args := request(nil, uint64(i), uint64(step), false)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListRedeemClaim(), args)
			require.NoError(t, err)
			var resp types.QueryAllRedeemClaimResponse
			require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
			require.LessOrEqual(t, len(resp.RedeemClaim), step)
			require.Subset(t,
				nullify.Fill(objs),
				nullify.Fill(resp.RedeemClaim),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		t.Parallel()
		step := 2
		var next []byte
		for i := 0; i < len(objs); i += step {
			args := request(next, 0, uint64(step), false)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListRedeemClaim(), args)
			require.NoError(t, err)
			var resp types.QueryAllRedeemClaimResponse
			require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
			require.LessOrEqual(t, len(resp.RedeemClaim), step)
			require.Subset(t,
				nullify.Fill(objs),
				nullify.Fill(resp.RedeemClaim),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		t.Parallel()
		args := request(nil, 0, uint64(len(objs)), true)
		out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListRedeemClaim(), args)
		require.NoError(t, err)
		var resp types.QueryAllRedeemClaimResponse
		require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
		require.NoError(t, err)
		require.Equal(t, len(objs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(objs),
			nullify.Fill(resp.RedeemClaim),
		)
	})
}
