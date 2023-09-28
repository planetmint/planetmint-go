package ante

import (
	assettypes "github.com/planetmint/planetmint-go/x/asset/types"
	daotypes "github.com/planetmint/planetmint-go/x/dao/types"
	machinetypes "github.com/planetmint/planetmint-go/x/machine/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	ante "github.com/cosmos/cosmos-sdk/x/auth/ante"
	authsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
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
		case "planetmintgo.dao.MsgReportPoPResult":
			popMsg, ok := msg.(*daotypes.MsgReportPopResult)
			if ok {
				_, found := cm.mk.GetMachineIndexByAddress(ctx, popMsg.GetCreator())
				if !found {
					return ctx, errorsmod.Wrapf(machinetypes.ErrMachineNotFound, "error during CheckTx or ReCheckTx")
				}
			}
		default:
			continue
		}
	}

	return next(ctx, tx, simulate)
}

// HandlerOptions are the options required for constructing a default SDK AnteHandler.
type HandlerOptions struct {
	AccountKeeper          AccountKeeper
	BankKeeper             BankKeeper
	ExtensionOptionChecker ante.ExtensionOptionChecker
	FeegrantKeeper         FeegrantKeeper
	SignModeHandler        authsigning.SignModeHandler
	SigGasConsumer         func(meter sdk.GasMeter, sig signing.SignatureV2, params authtypes.Params) error
	TxFeeChecker           ante.TxFeeChecker
	MachineKeeper          MachineKeeper
}

// NewAnteHandler returns an AnteHandler that checks and increments sequence
// numbers, checks signatures & account numbers, and deducts fees from the first
// signer.
func NewAnteHandler(options HandlerOptions) (sdk.AnteHandler, error) {
	if options.AccountKeeper == nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "account keeper is required for ante builder")
	}

	if options.BankKeeper == nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "bank keeper is required for ante builder")
	}

	if options.SignModeHandler == nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "sign mode handler is required for ante builder")
	}

	if options.MachineKeeper == nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "machine keeper is required for ante builder")
	}

	anteDecorators := []sdk.AnteDecorator{
		ante.NewSetUpContextDecorator(), // outermost AnteDecorator. SetUpContext must be called first
		ante.NewExtensionOptionsDecorator(options.ExtensionOptionChecker),
		ante.NewValidateBasicDecorator(),
		ante.NewTxTimeoutHeightDecorator(),
		ante.NewValidateMemoDecorator(options.AccountKeeper),
		NewCheckMachineDecorator(options.MachineKeeper),
		ante.NewConsumeGasForTxSizeDecorator(options.AccountKeeper),
		ante.NewDeductFeeDecorator(options.AccountKeeper, options.BankKeeper, options.FeegrantKeeper, options.TxFeeChecker),
		ante.NewSetPubKeyDecorator(options.AccountKeeper), // SetPubKeyDecorator must be called before all signature verification decorators
		ante.NewValidateSigCountDecorator(options.AccountKeeper),
		ante.NewSigGasConsumeDecorator(options.AccountKeeper, options.SigGasConsumer),
		ante.NewSigVerificationDecorator(options.AccountKeeper, options.SignModeHandler),
		ante.NewIncrementSequenceDecorator(options.AccountKeeper),
	}

	return sdk.ChainAnteDecorators(anteDecorators...), nil
}
