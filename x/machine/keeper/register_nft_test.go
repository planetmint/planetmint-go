package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/golang/mock/gomock"
	keepertest "github.com/planetmint/planetmint-go/testutil/keeper"
	"github.com/planetmint/planetmint-go/x/machine/keeper"
	"github.com/planetmint/planetmint-go/x/machine/testutil"
	"github.com/stretchr/testify/assert"
)

func TestRegisterNFT(t *testing.T) {
	_, ctx := keepertest.MachineKeeper(t)
	url := "https://testnet-assets.rddl.io/register_asset"
	contract := `{"entity":{"domain":"testnet-assets.rddl.io"},"issuer_pubkey":"020000000000000000000000000000000000000000000000000000000000000000","machine_addr":"plmnt10mq5nj8jhh27z7ejnz2ql3nh0qhzjnfvy50877","name":"machine","precision":0,"version":0}`
	asset := "0000000000000000000000000000000000000000000000001000000000000000"
	goctx := sdk.WrapSDKContext(ctx)

	// Call to set sync.Once
	_ = keeper.GetAssetServiceClient()

	ctrl := gomock.NewController(t)
	ascMock := testutil.NewMockIAssetServiceClient(ctrl)
	ascMock.EXPECT().RegisterAsset(goctx, asset, contract, url).Return(nil).AnyTimes()
	keeper.SetAssetServiceClient(ascMock)

	asc := keeper.GetAssetServiceClient()

	err := asc.RegisterAsset(goctx, asset, contract, url)
	assert.NoError(t, err)
}
