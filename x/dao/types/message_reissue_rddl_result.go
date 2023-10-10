package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgReissueRDDLResult = "reissue_rddl_result"

var _ sdk.Msg = &MsgReissueRDDLResult{}

func NewMsgReissueRDDLResult(creator string, proposer string, txId string, blockHeight uint64) *MsgReissueRDDLResult {
	return &MsgReissueRDDLResult{
		Creator:     creator,
		Proposer:    proposer,
		TxId:        txId,
		BlockHeight: blockHeight,
	}
}

func (msg *MsgReissueRDDLResult) Route() string {
	return RouterKey
}

func (msg *MsgReissueRDDLResult) Type() string {
	return TypeMsgReissueRDDLResult
}

func (msg *MsgReissueRDDLResult) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgReissueRDDLResult) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgReissueRDDLResult) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
