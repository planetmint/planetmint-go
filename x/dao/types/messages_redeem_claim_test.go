package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgCreateRedeemClaim_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateRedeemClaim
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateRedeemClaim{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
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
	tests := []struct {
		name string
		msg  MsgUpdateRedeemClaim
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateRedeemClaim{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
