package keeper_test

import (
	"testing"

	testkeeper "github.com/frostornge/ignite-example/testutil/keeper"
	"github.com/frostornge/ignite-example/x/igniteexample/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.IgniteexampleKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
