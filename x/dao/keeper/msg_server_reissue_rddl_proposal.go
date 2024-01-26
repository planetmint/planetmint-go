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

	var reissuance types.Reissuance
	reissuance.BlockHeight = msg.GetBlockHeight()
	reissuance.Proposer = msg.GetProposer()
	reissuance.Command = msg.GetCommand()
	reissuance.FirstIncludedPop = msg.GetFirstIncludedPop()
	reissuance.LastIncludedPop = msg.GetLastIncludedPop()
	k.StoreReissuance(ctx, reissuance)

	validatorIdentity, validResult := util.GetValidatorCometBFTIdentity(ctx)
	if !validResult || msg.Proposer != validatorIdentity {
		util.GetAppLogger().Info(ctx, reissueTag+"Not the proposer. valid result: %t proposer: %s validator identity: %s", validResult, msg.Proposer, validatorIdentity)
		return &types.MsgReissueRDDLProposalResponse{}, nil
	}

	util.GetAppLogger().Info(ctx, reissueTag+"asset: "+msg.GetCommand())
	txID, err := util.ReissueAsset(msg.Command)
	if err != nil {
		util.GetAppLogger().Error(ctx, reissueTag+"asset reissuance failed: "+err.Error())
	}
	util.SendReissuanceResult(goCtx, msg.GetProposer(), txID, msg.GetBlockHeight())

	return &types.MsgReissueRDDLProposalResponse{}, nil
}
