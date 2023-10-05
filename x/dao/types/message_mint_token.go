package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgMintToken = "mint_token"

var _ sdk.Msg = &MsgMintToken{}

func NewMsgMintToken(creator string, mintRequest *MintRequest) *MsgMintToken {
	return &MsgMintToken{
		Creator:     creator,
		MintRequest: mintRequest,
	}
}

func (msg *MsgMintToken) Route() string {
	return RouterKey
}

func (msg *MsgMintToken) Type() string {
	return TypeMsgMintToken
}

func (msg *MsgMintToken) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgMintToken) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgMintToken) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
