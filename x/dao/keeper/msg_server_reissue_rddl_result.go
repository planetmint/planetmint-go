package keeper

import (
	"context"
	"strconv"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/planetmint/planetmint-go/x/dao/types"
)

func (k msgServer) ReissueRDDLResult(goCtx context.Context, msg *types.MsgReissueRDDLResult) (*types.MsgReissueRDDLResultResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	reissuance, found := k.LookupReissuance(ctx, msg.GetBlockHeight())
	if !found {
		return nil, errorsmod.Wrapf(types.ErrReissuanceNotFound, " for provided block height %s", strconv.FormatInt(msg.GetBlockHeight(), 10))
	}
	if reissuance.GetBlockHeight() != msg.GetBlockHeight() {
		return nil, errorsmod.Wrapf(types.ErrWrongBlockHeight, " for provided block height %s", strconv.FormatInt(msg.GetBlockHeight(), 10))
	}
	if reissuance.GetProposer() != msg.GetProposer() {
		return nil, errorsmod.Wrapf(types.ErrInvalidProposer, " for provided block height %s", strconv.FormatInt(msg.GetBlockHeight(), 10))
	}
	if reissuance.GetTxID() != "" {
		return nil, errorsmod.Wrapf(types.ErrTXAlreadySet, " for provided block height %s", strconv.FormatInt(msg.GetBlockHeight(), 10))
	}
	reissuance.TxID = msg.GetTxID()

	k.StoreReissuance(ctx, reissuance)

	return &types.MsgReissueRDDLResultResponse{}, nil
}
