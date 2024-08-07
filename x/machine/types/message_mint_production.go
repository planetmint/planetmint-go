package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgMintProduction = "mint_production"

var _ sdk.Msg = &MsgMintProduction{}

func NewMsgMintProduction(creator string, production sdk.Coins) *MsgMintProduction {
	return &MsgMintProduction{
		Creator:    creator,
		Production: production,
	}
}

func (msg *MsgMintProduction) Route() string {
	return RouterKey
}

func (msg *MsgMintProduction) Type() string {
	return TypeMsgMintProduction
}

func (msg *MsgMintProduction) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgMintProduction) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgMintProduction) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
