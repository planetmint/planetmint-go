package keeper

import (
	"context"
	"fmt"
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
		fmt.Println("REISSUE: Asset")
		txID, err := util.ReissueAsset(msg.Tx)
		if err == nil {
			// 3. notarize result by notarizing the liquid tx-id
			_ = util.SendRDDLReissuanceResult(ctx, msg.GetProposer(), txID, msg.GetBlockheight())
			//TODO verify and  resolve error
		} else {
			fmt.Println("REISSUE: Asset reissuance failure")
		}
		//TODO: reissuance need to be initiated otherwise
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
			pop_distribution.FirstPop = msg.GetBlockHeight()
			pop_distribution.LastPop = msg.GetBlockHeight()
			pop_distribution.RddlAmount = cmd_args[2]
		} else {
			pop_distribution.LastPop = msg.GetBlockHeight()
			amount_to_add, err := strconv.ParseUint(cmd_args[2], 10, 64)
			if err != nil {

			}
			previous_sum, err := strconv.ParseUint(cmd_args[2], 10, 64)
			if err != nil {

			}
			new_value := previous_sum + amount_to_add
			pop_distribution.RddlAmount = strconv.FormatUint(new_value, 10)

		}

	} else {
		pop_distribution.FirstPop = msg.GetBlockHeight()
		pop_distribution.LastPop = msg.GetBlockHeight()
		pop_distribution.RddlAmount = cmd_args[2]
	}
	k.StorePoPDistribution(ctx, pop_distribution)

	return &types.MsgReissueRDDLProposalResponse{}, nil
}
