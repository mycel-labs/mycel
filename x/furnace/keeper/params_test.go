package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	testkeeper "github.com/mycel-domain/mycel/testutil/keeper"
	"github.com/mycel-domain/mycel/x/furnace/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.FurnaceKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
