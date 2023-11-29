package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/planetmint/planetmint-go/util"
	"github.com/planetmint/planetmint-go/x/dao/types"
)

func (k msgServer) ReportPopResult(goCtx context.Context, msg *types.MsgReportPopResult) (*types.MsgReportPopResultResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := util.ValidateStruct(*msg.Challenge)
	if err != nil {
		return nil, errorsmod.Wrapf(types.ErrInvalidChallenge, err.Error())
	}

	if isInitiator(*msg.Challenge) {
		err = k.issuePoPRewards(*msg.Challenge)
		if err != nil {
			return nil, errorsmod.Wrapf(types.ErrFailedPoPRewardsIssuance, err.Error())
		}
	}

	k.StoreChallenge(ctx, *msg.Challenge)

	return &types.MsgReportPopResultResponse{}, nil
}

// TODO: ensuer issuePoPrewards is only called once per PoP on all validators
func (k msgServer) issuePoPRewards(_ types.Challenge) (err error) {
	// cfg := config.GetConfig()
	// client := osc.NewClient(cfg.WatchmenEndpoint, 1234)

	// TODO will be reintegrated with by merging branch 184-implement-staged-claim
	// TODO: finalize message and endpoint
	// msg := osc.NewMessage("/rddl/token")
	// msg.Append(challenge.Challenger)
	// msg.Append(challenge.Challengee)
	// err := client.Send(msg)

	return err
}

// TODO: implement check if node is responsible for triggering issuance
func isInitiator(_ types.Challenge) bool {
	return false
}
