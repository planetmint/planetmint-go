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
	opWeightMsgReportPopResult = "op_weight_msg_report_pop_result"
	// TODO: Determine the simulation weight value
	defaultWeightMsgReportPopResult int = 100
	opWeightMsgReissueRDDLProposal      = "op_weight_msg_reissue_rddl_proposal"
	// TODO: Determine the simulation weight value
	defaultWeightMsgReissueRDDLProposal int = 100

	opWeightMsgReissueRDDLResult = "op_weight_msg_reissue_rddl_result"
	// TODO: Determine the simulation weight value
	defaultWeightMsgReissueRDDLResult int = 100

	opWeightMsgInitPop = "op_weight_msg_init_pop"
	// TODO: Determine the simulation weight value
	defaultWeightMsgInitPop int = 100

	opWeightMsgCreateRedeemClaim = "op_weight_msg_redeem_claim"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateRedeemClaim int = 100

	opWeightMsgUpdateRedeemClaim = "op_weight_msg_redeem_claim"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateRedeemClaim int = 100

	opWeightMsgDeleteRedeemClaim = "op_weight_msg_redeem_claim"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteRedeemClaim int = 100

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
		RedeemClaimList: []types.RedeemClaim{
			{
				Creator:      sample.AccAddress(),
				Beneficiary:  "0",
				LiquidTxHash: "0",
			},
			{
				Creator:      sample.AccAddress(),
				Beneficiary:  "1",
				LiquidTxHash: "1",
			},
		},
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

	var weightMsgReportPopResult int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgReportPopResult, &weightMsgReportPopResult, nil,
		func(_ *rand.Rand) {
			weightMsgReportPopResult = defaultWeightMsgReportPopResult
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgReportPopResult,
		daosimulation.SimulateMsgReportPopResult(am.accountKeeper, am.bankKeeper, am.keeper),
	))

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

	var weightMsgInitPop int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgInitPop, &weightMsgInitPop, nil,
		func(_ *rand.Rand) {
			weightMsgInitPop = defaultWeightMsgInitPop
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgInitPop,
		daosimulation.SimulateMsgInitPop(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateRedeemClaim int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateRedeemClaim, &weightMsgCreateRedeemClaim, nil,
		func(_ *rand.Rand) {
			weightMsgCreateRedeemClaim = defaultWeightMsgCreateRedeemClaim
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateRedeemClaim,
		daosimulation.SimulateMsgCreateRedeemClaim(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateRedeemClaim int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateRedeemClaim, &weightMsgUpdateRedeemClaim, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateRedeemClaim = defaultWeightMsgUpdateRedeemClaim
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateRedeemClaim,
		daosimulation.SimulateMsgUpdateRedeemClaim(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteRedeemClaim int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteRedeemClaim, &weightMsgDeleteRedeemClaim, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteRedeemClaim = defaultWeightMsgDeleteRedeemClaim
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteRedeemClaim,
		daosimulation.SimulateMsgDeleteRedeemClaim(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(_ module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgReportPopResult,
			defaultWeightMsgReportPopResult,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				daosimulation.SimulateMsgReportPopResult(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
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
		simulation.NewWeightedProposalMsg(
			opWeightMsgInitPop,
			defaultWeightMsgInitPop,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				daosimulation.SimulateMsgInitPop(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateRedeemClaim,
			defaultWeightMsgCreateRedeemClaim,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				daosimulation.SimulateMsgCreateRedeemClaim(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateRedeemClaim,
			defaultWeightMsgUpdateRedeemClaim,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				daosimulation.SimulateMsgUpdateRedeemClaim(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteRedeemClaim,
			defaultWeightMsgDeleteRedeemClaim,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				daosimulation.SimulateMsgDeleteRedeemClaim(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
