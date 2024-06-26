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
