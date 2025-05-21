package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/planetmint/planetmint-go/x/der/keeper"
	"github.com/planetmint/planetmint-go/x/der/types"
)

func SimulateMsgNotarizeLiquidDerAsset(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgNotarizeLiquidDerAsset{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the NotarizeLiquidDerAsset simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "NotarizeLiquidDerAsset simulation not implemented"), nil, nil
	}
}
