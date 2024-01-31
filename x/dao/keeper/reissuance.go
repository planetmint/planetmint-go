package keeper

import (
	"fmt"
	"math"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/planetmint/planetmint-go/util"
	"github.com/planetmint/planetmint-go/x/dao/types"
)

var ReissueCommand string

func init() {
	ReissueCommand = "reissueasset"
}

func GetReissuanceAsStringValue(blockHeight int64, popEpochs int64) string {
	PopNumber := util.GetPopNumber(blockHeight, popEpochs)
	exactCycleID := PopNumber / util.PopsPerCycle

	switch cycleID := math.Floor(exactCycleID); cycleID {
	case 0:
		return "998.85844748"
	case 1:
		return "499.42922374"
	case 2:
		return "249.71461187"
	case 3:
		return "124.85730593"
	case 4:
		return "62.42865296"
	default:
		return "0.0"
	}
}

func GetReissuanceCommand(assetID string, blockHeight int64, popsPerEpoch int64) string {
	return ReissueCommand + " " + assetID + " " + GetReissuanceAsStringValue(blockHeight, popsPerEpoch)
}

func IsValidReissuanceCommand(reissuanceStr string, assetID string, blockHeight int64, popsPerEpoch int64) bool {
	expected := ReissueCommand + " " + assetID + " " + GetReissuanceAsStringValue(blockHeight, popsPerEpoch)
	return reissuanceStr == expected
}

func GetReissuanceCommandForValue(assetID string, value uint64) string {
	return ReissueCommand + " " + assetID + " " + util.UintValueToRDDLTokenString(value)
}

func (k Keeper) CreateNextReissuanceObject(ctx sdk.Context, currentBlockHeight int64) (reissuance types.Reissuance, err error) {
	var lastReissuedPop int64
	lastReissuance, found := k.GetLastReissuance(ctx)
	if found {
		lastReissuedPop = lastReissuance.LastIncludedPop
	}
	reissuanceValue, firstIncludedPop, lastIncludedPop, err := k.ComputeReissuanceValue(ctx, lastReissuedPop, currentBlockHeight)
	if err != nil {
		return
	}

	reissuance.Command = GetReissuanceCommandForValue(k.GetParams(ctx).ReissuanceAsset, reissuanceValue)
	reissuance.BlockHeight = currentBlockHeight
	reissuance.FirstIncludedPop = firstIncludedPop
	reissuance.LastIncludedPop = lastIncludedPop
	return
}

func (k Keeper) IsValidReissuanceProposal(ctx sdk.Context, msg *types.MsgReissueRDDLProposal) (isValid bool) {
	reissuance, err := k.CreateNextReissuanceObject(ctx, msg.GetBlockHeight())
	if err != nil {
		return
	}
	if reissuance.GetBlockHeight() == msg.GetBlockHeight() &&
		reissuance.GetFirstIncludedPop() == msg.GetFirstIncludedPop() &&
		reissuance.GetLastIncludedPop() == msg.GetLastIncludedPop() &&
		reissuance.GetCommand() == msg.GetCommand() &&
		msg.GetProposer() != "" {
		isValid = true
	}
	return
}

func (k Keeper) StoreReissuance(ctx sdk.Context, reissuance types.Reissuance) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ReissuanceBlockHeightKey))
	appendValue := k.cdc.MustMarshal(&reissuance)
	store.Set(util.SerializeInt64(reissuance.BlockHeight), appendValue)
}

func (k Keeper) LookupReissuance(ctx sdk.Context, height int64) (val types.Reissuance, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ReissuanceBlockHeightKey))
	reissuance := store.Get(util.SerializeInt64(height))
	if reissuance == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(reissuance, &val)
	return val, true
}

func (k Keeper) getReissuancesRange(ctx sdk.Context, from int64) (reissuances []types.Reissuance) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ReissuanceBlockHeightKey))

	iterator := store.Iterator(util.SerializeInt64(from), nil)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		reissuance := iterator.Value()
		var reissuanceOrg types.Reissuance
		k.cdc.MustUnmarshal(reissuance, &reissuanceOrg)
		reissuances = append(reissuances, reissuanceOrg)
	}
	return
}

func (k Keeper) GetLastReissuance(ctx sdk.Context) (val types.Reissuance, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ReissuanceBlockHeightKey))

	iterator := store.ReverseIterator(nil, nil)
	defer iterator.Close()
	found = iterator.Valid()
	if found {
		reissuance := iterator.Value()
		k.cdc.MustUnmarshal(reissuance, &val)
	}
	return val, found
}

func (k Keeper) ComputeReissuanceValue(ctx sdk.Context, startHeight int64, endHeight int64) (reissuanceValue uint64, firstIncludedPop int64, lastIncludedPop int64, err error) {
	challenges, err := k.GetChallengeRange(ctx, startHeight, endHeight)
	if err != nil {
		util.GetAppLogger().Error(ctx, "unable to compute get challenges")
		return
	}
	var overallAmount uint64
	popEpochs := k.GetParams(ctx).PopEpochs
	for _, obj := range challenges {
		popString := fmt.Sprintf("firstPoP: %d, PoP height: %d, current height %d", startHeight, obj.GetHeight(), endHeight)
		// if (index == 0 && startHeight == 0 && obj.BlockHeight == 0) || // corner case (beginning of the chain)
		if startHeight < obj.GetHeight() && obj.GetHeight()+2*popEpochs <= endHeight {
			popReissuanceString := GetReissuanceAsStringValue(obj.GetHeight(), k.GetParams(ctx).PopEpochs)
			amount, err := util.RDDLTokenStringToUint(popReissuanceString)
			if err != nil {
				util.GetAppLogger().Error(ctx, "unable to compute PoP reissuance value: "+popString)
				continue
			}
			util.GetAppLogger().Debug(ctx, "PoP is part of the reissuance: "+popString)
			if firstIncludedPop == 0 {
				firstIncludedPop = obj.GetHeight()
			}
			lastIncludedPop = obj.GetHeight()
			overallAmount += amount
		} else {
			util.GetAppLogger().Info(ctx, "PoP is not part of the reissuance: "+popString)
			if obj.GetHeight()+2*popEpochs > endHeight {
				break
			}
		}
	}
	reissuanceValue = overallAmount
	return
}
