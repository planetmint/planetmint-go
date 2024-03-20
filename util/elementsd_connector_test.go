package util_test

import (
	"math/rand"
	"strconv"
	"sync"
	"testing"

	"github.com/planetmint/planetmint-go/testutil/moduleobject"
	"github.com/planetmint/planetmint-go/testutil/sample"
	"github.com/planetmint/planetmint-go/util"
	"github.com/planetmint/planetmint-go/x/machine/types"
	elements "github.com/rddl-network/elements-rpc"
	elementsmocks "github.com/rddl-network/elements-rpc/utils/mocks"
	"github.com/stretchr/testify/assert"
)

func TestReissueAsset(t *testing.T) {
	elements.Client = &elementsmocks.MockClient{}
	_, err := util.ReissueAsset("reissueasset 06c20c8de513527f1ae6c901f74a05126525ac2d7e89306f4a7fd5ec4e674403 900.000")
	assert.NoError(t, err)
}

func TestDistributeAsset(t *testing.T) {
	elements.Client = &elementsmocks.MockClient{}

	_, err := util.DistributeAsset(
		"tlq1qqt5078sef4aqls29c3j3pwfmukgjug70t37x26gwyhzpdxmtmjmphar88fwsl9qcm559jevve772prhtuyf9xkxdtrhvuce6a",
		"20",
		"06c20c8de513527f1ae6c901f74a05126525ac2d7e89306f4a7fd5ec4e674403")
	assert.NoError(t, err)
}

func TestIssueNFTAsset(t *testing.T) {
	elements.Client = &elementsmocks.MockClient{}

	params := types.DefaultParams()
	var wg sync.WaitGroup

	for i := 0; i < 1; i++ {
		wg.Add(1)
		go func() {
			randomInt := rand.Int()
			sk, pk := sample.KeyPair(randomInt)
			machine := moduleobject.MachineRandom(pk, pk, sk, "address "+strconv.Itoa(randomInt), randomInt)

			_, _, _, err := util.IssueNFTAsset(machine.Name, machine.Address, params.AssetRegistryDomain)
			assert.NoError(t, err)

			wg.Done()
		}()
	}
	wg.Wait()
}
