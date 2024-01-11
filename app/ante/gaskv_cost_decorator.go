package ante

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type GasKVCostDecorator struct {
	sk StakingKeeper
}

func NewGasKVCostDecorator(sk StakingKeeper) GasKVCostDecorator {
	return GasKVCostDecorator{sk: sk}
}

func (gc GasKVCostDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (_ sdk.Context, err error) {
	if simulate || ctx.BlockHeight() == 0 {
		return next(ctx, tx, simulate)
	}

	msgs := tx.GetMsgs()
	signers := msgs[0].GetSigners()
	signer := signers[0]

	valAddr := sdk.ValAddress(signer)
	_, found := gc.sk.GetValidator(ctx, valAddr)

	if !found {
		return next(ctx, tx, simulate)
	}

	ctx = ctx.WithKVGasConfig(sdk.GasConfig{
		HasCost:          0,
		DeleteCost:       0,
		ReadCostFlat:     0,
		ReadCostPerByte:  0,
		WriteCostFlat:    0,
		WriteCostPerByte: 0,
		IterNextCostFlat: 0,
	})

	return next(ctx, tx, simulate)
}
