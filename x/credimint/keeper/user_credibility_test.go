package keeper_test

import (
	"testing"

	keepertest "credimint/testutil/keeper"
	"credimint/testutil/nullify"
	"credimint/x/credimint/keeper"
	"credimint/x/credimint/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func createNUserCredibility(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.UserCredibility {
	items := make([]types.UserCredibility, n)
	for i := range items {
		items[i].Id = keeper.AppendUserCredibility(ctx, items[i])
	}
	return items
}

func TestUserCredibilityGet(t *testing.T) {
	keeper, ctx := keepertest.CredimintKeeper(t)
	items := createNUserCredibility(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetUserCredibility(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestUserCredibilityRemove(t *testing.T) {
	keeper, ctx := keepertest.CredimintKeeper(t)
	items := createNUserCredibility(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveUserCredibility(ctx, item.Id)
		_, found := keeper.GetUserCredibility(ctx, item.Id)
		require.False(t, found)
	}
}

func TestUserCredibilityGetAll(t *testing.T) {
	keeper, ctx := keepertest.CredimintKeeper(t)
	items := createNUserCredibility(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllUserCredibility(ctx)),
	)
}

func TestUserCredibilityCount(t *testing.T) {
	keeper, ctx := keepertest.CredimintKeeper(t)
	items := createNUserCredibility(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetUserCredibilityCount(ctx))
}
