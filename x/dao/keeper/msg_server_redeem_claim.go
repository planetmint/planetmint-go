package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/planetmint/planetmint-go/x/dao/types"
)

func (k msgServer) CreateRedeemClaim(goCtx context.Context, msg *types.MsgCreateRedeemClaim) (*types.MsgCreateRedeemClaimResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetRedeemClaim(
		ctx,
		msg.Beneficiary,
		msg.LiquidTxHash,
	)
	if isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var redeemClaim = types.RedeemClaim{
		Creator:      msg.Creator,
		Beneficiary:  msg.Beneficiary,
		LiquidTxHash: msg.LiquidTxHash,
		Amount:       msg.Amount,
		Confirmed:    msg.Confirmed,
	}

	k.SetRedeemClaim(
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
		msg.LiquidTxHash,
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
		Amount:       msg.Amount,
		Confirmed:    msg.Confirmed,
	}

	k.SetRedeemClaim(ctx, redeemClaim)

	return &types.MsgUpdateRedeemClaimResponse{}, nil
}
