package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/planetmint/planetmint-go/errormsg"
	"github.com/planetmint/planetmint-go/x/dao/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Challenges(goCtx context.Context, req *types.QueryChallengesRequest) (*types.QueryChallengesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, errormsg.InvalidRequest)
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	challenges := make([]types.Challenge, 0)

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ChallengeKey))
	pageRes, err := query.Paginate(store, req.Pagination, func(_ []byte, value []byte) (err error) {
		var challenge types.Challenge
		err = challenge.Unmarshal(value)
		challenges = append(challenges, challenge)
		return
	})

	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "paginate: %v", err)
	}

	return &types.QueryChallengesResponse{Challenges: challenges, Pagination: pageRes}, nil
}
