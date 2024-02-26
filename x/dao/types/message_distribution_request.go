package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/planetmint/planetmint-go/errormsg"
)

const TypeMsgDistributionRequest = "distribution_request"

var _ sdk.Msg = &MsgDistributionRequest{}

func NewMsgDistributionRequest(creator string, distribution *DistributionOrder) *MsgDistributionRequest {
	return &MsgDistributionRequest{
		Creator:      creator,
		Distribution: distribution,
	}
}

func (msg *MsgDistributionRequest) Route() string {
	return RouterKey
}

func (msg *MsgDistributionRequest) Type() string {
	return TypeMsgDistributionRequest
}

func (msg *MsgDistributionRequest) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDistributionRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDistributionRequest) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, errormsg.ErrorInvalidCreator, err)
	}
	return nil
}
