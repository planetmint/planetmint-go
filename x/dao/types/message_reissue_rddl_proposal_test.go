package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/planetmint/planetmint-go/testutil/sample"
	"github.com/stretchr/testify/require"
)

func TestMsgReissueRDDLProposal_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgReissueRDDLProposal
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgReissueRDDLProposal{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgReissueRDDLProposal{
				Creator: sample.AccAddress(),
			},
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
