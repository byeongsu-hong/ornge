package ornge_test

import (
	"testing"

	keepertest "github.com/frostornge/ornge/testutil/keeper"
	"github.com/frostornge/ornge/testutil/nullify"
	"github.com/frostornge/ornge/x/ornge"
	"github.com/frostornge/ornge/x/ornge/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.OrngeKeeper(t)
	ornge.InitGenesis(ctx, *k, genesisState)
	got := ornge.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
