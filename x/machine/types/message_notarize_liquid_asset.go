package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgNotarizeLiquidAsset{}

func NewMsgNotarizeLiquidAsset(creator string, notarization *LiquidAsset) *MsgNotarizeLiquidAsset {
	return &MsgNotarizeLiquidAsset{
		Creator:      creator,
		Notarization: notarization,
	}
}

func (msg *MsgNotarizeLiquidAsset) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
