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
			initPopMsg, ok := msg.(*types.MsgInitPop)
			if ok {
				ctx, err = cv.handleMsg(ctx, initPopMsg)
			}
		case "/planetmintgo.dao.MsgDistributionRequest":
			distributionRequestMsg, ok := msg.(*types.MsgDistributionRequest)
			if ok {
				ctx, err = cv.handleMsg(ctx, distributionRequestMsg)
			}
		case "/planetmintgo.dao.MsgDistributionResult":
			distributionResultMsg, ok := msg.(*types.MsgDistributionResult)
			if ok {
				ctx, err = cv.handleMsg(ctx, distributionResultMsg)
			}
		case "/planetmintgo.dao.MsgReissueRDDLProposal":
			reissueProposalMsg, ok := msg.(*types.MsgReissueRDDLProposal)
			if ok {
				ctx, err = cv.handleMsg(ctx, reissueProposalMsg)
			}
		case "/planetmintgo.dao.MsgReissueRDDLResult":
			reissueResultMsg, ok := msg.(*types.MsgReissueRDDLResult)
			if ok {
				ctx, err = cv.handleMsg(ctx, reissueResultMsg)
			}
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
		return ctx, errorsmod.Wrapf(types.ErrRestrictedMsg, "error during CheckTx or ReCheckTx")
	}
	return ctx, nil
}
