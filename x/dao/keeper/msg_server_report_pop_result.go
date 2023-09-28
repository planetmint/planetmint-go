package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hypebeast/go-osc/osc"
	"github.com/planetmint/planetmint-go/config"
	"github.com/planetmint/planetmint-go/util"
	"github.com/planetmint/planetmint-go/x/dao/types"
)

func (k msgServer) ReportPopResult(goCtx context.Context, msg *types.MsgReportPopResult) (*types.MsgReportPopResultResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := util.ValidateStruct(*msg.Challenge)
	if err != nil {
		return nil, errorsmod.Wrapf(types.ErrInvalidChallenge, err.Error())
	}

	k.issuePoPRewards(*msg.Challenge)
	k.StoreChallenge(ctx, *msg.Challenge)

	return &types.MsgReportPopResultResponse{}, nil
}

// TODO: ensuer issuePoPrewards is only called once per PoP on all validators
func (k msgServer) issuePoPRewards(challenge types.Challenge) error {
	cfg := config.GetConfig()
	client := osc.NewClient(cfg.WatchmenEndpoint, 1234)

	// TODO: finalize message and endpoint
	msg := osc.NewMessage("/rddl/token")
	msg.Append(challenge.Challenger)
	msg.Append(challenge.Challengee)
	err := client.Send(msg)

	return err
}
