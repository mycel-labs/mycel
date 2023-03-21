package keeper_test

import (
	"context"
	"testing"

	keepertest "mycel/testutil/keeper"
	"mycel/x/registry"
	"mycel/x/registry/keeper"
	"mycel/x/registry/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, keeper.Keeper, context.Context) {
	k, ctx := keepertest.RegistryKeeper(t)
	registry.InitGenesis(ctx, *k, *types.DefaultGenesis())
	return keeper.NewMsgServerImpl(*k), *k, sdk.WrapSDKContext(ctx)
}
