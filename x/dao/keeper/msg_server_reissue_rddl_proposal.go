package keeper

import (
	"context"
	"strings"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/planetmint/planetmint-go/clients/shamir/coordinator"
	"github.com/planetmint/planetmint-go/errormsg"
	"github.com/planetmint/planetmint-go/util"
	"github.com/planetmint/planetmint-go/x/dao/types"
)

var (
	reissueTag = "reissue: "
)

func (k msgServer) ReissueRDDLProposal(goCtx context.Context, msg *types.MsgReissueRDDLProposal) (*types.MsgReissueRDDLProposalResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := k.validateReissuanceProposal(ctx, msg)
	if err != nil {
		return nil, err
	}

	var reissuance types.Reissuance
	reissuance.BlockHeight = msg.GetBlockHeight()
	reissuance.Proposer = msg.GetProposer()
	reissuance.Command = msg.GetCommand()
	reissuance.FirstIncludedPop = msg.GetFirstIncludedPop()
	reissuance.LastIncludedPop = msg.GetLastIncludedPop()
	k.StoreReissuance(ctx, reissuance)

	validatorIdentity, err := util.GetValidatorCometBFTIdentity(ctx, k.RootDir)
	if err != nil {
		util.GetAppLogger().Error(ctx, err, reissueTag+errormsg.CouldNotGetValidatorIdentity)
		return nil, err
	}
	if msg.Proposer != validatorIdentity {
		util.GetAppLogger().Info(ctx, reissueTag+"Not the proposer. proposer: %s validator identity: %s", msg.Proposer, validatorIdentity)
		return &types.MsgReissueRDDLProposalResponse{}, nil
	}

	util.GetAppLogger().Info(ctx, reissueTag+"asset: "+msg.GetCommand())
	cmdArgs := strings.Split(msg.Command, " ")
	txID, err := coordinator.ReIssueAsset(goCtx, cmdArgs[1], cmdArgs[2])
	if err != nil {
		util.GetAppLogger().Error(ctx, err, reissueTag+"asset reissuance failed")
	}
	util.SendReissuanceResult(goCtx, msg.GetProposer(), txID, msg.GetBlockHeight())

	return &types.MsgReissueRDDLProposalResponse{}, nil
}

func (k msgServer) validateReissuanceProposal(ctx sdk.Context, msg *types.MsgReissueRDDLProposal) (err error) {
	util.GetAppLogger().Debug(ctx, reissueTag+"received reissuance proposal: "+msg.String())
	isValid := k.IsValidReissuanceProposal(ctx, msg)
	if !isValid {
		util.GetAppLogger().Info(ctx, reissueTag+"rejected reissuance proposal")
		return errorsmod.Wrap(types.ErrReissuanceProposal, reissueTag)
	}
	util.GetAppLogger().Debug(ctx, reissueTag+"accepted reissuance proposal: "+msg.String())
	return
}
