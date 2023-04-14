package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/mycel-domain/mycel/testutil/keeper"
	"github.com/mycel-domain/mycel/x/registry"
	"github.com/mycel-domain/mycel/x/registry/keeper"
	"github.com/mycel-domain/mycel/x/registry/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, keeper.Keeper, context.Context) {
	k, ctx := keepertest.RegistryKeeper(t)
	registry.InitGenesis(ctx, *k, *types.DefaultGenesis())
	return keeper.NewMsgServerImpl(*k), *k, sdk.WrapSDKContext(ctx)
}
