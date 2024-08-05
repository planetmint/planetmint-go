package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/planetmint/planetmint-go/x/machine/types"
)

func (k msgServer) MintProduction(goCtx context.Context, msg *types.MsgMintProduction) (*types.MsgMintProductionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	addr, err := sdk.AccAddressFromBech32(msg.GetCreator())
	if err != nil {
		return nil, errorsmod.Wrapf(types.ErrInvalidAddress, "invalid address: %s", msg.GetCreator())
	}

	err = k.bankKeeper.MintCoins(ctx, types.ModuleName, msg.GetProduction())
	if err != nil {
		return nil, errorsmod.Wrapf(types.ErrMintFailed, "error while minting production %v for address %s", msg.GetProduction(), msg.GetCreator())
	}

	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, addr, msg.GetProduction())
	if err != nil {
		return nil, errorsmod.Wrapf(types.ErrTransferFailed, "error while transferring %v production to address %s", msg.GetProduction(), msg.GetCreator())
	}

	return &types.MsgMintProductionResponse{}, nil
}
