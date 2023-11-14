package ante

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/planetmint/planetmint-go/config"
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

func (cmad CheckMintAddressDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (sdk.Context, error) {
	for _, msg := range tx.GetMsgs() {
		if sdk.MsgTypeURL(msg) == "/planetmintgo.dao.MsgMintToken" {
			mintMsg, ok := msg.(*daotypes.MsgMintToken)
			if ok {
				cfg := config.GetConfig()
				if mintMsg.Creator != cfg.MintAddress {
					return ctx, errorsmod.Wrapf(daotypes.ErrInvalidMintAddress, "expected: %s; got: %s", cfg.MintAddress, mintMsg.Creator)
				}
				_, found := cmad.dk.GetMintRequestByHash(ctx, mintMsg.GetMintRequest().GetLiquidTxHash())
				if found {
					return ctx, errorsmod.Wrapf(daotypes.ErrAlreadyMinted, "liquid tx hash %s has already been minted", mintMsg.GetMintRequest().GetLiquidTxHash())
				}
			}
		}
	}

	return next(ctx, tx, simulate)
}
