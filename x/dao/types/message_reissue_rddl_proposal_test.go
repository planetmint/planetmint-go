package types

import (
	"testing"

	"github.com/cometbft/cometbft/crypto/ed25519"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func AccAddress() string {
	pk := ed25519.GenPrivKey().PubKey()
	addr := pk.Address()
	return sdk.AccAddress(addr).String()
}

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
				Creator: AccAddress(),
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
