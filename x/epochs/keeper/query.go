package keeper

import (
	"github.com/mycel-domain/mycel/x/epochs/types"
)

var _ types.QueryServer = Keeper{}
