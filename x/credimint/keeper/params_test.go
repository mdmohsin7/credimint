package keeper_test

import (
	"testing"

	testkeeper "credimint/testutil/keeper"
	"credimint/x/credimint/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.CredimintKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
