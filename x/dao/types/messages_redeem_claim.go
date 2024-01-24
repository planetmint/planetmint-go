package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateRedeemClaim  = "create_redeem_claim"
	TypeMsgUpdateRedeemClaim  = "update_redeem_claim"
	TypeMsgConfirmRedeemClaim = "confirm_redeem_claim"
)

var _ sdk.Msg = &MsgCreateRedeemClaim{}

func NewMsgCreateRedeemClaim(
	creator string,
	beneficiary string,
	amount uint64,

) *MsgCreateRedeemClaim {
	return &MsgCreateRedeemClaim{
		Creator:     creator,
		Beneficiary: beneficiary,
		Amount:      amount,
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
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateRedeemClaim{}

func NewMsgUpdateRedeemClaim(
	creator string,
	beneficiary string,
	liquidTxHash string,
	id uint64,

) *MsgUpdateRedeemClaim {
	return &MsgUpdateRedeemClaim{
		Creator:      creator,
		Beneficiary:  beneficiary,
		LiquidTxHash: liquidTxHash,
		Id:           id,
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
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgConfirmRedeemClaim{}

func NewMsgConfirmRedeemClaim(creator string, id uint64, beneficiary string) *MsgConfirmRedeemClaim {
	return &MsgConfirmRedeemClaim{
		Creator:     creator,
		Id:          id,
		Beneficiary: beneficiary,
	}
}

func (msg *MsgConfirmRedeemClaim) Route() string {
	return RouterKey
}

func (msg *MsgConfirmRedeemClaim) Type() string {
	return TypeMsgConfirmRedeemClaim
}

func (msg *MsgConfirmRedeemClaim) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgConfirmRedeemClaim) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgConfirmRedeemClaim) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
