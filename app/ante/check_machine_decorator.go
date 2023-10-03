package ante

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	assettypes "github.com/planetmint/planetmint-go/x/asset/types"
	machinetypes "github.com/planetmint/planetmint-go/x/machine/types"
)

type CheckMachineDecorator struct {
	mk MachineKeeper
}

func NewCheckMachineDecorator(mk MachineKeeper) CheckMachineDecorator {
	return CheckMachineDecorator{
		mk: mk,
	}
}

func (cm CheckMachineDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (sdk.Context, error) {
	for _, msg := range tx.GetMsgs() {
		switch sdk.MsgTypeURL(msg) {
		case "/planetmintgo.asset.MsgNotarizeAsset":
			notarizeMsg, ok := msg.(*assettypes.MsgNotarizeAsset)
			if ok {
				_, found := cm.mk.GetMachineIndexByAddress(ctx, notarizeMsg.GetCreator())
				if !found {
					return ctx, errorsmod.Wrapf(machinetypes.ErrMachineNotFound, "error during CheckTx or ReCheckTx")
				}
			}
		case "/planetmintgo.machine.MsgAttestMachine":
			attestMsg, ok := msg.(*machinetypes.MsgAttestMachine)
			if ok {
				if attestMsg.GetCreator() != attestMsg.Machine.GetAddress() {
					return ctx, errorsmod.Wrapf(machinetypes.ErrMachineIsNotCreator, "error during CheckTx or ReCheckTx")
				}
				_, activated, found := cm.mk.GetTrustAnchor(ctx, attestMsg.Machine.MachineId)
				if !found {
					return ctx, errorsmod.Wrapf(machinetypes.ErrTrustAnchorNotFound, "error during CheckTx or ReCheckTx")
				}
				if activated {
					return ctx, errorsmod.Wrapf(machinetypes.ErrTrustAnchorAlreadyInUse, "error during CheckTx or ReCheckTx")
				}
			}
		default:
			continue
		}
	}

	return next(ctx, tx, simulate)
}
