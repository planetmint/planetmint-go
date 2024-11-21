package simulation

import (
	"math"
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	moduletestutil "github.com/cosmos/cosmos-sdk/types/module/testutil"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/planetmint/planetmint-go/testutil/sample"
	"github.com/planetmint/planetmint-go/x/machine/keeper"
	"github.com/planetmint/planetmint-go/x/machine/types"
)

func SimulateMsgRegisterTrustAnchor(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	_ keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)

		_, pubKey := sample.KeyPair(simtypes.RandIntBetween(r, math.MinInt, math.MaxInt))
		msg := &types.MsgRegisterTrustAnchor{
			Creator:     simAccount.Address.String(),
			TrustAnchor: &types.TrustAnchor{Pubkey: pubKey},
		}

		account := ak.GetAccount(ctx, simAccount.Address)
		spendable := bk.SpendableCoins(ctx, account.GetAddress())

		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           moduletestutil.MakeTestEncodingConfig().TxConfig,
			Cdc:             nil,
			Msg:             msg,
			MsgType:         msg.Type(),
			Context:         ctx,
			SimAccount:      simAccount,
			AccountKeeper:   ak,
			Bankkeeper:      bk,
			ModuleName:      types.ModuleName,
			CoinsSpentInMsg: spendable,
		}

		fees := sdk.NewCoins()

		return simulation.GenAndDeliverTx(txCtx, fees)
	}
}
