package credimint_test

import (
	"testing"

	keepertest "credimint/testutil/keeper"
	"credimint/testutil/nullify"
	"credimint/x/credimint"
	"credimint/x/credimint/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
		PortId: types.PortID,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CredimintKeeper(t)
	credimint.InitGenesis(ctx, *k, genesisState)
	got := credimint.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.PortId, got.PortId)

	// this line is used by starport scaffolding # genesis/test/assert
}
