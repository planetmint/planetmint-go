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

	validatorIdentity, validResult := util.GetValidatorCometBFTIdentity(ctx)
	if validResult && msg.Proposer == validatorIdentity {
		// 1. sign tx
		// 2. broadcast tx
		logger.Debug("REISSUE: Asset")
		txID, err := util.ReissueAsset(msg.Tx)
		if err == nil {
			// 3. notarize result by notarizing the liquid tx-id
			_ = util.SendRDDLReissuanceResult(goCtx, msg.GetProposer(), txID, msg.GetBlockHeight())
			// TODO verify and  resolve error
		} else {
			logger.Error("REISSUE: Asset reissuance failure: " + err.Error())
		}
	}

	var reissuance types.Reissuance
	reissuance.BlockHeight = msg.GetBlockHeight()
	reissuance.Proposer = msg.GetProposer()
	reissuance.Rawtx = msg.GetTx()
	k.StoreReissuance(ctx, reissuance)
	return &types.MsgReissueRDDLProposalResponse{}, nil
}
