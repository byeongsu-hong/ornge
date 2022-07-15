package keeper

import (
	"github.com/frostornge/ignite-example/x/igniteexample/types"
)

var _ types.QueryServer = Keeper{}
