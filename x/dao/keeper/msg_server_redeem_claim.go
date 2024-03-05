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
	params := k.GetParams(ctx)

	var redeemClaim = types.RedeemClaim{
		Creator:     msg.Creator,
		Beneficiary: msg.Beneficiary,
		Amount:      msg.Amount,
	}

	err := k.burnClaimAmount(ctx, sdk.MustAccAddressFromBech32(msg.Creator), msg.Amount)
	if err != nil {
		util.GetAppLogger().Error(ctx, createRedeemClaimTag+"could not burn claim")
	}

	id := k.CreateNewRedeemClaim(
		ctx,
		redeemClaim,
	)

	if util.IsValidatorBlockProposer(ctx, ctx.BlockHeader().ProposerAddress, k.RootDir) {
		util.GetAppLogger().Info(ctx, fmt.Sprintf("Issuing RDDL claim: %s/%d", msg.Beneficiary, id))
		txID, err := util.DistributeAsset(msg.Beneficiary, util.UintValueToRDDLTokenString(msg.Amount), params.ReissuanceAsset)
		if err != nil {
			util.GetAppLogger().Error(ctx, createRedeemClaimTag+"could not issue claim to beneficiary: "+msg.GetBeneficiary())
		}
		util.SendUpdateRedeemClaim(goCtx, msg.Beneficiary, id, txID)
	}

	return &types.MsgCreateRedeemClaimResponse{}, nil
}

func (k msgServer) UpdateRedeemClaim(goCtx context.Context, msg *types.MsgUpdateRedeemClaim) (*types.MsgUpdateRedeemClaimResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

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

func (k msgServer) ConfirmRedeemClaim(goCtx context.Context, msg *types.MsgConfirmRedeemClaim) (*types.MsgConfirmRedeemClaimResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

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
		LiquidTxHash: valFound.LiquidTxHash,
		Amount:       valFound.Amount,
		Id:           valFound.Id,
		Confirmed:    true,
	}

	k.SetRedeemClaim(ctx, redeemClaim)

	return &types.MsgConfirmRedeemClaimResponse{}, nil
}

func (k msgServer) burnClaimAmount(ctx sdk.Context, addr sdk.AccAddress, amount uint64) (err error) {
	params := k.GetParams(ctx)
	burnCoins := sdk.NewCoins(sdk.NewCoin(params.ClaimDenom, sdk.NewIntFromUint64(amount)))
	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, addr, types.ModuleName, burnCoins)
	if err != nil {
		return err
	}
	err = k.bankKeeper.BurnCoins(ctx, types.ModuleName, burnCoins)
	if err != nil {
		return err
	}
	return
}
