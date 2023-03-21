package keeper

import (
	"mycel/x/registry/types"
)

var _ types.QueryServer = Keeper{}
