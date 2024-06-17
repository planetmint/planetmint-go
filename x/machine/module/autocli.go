package machine

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "github.com/planetmint/planetmint-go/api/planetmintgo/machine"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: modulev1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "Shows the parameters of the module",
				},
				{
					RpcMethod:      "GetMachineByPublicKey",
					Use:            "get-machine-by-public-key [public-key]",
					Short:          "Query get-machine-by-public-key",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "publicKey"}},
				},
				{
					RpcMethod:      "GetTrustAnchorStatus",
					Use:            "get-trust-anchor-status [machine-id]",
					Short:          "Query get-trust-anchor-status",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "machineId"}},
				},
				{
					RpcMethod:      "GetMachineByAddress",
					Use:            "get-machine-by-address [address]",
					Short:          "Query get-machine-by-address",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "address"}},
				},
				{
					RpcMethod:      "GetLiquidAssetsByMachineId",
					Use:            "get-liquid-assets-by-machine-id [machine-id]",
					Short:          "Query get-liquid-assets-by-machine-id",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "machineId"}},
				},
				// this line is used by ignite scaffolding # autocli/query
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              modulev1.Msg_ServiceDesc.ServiceName,
			EnhanceCustomCommand: true, // only required if you want to use the custom command
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "UpdateParams",
					Skip:      true, // skipped because authority gated
				},
				{
					RpcMethod:      "AttestMachine",
					Use:            "attest-machine [machine]",
					Short:          "Send a attest-machine tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "machine"}},
				},
				{
					RpcMethod:      "NotarizeLiquidAsset",
					Use:            "notarize-liquid-asset [notarization]",
					Short:          "Send a notarize-liquid-asset tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "notarization"}},
				},
				{
					RpcMethod:      "RegisterTrustAnchor",
					Use:            "register-trust-anchor [trust-anchor]",
					Short:          "Send a register-trust-anchor tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "trustAnchor"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
