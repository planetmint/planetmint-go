package ante

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	ante "github.com/cosmos/cosmos-sdk/x/auth/ante"
	authsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

// HandlerOptions are the options required for constructing a default SDK AnteHandler.
type HandlerOptions struct {
	AccountKeeper          AccountKeeper
	BankKeeper             BankKeeper
	ExtensionOptionChecker ante.ExtensionOptionChecker
	FeegrantKeeper         FeegrantKeeper
	SignModeHandler        authsigning.SignModeHandler
	SigGasConsumer         func(meter sdk.GasMeter, sig signing.SignatureV2, params authtypes.Params) error
	TxFeeChecker           TxFeeChecker
	MachineKeeper          MachineKeeper
	DaoKeeper              DaoKeeper
	StakingKeeper          StakingKeeper
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
	if options.DaoKeeper == nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "dao keeper is required for ante builder")
	}
	if options.StakingKeeper == nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "staking keeper is required for ante builder")
	}

	anteDecorators := []sdk.AnteDecorator{
		NewSetUpContextDecorator(options.DaoKeeper), // outermost AnteDecorator. SetUpContext must be called first
		ante.NewExtensionOptionsDecorator(options.ExtensionOptionChecker),
		NewGasKVCostDecorator(options.StakingKeeper),
		ante.NewValidateBasicDecorator(),
		ante.NewTxTimeoutHeightDecorator(),
		ante.NewValidateMemoDecorator(options.AccountKeeper),
		NewCheckValidatorDecorator(options.StakingKeeper),
		NewCheckMachineDecorator(options.MachineKeeper),
		NewCheckMintAddressDecorator(options.DaoKeeper),
		ante.NewConsumeGasForTxSizeDecorator(options.AccountKeeper),
		NewDeductFeeDecorator(options.AccountKeeper, options.BankKeeper, options.FeegrantKeeper, options.TxFeeChecker),
		ante.NewSetPubKeyDecorator(options.AccountKeeper), // SetPubKeyDecorator must be called before all signature verification decorators
		ante.NewValidateSigCountDecorator(options.AccountKeeper),
		ante.NewSigGasConsumeDecorator(options.AccountKeeper, options.SigGasConsumer),
		ante.NewSigVerificationDecorator(options.AccountKeeper, options.SignModeHandler),
		ante.NewIncrementSequenceDecorator(options.AccountKeeper),
	}

	return sdk.ChainAnteDecorators(anteDecorators...), nil
}
