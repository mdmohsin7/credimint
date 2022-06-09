package types_test

import (
	"testing"

	"credimint/x/credimint/types"
	"github.com/stretchr/testify/require"
)

func TestGenesisState_Validate(t *testing.T) {
	for _, tc := range []struct {
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

				UserCredibilityList: []types.UserCredibility{
					{
						Id: 0,
					},
					{
						Id: 1,
					},
				},
				UserCredibilityCount: 2,
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated userCredibility",
			genState: &types.GenesisState{
				UserCredibilityList: []types.UserCredibility{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid userCredibility count",
			genState: &types.GenesisState{
				UserCredibilityList: []types.UserCredibility{
					{
						Id: 1,
					},
				},
				UserCredibilityCount: 0,
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	} {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
