package types

const (
	// ModuleName defines the module name
	ModuleName = "machine"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_machine"

	MachineKey = "Machine/value/"

	TAIndexKey = "Machine/TAIndex/"

	IssuerPlanetmintIndexKey = "Machine/IssuerPlanetmintIndex/"

	IssuerLiquidIndexKey = "Machine/IssuerLiquidIndex/"

	TrustAnchorKey = "Machine/trustAnchor/"

	AddressIndexKey = "Machine/address"

	LiquidAssetKey = "Machine/LiquidAsset/"

	ParamsKey = "Machine/Params"

	ActivatedTACounterPrefix = "ActivatedTACounter"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
