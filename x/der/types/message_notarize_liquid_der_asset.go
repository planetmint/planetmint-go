package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgNotarizeLiquidDerAsset = "notarize_liquid_der_asset"

var _ sdk.Msg = &MsgNotarizeLiquidDerAsset{}

func NewMsgNotarizeLiquidDerAsset(creator string, derAsset *LiquidDerAsset) *MsgNotarizeLiquidDerAsset {
	return &MsgNotarizeLiquidDerAsset{
		Creator:  creator,
		DerAsset: derAsset,
	}
}

func (msg *MsgNotarizeLiquidDerAsset) Route() string {
	return RouterKey
}

func (msg *MsgNotarizeLiquidDerAsset) Type() string {
	return TypeMsgNotarizeLiquidDerAsset
}

func (msg *MsgNotarizeLiquidDerAsset) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgNotarizeLiquidDerAsset) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgNotarizeLiquidDerAsset) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
