package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/planetmint/planetmint-go/util"
	"github.com/planetmint/planetmint-go/x/dao/types"
)

var (
	reissueTag = "reissue: "
)

func (k msgServer) ReissueRDDLProposal(goCtx context.Context, msg *types.MsgReissueRDDLProposal) (*types.MsgReissueRDDLProposalResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	validatorIdentity, validResult := util.GetValidatorCometBFTIdentity(ctx)
	if validResult && msg.Proposer == validatorIdentity {
		util.GetAppLogger().Info(ctx, reissueTag+"asset: "+msg.GetTx())
		txID, err := util.ReissueAsset(msg.Tx)
		if err != nil {
			util.GetAppLogger().Error(ctx, reissueTag+"asset reissuance failed: "+err.Error())
		}
		// 3. notarize result by notarizing the liquid tx-id
		util.SendReissuanceResult(goCtx, msg.GetProposer(), txID, msg.GetBlockHeight())
	} else {
		util.GetAppLogger().Error(ctx, reissueTag+"failed. valid result: %v proposer: %s validator identity: %s", validResult, msg.Proposer, validatorIdentity)
	}

	var reissuance types.Reissuance
	reissuance.BlockHeight = msg.GetBlockHeight()
	reissuance.Proposer = msg.GetProposer()
	reissuance.RawTx = msg.GetTx()
	reissuance.FirstIncludedPop = msg.GetFirstIncludedPop()
	reissuance.LastIncludedPop = msg.GetLastIncludedPop()
	k.StoreReissuance(ctx, reissuance)
	return &types.MsgReissueRDDLProposalResponse{}, nil
}
