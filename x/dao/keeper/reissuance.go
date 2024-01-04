package keeper

import (
	"fmt"
	"math"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/planetmint/planetmint-go/config"
	"github.com/planetmint/planetmint-go/util"
	"github.com/planetmint/planetmint-go/x/dao/types"
)

var ReIssueCommand string

func init() {
	ReIssueCommand = "reissueasset"
}

func GetReissuanceAsStringValue(blockHeight int64) string {
	PopNumber := util.GetPopNumber(blockHeight)
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

func GetReissuanceCommand(assetID string, blockHeight int64) string {
	return ReIssueCommand + " " + assetID + " " + GetReissuanceAsStringValue(blockHeight)
}

func IsValidReissuanceCommand(reissuanceStr string, assetID string, blockHeight int64) bool {
	expected := ReIssueCommand + " " + assetID + " " + GetReissuanceAsStringValue(blockHeight)
	return reissuanceStr == expected
}

func GetReissuanceCommandForValue(assetID string, value uint64) string {
	return ReIssueCommand + " " + assetID + " " + util.UintValueToRDDLTokenString(value)
}

func (k Keeper) CreateNextReIssuanceObject(ctx sdk.Context, currentBlockHeight int64) (reIssuance types.Reissuance, err error) {
	var lastReissuedPop int64
	lastReIssuance, found := k.GetLastReIssuance(ctx)
	if found {
		lastReissuedPop = lastReIssuance.LastIncludedPop
	}
	reIssuanceValue, firstIncludedPop, lastIncludedPop, err := k.ComputeReIssuanceValue(ctx, lastReissuedPop, currentBlockHeight)
	if err != nil {
		return
	}

	reIssuance.RawTx = GetReissuanceCommandForValue(config.GetConfig().ReissuanceAsset, reIssuanceValue)
	reIssuance.BlockHeight = currentBlockHeight
	reIssuance.FirstIncludedPop = firstIncludedPop
	reIssuance.LastIncludedPop = lastIncludedPop
	return
}

func (k Keeper) IsValidReIssuanceProposal(ctx sdk.Context, msg *types.MsgReissueRDDLProposal) (isValid bool) {
	reIssuance, err := k.CreateNextReIssuanceObject(ctx, msg.GetBlockHeight())
	if err != nil {
		return
	}
	if reIssuance.GetBlockHeight() == msg.GetBlockHeight() &&
		reIssuance.GetFirstIncludedPop() == msg.GetFirstIncludedPop() &&
		reIssuance.GetLastIncludedPop() == msg.GetLastIncludedPop() &&
		reIssuance.GetRawTx() == msg.GetTx() &&
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

func (k Keeper) GetLastReIssuance(ctx sdk.Context) (val types.Reissuance, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ReissuanceBlockHeightKey))

	iterator := store.ReverseIterator(nil, nil)
	defer iterator.Close()
	found = iterator.Valid()
	if found {
		reIssuance := iterator.Value()
		k.cdc.MustUnmarshal(reIssuance, &val)
	}
	return val, found
}

func (k Keeper) ComputeReIssuanceValue(ctx sdk.Context, startHeight int64, endHeight int64) (reIssuanceValue uint64, firstIncludedPop int64, lastIncludedPop int64, err error) {
	challenges, err := k.GetChallengeRange(ctx, startHeight, endHeight)
	if err != nil {
		util.GetAppLogger().Error(ctx, "unable to compute get challenges")
		return
	}
	var overallAmount uint64
	popEpochs := int64(config.GetConfig().PopEpochs)
	for _, obj := range challenges {
		popString := fmt.Sprintf("firstPoP: %d, PoP height: %d, current height %d", startHeight, obj.GetHeight(), endHeight)
		// if (index == 0 && startHeight == 0 && obj.BlockHeight == 0) || // corner case (beginning of the chain)
		if startHeight < obj.GetHeight() && obj.GetHeight()+2*popEpochs <= endHeight {
			popReIssuanceString := GetReissuanceAsStringValue(obj.GetHeight())
			amount, err := util.RDDLTokenStringToUint(popReIssuanceString)
			if err != nil {
				util.GetAppLogger().Error(ctx, "unable to compute PoP re-issuance value: "+popString)
				continue
			}
			util.GetAppLogger().Info(ctx, "PoP is part of the reissuance: "+popString)
			if firstIncludedPop == 0 {
				firstIncludedPop = obj.GetHeight()
			}
			lastIncludedPop = obj.GetHeight()
			overallAmount += amount
		} else {
			util.GetAppLogger().Debug(ctx, "PoP is not part of the reissuance: "+popString)
			if obj.GetHeight()+2*popEpochs > endHeight {
				break
			}
		}
	}
	reIssuanceValue = overallAmount
	return
}
