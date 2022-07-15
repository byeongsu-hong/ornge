package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/frostornge/ignite-example/testutil/keeper"
	"github.com/frostornge/ignite-example/x/igniteexample/keeper"
	"github.com/frostornge/ignite-example/x/igniteexample/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.IgniteexampleKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
