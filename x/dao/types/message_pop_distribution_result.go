package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgPopDistributionResult = "pop_distribution_result"

var _ sdk.Msg = &MsgPopDistributionResult{}

func NewMsgPopDistributionResult(creator string, lastPop uint64, daoTx string, investorTx string, popTx string) *MsgPopDistributionResult {
	return &MsgPopDistributionResult{
		Creator:    creator,
		LastPop:    lastPop,
		DaoTx:      daoTx,
		InvestorTx: investorTx,
		PopTx:      popTx,
	}
}

func (msg *MsgPopDistributionResult) Route() string {
	return RouterKey
}

func (msg *MsgPopDistributionResult) Type() string {
	return TypeMsgPopDistributionResult
}

func (msg *MsgPopDistributionResult) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgPopDistributionResult) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgPopDistributionResult) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
