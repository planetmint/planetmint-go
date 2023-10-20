package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/planetmint/planetmint-go/util"
	"github.com/planetmint/planetmint-go/x/dao/types"
)

func (k msgServer) ReissueRDDLProposal(goCtx context.Context, msg *types.MsgReissueRDDLProposal) (*types.MsgReissueRDDLProposalResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	logger := ctx.Logger()

	validator_identity, valid_result := util.GetValidatorCometBFTIdentity(ctx)
	if valid_result && msg.Proposer == validator_identity {
		// 1. sign tx
		// 2. broadcast tx
		logger.Debug("REISSUE: Asset")
		txID, err := util.ReissueAsset(msg.Tx)
		if err == nil {
			// 3. notarize result by notarizing the liquid tx-id
			_ = util.SendRDDLReissuanceResult(ctx, msg.GetProposer(), txID, msg.GetBlockheight())
			//TODO verify and  resolve error
		} else {
			logger.Debug("REISSUE: Asset reissuance failure")
		}
		//TODO: reissuance need to be initiated otherwise
	}

	var reissuance types.Reissuance
	reissuance.BlockHeight = msg.GetBlockheight()
	reissuance.Proposer = msg.GetProposer()
	reissuance.Rawtx = msg.GetTx()
	k.StoreReissuance(ctx, reissuance)

	return &types.MsgReissueRDDLProposalResponse{}, nil
}
