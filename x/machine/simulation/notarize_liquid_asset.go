package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/planetmint/planetmint-go/x/machine/keeper"
	"github.com/planetmint/planetmint-go/x/machine/types"
)

func SimulateMsgNotarizeLiquidAsset(
	_ types.AccountKeeper,
	_ types.BankKeeper,
	_ keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgNotarizeLiquidAsset{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the NotarizeLiquidAsset simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "NotarizeLiquidAsset simulation not implemented"), nil, nil
	}
}
