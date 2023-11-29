package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgInitPop = "init_pop"

var _ sdk.Msg = &MsgInitPop{}

func NewMsgInitPop(creator string, initiator string, challenger string, challengee string, height uint64) *MsgInitPop {
	return &MsgInitPop{
		Creator:    creator,
		Initiator:  initiator,
		Challenger: challenger,
		Challengee: challengee,
		Height:     height,
	}
}

func (msg *MsgInitPop) Route() string {
	return RouterKey
}

func (msg *MsgInitPop) Type() string {
	return TypeMsgInitPop
}

func (msg *MsgInitPop) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgInitPop) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgInitPop) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
