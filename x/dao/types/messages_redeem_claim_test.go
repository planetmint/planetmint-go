package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/planetmint/planetmint-go/testutil/errormsg"
	"github.com/stretchr/testify/require"
)

func TestMsgCreateRedeemClaim_ValidateBasic(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		msg  MsgCreateRedeemClaim
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateRedeemClaim{
				Creator: errormsg.ErrorInvalidAddress,
			},
			err: sdkerrors.ErrInvalidAddress,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestMsgUpdateRedeemClaim_ValidateBasic(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		msg  MsgUpdateRedeemClaim
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateRedeemClaim{
				Creator: errormsg.ErrorInvalidAddress,
			},
			err: sdkerrors.ErrInvalidAddress,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestMsgConfirmRedeemClaim_ValidateBasic(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		msg  MsgConfirmRedeemClaim
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgConfirmRedeemClaim{
				Creator: errormsg.ErrorInvalidAddress,
			},
			err: sdkerrors.ErrInvalidAddress,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
