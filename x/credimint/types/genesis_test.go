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
				PortId: types.PortID,
				LoanList: []types.Loan{
					{
						Id: 0,
					},
					{
						Id: 1,
					},
				},
				LoanCount: 2,
				UserList: []types.User{
					{
						Index: "0",
					},
					{
						Index: "1",
					},
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated loan",
			genState: &types.GenesisState{
				LoanList: []types.Loan{
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
			desc: "invalid loan count",
			genState: &types.GenesisState{
				LoanList: []types.Loan{
					{
						Id: 1,
					},
				},
				LoanCount: 0,
			},
			valid: false,
		},
		{
			desc: "duplicated user",
			genState: &types.GenesisState{
				UserList: []types.User{
					{
						Index: "0",
					},
					{
						Index: "0",
					},
				},
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
