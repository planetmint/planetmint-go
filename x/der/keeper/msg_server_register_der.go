package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/planetmint/planetmint-go/util"
	"github.com/planetmint/planetmint-go/x/der/types"
	machinesTypes "github.com/planetmint/planetmint-go/x/machine/types"
)

func (k msgServer) RegisterDER(goCtx context.Context, msg *types.MsgRegisterDER) (*types.MsgRegisterDERResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	k.StoreDerAsset(ctx, *msg.Der)

	// Get machine params from MachineKeeper
	params := k.MachineKeeper.GetParams(ctx)

	// Process NFT issuance if validator is block proposer
	if util.IsValidatorBlockProposer(ctx, k.rootDir) {
		k.handleDERNFTIssuance(goCtx, msg.Der, params)
	} else {
		util.GetAppLogger().Info(ctx, "Not block proposer: skipping DER NFT issuance")
	}

	return &types.MsgRegisterDERResponse{}, nil
}

func (k msgServer) handleDERNFTIssuance(goCtx context.Context, der *types.DER, params machinesTypes.Params) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	logger := util.GetAppLogger()
	logger.Info(ctx, "Issuing DER NFT: "+der.String())

	err := util.IssueDerNFT(goCtx, der,
		params.AssetRegistryScheme,
		params.AssetRegistryDomain,
		params.AssetRegistryPath,
	)

	if err != nil {
		logger.Error(ctx, err, "DER NFT issuance failed")
	} else {
		logger.Info(ctx, "DER NFT issuance successful: "+der.ZigbeeID)
	}
}
