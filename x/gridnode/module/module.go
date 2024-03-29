// package gridnode

// import (
// 	"context"
// 	"encoding/json"
// 	"fmt"

// 	// this line is used by starport scaffolding # 1

// 	"github.com/grpc-ecosystem/grpc-gateway/runtime"
// 	"github.com/spf13/cobra"

// 	abci "github.com/cometbft/cometbft/abci/types"

// 	"cosmossdk.io/core/appmodule"
// 	"github.com/cosmos/cosmos-sdk/client"
// 	"github.com/cosmos/cosmos-sdk/codec"
// 	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
// 	sdk "github.com/cosmos/cosmos-sdk/types"
// 	"github.com/cosmos/cosmos-sdk/types/module"
// 	"github.com/unigrid-project/cosmos-gridnode/x/gridnode/client/cli"
// 	"github.com/unigrid-project/cosmos-gridnode/x/gridnode/keeper"
// 	"github.com/unigrid-project/cosmos-gridnode/x/gridnode/types"
// )

// var (
// 	_ module.AppModuleBasic = (*AppModule)(nil)

// 	_ appmodule.AppModule = (*AppModule)(nil)
// )

// // ----------------------------------------------------------------------------
// // AppModuleBasic
// // ----------------------------------------------------------------------------

// // AppModuleBasic implements the AppModuleBasic interface that defines the independent methods a Cosmos SDK module needs to implement.
// type AppModuleBasic struct {
// 	cdc codec.BinaryCodec
// }

// func NewAppModuleBasic(cdc codec.BinaryCodec) AppModuleBasic {
// 	return AppModuleBasic{cdc: cdc}
// }

// // Name returns the name of the module as a string
// func (AppModuleBasic) Name() string {
// 	return types.ModuleName
// }

// // RegisterLegacyAminoCodec registers the amino codec for the module, which is used to marshal and unmarshal structs to/from []byte in order to persist them in the module's KVStore
// func (AppModuleBasic) RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
// 	types.RegisterCodec(cdc)
// }

// // RegisterInterfaces registers a module's interface types and their concrete implementations as proto.Message
// func (a AppModuleBasic) RegisterInterfaces(reg cdctypes.InterfaceRegistry) {
// 	types.RegisterInterfaces(reg)
// }

// // DefaultGenesis returns a default GenesisState for the module, marshalled to json.RawMessage. The default GenesisState need to be defined by the module developer and is primarily used for testing
// func (AppModuleBasic) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
// 	return cdc.MustMarshalJSON(types.DefaultGenesis())
// }

// // ValidateGenesis used to validate the GenesisState, given in its json.RawMessage form
// func (AppModuleBasic) ValidateGenesis(cdc codec.JSONCodec, config client.TxEncodingConfig, bz json.RawMessage) error {
// 	var genState types.GenesisState
// 	if err := cdc.UnmarshalJSON(bz, &genState); err != nil {
// 		return fmt.Errorf("failed to unmarshal %s genesis state: %w", types.ModuleName, err)
// 	}
// 	return genState.Validate()
// }

// // RegisterGRPCGatewayRoutes registers the gRPC Gateway routes for the module
// func (AppModuleBasic) RegisterGRPCGatewayRoutes(clientCtx client.Context, mux *runtime.ServeMux) {
// 	types.RegisterQueryHandlerClient(context.Background(), mux, types.NewQueryClient(clientCtx))
// }

// // GetTxCmd returns the root Tx command for the module. The subcommands of this root command are used by end-users to generate new transactions containing messages defined in the module
// func (a AppModuleBasic) GetTxCmd() *cobra.Command {
// 	return cli.GetTxCmd()
// }

// // GetQueryCmd returns the root query command for the module. The subcommands of this root command are used by end-users to generate new queries to the subset of the state defined by the module
// func (AppModuleBasic) GetQueryCmd() *cobra.Command {
// 	return cli.GetQueryCmd(types.StoreKey)
// }

// // ----------------------------------------------------------------------------
// // AppModule
// // ----------------------------------------------------------------------------

// // AppModule implements the AppModule interface that defines the inter-dependent methods that modules need to implement
// type AppModule struct {
// 	AppModuleBasic

// 	keeper        keeper.Keeper
// 	accountKeeper types.AccountKeeper
// 	bankKeeper    types.BankKeeper
// }

// func (am AppModule) IsAppModule() {}

// func (am AppModule) IsOnePerModuleType() {}

// func NewAppModule(
// 	cdc codec.Codec,
// 	keeper keeper.Keeper,
// 	accountKeeper types.AccountKeeper,
// 	bankKeeper types.BankKeeper,
// ) AppModule {
// 	return AppModule{
// 		AppModuleBasic: NewAppModuleBasic(cdc),
// 		keeper:         keeper,
// 		accountKeeper:  accountKeeper,
// 		bankKeeper:     bankKeeper,
// 	}
// }

// // RegisterServices registers a gRPC query service to respond to the module-specific gRPC queries
// func (am AppModule) RegisterServices(cfg module.Configurator) {
// 	types.RegisterMsgServer(cfg.MsgServer(), keeper.NewMsgServerImpl(am.keeper))
// 	types.RegisterQueryServer(cfg.QueryServer(), am.keeper)
// }

// // RegisterInvariants registers the invariants of the module. If an invariant deviates from its predicted value, the InvariantRegistry triggers appropriate logic (most often the chain will be halted)
// func (am AppModule) RegisterInvariants(_ sdk.InvariantRegistry) {}

// // InitGenesis performs the module's genesis initialization. It returns no validator updates.
// func (am AppModule) InitGenesis(ctx sdk.Context, cdc codec.JSONCodec, gs json.RawMessage) []abci.ValidatorUpdate {
// 	var genState types.GenesisState
// 	// Initialize global index to index in genesis state
// 	cdc.MustUnmarshalJSON(gs, &genState)

// 	InitGenesis(ctx, am.keeper, genState)

// 	return []abci.ValidatorUpdate{}
// }

// // ExportGenesis returns the module's exported genesis state as raw JSON bytes.
// func (am AppModule) ExportGenesis(ctx sdk.Context, cdc codec.JSONCodec) json.RawMessage {
// 	genState := ExportGenesis(ctx, am.keeper)
// 	return cdc.MustMarshalJSON(genState)
// }

// // ConsensusVersion is a sequence number for state-breaking change of the module. It should be incremented on each consensus-breaking change introduced by the module. To avoid wrong/empty versions, the initial version should be set to 1
// func (AppModule) ConsensusVersion() uint64 { return 1 }

// // BeginBlock contains the logic that is automatically triggered at the beginning of each block
// func (am AppModule) BeginBlock(ctx sdk.Context) {
// 	BeginBlocker(ctx, am.keeper)
// }

// // EndBlock contains the logic that is automatically triggered at the end of each block
// func (am AppModule) EndBlock(_ context.Context) ([]abci.ValidatorUpdate, error) {
// 	return []abci.ValidatorUpdate{}, nil
// }

package gridnode

import (
	"context"
	"encoding/json"
	"fmt"

	"cosmossdk.io/core/appmodule"
	"cosmossdk.io/core/store"
	"cosmossdk.io/depinject"
	"cosmossdk.io/log"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	typesparams "github.com/cosmos/cosmos-sdk/x/params/types"

	"github.com/spf13/cobra"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"

	// this line is used by starport scaffolding # 1

	modulev1 "github.com/unigrid-project/cosmos-gridnode/api/gridnode/gridnode/v1/module"
	"github.com/unigrid-project/cosmos-gridnode/x/gridnode/client/cli"
	"github.com/unigrid-project/cosmos-gridnode/x/gridnode/keeper"
	"github.com/unigrid-project/cosmos-gridnode/x/gridnode/types"
)

var (
	_ module.AppModuleBasic      = (*AppModule)(nil)
	_ module.AppModuleSimulation = (*AppModule)(nil)
	_ module.HasGenesis          = (*AppModule)(nil)
	_ module.HasInvariants       = (*AppModule)(nil)
	_ module.HasConsensusVersion = (*AppModule)(nil)

	_ appmodule.AppModule       = (*AppModule)(nil)
	_ appmodule.HasBeginBlocker = (*AppModule)(nil)
	_ appmodule.HasEndBlocker   = (*AppModule)(nil)
)

// ----------------------------------------------------------------------------
// AppModuleBasic
// ----------------------------------------------------------------------------

// AppModuleBasic implements the AppModuleBasic interface that defines the
// independent methods a Cosmos SDK module needs to implement.
type AppModuleBasic struct {
	cdc codec.BinaryCodec
	ak  types.AccountKeeper
}

func NewAppModuleBasic(cdc codec.BinaryCodec, ak types.AccountKeeper) AppModuleBasic {
	return AppModuleBasic{cdc: cdc, ak: ak}
}

// Name returns the name of the module as a string.
func (AppModuleBasic) Name() string {
	return types.ModuleName
}

// RegisterLegacyAminoCodec registers the amino codec for the module, which is used
// to marshal and unmarshal structs to/from []byte in order to persist them in the module's KVStore.
func (AppModuleBasic) RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	types.RegisterCodec(cdc)
}

// RegisterInterfaces registers a module's interface types and their concrete implementations as proto.Message.
func (a AppModuleBasic) RegisterInterfaces(reg cdctypes.InterfaceRegistry) {
	types.RegisterInterfaces(reg)
}

// DefaultGenesis returns a default GenesisState for the module, marshalled to json.RawMessage.
// The default GenesisState need to be defined by the module developer and is primarily used for testing.
func (AppModuleBasic) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	return cdc.MustMarshalJSON(types.DefaultGenesis())
}

// ValidateGenesis used to validate the GenesisState, given in its json.RawMessage form.
func (AppModuleBasic) ValidateGenesis(cdc codec.JSONCodec, config client.TxEncodingConfig, bz json.RawMessage) error {
	var genState types.GenesisState
	if err := cdc.UnmarshalJSON(bz, &genState); err != nil {
		return fmt.Errorf("failed to unmarshal %s genesis state: %w", types.ModuleName, err)
	}
	return genState.Validate()
}

// // GetQueryCmd returns the root query command for the module. The subcommands of this root command are used by end-users to generate new queries to the subset of the state defined by the module
func (AppModuleBasic) GetQueryCmd() *cobra.Command {
	return cli.GetQueryCmd(types.StoreKey)
}

// // GetTxCmd returns the root Tx command for the module. The subcommands of this root command are used by end-users to generate new transactions containing messages defined in the module
func (a AppModuleBasic) GetTxCmd() *cobra.Command {
	return cli.GetTxCmd()
}

// RegisterGRPCGatewayRoutes registers the gRPC Gateway routes for the module.
func (AppModuleBasic) RegisterGRPCGatewayRoutes(clientCtx client.Context, mux *runtime.ServeMux) {
	if err := types.RegisterQueryHandlerClient(context.Background(), mux, types.NewQueryClient(clientCtx)); err != nil {
		panic(err)
	}
}

// ----------------------------------------------------------------------------
// AppModule
// ----------------------------------------------------------------------------

// AppModule implements the AppModule interface that defines the inter-dependent methods that modules need to implement
type AppModule struct {
	AppModuleBasic

	keeper        keeper.Keeper
	accountKeeper types.AccountKeeper
	bankKeeper    types.BankKeeper
}

func NewAppModule(
	cdc codec.Codec,
	keeper keeper.Keeper,
	accountKeeper types.AccountKeeper,
	bankKeeper types.BankKeeper,
) AppModule {
	return AppModule{
		AppModuleBasic: NewAppModuleBasic(cdc, accountKeeper),
		keeper:         keeper,
		accountKeeper:  accountKeeper,
		bankKeeper:     bankKeeper,
	}
}

// RegisterServices registers a gRPC query service to respond to the module-specific gRPC queries
func (am AppModule) RegisterServices(cfg module.Configurator) {
	types.RegisterMsgServer(cfg.MsgServer(), keeper.NewMsgServerImpl(am.keeper))
	types.RegisterQueryServer(cfg.QueryServer(), am.keeper)
}

// RegisterInvariants registers the invariants of the module. If an invariant deviates from its predicted value, the InvariantRegistry triggers appropriate logic (most often the chain will be halted)
func (am AppModule) RegisterInvariants(_ sdk.InvariantRegistry) {}

// InitGenesis performs the module's genesis initialization. It returns no validator updates.
func (am AppModule) InitGenesis(ctx sdk.Context, cdc codec.JSONCodec, gs json.RawMessage) {
	var genState types.GenesisState
	// Initialize global index to index in genesis state
	cdc.MustUnmarshalJSON(gs, &genState)

	InitGenesis(ctx, am.keeper, genState)
}

// ExportGenesis returns the module's exported genesis state as raw JSON bytes.
func (am AppModule) ExportGenesis(ctx sdk.Context, cdc codec.JSONCodec) json.RawMessage {
	genState := ExportGenesis(ctx, am.keeper)
	return cdc.MustMarshalJSON(genState)
}

// ConsensusVersion is a sequence number for state-breaking change of the module.
// It should be incremented on each consensus-breaking change introduced by the module.
// To avoid wrong/empty versions, the initial version should be set to 1.
func (AppModule) ConsensusVersion() uint64 { return 1 }

// BeginBlock contains the logic that is automatically triggered at the beginning of each block.
// The begin block implementation is optional.
func (am AppModule) BeginBlock(ctx context.Context) error {
	fmt.Println("BeginBlock: start")

	BeginBlocker(ctx, am.keeper) // Call BeginBlocker without expecting a return value

	fmt.Println("BeginBlock: completed successfully")
	return nil
}

// EndBlock contains the logic that is automatically triggered at the end of each block.
// The end block implementation is optional.
func (am AppModule) EndBlock(_ context.Context) error {
	return nil
}

// IsOnePerModuleType implements the depinject.OnePerModuleType interface.
func (am AppModule) IsOnePerModuleType() {}

// IsAppModule implements the appmodule.AppModule interface.
func (am AppModule) IsAppModule() {}

// ----------------------------------------------------------------------------
// App Wiring Setup
// ----------------------------------------------------------------------------

func init() {
	appmodule.Register(
		&modulev1.Module{},
		appmodule.Provide(ProvideModule),
	)
}

type ModuleInputs struct {
	depinject.In

	StoreService store.KVStoreService
	Cdc          codec.Codec
	Config       *modulev1.Module
	Logger       log.Logger

	AccountKeeper types.AccountKeeper
	BankKeeper    types.BankKeeper
}

type ModuleOutputs struct {
	depinject.Out

	GridnodeKeeper keeper.Keeper
	Module         appmodule.AppModule
}

func ProvideModule(in ModuleInputs) ModuleOutputs {
	// default to governance authority if not provided
	authority := authtypes.NewModuleAddress(govtypes.ModuleName)
	if in.Config.Authority != "" {
		authority = authtypes.NewModuleAddressOrBech32Address(in.Config.Authority)
	}

	// storeKey := storetypes.NewKVStoreKey(types.StoreKey)
	// memStoreKey := storetypes.NewMemoryStoreKey(types.MemStoreKey)
	paramKey := storetypes.NewKVStoreKey(types.ParamsKey)
	paramMemKey := storetypes.NewMemoryStoreKey(types.PramsMemKey)
	paramsSubspace := typesparams.NewSubspace(
		in.Cdc,
		types.Amino,
		paramKey,
		paramMemKey,
		"GridnodeParams",
	)

	k := keeper.NewKeeper(
		in.Cdc,
		//storeKey,
		//memStoreKey,
		paramsSubspace,
		in.BankKeeper,
		authority.String(),
		in.StoreService,
		in.AccountKeeper,
	)
	m := NewAppModule(
		in.Cdc,
		*k,
		in.AccountKeeper,
		in.BankKeeper,
	)

	return ModuleOutputs{GridnodeKeeper: *k, Module: m}
}
