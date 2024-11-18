package util_test

import (
	"context"
	"encoding/json"
	"math/rand"
	"strconv"
	"sync"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/golang/mock/gomock"
	"github.com/planetmint/planetmint-go/clients"
	"github.com/planetmint/planetmint-go/config"
	"github.com/planetmint/planetmint-go/testutil/keeper"
	clientmocks "github.com/planetmint/planetmint-go/testutil/mocks"
	"github.com/planetmint/planetmint-go/testutil/moduleobject"
	"github.com/planetmint/planetmint-go/testutil/sample"
	"github.com/planetmint/planetmint-go/util"
	"github.com/planetmint/planetmint-go/util/mocks"
	"github.com/planetmint/planetmint-go/x/machine/types"
	elements "github.com/rddl-network/elements-rpc"
	elementsmocks "github.com/rddl-network/elements-rpc/utils/mocks"
	scctypes "github.com/rddl-network/shamir-coordinator-service/types"
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
	t.Setenv(config.ValAddr, "plmnt10mq5nj8jhh27z7ejnz2ql3nh0qhzjnfvy50877")
	ctrl := gomock.NewController(t)
	elements.Client = &elementsmocks.MockClient{}
	shamirMock := clientmocks.NewMockIShamirCoordinatorClient(ctrl)
	shamirMock.EXPECT().IssueMachineNFT(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(scctypes.IssueMachineNFTResponse{
		HexTX:    "0000000000000000000000000000000000000000000000000000000000000000",
		Contract: `{"entity":{"domain":"testnet-assets.rddl.io"}, "issuer_pubkey":"02", "machine_addr":"addr","name":"machine","precicion":8,"version":1}`,
		Asset:    "0000000000000000000000000000000000000000000000000000000000000000",
	}, nil)
	clients.ShamirCoordinatorServiceClient = shamirMock
	util.RegisterAssetServiceHTTPClient = &mocks.MockClient{}
	_, ctx := keeper.MachineKeeper(t)
	params := types.DefaultParams()
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			randomInt := rand.Int()
			sk, pk := sample.KeyPair(randomInt)
			machine := moduleobject.MachineRandom(pk, pk, sk, "address "+strconv.Itoa(randomInt), randomInt)
			goCtx := sdk.WrapSDKContext(ctx)

			err := util.IssueMachineNFT(goCtx, &machine, params.AssetRegistryScheme, params.AssetRegistryDomain, params.AssetRegistryPath)
			assert.NoError(t, err)
			wg.Done()
		}()
	}
	wg.Wait()
}
