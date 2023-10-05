package ante

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/planetmint/planetmint-go/config"
	daotypes "github.com/planetmint/planetmint-go/x/dao/types"
)

type CheckMintAddressDecorator struct {
	MintAddress string
}

func NewCheckMintAddressDecorator() CheckMintAddressDecorator {
	return CheckMintAddressDecorator{}
}

func (cmad CheckMintAddressDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (sdk.Context, error) {
	for _, msg := range tx.GetMsgs() {
		if sdk.MsgTypeURL(msg) == "/planetmintgo.dao.MsgMintToken" {
			mintMsg, ok := msg.(*daotypes.MsgMintToken)
			if ok {
				cfg := config.GetConfig()
				if mintMsg.Creator != cfg.MintAddress {
					return ctx, errorsmod.Wrapf(daotypes.ErrInvalidMintAddress, "expected: %s; got: %s", cmad.MintAddress, mintMsg.Creator)
				}
			}
		}
	}

	return next(ctx, tx, simulate)
}
