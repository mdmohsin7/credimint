package keeper_test

import (
	"context"
	"testing"

	keepertest "credimint/testutil/keeper"
	"credimint/x/lqs/keeper"
	"credimint/x/lqs/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.LqsKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
