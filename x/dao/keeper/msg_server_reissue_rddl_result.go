package keeper

import (
	"context"
	"strconv"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/planetmint/planetmint-go/config"
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

	k.resolveStagedClaims(ctx, uint64(reissuance.BlockHeight))
	k.StoreReissuance(ctx, reissuance)

	return &types.MsgReissueRDDLResultResponse{}, nil
}

func (k msgServer) resolveStagedClaims(ctx sdk.Context, height uint64) (err error) {
	cfg := config.GetConfig()
	// lookup all challenges since the last reissuance
	challenges, err := k.GetChallengeRange(ctx, height-uint64(cfg.DistributionEpochs))
	if err != nil {
		return err
	}

	for _, challenge := range challenges {
		err = k.convertClaim(ctx, challenge.Challengee)
		if err != nil {
			return err
		}
		err = k.convertClaim(ctx, challenge.Challenger)
		if err != nil {
			return err
		}
	}

	return
}

func (k msgServer) convertClaim(ctx sdk.Context, addr string) (err error) {
	accAddress, err := sdk.AccAddressFromBech32(addr)
	if err != nil {
		return err
	}

	stagedClaim := k.bankKeeper.GetBalance(ctx, accAddress, "stagedCRDDL")
	claim := sdk.NewCoins(sdk.NewCoin("cRDDL", stagedClaim.Amount))

	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, accAddress, types.ModuleName, sdk.NewCoins(stagedClaim))
	if err != nil {
		return err
	}

	err = k.bankKeeper.BurnCoins(ctx, types.ModuleName, sdk.NewCoins(stagedClaim))
	if err != nil {
		return err
	}

	err = k.bankKeeper.MintCoins(ctx, types.ModuleName, claim)
	if err != nil {
		return err
	}

	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, accAddress, claim)
	if err != nil {
		return err
	}

	return
}
