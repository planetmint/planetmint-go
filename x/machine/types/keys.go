package types

const (
	// ModuleName defines the module name
	ModuleName = "machine"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_machine"

	MachineKey = "Machine/value/"

	TAIndexKey = "Machine/TAIndex/"

	IssuerPlanetmintIndexKey = "Machine/IssuerPlanetmintIndex/"

	IssuerLiquidIndexKey = "Machine/IssuerLiquidIndex/"

	TrustAnchorKey = "Machine/trustAnchor/"

	AddressIndexKey = "Machine/address"

	LiquidAssetKey = "Machine/LiquidAsset/"
)

var (
	ParamsKey = []byte("p_machine")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
