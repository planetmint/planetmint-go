package keeper

import (
	"context"
	"errors"

	"planetmint-go/x/machine/types"

	"github.com/btcsuite/btcd/btcutil/hdkeychain"
	"github.com/btcsuite/btcd/chaincfg"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) AttestMachine(goCtx context.Context, msg *types.MsgAttestMachine) (*types.MsgAttestMachineResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	isValidIssuerLiquid := validateIssuerLiquid(msg.Machine.IssuerLiquid)
	if !isValidIssuerLiquid {
		return nil, errors.New("invalid liquid key")
	}

	if msg.Machine.Reissue {
		k.Logger(ctx).Info("TODO Implement handle on reissue == true")
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
