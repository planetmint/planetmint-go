package ante

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/planetmint/planetmint-go/util"
	daotypes "github.com/planetmint/planetmint-go/x/dao/types"
)

var (
	anteHandlerTag = "ante handler: "
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
				util.GetAppLogger().Debug(ctx, anteHandlerTag+"received reissuance proposal: "+MsgProposal.String())
				isValid := cmad.dk.IsValidReissuanceProposal(ctx, MsgProposal)
				if !isValid {
					util.GetAppLogger().Info(ctx, anteHandlerTag+"rejected reissuance proposal")
					return ctx, errorsmod.Wrapf(daotypes.ErrReissuanceProposal, ErrorAnteContext)
				}
				util.GetAppLogger().Debug(ctx, anteHandlerTag+"accepted reissuance proposal: "+MsgProposal.String())
			}
		}
	}

	return next(ctx, tx, simulate)
}
