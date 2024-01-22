package keeper

import (
	"context"
	"fmt"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/planetmint/planetmint-go/util"
	"github.com/planetmint/planetmint-go/x/dao/types"
)

var (
	createRedeemClaimTag = "create redeem claim: "
)

func (k msgServer) CreateRedeemClaim(goCtx context.Context, msg *types.MsgCreateRedeemClaim) (*types.MsgCreateRedeemClaimResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: check for sufficient claim => claim ante handler
	var redeemClaim = types.RedeemClaim{
		Creator:     msg.Creator,
		Beneficiary: msg.Beneficiary,
		Amount:      msg.Amount,
	}

	id := k.CreateNewRedeemClaim(
		ctx,
		redeemClaim,
	)

	if util.IsValidatorBlockProposer(ctx, ctx.BlockHeader().ProposerAddress) {
		util.GetAppLogger().Info(ctx, fmt.Sprintf("Issuing RDDL claim: %s/%d", msg.Beneficiary, id))
		txID, err := util.DistributeAsset(msg.Beneficiary, "TODO: convert uint to str")
		if err != nil {
			util.GetAppLogger().Error(ctx, createRedeemClaimTag+"could not issue claim to beneficiary: "+msg.GetBeneficiary())
		}
		util.SendUpdateRedeemClaim(goCtx, msg.Beneficiary, id, txID)
	}

	return &types.MsgCreateRedeemClaimResponse{}, nil
}

func (k msgServer) UpdateRedeemClaim(goCtx context.Context, msg *types.MsgUpdateRedeemClaim) (*types.MsgUpdateRedeemClaimResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: move check to ante handler
	// Check if the value exists
	valFound, isFound := k.GetRedeemClaim(
		ctx,
		msg.Beneficiary,
		msg.Id,
	)
	if !isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	var redeemClaim = types.RedeemClaim{
		Creator:      valFound.Creator,
		Beneficiary:  msg.Beneficiary,
		LiquidTxHash: msg.LiquidTxHash,
		Amount:       valFound.Amount,
		Id:           valFound.Id,
	}

	k.SetRedeemClaim(ctx, redeemClaim)

	return &types.MsgUpdateRedeemClaimResponse{}, nil
}

// TODO: add msg handler for confirm claim
