package util_test

import (
	"context"
	"encoding/json"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/planetmint/planetmint-go/testutil/keeper"
	"github.com/planetmint/planetmint-go/testutil/moduleobject"
	"github.com/planetmint/planetmint-go/testutil/sample"
	"github.com/planetmint/planetmint-go/util"
	"github.com/planetmint/planetmint-go/util/mocks"
	"github.com/planetmint/planetmint-go/x/machine/types"
	elements "github.com/rddl-network/elements-rpc"
	elementsmocks "github.com/rddl-network/elements-rpc/utils/mocks"
	"github.com/stretchr/testify/assert"
)

func TestRegisterNFT(t *testing.T) {
	params := types.DefaultParams()
	url := params.AssetRegistryScheme + "://" + params.AssetRegistryDomain + "/" + params.AssetRegistryPath
	entity := types.Entity{
		Domain: params.AssetRegistryDomain,
	}
	c := types.Contract{
		Entity:       entity,
		IssuerPubkey: "020000000000000000000000000000000000000000000000000000000000000000",
		MachineAddr:  "plmnt10mq5nj8jhh27z7ejnz2ql3nh0qhzjnfvy50877",
		Name:         "machine",
		Precision:    0,
		Version:      0,
	}
	contractBytes, err := json.Marshal(c)
	assert.NoError(t, err)

	contract := string(contractBytes)
	asset := "0000000000000000000000000000000000000000000000000000000000000000"
	goctx := context.Background()

	util.RegisterAssetServiceHTTPClient = &mocks.MockClient{}
	err = util.RegisterAsset(goctx, asset, contract, url)
	assert.NoError(t, err)
}

func TestMachineNFTIssuance(t *testing.T) {
	elements.Client = &elementsmocks.MockClient{}
	util.RegisterAssetServiceHTTPClient = &mocks.MockClient{}
	_, ctx := keeper.MachineKeeper(t)
	sk, pk := sample.KeyPair()
	machine := moduleobject.Machine(pk, pk, sk, "")
	goCtx := sdk.WrapSDKContext(ctx)

	params := types.DefaultParams()
	err := util.IssueMachineNFT(goCtx, &machine, params.AssetRegistryScheme, params.AssetRegistryDomain, params.AssetRegistryPath)
	assert.NoError(t, err)
}
