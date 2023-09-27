package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hypebeast/go-osc/osc"
	"github.com/planetmint/planetmint-go/config"
	"github.com/planetmint/planetmint-go/x/dao/types"
)

func (k msgServer) ReportPopResult(goCtx context.Context, msg *types.MsgReportPopResult) (*types.MsgReportPopResultResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO validate only machine node can be creator
	k.issuePoPRewards(*msg.Challenge)
	k.StoreChallenge(ctx, *msg.Challenge)

	return &types.MsgReportPopResultResponse{}, nil
}

func (k msgServer) issuePoPRewards(challenge types.Challenge) error {
	cfg := config.GetConfig()
	client := osc.NewClient(cfg.WatchmenEndpoint, 1234)

	msg := osc.NewMessage("/rddl/token")
	msg.Append(challenge.Challenger)
	msg.Append(challenge.Challengee)
	err := client.Send(msg)

	return err
}
