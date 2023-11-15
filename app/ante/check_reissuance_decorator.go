package ante

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/planetmint/planetmint-go/config"
	"github.com/planetmint/planetmint-go/x/dao/keeper"
	daotypes "github.com/planetmint/planetmint-go/x/dao/types"
)

type CheckReissuanceDecorator struct{}

func NewCheckReissuanceDecorator() CheckReissuanceDecorator {
	return CheckReissuanceDecorator{}
}

func (cmad CheckReissuanceDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (sdk.Context, error) {
	logger := ctx.Logger()
	for _, msg := range tx.GetMsgs() {
		if sdk.MsgTypeURL(msg) == "/planetmintgo.dao.MsgReissueRDDLProposal" {
			MsgProposal, ok := msg.(*daotypes.MsgReissueRDDLProposal)
			if ok {
				logger.Debug("REISSUE: receive Proposal")
				conf := config.GetConfig()
				isValid := keeper.IsValidReissuanceCommand(MsgProposal.GetTx(), conf.ReissuanceAsset, MsgProposal.GetBlockHeight())
				if !isValid {
					logger.Debug("REISSUE: Invalid Proposal")
					return ctx, errorsmod.Wrapf(daotypes.ErrReissuanceProposal, "error during CheckTx or ReCheckTx")
				}
			}
		}
	}

	return next(ctx, tx, simulate)
}
