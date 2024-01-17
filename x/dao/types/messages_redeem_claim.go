package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateRedeemClaim = "create_redeem_claim"
	TypeMsgUpdateRedeemClaim = "update_redeem_claim"
	TypeMsgDeleteRedeemClaim = "delete_redeem_claim"
)

var _ sdk.Msg = &MsgCreateRedeemClaim{}

func NewMsgCreateRedeemClaim(
	creator string,
	beneficiary string,
	liquidTxHash string,
	amount string,
	issued bool,

) *MsgCreateRedeemClaim {
	return &MsgCreateRedeemClaim{
		Creator:      creator,
		Beneficiary:  beneficiary,
		LiquidTxHash: liquidTxHash,
		Amount:       amount,
		Issued:       issued,
	}
}

func (msg *MsgCreateRedeemClaim) Route() string {
	return RouterKey
}

func (msg *MsgCreateRedeemClaim) Type() string {
	return TypeMsgCreateRedeemClaim
}

func (msg *MsgCreateRedeemClaim) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateRedeemClaim) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateRedeemClaim) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateRedeemClaim{}

func NewMsgUpdateRedeemClaim(
	creator string,
	beneficiary string,
	liquidTxHash string,
	amount string,
	issued bool,

) *MsgUpdateRedeemClaim {
	return &MsgUpdateRedeemClaim{
		Creator:      creator,
		Beneficiary:  beneficiary,
		LiquidTxHash: liquidTxHash,
		Amount:       amount,
		Issued:       issued,
	}
}

func (msg *MsgUpdateRedeemClaim) Route() string {
	return RouterKey
}

func (msg *MsgUpdateRedeemClaim) Type() string {
	return TypeMsgUpdateRedeemClaim
}

func (msg *MsgUpdateRedeemClaim) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateRedeemClaim) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateRedeemClaim) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteRedeemClaim{}

func NewMsgDeleteRedeemClaim(
	creator string,
	beneficiary string,
	liquidTxHash string,

) *MsgDeleteRedeemClaim {
	return &MsgDeleteRedeemClaim{
		Creator:      creator,
		Beneficiary:  beneficiary,
		LiquidTxHash: liquidTxHash,
	}
}
func (msg *MsgDeleteRedeemClaim) Route() string {
	return RouterKey
}

func (msg *MsgDeleteRedeemClaim) Type() string {
	return TypeMsgDeleteRedeemClaim
}

func (msg *MsgDeleteRedeemClaim) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteRedeemClaim) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteRedeemClaim) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
