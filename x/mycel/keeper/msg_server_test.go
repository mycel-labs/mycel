package keeper_test

import (
	"context"
	"fmt"
	"testing"

	"mycel/testutil"
	keepertest "mycel/testutil/keeper"
	"mycel/x/mycel"
	"mycel/x/mycel/keeper"
	"mycel/x/mycel/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func GetMsgCreateDomain() *types.MsgCreateDomain {
	return &types.MsgCreateDomain{
		Creator:                  testutil.Alice,
		Name:                     "poyo",
		Parent:                   "ninniku",
		RegistrationPeriodInYear: 1,
	}
}

func setupMsgServer(t testing.TB) (types.MsgServer, keeper.Keeper, context.Context) {
	k, ctx := keepertest.MycelKeeper(t)
	mycel.InitGenesis(ctx, *k, *types.DefaultGenesis())
	return keeper.NewMsgServerImpl(*k), *k, sdk.WrapSDKContext(ctx)
}

func TestCreateDomainSuccess(t *testing.T) {
	msgServer, _, context := setupMsgServer(t)
	_, err := msgServer.CreateDomain(context, GetMsgCreateDomain())
	require.Nil(t, err)

}

func TestCreateDomainIsDomainAlreadyTakenFailure(t *testing.T) {
	msgServer, _, context := setupMsgServer(t)
	domain := GetMsgCreateDomain()
	_, err1 := msgServer.CreateDomain(context, domain)
	require.Nil(t, err1)
	_, err2 := msgServer.CreateDomain(context, domain)
	require.EqualError(t, err2, fmt.Sprintf("domain is already taken: %s.%s", domain.Name, domain.Parent))
}
