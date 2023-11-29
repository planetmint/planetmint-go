package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/planetmint/planetmint-go/x/dao/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetChallenge(goCtx context.Context, req *types.QueryGetChallengeRequest) (*types.QueryGetChallengeResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	challenge, found := k.LookupChallenge(ctx, req.GetHeight())
	if found == false {
		return nil, status.Error(codes.NotFound, "challenge not found")
	}

	return &types.QueryGetChallengeResponse{Challenge: &challenge}, nil
}
