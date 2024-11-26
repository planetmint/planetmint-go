package module

// import (
// 	"cosmossdk.io/core/appmodule"
// 	"cosmossdk.io/depinject"
// 	"github.com/cosmos/cosmos-sdk/baseapp"
// 	"github.com/cosmos/cosmos-sdk/x/group"
// 	"github.com/gogo/protobuf/codec"
// 	"github.com/planetmint/planetmint-go/x/machine/keeper"
// 	modulev1 "github.com/planetmint/planetmint-go/x/machine/types"
// )

// func init() {
// 	appmodule.Register(
// 		&modulev1.Module{},
// 		appmodule.Provide(ProvideModule),
// 	)
// }

// type MachineInputs struct {
// 	depinject.In
// 	Config           *modulev1.Module
// 	Key              *store.KVStoreKey
// 	Cdc              codec.Codec
// 	AccountKeeper    machine.AccountKeeper
// 	BankKeeper       machine.BankKeeper
// 	Registry         cdctypes.InterfaceRegistry
// 	MsgServiceRouter *baseapp.MsgServiceRouter
// }
// type MachineOutputs struct {
// 	depinject.Out
// 	MachineKeeper keeper.Keeper
// 	Module        appmodule.AppModule
// }

// func ProvideModule(in MachineInputs) MachineOutputs {
// 	/*
// 		Example of setting machine params:
// 		in.Config.MaxMetadataLen = 1000
// 		in.Config.MaxExecutionPeriod = "1209600s"
// 	*/
// 	k := keeper.NewKeeper(in.Key, in.Cdc, in.MsgServiceRouter, in.AccountKeeper, group.Config{MaxExecutionPeriod: in.Config.MaxExecutionPeriod.AsDuration(), MaxMetadataLen: in.Config.MaxMetadataLen})
// 	m := NewAppModule(in.Cdc, k, in.AccountKeeper, in.BankKeeper, in.Registry)
// 	return MachineOutputs{MachineKeeper: k, Module: m}
// }
