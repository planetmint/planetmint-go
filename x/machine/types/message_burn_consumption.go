package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgBurnConsumption = "burn_consumption"

var _ sdk.Msg = &MsgBurnConsumption{}

func NewMsgBurnConsumption(creator string, consumption sdk.Coins) *MsgBurnConsumption {
	return &MsgBurnConsumption{
		Creator:     creator,
		Consumption: consumption,
	}
}

func (msg *MsgBurnConsumption) Route() string {
	return RouterKey
}

func (msg *MsgBurnConsumption) Type() string {
	return TypeMsgBurnConsumption
}

func (msg *MsgBurnConsumption) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgBurnConsumption) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgBurnConsumption) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
