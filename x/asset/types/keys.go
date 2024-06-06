package types

const (
	// ModuleName defines the module name
	ModuleName = "asset"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_asset"

	AssetKey = "Asset/value/"
	CountKey = "count/"
)

func AddressCountKey(address string) (prefix []byte) {
	addressPrefix := AddressPrefix(address)
	prefix = append(prefix, addressPrefix...)
	prefix = append(prefix, []byte(CountKey)...)
	return
}

func AddressPrefix(address string) (prefix []byte) {
	addressBytes := []byte(address)
	prefix = append(prefix, addressBytes...)
	prefix = append(prefix, []byte("/")...)
	return
}

func KeyPrefix(p string) []byte {
	return []byte(p)
}
