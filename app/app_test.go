package app

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestAccountAddressPrefix makes sure that the account address prefix has a certain value.
func TestAccountAddressPrefix(t *testing.T) {
	t.Parallel()
	accountAddressPrefix := "plmnt"
	assert.Equal(t, AccountAddressPrefix, accountAddressPrefix, "The account address prefix should be 'plmnt'.")
}
