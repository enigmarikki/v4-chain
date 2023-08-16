package types

const (
	// ModuleName defines the module name
	ModuleName = "perpetuals"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}