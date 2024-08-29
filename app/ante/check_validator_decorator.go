package ante

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/planetmint/planetmint-go/x/dao/types"
)

type CheckValidatorDecorator struct {
	sk StakingKeeper
}

func NewCheckValidatorDecorator(sk StakingKeeper) CheckValidatorDecorator {
	return CheckValidatorDecorator{sk: sk}
}

func (cv CheckValidatorDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (_ sdk.Context, err error) {
	if simulate || ctx.BlockHeight() == 0 {
		return next(ctx, tx, simulate)
	}

	for _, msg := range tx.GetMsgs() {
		switch sdk.MsgTypeURL(msg) {
		case "/planetmintgo.dao.MsgInitPop":
			fallthrough
		case "/planetmintgo.dao.MsgDistributionRequest":
			fallthrough
		case "/planetmintgo.dao.MsgDistributionResult":
			fallthrough
		case "/planetmintgo.dao.MsgReissueRDDLProposal":
			fallthrough
		case "/planetmintgo.dao.MsgReissueRDDLResult":
			fallthrough
		case "/planetmintgo.dao.MsgUpdateRedeemClaim":
			fallthrough
		case "/planetmintgo.machine.MsgNotarizeLiquidAsset":
			fallthrough
		case "/planetmintgo.machine.MsgRegisterTrustAnchor":
			ctx, err = cv.handleMsg(ctx, msg)
		default:
			continue
		}
	}

	if err != nil {
		return ctx, err
	}

	return next(ctx, tx, simulate)
}

func (cv CheckValidatorDecorator) handleMsg(ctx sdk.Context, msg sdk.Msg) (_ sdk.Context, err error) {
	signer := msg.GetSigners()[0]
	_, found := cv.sk.GetValidator(ctx, sdk.ValAddress(signer))
	if !found {
		return ctx, errorsmod.Wrap(types.ErrRestrictedMsg, ErrorAnteContext)
	}
	return ctx, nil
}
