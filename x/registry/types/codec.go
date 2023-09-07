package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgUpdateWalletRecord{}, "registry/UpdateWalletRecord", nil)
	cdc.RegisterConcrete(&MsgRegisterDomain{}, "registry/RegisterDomain", nil)
	cdc.RegisterConcrete(&MsgRegisterTopLevelDomain{}, "registry/RegisterTopLevelDomain", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdateWalletRecord{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRegisterDomain{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRegisterTopLevelDomain{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
