package types

import (
	"testing"

	"github.com/planetmint/planetmint-go/testutil/errormsg"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgNotarizeAssetValidateBasic(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		msg  MsgNotarizeAsset
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgNotarizeAsset{
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
