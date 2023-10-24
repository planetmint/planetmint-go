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

	validator_identity, valid_result := util.GetValidatorCometBFTIdentity(ctx)
	if valid_result && msg.Proposer == validator_identity {
		// 1. sign tx
		// 2. broadcast tx
		ctx.Logger().Debug("REISSUE: Asset")
		txID, err := util.ReissueAsset(msg.Tx)
		if err == nil {
			// 3. notarize result by notarizing the liquid tx-id
			_ = util.SendRDDLReissuanceResult(ctx, msg.GetProposer(), txID, msg.GetBlockheight())
			//TODO verify and  resolve error
		} else {
			ctx.Logger().Debug("REISSUE: Asset reissuance failure")
		}
	}

	var reissuance types.Reissuance
	reissuance.BlockHeight = msg.GetBlockheight()
	reissuance.Proposer = msg.GetProposer()
	reissuance.Rawtx = msg.GetTx()
	k.StoreReissuance(ctx, reissuance)

	cmd_args := strings.Split(msg.Tx, " ")
	pop_distribution, found := k.LookupPoPDistribution(ctx)
	if found {
		if pop_distribution.GetFirstPop() == 0 && pop_distribution.GetLastPop() == 0 {
			pop_distribution.FirstPop = msg.GetBlockheight()
			pop_distribution.LastPop = msg.GetBlockheight()
			pop_distribution.RddlAmount = cmd_args[2]
		} else {
			pop_distribution.LastPop = msg.GetBlockheight()
			amount_to_add, err := strconv.ParseUint(cmd_args[2], 10, 64)
			if err != nil {
				ctx.Logger().Error("ReissueProposal: could not parse string to integer: ", err.Error())
			}
			previous_sum, err := strconv.ParseUint(cmd_args[2], 10, 64)
			if err != nil {
				ctx.Logger().Error("ReissueProposal: could not parse string to integer: ", err.Error())
			}
			new_value := previous_sum + amount_to_add
			pop_distribution.RddlAmount = strconv.FormatUint(new_value, 10)

		}

	} else {
		pop_distribution.FirstPop = msg.GetBlockheight()
		pop_distribution.LastPop = msg.GetBlockheight()
		pop_distribution.RddlAmount = cmd_args[2]
	}
	k.StorePoPDistribution(ctx, pop_distribution)

	return &types.MsgReissueRDDLProposalResponse{}, nil
}
