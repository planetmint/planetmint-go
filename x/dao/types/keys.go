package types

const (
	// ModuleName defines the module name
	ModuleName = "dao"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_dao"

	ChallengeKey = "Dao/Challenge"

	MintRequestAddressKey = "Dao/MintRequestAddress"

	MintRequestHashKey = "Dao/MintRequestHash"

	ReissuanceBlockHeightKey = "Dao/ReissuanceBlockHeight"

	DistributionKey = "Dao/Distribution"

	PoPDistributionKey = "Dao/PoPDistribution"

	PoPInitiatorReward = "Dao/PoPInitiatorReward"

	ParamsKey = "Dao/Params"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
