package keeper

import (
	"mycel/x/epochs/types"
)

var _ types.QueryServer = Keeper{}
