package dao

func GetReissuanceCommand(assetID string, BlockHeight int64) string {
	return "reissueasset " + assetID + " 99869000000"
}

func IsValidReissuanceCommand(reissuanceStr string, assetID string, BlockHeight int64) bool {
	expected := "reissueasset " + assetID + " 99869000000"
	return reissuanceStr == expected
}
