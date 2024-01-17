package types_test

import (
	"testing"

	"github.com/planetmint/planetmint-go/x/dao/types"
	"github.com/stretchr/testify/require"
)

func TestGenesisState_Validate(t *testing.T) {
	t.Parallel()
	tests := []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{

				RedeemClaimList: []types.RedeemClaim{
					{
						Beneficiary:  "0",
						LiquidTxHash: "0",
					},
					{
						Beneficiary:  "1",
						LiquidTxHash: "1",
					},
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated redeemClaim",
			genState: &types.GenesisState{
				RedeemClaimList: []types.RedeemClaim{
					{
						Beneficiary:  "0",
						LiquidTxHash: "0",
					},
					{
						Beneficiary:  "0",
						LiquidTxHash: "0",
					},
				},
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			t.Parallel()
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
