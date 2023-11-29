package keeper

import (
	"math"
	"math/big"
	"strconv"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/planetmint/planetmint-go/config"
	"github.com/planetmint/planetmint-go/x/dao/types"
)

func GetPopNumber(blockHeight int64) float64 {
	return float64(blockHeight) / float64(config.GetConfig().PopEpochs)
}

var PopsPerCycle float64

func init() {
	PopsPerCycle = 1051200.0
}

func GetReissuanceAsStringValue(blockHeight int64) string {
	PopNumber := GetPopNumber(blockHeight)
	exactCycleID := PopNumber / PopsPerCycle

	switch cycleID := math.Floor(exactCycleID); cycleID {
	case 0:
		return "998.69000000"
	case 1:
		return "499.34000000"
	case 2:
		return "249.67000000"
	case 3:
		return "124.83000000"
	case 4:
		return "62.42000000"
	default:
		return "0.0"
	}
}

func GetReissuanceCommand(assetID string, blockHeight int64) string {
	return "reissueasset " + assetID + " " + GetReissuanceAsStringValue(blockHeight)
}

func IsValidReissuanceCommand(reissuanceStr string, assetID string, blockHeight int64) bool {
	expected := "reissueasset " + assetID + " " + GetReissuanceAsStringValue(blockHeight)
	return reissuanceStr == expected
}

func GetReissuanceCommandForValue(assetID string, value uint64) string {
	return "reissueasset " + assetID + " " + strconv.FormatUint(value, 10)
}

func (k Keeper) StoreReissuance(ctx sdk.Context, reissuance types.Reissuance) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ReissuanceBlockHeightKey))
	appendValue := k.cdc.MustMarshal(&reissuance)
	store.Set(getReissuanceBytes(reissuance.BlockHeight), appendValue)
}

func (k Keeper) LookupReissuance(ctx sdk.Context, height int64) (val types.Reissuance, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ReissuanceBlockHeightKey))
	reissuance := store.Get(getReissuanceBytes(height))
	if reissuance == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(reissuance, &val)
	return val, true
}

func (k Keeper) getReissuancesRange(ctx sdk.Context, from int64) (reissuances []types.Reissuance) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ReissuanceBlockHeightKey))

	iterator := store.Iterator(getReissuanceBytes(from), nil)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		reissuance := iterator.Value()
		var reissuanceOrg types.Reissuance
		k.cdc.MustUnmarshal(reissuance, &reissuanceOrg)
		reissuances = append(reissuances, reissuanceOrg)
	}
	return reissuances
}

func (k Keeper) ComputeReIssuanceValue(blockHeight int64) (reIssuanceValue uint64) {
	return
}

func (k Keeper) getReissuancesPage(ctx sdk.Context, _ []byte, _ uint64, _ uint64, _ bool, reverse bool) (reissuances []types.Reissuance) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ReissuanceBlockHeightKey))

	iterator := store.Iterator(nil, nil)
	defer iterator.Close()
	if reverse {
		iterator = store.ReverseIterator(nil, nil)
		defer iterator.Close()
	}

	for ; iterator.Valid(); iterator.Next() {
		reissuance := iterator.Value()
		var reissuanceOrg types.Reissuance
		k.cdc.MustUnmarshal(reissuance, &reissuanceOrg)
		reissuances = append(reissuances, reissuanceOrg)
	}
	return reissuances
}

func getReissuanceBytes(height int64) []byte {
	// Adding 1 because 0 will be interpreted as nil, which is an invalid key
	return big.NewInt(height + 1).Bytes()
}
