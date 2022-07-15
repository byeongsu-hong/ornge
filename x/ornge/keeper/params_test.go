package keeper_test

import (
	"testing"

	testkeeper "github.com/frostornge/ornge/testutil/keeper"
	"github.com/frostornge/ornge/x/ornge/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.OrngeKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
