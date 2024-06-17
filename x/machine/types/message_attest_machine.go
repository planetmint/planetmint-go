package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgAttestMachine{}

func NewMsgAttestMachine(creator string, machine *Machine) *MsgAttestMachine {
	return &MsgAttestMachine{
		Creator: creator,
		Machine: machine,
	}
}

func (msg *MsgAttestMachine) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
