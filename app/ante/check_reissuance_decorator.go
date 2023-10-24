package ante

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/planetmint/planetmint-go/config"
	"github.com/planetmint/planetmint-go/x/dao"
	daotypes "github.com/planetmint/planetmint-go/x/dao/types"
)

type CheckReissuanceDecorator struct{}

func NewCheckReissuanceDecorator() CheckReissuanceDecorator {
	return CheckReissuanceDecorator{}
}

func (cmad CheckReissuanceDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (sdk.Context, error) {
	for _, msg := range tx.GetMsgs() {
		if sdk.MsgTypeURL(msg) == "/planetmintgo.dao.MsgReissueRDDLProposal" {
			MsgProposal, ok := msg.(*daotypes.MsgReissueRDDLProposal)
			if ok {
				ctx.Logger.Debug("REISSUE: receive Proposal")
				conf := config.GetConfig()
				isValid := dao.IsValidReissuanceCommand(MsgProposal.GetTx(), conf.ReissuanceAsset, MsgProposal.GetBlockheight())
				if !isValid {
					ctx.Logger.Debug("REISSUE: Invalid Proposal")
					return ctx, errorsmod.Wrapf(daotypes.ErrReissuanceProposal, "error during CheckTx or ReCheckTx")
				}
			}
		}
	}

	return next(ctx, tx, simulate)
}
