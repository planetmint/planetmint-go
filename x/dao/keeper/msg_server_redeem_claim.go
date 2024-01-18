package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/planetmint/planetmint-go/util"
	"github.com/planetmint/planetmint-go/x/dao/types"
)

func (k msgServer) CreateRedeemClaim(goCtx context.Context, msg *types.MsgCreateRedeemClaim) (*types.MsgCreateRedeemClaimResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: check for sufficient claim

	if util.IsValidatorBlockProposer(ctx, ctx.BlockHeader().ProposerAddress) {
		util.GetAppLogger().Info(ctx, "Issuing RDDL claim: ") // Add Beneficiary and Claim Index
		// TODO: send message to elements service
	}

	var redeemClaim = types.RedeemClaim{
		Creator:     msg.Creator,
		Beneficiary: msg.Beneficiary,
		Amount:      msg.Amount,
	}

	k.CreateNewRedeemClaim(
		ctx,
		redeemClaim,
	)
	return &types.MsgCreateRedeemClaimResponse{}, nil
}

func (k msgServer) UpdateRedeemClaim(goCtx context.Context, msg *types.MsgUpdateRedeemClaim) (*types.MsgUpdateRedeemClaimResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetRedeemClaim(
		ctx,
		msg.Beneficiary,
		msg.Id,
	)
	if !isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	var redeemClaim = types.RedeemClaim{
		Creator:      msg.Creator,
		Beneficiary:  msg.Beneficiary,
		LiquidTxHash: msg.LiquidTxHash,
		Amount:       valFound.Amount,
		Id:           valFound.Id,
	}

	k.SetRedeemClaim(ctx, redeemClaim)

	return &types.MsgUpdateRedeemClaimResponse{}, nil
}

// TODO: add msg handler for confirm claim
