package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestDerivationPath makes sure that purpose and cointype are set to PLMNT (see https://github.com/satoshilabs/slips/blob/master/slip-0044.md)
func TestDerivationPath(t *testing.T) {
	t.Parallel()
	sdkConfig := initSDKConfig()

	purpose := uint32(44)
	assert.Equal(t, purpose, sdkConfig.GetPurpose())

	coinType := uint32(8680)
	assert.Equal(t, coinType, sdkConfig.GetCoinType())
}
