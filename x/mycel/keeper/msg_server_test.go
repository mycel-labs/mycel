package keeper_test

import (
	"context"
	"testing"

	"mycel/testutil"
	keepertest "mycel/testutil/keeper"
	"mycel/x/mycel"
	"mycel/x/mycel/keeper"
	"mycel/x/mycel/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func setupMsgServer(t testing.TB) (types.MsgServer, keeper.Keeper, context.Context) {
	k, ctx := keepertest.MycelKeeper(t)
	mycel.InitGenesis(ctx, *k, *types.DefaultGenesis())
	return keeper.NewMsgServerImpl(*k), *k, sdk.WrapSDKContext(ctx)
}

func TestCreateDomainSuccess(t *testing.T) {
	msgServer, _, context := setupMsgServer(t)
	_, err := msgServer.CreateDomain(context, &types.MsgCreateDomain{
		Creator:                  testutil.Alice,
		Name:                     "poyo",
		Parent:                   "ninniku",
		RegistrationPeriodInYear: 1,
	})
	require.Nil(t, err)

}
