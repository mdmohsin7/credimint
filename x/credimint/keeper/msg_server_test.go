package keeper_test

import (
	"context"
	"testing"

	keepertest "credimint/testutil/keeper"
	"credimint/x/credimint/keeper"
	"credimint/x/credimint/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.CredimintKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
