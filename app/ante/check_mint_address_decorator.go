package ante

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	daotypes "github.com/planetmint/planetmint-go/x/dao/types"
)

type CheckMintAddressDecorator struct {
	dk DaoKeeper
}

func NewCheckMintAddressDecorator(dk DaoKeeper) CheckMintAddressDecorator {
	return CheckMintAddressDecorator{
		dk: dk,
	}
}

// TODO: repurpose for all dao param addresses
func (cmad CheckMintAddressDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (sdk.Context, error) {
	for _, msg := range tx.GetMsgs() {
		if sdk.MsgTypeURL(msg) != "/planetmintgo.dao.MsgMintToken" {
			continue
		}
		mintMsg, ok := msg.(*daotypes.MsgMintToken)
		if !ok {
			continue
		}
		if mintMsg.Creator != cmad.dk.GetMintAddress(ctx) {
			return ctx, errorsmod.Wrapf(daotypes.ErrInvalidMintAddress, "expected: %s; got: %s", cmad.dk.GetMintAddress(ctx), mintMsg.Creator)
		}
		_, found := cmad.dk.GetMintRequestByHash(ctx, mintMsg.GetMintRequest().GetLiquidTxHash())
		if found {
			return ctx, errorsmod.Wrapf(daotypes.ErrAlreadyMinted, "liquid tx hash %s has already been minted", mintMsg.GetMintRequest().GetLiquidTxHash())
		}
	}

	return next(ctx, tx, simulate)
}
