package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgAttestMachine{}, "machine/AttestMachine", nil)
	cdc.RegisterConcrete(&MsgRegisterTrustAnchor{}, "machine/RegisterTrustAnchor", nil)
	cdc.RegisterConcrete(&MsgNotarizeLiquidAsset{}, "machine/NotarizeLiquidAsset", nil)
	cdc.RegisterConcrete(&MsgUpdateParams{}, "machine/UpdateParams", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgAttestMachine{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRegisterTrustAnchor{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgNotarizeLiquidAsset{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdateParams{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
