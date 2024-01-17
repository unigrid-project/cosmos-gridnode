package types

const (
	// ModuleName defines the module name
	ModuleName = "gridnode"

	// ModuleAccountName defines the module account name
	ModuleAccountName = "module_account_gridnode"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	ParamsKey = ModuleName + "/ParamsKey"
	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_gridnode"
	PramsMemKey = "mem_gridnode_params"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
