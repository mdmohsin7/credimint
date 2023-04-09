package types

const (
	// ModuleName defines the module name
	ModuleName = "credimint"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_credimint"

	// Version defines the current version the IBC module supports
	Version = "credimint-1"

	// PortID is the default port id that module binds to
	PortID = "credimint"
)

var (
	// PortKey defines the key to store the port ID in store
	PortKey = KeyPrefix("credimint-port-")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
