package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/mycel-domain/mycel/testutil/keeper"
	"github.com/mycel-domain/mycel/x/epochs/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.EpochsKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
