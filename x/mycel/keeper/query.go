package keeper

import (
	"mycel/x/mycel/types"
)

var _ types.QueryServer = Keeper{}
