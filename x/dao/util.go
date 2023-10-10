package dao

func GetReissuanceCommand(asset_id string, BlockHeight int64) string {
	return "reissueasset " + asset_id + " 998.69"
}

func IsValidReissuanceCommand(reissuance_str string, asset_id string, BlockHeight uint64) bool {
	expected := "reissueasset " + asset_id + " 998.69"
	return reissuance_str == expected
}
