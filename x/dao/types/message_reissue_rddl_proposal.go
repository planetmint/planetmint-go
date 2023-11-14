package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgReissueRDDLProposal = "reissue_rddl_proposal"

var _ sdk.Msg = &MsgReissueRDDLProposal{}

func NewMsgReissueRDDLProposal(creator string, proposer string, tx string, blockHeight int64) *MsgReissueRDDLProposal {
	return &MsgReissueRDDLProposal{
		Creator:     creator,
		Proposer:    proposer,
		Tx:          tx,
		BlockHeight: blockHeight,
	}
}

func (msg *MsgReissueRDDLProposal) Route() string {
	return RouterKey
}

func (msg *MsgReissueRDDLProposal) Type() string {
	return TypeMsgReissueRDDLProposal
}

func (msg *MsgReissueRDDLProposal) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgReissueRDDLProposal) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgReissueRDDLProposal) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
