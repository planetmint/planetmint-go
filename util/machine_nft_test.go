package util_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/planetmint/planetmint-go/testutil/keeper"
	"github.com/planetmint/planetmint-go/testutil/moduleobject"
	"github.com/planetmint/planetmint-go/testutil/sample"
	"github.com/planetmint/planetmint-go/util"
	"github.com/planetmint/planetmint-go/util/mocks"
	elements "github.com/rddl-network/elements-rpc"
	elementsmocks "github.com/rddl-network/elements-rpc/utils/mocks"
	"github.com/stretchr/testify/assert"
)

func TestRegisterNFT(t *testing.T) {
	t.Parallel()
	url := "https://testnet-assets.rddl.io/register_asset"
	contract := `{"entity":{"domain":"testnet-assets.rddl.io"},"issuer_pubkey":"020000000000000000000000000000000000000000000000000000000000000000","machine_addr":"plmnt10mq5nj8jhh27z7ejnz2ql3nh0qhzjnfvy50877","name":"machine","precision":0,"version":0}`
	asset := "0000000000000000000000000000000000000000000000000000000000000000"
	goctx := context.Background()

	util.RegisterAssetServiceHTTPClient = &mocks.MockClient{}
	err := util.RegisterAsset(goctx, asset, contract, url)
	assert.NoError(t, err)
}

func TestMachineNFTIssuance(t *testing.T) {
	t.Parallel()

	elements.Client = &elementsmocks.MockClient{}
	util.RegisterAssetServiceHTTPClient = &mocks.MockClient{}
	_, ctx := keeper.MachineKeeper(t)
	sk, pk := sample.KeyPair()
	machine := moduleobject.Machine(pk, pk, sk, "")
	goCtx := sdk.WrapSDKContext(ctx)

	err := util.IssueMachineNFT(goCtx, &machine, "https", "testnet-asset.rddl.io", "/register_asset")

	assert.NoError(t, err)
}
