package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRegisterTrustAnchor = "register_trust_anchor"

var _ sdk.Msg = &MsgRegisterTrustAnchor{}

func NewMsgRegisterTrustAnchor(creator string, trustAnchor *TrustAnchor) *MsgRegisterTrustAnchor {
	return &MsgRegisterTrustAnchor{
		Creator:     creator,
		TrustAnchor: trustAnchor,
	}
}

func (msg *MsgRegisterTrustAnchor) Route() string {
	return RouterKey
}

func (msg *MsgRegisterTrustAnchor) Type() string {
	return TypeMsgRegisterTrustAnchor
}

func (msg *MsgRegisterTrustAnchor) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRegisterTrustAnchor) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRegisterTrustAnchor) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
