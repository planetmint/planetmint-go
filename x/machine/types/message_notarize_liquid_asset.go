package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/planetmint/planetmint-go/errormsg"
)

const TypeMsgNotarizeLiquidAsset = "notarize_liquid_asset"

var _ sdk.Msg = &MsgNotarizeLiquidAsset{}

func NewMsgNotarizeLiquidAsset(creator string, notarization *LiquidAsset) *MsgNotarizeLiquidAsset {
	return &MsgNotarizeLiquidAsset{
		Creator:      creator,
		Notarization: notarization,
	}
}

func (msg *MsgNotarizeLiquidAsset) Route() string {
	return RouterKey
}

func (msg *MsgNotarizeLiquidAsset) Type() string {
	return TypeMsgNotarizeLiquidAsset
}

func (msg *MsgNotarizeLiquidAsset) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgNotarizeLiquidAsset) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgNotarizeLiquidAsset) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, errormsg.ErrorInvalidCreator, err)
	}
	return nil
}
