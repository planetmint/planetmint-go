package ante

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/planetmint/planetmint-go/util"
	daotypes "github.com/planetmint/planetmint-go/x/dao/types"
)

type CheckReissuanceDecorator struct {
	dk DaoKeeper
}

func NewCheckReissuanceDecorator(dk DaoKeeper) CheckReissuanceDecorator {
	return CheckReissuanceDecorator{
		dk: dk,
	}
}

func (cmad CheckReissuanceDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (sdk.Context, error) {
	for _, msg := range tx.GetMsgs() {
		if sdk.MsgTypeURL(msg) == "/planetmintgo.dao.MsgReissueRDDLProposal" {
			MsgProposal, ok := msg.(*daotypes.MsgReissueRDDLProposal)
			if ok {
				util.GetAppLogger().Debug(ctx, "ante handler - received re-issuance roposal")
				isValid := cmad.dk.IsValidReIssuanceProposal(ctx, MsgProposal)
				if !isValid {
					util.GetAppLogger().Info(ctx, "ante handler: rejected re-issuance proposal")
					return ctx, errorsmod.Wrapf(daotypes.ErrReissuanceProposal, "error during CheckTx or ReCheckTx")
				}
				util.GetAppLogger().Debug(ctx, "ante handler - accepted re-issuance proposal")
			}
		}
	}

	return next(ctx, tx, simulate)
}
