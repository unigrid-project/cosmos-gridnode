package keeper

import (
	"testing"

	"cosmossdk.io/log"
	"cosmossdk.io/store"

	"cosmossdk.io/store/metrics"
	storetypes "cosmossdk.io/store/types"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	cdb "github.com/cosmos/cosmos-db"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	typesparams "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/stretchr/testify/require"
	"github.com/unigrid-project/cosmos-gridnode/x/gridnode/keeper"
	"github.com/unigrid-project/cosmos-gridnode/x/gridnode/types"
)

func GridnodeKeeper(t testing.TB) (*keeper.Keeper, sdk.Context) {
	authority := sdk.AccAddress("authority")
	storeKey := storetypes.NewKVStoreKey(types.StoreKey)
	memStoreKey := storetypes.NewMemoryStoreKey(types.MemStoreKey)

	db := cdb.NewMemDB()
	stateStore := store.NewCommitMultiStore(db, log.NewNopLogger(), metrics.NewNoOpMetrics())
	stateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(memStoreKey, storetypes.StoreTypeMemory, nil)
	require.NoError(t, stateStore.LoadLatestVersion())

	registry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(registry)

	paramsSubspace := typesparams.NewSubspace(cdc,
		types.Amino,
		storeKey,
		memStoreKey,
		"GridnodeParams",
	)

	// Create a mock BankKeeper for testing purposes
	bankKeeper := mockBankKeeper()
	accountKeeper := mockAccountKeeper()

	k := keeper.NewKeeper(
		cdc,
		paramsSubspace,
		bankKeeper, // Pass the bankKeeper as an argument
		authority.String(),
		runtime.NewKVStoreService(storeKey),
		accountKeeper,
	)

	ctx := sdk.NewContext(stateStore, tmproto.Header{}, false, log.NewNopLogger())

	// Initialize params
	k.SetParams(ctx, types.DefaultParams())

	return k, ctx
}

// Create a mock BankKeeper for testing purposes
func mockBankKeeper() types.BankKeeper {
	// Implement your mock BankKeeper here
	// This can be a simple struct that fulfills the types.BankKeeper interface
	// with the necessary methods for your tests
	return &MockBankKeeper{}
}

// Create a mock BankKeeper for testing purposes
func mockAccountKeeper() types.AccountKeeper {
	// Implement your mock BankKeeper here
	// This can be a simple struct that fulfills the types.BankKeeper interface
	// with the necessary methods for your tests
	return &MockAccountKeeper{}
}
