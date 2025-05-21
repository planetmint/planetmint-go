package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/planetmint/planetmint-go/util"
	"github.com/planetmint/planetmint-go/x/der/types"
)

func (k msgServer) RegisterDER(goCtx context.Context, msg *types.MsgRegisterDER) (*types.MsgRegisterDERResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	k.StoreDerAttest(ctx, *msg.Der)

	//TODO: init NFT creation and storag of NFT to DER associations
	// Process NFT issuance if validator is block proposer
	if util.IsValidatorBlockProposer(ctx, k.rootDir) {
		if err := k.handleDERNFTIssuance(goCtx, *msg.Der, params); err != nil {
			return err
		}
	} else {
		util.GetAppLogger().Info(ctx, "Not block proposer: skipping DER NFT issuance")
	}

	return &types.MsgRegisterDERResponse{}, nil
}

func (k msgServer) handleDERNFTIssuance(goCtx context.Context, machine *types.Machine, params types.Params) error {
	ctx := sdk.UnwrapSDKContext(goCtx)
	logger := util.GetAppLogger()
	logger.Info(ctx, "Issuing Machine NFT: "+machine.String())

	err := util.IssueMachineNFT(goCtx, machine,
		params.AssetRegistryScheme,
		params.AssetRegistryDomain,
		params.AssetRegistryPath,
	)

	if err != nil {
		logger.Error(ctx, err, "Machine NFT issuance failed")
		return err
	}

	logger.Info(ctx, "Machine NFT issuance successful: "+machine.String())
	return nil
}
