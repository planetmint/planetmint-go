package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/planetmint/planetmint-go/testutil/errormsg"
	"github.com/stretchr/testify/require"
)

func TestMsgInitPopValidateBasic(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		msg  MsgInitPop
		err  error
	}{
		{
			name: sdkerrors.ErrInvalidAddress.Error(),
			msg: MsgInitPop{
				Creator: errormsg.ErrorInvalidAddress,
			},
			err: sdkerrors.ErrInvalidAddress,
		},
	}
	for _, tt := range tests {
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
