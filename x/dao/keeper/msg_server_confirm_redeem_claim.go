package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/planetmint/planetmint-go/x/dao/types"
)

func (k msgServer) ConfirmRedeemClaim(goCtx context.Context, msg *types.MsgConfirmRedeemClaim) (*types.MsgConfirmRedeemClaimResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgConfirmRedeemClaimResponse{}, nil
}
