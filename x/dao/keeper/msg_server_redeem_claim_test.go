package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	keepertest "github.com/planetmint/planetmint-go/testutil/keeper"
	"github.com/planetmint/planetmint-go/testutil/sample"
	"github.com/planetmint/planetmint-go/x/dao/keeper"
	"github.com/planetmint/planetmint-go/x/dao/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestRedeemClaimMsgServerCreate(t *testing.T) {
	t.Parallel()
	k, ctx := keepertest.DaoKeeper(t)
	srv := keeper.NewMsgServerImpl(*k)
	wctx := sdk.WrapSDKContext(ctx)
	creator := sample.ConstBech32Addr
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreateRedeemClaim{Creator: creator,
			Beneficiary: strconv.Itoa(i),
		}
		_, err := srv.CreateRedeemClaim(wctx, expected)
		require.NoError(t, err)
		rst, found := k.GetRedeemClaim(ctx,
			expected.Beneficiary,
			uint64(0),
		)
		require.True(t, found)
		require.Equal(t, expected.Creator, rst.Creator)
	}
}

func TestRedeemClaimMsgServerUpdate(t *testing.T) {
	t.Parallel()
	creator := sample.ConstBech32Addr

	tests := []struct {
		desc    string
		request *types.MsgUpdateRedeemClaim
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgUpdateRedeemClaim{Creator: creator,
				Beneficiary:  strconv.Itoa(0),
				LiquidTxHash: strconv.Itoa(0),
			},
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgUpdateRedeemClaim{Creator: creator,
				Beneficiary:  strconv.Itoa(100000),
				LiquidTxHash: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			t.Parallel()
			k, ctx := keepertest.DaoKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)
			expected := &types.MsgCreateRedeemClaim{Creator: creator,
				Beneficiary: strconv.Itoa(0),
			}
			_, err := srv.CreateRedeemClaim(wctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdateRedeemClaim(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetRedeemClaim(ctx,
					expected.Beneficiary,
					uint64(0),
				)
				require.True(t, found)
				require.Equal(t, expected.Creator, rst.Creator)
			}
		})
	}
}
