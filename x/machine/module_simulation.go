package machine

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/planetmint/planetmint-go/testutil/sample"
	machinesimulation "github.com/planetmint/planetmint-go/x/machine/simulation"
	"github.com/planetmint/planetmint-go/x/machine/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = machinesimulation.FindAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
	_ = rand.Rand{}
)

const (
	opWeightMsgAttestMachine = "op_weight_msg_attest_machine"
	// TODO: Determine the simulation weight value
	defaultWeightMsgAttestMachine int = 100

	opWeightMsgRegisterTrustAnchor = "op_weight_msg_register_trust_anchor"
	// TODO: Determine the simulation weight value
	defaultWeightMsgRegisterTrustAnchor int = 100

	opWeightMsgNotarizeLiquidAsset = "op_weight_msg_notarize_liquid_asset"
	// TODO: Determine the simulation weight value
	defaultWeightMsgNotarizeLiquidAsset int = 100

	opWeightMsgUpdateParams = "op_weight_msg_update_params"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateParams int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	machineGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&machineGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {
	// Implement if needed
}

// ProposalContents doesn't return any content functions for governance proposals.
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalMsg {
	return nil
}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgAttestMachine int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgAttestMachine, &weightMsgAttestMachine, nil,
		func(_ *rand.Rand) {
			weightMsgAttestMachine = defaultWeightMsgAttestMachine
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgAttestMachine,
		machinesimulation.SimulateMsgAttestMachine(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgRegisterTrustAnchor int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgRegisterTrustAnchor, &weightMsgRegisterTrustAnchor, nil,
		func(_ *rand.Rand) {
			weightMsgRegisterTrustAnchor = defaultWeightMsgRegisterTrustAnchor
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRegisterTrustAnchor,
		machinesimulation.SimulateMsgRegisterTrustAnchor(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgNotarizeLiquidAsset int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgNotarizeLiquidAsset, &weightMsgNotarizeLiquidAsset, nil,
		func(_ *rand.Rand) {
			weightMsgNotarizeLiquidAsset = defaultWeightMsgNotarizeLiquidAsset
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgNotarizeLiquidAsset,
		machinesimulation.SimulateMsgNotarizeLiquidAsset(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateParams int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateParams, &weightMsgUpdateParams, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateParams = defaultWeightMsgUpdateParams
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateParams,
		machinesimulation.SimulateMsgUpdateParams(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(_ module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgAttestMachine,
			defaultWeightMsgAttestMachine,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				machinesimulation.SimulateMsgAttestMachine(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgRegisterTrustAnchor,
			defaultWeightMsgRegisterTrustAnchor,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				machinesimulation.SimulateMsgRegisterTrustAnchor(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgNotarizeLiquidAsset,
			defaultWeightMsgNotarizeLiquidAsset,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				machinesimulation.SimulateMsgNotarizeLiquidAsset(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateParams,
			defaultWeightMsgUpdateParams,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				machinesimulation.SimulateMsgUpdateParams(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
