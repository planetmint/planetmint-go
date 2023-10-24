package keeper

import (
	"context"
	"strconv"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/planetmint/planetmint-go/util"
	"github.com/planetmint/planetmint-go/x/dao/types"
)

func (k msgServer) ReissueRDDLProposal(goCtx context.Context, msg *types.MsgReissueRDDLProposal) (*types.MsgReissueRDDLProposalResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	logger := ctx.Logger()

	validatorIdentity, validResult := util.GetValidatorCometBFTIdentity(ctx)
	if validResult && msg.Proposer == validatorIdentity {
		// 1. sign tx
		// 2. broadcast tx
		logger.Debug("REISSUE: Asset")
		txID, err := util.ReissueAsset(msg.Tx)
		if err == nil {
			// 3. notarize result by notarizing the liquid tx-id
			_ = util.SendRDDLReissuanceResult(ctx, msg.GetProposer(), txID, msg.GetBlockHeight())
			//TODO verify and  resolve error
		} else {
			logger.Debug("REISSUE: Asset reissuance failure")
		}
	}

	var reissuance types.Reissuance
	reissuance.BlockHeight = msg.GetBlockHeight()
	reissuance.Proposer = msg.GetProposer()
	reissuance.Rawtx = msg.GetTx()
	k.StoreReissuance(ctx, reissuance)

	cmdArgs := strings.Split(msg.Tx, " ")
	popDistribution, found := k.LookupPoPDistribution(ctx)
	if found {
		if popDistribution.GetFirstPop() == 0 && popDistribution.GetLastPop() == 0 {
			popDistribution.FirstPop = msg.GetBlockHeight()
			popDistribution.LastPop = msg.GetBlockHeight()
			popDistribution.RddlAmount = cmdArgs[2]
		} else {
			popDistribution.LastPop = msg.GetBlockHeight()
			addedAmount, err := strconv.ParseUint(cmdArgs[2], 10, 64)
			if err != nil {
				ctx.Logger().Error("ReissueProposal: could not parse string to integer: ", err.Error())
			}
			previousSum, err := strconv.ParseUint(cmdArgs[2], 10, 64)
			if err != nil {
				ctx.Logger().Error("ReissueProposal: could not parse string to integer: ", err.Error())
			}
			newValue := previousSum + addedAmount
			popDistribution.RddlAmount = strconv.FormatUint(newValue, 10)

		}

	} else {
		popDistribution.FirstPop = msg.GetBlockHeight()
		popDistribution.LastPop = msg.GetBlockHeight()
		popDistribution.RddlAmount = cmdArgs[2]
	}
	k.StorePoPDistribution(ctx, popDistribution)

	return &types.MsgReissueRDDLProposalResponse{}, nil
}
