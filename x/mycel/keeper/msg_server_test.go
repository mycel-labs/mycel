package keeper_test

import (
	"context"
	"testing"

	keepertest "mycel/testutil/keeper"
	"mycel/x/mycel"
	"mycel/x/mycel/keeper"
	"mycel/x/mycel/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, keeper.Keeper, context.Context) {
	k, ctx := keepertest.MycelKeeper(t)
	mycel.InitGenesis(ctx, *k, *types.DefaultGenesis())
	return keeper.NewMsgServerImpl(*k), *k, sdk.WrapSDKContext(ctx)
}
