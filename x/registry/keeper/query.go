package keeper

import (
	"github.com/mycel-domain/mycel/x/registry/types"
)

var _ types.QueryServer = Keeper{}
