package ante

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	daotypes "github.com/planetmint/planetmint-go/x/dao/types"
)

type CheckReissuanceDecorator struct {
	MintAddress string
}

func NewCheckReissuanceDecorator() CheckReissuanceDecorator {
	return CheckReissuanceDecorator{}
}

func (cmad CheckReissuanceDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (sdk.Context, error) {
	for _, msg := range tx.GetMsgs() {
		if sdk.MsgTypeURL(msg) == "/planetmintgo.dao.NewMsgReissueRDDLProposal" {
			_, ok := msg.(*daotypes.MsgReissueRDDLProposal)
			//reissueMsg, ok := msg.(*daotypes.NewMsgReissueRDDLProposal)
			if ok {
				// TODO: verify if the messages related PoP (BlockHeight) reflects
				//       what is actually traded within the raw transaction
			}
		}
	}

	return next(ctx, tx, simulate)
}
