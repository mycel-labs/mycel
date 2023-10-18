package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgUpdateWalletRecord{}, "registry/UpdateWalletRecord", nil)
	cdc.RegisterConcrete(&MsgRegisterSecondLevelDomain{}, "registry/RegisterSecondLevelDomain", nil)
	cdc.RegisterConcrete(&MsgRegisterTopLevelDomain{}, "registry/RegisterTopLevelDomain", nil)
	cdc.RegisterConcrete(&MsgWithdrawRegistrationFee{}, "registry/WithdrawRegistrationFee", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdateWalletRecord{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRegisterSecondLevelDomain{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRegisterTopLevelDomain{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgWithdrawRegistrationFee{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
