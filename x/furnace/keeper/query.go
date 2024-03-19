package keeper

import (
	"github.com/mycel-domain/mycel/x/furnace/types"
)

var _ types.QueryServer = Keeper{}
