package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/planetmint/planetmint-go/util"
	"github.com/planetmint/planetmint-go/x/dao/types"
)

func (k msgServer) ReissueRDDLProposal(goCtx context.Context, msg *types.MsgReissueRDDLProposal) (*types.MsgReissueRDDLProposalResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	validator_identity, valid_result := util.GetValidatorCometBFTIdentity(ctx)
	if valid_result && msg.Proposer == validator_identity {
		// 1. sign tx
		// 2. broadcast tx
		// 3. notarize result by notarizing the liquid tx-id
	}
	// 4. notarize msg in any case

	return &types.MsgReissueRDDLProposalResponse{}, nil
}
