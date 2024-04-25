package keeper

import (
	"context"
	"fmt"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/planetmint/planetmint-go/clients"
	"github.com/planetmint/planetmint-go/util"
	"github.com/planetmint/planetmint-go/x/dao/types"
)

var (
	createRedeemClaimTag = "create redeem claim: "
)

func (k msgServer) CreateRedeemClaim(goCtx context.Context, msg *types.MsgCreateRedeemClaim) (*types.MsgCreateRedeemClaimResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	params := k.GetParams(ctx)

	addr := sdk.MustAccAddressFromBech32(msg.Creator)
	burnCoins := k.bankKeeper.GetBalance(ctx, addr, params.ClaimDenom)

	var redeemClaim = types.RedeemClaim{
		Creator:     msg.Creator,
		Beneficiary: msg.Beneficiary,
		Amount:      burnCoins.Amount.Uint64(),
	}

	err := k.burnClaimAmount(ctx, sdk.MustAccAddressFromBech32(msg.Creator), sdk.NewCoins(burnCoins))
	if err != nil {
		util.GetAppLogger().Error(ctx, createRedeemClaimTag+"could not burn claim")
	}

	id := k.CreateNewRedeemClaim(
		ctx,
		redeemClaim,
	)

	if util.IsValidatorBlockProposer(ctx, ctx.BlockHeader().ProposerAddress, k.RootDir) {
		go k.postClaimToService(ctx, msg.GetBeneficiary(), burnCoins.Amount.Uint64(), id)
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

	err := k.validateConfirmRedeemClaim(ctx, msg)
	if err != nil {
		return nil, err
	}

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

func (k msgServer) validateConfirmRedeemClaim(ctx sdk.Context, msg *types.MsgConfirmRedeemClaim) (err error) {
	if msg.Creator != k.GetClaimAddress(ctx) {
		return errorsmod.Wrapf(types.ErrInvalidClaimAddress, "expected: %s; got: %s", k.GetClaimAddress(ctx), msg.Creator)
	}
	_, found := k.GetRedeemClaim(ctx, msg.Beneficiary, msg.Id)
	if !found {
		return errorsmod.Wrapf(sdkerrors.ErrNotFound, "no redeem claim found for beneficiary: %s; id: %d", msg.Beneficiary, msg.Id)
	}

	return nil
}

func (k msgServer) burnClaimAmount(ctx sdk.Context, addr sdk.AccAddress, burnCoins sdk.Coins) (err error) {
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

func (k msgServer) postClaimToService(ctx sdk.Context, beneficiary string, amount uint64, id uint64) {
	goCtx := sdk.WrapSDKContext(ctx)
	util.GetAppLogger().Info(ctx, fmt.Sprintf("Issuing RDDL claim: %s/%d", beneficiary, id))
	txID, err := clients.PostClaim(goCtx, beneficiary, amount, id)
	if err != nil {
		util.GetAppLogger().Error(ctx, createRedeemClaimTag+"could not issue claim to beneficiary: "+beneficiary)
	}
	util.SendUpdateRedeemClaim(goCtx, beneficiary, id, txID)
}
