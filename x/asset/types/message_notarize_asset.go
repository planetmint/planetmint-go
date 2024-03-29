package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/planetmint/planetmint-go/errormsg"
)

const TypeMsgNotarizeAsset = "notarize_asset"

var _ sdk.Msg = &MsgNotarizeAsset{}

func NewMsgNotarizeAsset(creator string, cid string) *MsgNotarizeAsset {
	return &MsgNotarizeAsset{
		Creator: creator,
		Cid:     cid,
	}
}

func (msg *MsgNotarizeAsset) Route() string {
	return RouterKey
}

func (msg *MsgNotarizeAsset) Type() string {
	return TypeMsgNotarizeAsset
}

func (msg *MsgNotarizeAsset) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgNotarizeAsset) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgNotarizeAsset) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, errormsg.ErrorInvalidCreator, err)
	}
	return nil
}
