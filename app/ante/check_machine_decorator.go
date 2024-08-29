package ante

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	assettypes "github.com/planetmint/planetmint-go/x/asset/types"
	daotypes "github.com/planetmint/planetmint-go/x/dao/types"
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

func (cm CheckMachineDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (_ sdk.Context, err error) {
	for _, msg := range tx.GetMsgs() {
		switch sdk.MsgTypeURL(msg) {
		case "/planetmintgo.asset.MsgNotarizeAsset":
			notarizeMsg, ok := msg.(*assettypes.MsgNotarizeAsset)
			if ok {
				ctx, err = cm.handleNotarizeAsset(ctx, notarizeMsg)
			}
		case "/planetmintgo.machine.MsgAttestMachine":
			attestMsg, ok := msg.(*machinetypes.MsgAttestMachine)
			if ok {
				ctx, err = cm.handleAttestMachine(ctx, attestMsg)
			}
		case "planetmintgo.dao.MsgReportPoPResult":
			popMsg, ok := msg.(*daotypes.MsgReportPopResult)
			if ok {
				ctx, err = cm.handlePopResult(ctx, popMsg)
			}
		default:
			continue
		}
	}

	if err != nil {
		return ctx, err
	}

	return next(ctx, tx, simulate)
}

func (cm CheckMachineDecorator) handleNotarizeAsset(ctx sdk.Context, notarizeMsg *assettypes.MsgNotarizeAsset) (sdk.Context, error) {
	_, found := cm.mk.GetMachineIndexByAddress(ctx, notarizeMsg.GetCreator())
	if !found {
		return ctx, errorsmod.Wrap(machinetypes.ErrMachineNotFound, ErrorAnteContext)
	}
	return ctx, nil
}

func (cm CheckMachineDecorator) handleAttestMachine(ctx sdk.Context, attestMsg *machinetypes.MsgAttestMachine) (sdk.Context, error) {
	if attestMsg.GetCreator() != attestMsg.Machine.GetAddress() {
		return ctx, errorsmod.Wrap(machinetypes.ErrMachineIsNotCreator, ErrorAnteContext)
	}
	_, activated, found := cm.mk.GetTrustAnchor(ctx, attestMsg.Machine.MachineId)
	if !found {
		return ctx, errorsmod.Wrap(machinetypes.ErrTrustAnchorNotFound, ErrorAnteContext)
	}
	if activated {
		return ctx, errorsmod.Wrap(machinetypes.ErrTrustAnchorAlreadyInUse, ErrorAnteContext)
	}
	return ctx, nil
}

func (cm CheckMachineDecorator) handlePopResult(ctx sdk.Context, popMsg *daotypes.MsgReportPopResult) (sdk.Context, error) {
	_, found := cm.mk.GetMachineIndexByAddress(ctx, popMsg.GetCreator())
	if !found {
		return ctx, errorsmod.Wrap(machinetypes.ErrMachineNotFound, ErrorAnteContext)
	}
	return ctx, nil
}
