package keeper

import (
	"context"
	"errors"
	"strconv"

	config "planetmint-go/config"
	"planetmint-go/util"
	"planetmint-go/x/machine/types"

	"github.com/btcsuite/btcd/btcutil/hdkeychain"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/crgimenes/go-osc"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) isNFTCreationRequest(machine *types.Machine) bool {
	if !machine.GetReissue() && machine.GetAmount() == 1 && machine.GetPrecision() == 8 {
		return true
	}
	return false
}
func (k msgServer) AttestMachine(goCtx context.Context, msg *types.MsgAttestMachine) (*types.MsgAttestMachineResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	ta, activated, found := k.GetTrustAnchor(ctx, msg.Machine.MachineId)
	if !found {
		return nil, errors.New("no preregistered trust anchor found for machine id")
	}
	if activated {
		return nil, errors.New("trust anchor has already been used for attestation")
	}

	isValidMachineId := util.ValidateSignature_hexstring(msg.Machine.MachineId, msg.Machine.MachineIdSignature, msg.Machine.MachineId)
	if !isValidMachineId {
		return nil, errors.New("invalid machine id")
	}

	isValidIssuerPlanetmint := validateExtendedPublicKey(msg.Machine.IssuerPlanetmint, config.PlmntNetParams)
	if !isValidIssuerPlanetmint {
		return nil, errors.New("invalid planetmint key")
	}
	isValidIssuerLiquid := validateExtendedPublicKey(msg.Machine.IssuerLiquid, config.LiquidNetParams)
	if !isValidIssuerLiquid {
		return nil, errors.New("invalid liquid key")
	}
	if k.isNFTCreationRequest(msg.Machine) {
		err := k.issueMachineNFT(msg.Machine)
		if err != nil {
			return nil, errors.New("an error occurred while issuing the machine NFT")
		}
	}

	if msg.Machine.GetType() == 0 { // 0 == RDDL_MACHINE_UNDEFINED
		return nil, errors.New("the machine type has to be defined")
	}

	k.StoreMachine(ctx, *msg.Machine)
	k.StoreMachineIndex(ctx, *msg.Machine)
	k.StoreTrustAnchor(ctx, ta, true)

	return &types.MsgAttestMachineResponse{}, nil
}

func validateExtendedPublicKey(issuer string, cfg chaincfg.Params) bool {
	xpubKey, err := hdkeychain.NewKeyFromString(issuer)
	if err != nil {
		return false
	}
	isValidExtendedPublicKey := xpubKey.IsForNet(&cfg)
	return isValidExtendedPublicKey
}

func (k msgServer) issueMachineNFT(machine *types.Machine) error {
	conf := config.GetConfig()
	client := osc.NewClient(conf.WatchmenEndpoint, conf.WatchmenPort)
	machine_precision := strconv.FormatInt(int64(machine.Precision), 10)
	machine_amount := strconv.FormatInt(int64(machine.Amount), 10)
	machine_type := strconv.FormatUint(uint64(machine.GetType()), 10)

	msg := osc.NewMessage("/rddl/issue")
	msg.Append(machine.Name)
	msg.Append(machine.Ticker)
	msg.Append(machine.Domain)
	msg.Append(machine_amount)
	msg.Append("1")
	msg.Append(machine_precision)
	msg.Append(machine.Metadata.GetAdditionalDataCID())
	msg.Append(machine.GetIssuerPlanetmint())
	msg.Append(machine_type)
	err := client.Send(msg)

	return err
}
