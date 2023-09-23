package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	// this line is used by starport scaffolding # 1
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	// this line is used by starport scaffolding # 2
}

// RegisterInterfaces registers the interfaces types with the interface registry.
// func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
// 	registry.RegisterImplementations(
// 		(*sdk.Msg)(nil),
// 		&MsgGridnodeDelegate{},
// 		&MsgGridnodeUndelegate{},
// 	)

// 	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
// }

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
