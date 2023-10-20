package dao

func GetReissuanceCommand(asset_id string, BlockHeight int64) string {
	return "reissueasset " + asset_id + " 99869000000"
}

func IsValidReissuanceCommand(reissuance_str string, asset_id string, BlockHeight int64) bool {
	expected := "reissueasset " + asset_id + " 99869000000"
	return reissuance_str == expected
}
