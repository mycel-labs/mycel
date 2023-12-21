package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"

	keepertest "github.com/mycel-domain/mycel/testutil/keeper"
	"github.com/mycel-domain/mycel/x/epochs/keeper"
	"github.com/mycel-domain/mycel/x/epochs/types"
)

func setupMsgServer(tb testing.TB) (types.MsgServer, context.Context) {
	tb.Helper()
	k, ctx := keepertest.EpochsKeeper(tb)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
