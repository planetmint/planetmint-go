package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgDistributionResult = "distribution_result"

var _ sdk.Msg = &MsgDistributionResult{}

func NewMsgDistributionResult(creator string, lastPop int64, daoTxid string, investorTxid string, popTxid string) *MsgDistributionResult {
	return &MsgDistributionResult{
		Creator:      creator,
		LastPop:      lastPop,
		DaoTxid:      daoTxid,
		InvestorTxid: investorTxid,
		PopTxid:      popTxid,
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
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
