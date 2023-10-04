package ante

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	daotypes "github.com/planetmint/planetmint-go/x/dao/types"
)

type CheckMintAddressDecorator struct {
	MintAddress string
}

func NewCheckMintAddressDecorator(mintAddress string) CheckMintAddressDecorator {
	return CheckMintAddressDecorator{
		MintAddress: mintAddress,
	}
}

func (cmad CheckMintAddressDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (sdk.Context, error) {
	for _, msg := range tx.GetMsgs() {
		if sdk.MsgTypeURL(msg) == "/planetmintgo.dao.MsgMintToken" {
			mintMsg, ok := msg.(*daotypes.MsgMintToken)
			if ok {
				if mintMsg.Creator != cmad.MintAddress {
					return ctx, errorsmod.Wrapf(daotypes.ErrInvalidMintAddress, "expected: %s; got: %s", cmad.MintAddress, mintMsg.Creator)
				}
			}
		}
	}

	return next(ctx, tx, simulate)
}
