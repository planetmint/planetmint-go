package keeper

import (
	"context"
	"errors"

	"planetmint-go/x/machine/types"

	"github.com/btcsuite/btcd/btcutil/hdkeychain"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/crgimenes/go-osc"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) AttestMachine(goCtx context.Context, msg *types.MsgAttestMachine) (*types.MsgAttestMachineResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	isValidIssuerLiquid := validateIssuerLiquid(msg.Machine.IssuerLiquid)
	if !isValidIssuerLiquid {
		return nil, errors.New("invalid liquid key")
	}

	if msg.Machine.Reissue {
		err := k.reissueMachineNFT(msg.Machine)
		if err != nil {
			return nil, errors.New("an error occured while reissuning the machine")
		}
	}

	k.StoreMachine(ctx, *msg.Machine)
	k.StoreMachineIndex(ctx, *msg.Machine)

	return &types.MsgAttestMachineResponse{}, nil
}

func validateIssuerLiquid(issuerLiquid string) bool {
	xpubKeyLiquid, err := hdkeychain.NewKeyFromString(issuerLiquid)
	if err != nil {
		return false
	}
	isValidLiquidKey := xpubKeyLiquid.IsForNet(&chaincfg.MainNetParams)
	return isValidLiquidKey
}

func (k msgServer) reissueMachineNFT(machine *types.Machine) error {
	client := osc.NewClient("localhost", 8765)
	msg := osc.NewMessage("/rddl/*")
	msg.Append(machine.Name)
	msg.Append(machine.Ticker)
	msg.Append(machine.Domain)
	msg.Append(machine.Amount)
	msg.Append("1")
	msg.Append(machine.Precision)
	err := client.Send(msg)
	return err
}
