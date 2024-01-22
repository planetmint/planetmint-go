package ante

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/planetmint/planetmint-go/config"
	daotypes "github.com/planetmint/planetmint-go/x/dao/types"
)

type RedeemClaimDecorator struct {
	dk DaoKeeper
	bk BankKeeper
}

func NewRedeemClaimDecorator(dk DaoKeeper, bk BankKeeper) RedeemClaimDecorator {
	return RedeemClaimDecorator{
		dk: dk,
		bk: bk,
	}
}

func (rcd RedeemClaimDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (_ sdk.Context, err error) {
	for _, msg := range tx.GetMsgs() {
		switch sdk.MsgTypeURL(msg) {
		case "/planetmintgo.dao.MsgCreateRedeemClaim":
			ctx, err = rcd.handleCreateRedeemClaim(ctx, msg)
		case "/planetmintgo.dao.MsgConfirmRedeemClaim":
			ctx, err = rcd.handleConfirmRedeemClaim(ctx, msg)
		default:
			continue
		}
	}

	if err != nil {
		return ctx, err
	}

	return next(ctx, tx, simulate)
}

func (rcd RedeemClaimDecorator) handleCreateRedeemClaim(ctx sdk.Context, msg sdk.Msg) (sdk.Context, error) {
	cfg := config.GetConfig()

	createRedeemClaimMsg, ok := msg.(*daotypes.MsgCreateRedeemClaim)
	if !ok {
		return ctx, errorsmod.Wrapf(sdkerrors.ErrLogic, "could not cast to MsgCreateRedeemClaim")
	}

	addr := sdk.MustAccAddressFromBech32(createRedeemClaimMsg.Creator)

	balance := rcd.bk.GetBalance(ctx, addr, cfg.ClaimDenom)

	if !balance.Amount.GTE(sdk.NewIntFromUint64(createRedeemClaimMsg.Amount)) {
		return ctx, errorsmod.Wrap(sdkerrors.ErrInsufficientFunds, "error during checkTx or reChec")
	}

	return ctx, nil
}

// TODO: check for daokeeper param
func (rcd RedeemClaimDecorator) handleConfirmRedeemClaim(ctx sdk.Context, msg sdk.Msg) (sdk.Context, error) {
	return ctx, nil
}
