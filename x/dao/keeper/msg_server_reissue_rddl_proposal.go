package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/planetmint/planetmint-go/util"
	"github.com/planetmint/planetmint-go/x/dao/types"
)

func (k msgServer) ReissueRDDLProposal(goCtx context.Context, msg *types.MsgReissueRDDLProposal) (*types.MsgReissueRDDLProposalResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	validatorIdentity, validResult := util.GetValidatorCometBFTIdentity(ctx)
	if validResult && msg.Proposer == validatorIdentity {
		util.GetAppLogger().Info(ctx, "REISSUE: Asset")
		txID, err := util.ReissueAsset(msg.Tx)
		if err != nil {
			util.GetAppLogger().Error(ctx, "REISSUE: Asset reissuance failed: "+err.Error())
		}
		// 3. notarize result by notarizing the liquid tx-id
		util.SendRDDLReissuanceResult(goCtx, msg.GetProposer(), txID, msg.GetBlockHeight())
	}

	var reissuance types.Reissuance
	reissuance.BlockHeight = msg.GetBlockHeight()
	reissuance.Proposer = msg.GetProposer()
	reissuance.Rawtx = msg.GetTx()
	k.StoreReissuance(ctx, reissuance)
	return &types.MsgReissueRDDLProposalResponse{}, nil
}
