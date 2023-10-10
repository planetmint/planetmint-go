package dao

import (
	"math/rand"

	"github.com/planetmint/planetmint-go/testutil/sample"
	daosimulation "github.com/planetmint/planetmint-go/x/dao/simulation"
	"github.com/planetmint/planetmint-go/x/dao/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = daosimulation.FindAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
	_ = rand.Rand{}
)

const (
	opWeightMsgReissueRDDLProposal = "op_weight_msg_reissue_rddl_proposal"
	// TODO: Determine the simulation weight value
	defaultWeightMsgReissueRDDLProposal int = 100

	opWeightMsgReissueRDDLResult = "op_weight_msg_reissue_rddl_result"
	// TODO: Determine the simulation weight value
	defaultWeightMsgReissueRDDLResult int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	daoGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&daoGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// ProposalContents doesn't return any content functions for governance proposals.
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalMsg {
	return nil
}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgReissueRDDLProposal int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgReissueRDDLProposal, &weightMsgReissueRDDLProposal, nil,
		func(_ *rand.Rand) {
			weightMsgReissueRDDLProposal = defaultWeightMsgReissueRDDLProposal
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgReissueRDDLProposal,
		daosimulation.SimulateMsgReissueRDDLProposal(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgReissueRDDLResult int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgReissueRDDLResult, &weightMsgReissueRDDLResult, nil,
		func(_ *rand.Rand) {
			weightMsgReissueRDDLResult = defaultWeightMsgReissueRDDLResult
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgReissueRDDLResult,
		daosimulation.SimulateMsgReissueRDDLResult(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgReissueRDDLProposal,
			defaultWeightMsgReissueRDDLProposal,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				daosimulation.SimulateMsgReissueRDDLProposal(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgReissueRDDLResult,
			defaultWeightMsgReissueRDDLResult,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				daosimulation.SimulateMsgReissueRDDLResult(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
