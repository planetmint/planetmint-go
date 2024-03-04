package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/planetmint/planetmint-go/testutil/keeper"
	"github.com/planetmint/planetmint-go/x/machine/keeper"
	"github.com/stretchr/testify/assert"
)

func TestRegisterNFT(t *testing.T) {
	_, ctx := keepertest.MachineKeeper(t)
	url := "https://testnet-assets.rddl.io"
	contract := `{"entity":{"domain":"testnet-assets.rddl.io"},"issuer_pubkey":"020000000000000000000000000000000000000000000000000000000000000000","machine_addr":"plmnt10mq5nj8jhh27z7ejnz2ql3nh0qhzjnfvy50877","name":"machine","precision":0,"version":0}`
	asset := "0000000000000000000000000000000000000000000000001000000000000000"
	goctx := sdk.WrapSDKContext(ctx)
	err := keeper.RegisterAsset(goctx, asset, contract, url)
	assert.NoError(t, err)
}
