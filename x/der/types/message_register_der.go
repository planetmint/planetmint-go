package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRegisterDER = "register_der"

var _ sdk.Msg = &MsgRegisterDER{}

func NewMsgRegisterDER(creator string, der *DER) *MsgRegisterDER {
	return &MsgRegisterDER{
		Creator: creator,
		Der:     der,
	}
}

func (msg *MsgRegisterDER) Route() string {
	return RouterKey
}

func (msg *MsgRegisterDER) Type() string {
	return TypeMsgRegisterDER
}

func (msg *MsgRegisterDER) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRegisterDER) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRegisterDER) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
