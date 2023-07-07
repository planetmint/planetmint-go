package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgAttestMachine = "attest_machine"

var _ sdk.Msg = &MsgAttestMachine{}

func NewMsgAttestMachine(creator string, machine *Machine) *MsgAttestMachine {
	return &MsgAttestMachine{
		Creator: creator,
		Machine: machine,
	}
}

func (msg *MsgAttestMachine) Route() string {
	return RouterKey
}

func (msg *MsgAttestMachine) Type() string {
	return TypeMsgAttestMachine
}

func (msg *MsgAttestMachine) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgAttestMachine) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAttestMachine) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
