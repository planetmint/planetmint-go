package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgReportPopResult{}, "dao/ReportPopResult", nil)
	cdc.RegisterConcrete(&MsgReissueRDDLProposal{}, "dao/ReissueRDDLProposal", nil)
	cdc.RegisterConcrete(&MsgMintToken{}, "dao/MintToken", nil)
	cdc.RegisterConcrete(&MsgReissueRDDLResult{}, "dao/ReissueRDDLResult", nil)
	cdc.RegisterConcrete(&MsgDistributionResult{}, "dao/DistributionResult", nil)
	cdc.RegisterConcrete(&MsgDistributionRequest{}, "dao/DistributionRequest", nil)
	cdc.RegisterConcrete(&MsgUpdateParams{}, "dao/UpdateParams", nil)
	cdc.RegisterConcrete(&MsgInitPop{}, "dao/InitPop", nil)
	cdc.RegisterConcrete(&MsgCreateRedeemClaim{}, "dao/CreateRedeemClaim", nil)
	cdc.RegisterConcrete(&MsgUpdateRedeemClaim{}, "dao/UpdateRedeemClaim", nil)
	cdc.RegisterConcrete(&MsgConfirmRedeemClaim{}, "dao/ConfirmRedeemClaim", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgReportPopResult{},
		&MsgReissueRDDLProposal{},
		&MsgMintToken{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgReissueRDDLResult{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgDistributionResult{},
		&MsgDistributionRequest{},
		&MsgUpdateParams{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgInitPop{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateRedeemClaim{},
		&MsgUpdateRedeemClaim{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgConfirmRedeemClaim{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
