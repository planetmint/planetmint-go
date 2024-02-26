package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/planetmint/planetmint-go/errormsg"
)

const TypeMsgReportPopResult = "report_pop_result"

var _ sdk.Msg = &MsgReportPopResult{}

func NewMsgReportPopResult(creator string, challenge *Challenge) *MsgReportPopResult {
	return &MsgReportPopResult{
		Creator:   creator,
		Challenge: challenge,
	}
}

func (msg *MsgReportPopResult) Route() string {
	return RouterKey
}

func (msg *MsgReportPopResult) Type() string {
	return TypeMsgReportPopResult
}

func (msg *MsgReportPopResult) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgReportPopResult) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgReportPopResult) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, errormsg.ErrorInvalidCreator, err)
	}
	return nil
}
