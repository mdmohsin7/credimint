package lqs_test

import (
	"testing"

	keepertest "credimint/testutil/keeper"
	"credimint/testutil/nullify"
	"credimint/x/lqs"
	"credimint/x/lqs/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.LqsKeeper(t)
	lqs.InitGenesis(ctx, *k, genesisState)
	got := lqs.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
