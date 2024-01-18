package simulation

import (
	"math/rand"
	"strconv"

	simappparams "cosmossdk.io/simapp/params"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/planetmint/planetmint-go/x/dao/keeper"
	"github.com/planetmint/planetmint-go/x/dao/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func SimulateMsgCreateRedeemClaim(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)

		i := r.Int()
		msg := &types.MsgCreateRedeemClaim{
			Creator:     simAccount.Address.String(),
			Beneficiary: strconv.Itoa(i),
		}

		_, found := k.GetRedeemClaim(ctx, msg.Beneficiary, uint64(i))
		if found {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "RedeemClaim already exist"), nil, nil
		}

		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           simappparams.MakeTestEncodingConfig().TxConfig,
			Cdc:             nil,
			Msg:             msg,
			MsgType:         msg.Type(),
			Context:         ctx,
			SimAccount:      simAccount,
			ModuleName:      types.ModuleName,
			CoinsSpentInMsg: sdk.NewCoins(),
			AccountKeeper:   ak,
			Bankkeeper:      bk,
		}
		return simulation.GenAndDeliverTxWithRandFees(txCtx)
	}
}

func SimulateMsgUpdateRedeemClaim(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		var (
			simAccount     = simtypes.Account{}
			redeemClaim    = types.RedeemClaim{}
			msg            = &types.MsgUpdateRedeemClaim{}
			allRedeemClaim = k.GetAllRedeemClaim(ctx)
			found          = false
		)
		for _, obj := range allRedeemClaim {
			simAccount, found = FindAccount(accs, obj.Creator)
			if found {
				redeemClaim = obj
				break
			}
		}
		if !found {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "redeemClaim creator not found"), nil, nil
		}
		msg.Creator = simAccount.Address.String()

		msg.Beneficiary = redeemClaim.Beneficiary
		msg.LiquidTxHash = redeemClaim.LiquidTxHash

		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           simappparams.MakeTestEncodingConfig().TxConfig,
			Cdc:             nil,
			Msg:             msg,
			MsgType:         msg.Type(),
			Context:         ctx,
			SimAccount:      simAccount,
			ModuleName:      types.ModuleName,
			CoinsSpentInMsg: sdk.NewCoins(),
			AccountKeeper:   ak,
			Bankkeeper:      bk,
		}
		return simulation.GenAndDeliverTxWithRandFees(txCtx)
	}
}

func SimulateMsgDeleteRedeemClaim(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		var (
			simAccount     = simtypes.Account{}
			redeemClaim    = types.RedeemClaim{}
			msg            = &types.MsgUpdateRedeemClaim{}
			allRedeemClaim = k.GetAllRedeemClaim(ctx)
			found          = false
		)
		for _, obj := range allRedeemClaim {
			simAccount, found = FindAccount(accs, obj.Creator)
			if found {
				redeemClaim = obj
				break
			}
		}
		if !found {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "redeemClaim creator not found"), nil, nil
		}
		msg.Creator = simAccount.Address.String()

		msg.Beneficiary = redeemClaim.Beneficiary
		msg.LiquidTxHash = redeemClaim.LiquidTxHash

		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           simappparams.MakeTestEncodingConfig().TxConfig,
			Cdc:             nil,
			Msg:             msg,
			MsgType:         msg.Type(),
			Context:         ctx,
			SimAccount:      simAccount,
			ModuleName:      types.ModuleName,
			CoinsSpentInMsg: sdk.NewCoins(),
			AccountKeeper:   ak,
			Bankkeeper:      bk,
		}
		return simulation.GenAndDeliverTxWithRandFees(txCtx)
	}
}
