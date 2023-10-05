package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/planetmint/planetmint-go/x/dao/types"
)

func (k msgServer) ReissueRDDLResult(goCtx context.Context, msg *types.MsgReissueRDDLResult) (*types.MsgReissueRDDLResultResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	reissuance, found := k.GetReissuance(ctx, msg.GetBlockHeight())
	if found != true {
		return nil, errorsmod.Wrapf(types.ErrReissuanceNotFound, " for provided block height %u", msg.GetBlockHeight())
	}
	if reissuance.GetBlockHeight() != msg.GetBlockHeight() {
		return nil, errorsmod.Wrapf(types.ErrWrongBlockHeight, " for provided block height %u", msg.GetBlockHeight())
	}
	if reissuance.GetProposer() != msg.GetProposer() {
		return nil, errorsmod.Wrapf(types.ErrInvalidProposer, " for provided block height %u", msg.GetBlockHeight())
	}
	if reissuance.GetTxId() != "" {
		return nil, errorsmod.Wrapf(types.ErrTXAlreadySet, " for provided block height %u", msg.GetBlockHeight())
	}
	reissuance.TxId = msg.GetTxId()
	k.StoreReissuance(ctx, reissuance)

	return &types.MsgReissueRDDLResultResponse{}, nil
}
