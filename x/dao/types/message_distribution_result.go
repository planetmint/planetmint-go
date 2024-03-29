package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/planetmint/planetmint-go/errormsg"
)

const TypeMsgDistributionResult = "distribution_result"

var _ sdk.Msg = &MsgDistributionResult{}

func NewMsgDistributionResult(creator string, lastPop int64, daoTxID string, investorTxID string,
	popTxID string, earlyInvestorTxID string, strategicTxID string) *MsgDistributionResult {
	return &MsgDistributionResult{
		Creator:           creator,
		LastPop:           lastPop,
		DaoTxID:           daoTxID,
		InvestorTxID:      investorTxID,
		PopTxID:           popTxID,
		EarlyInvestorTxID: earlyInvestorTxID,
		StrategicTxID:     strategicTxID,
	}
}

func (msg *MsgDistributionResult) Route() string {
	return RouterKey
}

func (msg *MsgDistributionResult) Type() string {
	return TypeMsgDistributionResult
}

func (msg *MsgDistributionResult) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDistributionResult) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDistributionResult) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, errormsg.ErrorInvalidCreator, err)
	}
	return nil
}
