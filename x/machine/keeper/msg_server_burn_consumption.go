package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/planetmint/planetmint-go/x/machine/types"
)

func (k msgServer) BurnConsumption(goCtx context.Context, msg *types.MsgBurnConsumption) (*types.MsgBurnConsumptionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	addr, err := sdk.AccAddressFromBech32(msg.GetCreator())
	if err != nil {
		return nil, errorsmod.Wrapf(types.ErrInvalidAddress, "invalid address: %s", msg.GetCreator())
	}

	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, addr, types.ModuleName, msg.GetConsumption())
	if err != nil {
		return nil, errorsmod.Wrapf(types.ErrTransferFailed, "error while transferring %v consumption from address %s", msg.GetConsumption(), msg.GetCreator())
	}

	err = k.bankKeeper.BurnCoins(ctx, types.ModuleName, msg.GetConsumption())
	if err != nil {
		return nil, errorsmod.Wrapf(types.ErrBurnFailed, "error while burning consumption %v for address %s", msg.GetConsumption(), msg.GetCreator())
	}

	return &types.MsgBurnConsumptionResponse{}, nil
}
