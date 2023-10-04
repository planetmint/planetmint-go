package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/planetmint/planetmint-go/config"
	"github.com/planetmint/planetmint-go/x/dao/types"
)

func (k msgServer) MintToken(goCtx context.Context, msg *types.MsgMintToken) (*types.MsgMintTokenResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	cfg := config.GetConfig()

	amt := msg.GetMintRequest().GetAmount()
	beneficiary := msg.GetMintRequest().GetBeneficiary()

	coin := sdk.NewCoin(cfg.TokenDenom, sdk.NewIntFromUint64(amt))
	coins := sdk.NewCoins(coin)

	err := k.bankKeeper.MintCoins(ctx, types.ModuleName, coins)
	if err != nil {
		return nil, errorsmod.Wrapf(types.ErrMintFailed, "error while minting %v token for address %s", amt, beneficiary)
	}

	beneficiaryAddr, err := sdk.AccAddressFromBech32(beneficiary)
	if err != nil {
		return nil, errorsmod.Wrapf(types.ErrInvalidAddress, "for provided address %s", beneficiary)
	}

	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, beneficiaryAddr, coins)
	if err != nil {
		return nil, errorsmod.Wrapf(types.ErrTransferFailed, "error while transferring %v token to address %s", amt, beneficiary)
	}

	k.StoreMintRequest(ctx, *msg.MintRequest)

	return &types.MsgMintTokenResponse{}, nil
}
