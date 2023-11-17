package keeper_test

import (
	"fmt"
	"strconv"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/mycel-domain/mycel/testutil"
	"github.com/mycel-domain/mycel/x/registry/keeper"
	"github.com/mycel-domain/mycel/x/registry/types"
)

func registerNTopLevelDomain(k *keeper.Keeper, ctx sdk.Context, creator string, n int) ([]types.TopLevelDomain, error) {
	items := make([]types.TopLevelDomain, n)
	for i := range items {
		creator := testutil.Alice
		name := "cel" + strconv.Itoa(i)

		tld, _, err := k.RegisterTopLevelDomain(ctx, creator, name, 1)
		if err != nil {
			return nil, err
		}
		items[i] = tld
	}
	return items, nil
}

func registerNSecondLevelDomain(k *keeper.Keeper, ctx sdk.Context, creator string, n int) ([]types.SecondLevelDomain, error) {
	items := make([]types.SecondLevelDomain, n)
	for i := range items {
		creator, err := sdk.AccAddressFromBech32(testutil.Alice)
		if err != nil {
			return nil, err
		}
		name := strconv.Itoa(i)
		accessControl := map[string]types.DomainRole{
			creator.String(): types.DomainRole_OWNER,
		}
		sld := types.SecondLevelDomain{
			Name:           name,
			Parent:         "cel" + name,
			Owner:          creator.String(),
			ExpirationDate: time.Time{},
			Records:        nil,
			AccessControl:  accessControl,
		}

		if err := k.RegisterSecondLevelDomain(ctx, sld, creator, 1); err != nil {
			return nil, err
		}
		items[i] = sld
	}
	return items, nil
}

func (suite *KeeperTestSuite) TestRole() {
	suite.SetupTest()
	k := suite.app.RegistryKeeper
	ctx := suite.ctx
	creator := testutil.Alice

	tlds, err := registerNTopLevelDomain(&k, ctx, creator, 1)
	if err != nil {
		suite.FailNow(fmt.Sprintf("%v", err))
	}

	slds, err := registerNSecondLevelDomain(&k, ctx, creator, 1)
	if err != nil {
		suite.FailNow(fmt.Sprintf("%v", err))
	}

	tcs := []struct {
		desc     string
		request  *types.QueryRoleRequest
		response *types.QueryRoleResponse
		err      error
	}{
		{
			desc: "TLD",
			request: &types.QueryRoleRequest{
				DomainName: tlds[0].Name,
				Address:    creator,
			},
			response: &types.QueryRoleResponse{
				Role: types.DomainRole.String(1),
			},
		},
		{
			desc: "SLD",
			request: &types.QueryRoleRequest{
				DomainName: fmt.Sprintf("%s.%s", slds[0].Name, tlds[0].Name),
				Address:    creator,
			},
			response: &types.QueryRoleResponse{
				Role: types.DomainRole.String(1),
			},
		},
		{
			desc: "Not owner of TLD",
			request: &types.QueryRoleRequest{
				DomainName: tlds[0].Name,
				Address:    testutil.Bob,
			},
			response: &types.QueryRoleResponse{
				Role: types.DomainRole.String(0),
			},
		},
		{
			desc: "Not owner of SLD",
			request: &types.QueryRoleRequest{
				DomainName: fmt.Sprintf("%s.%s", slds[0].Name, tlds[0].Name),
				Address:    testutil.Bob,
			},
			response: &types.QueryRoleResponse{
				Role: types.DomainRole.String(0),
			},
		},
		{
			desc: "Domain not found",
			request: &types.QueryRoleRequest{
				DomainName: "notexist",
				Address:    creator,
			},
			err: status.Error(codes.NotFound, "domain not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	}

	for i, tc := range tcs {
		suite.Run(fmt.Sprintf("Case %d", i), func() {

			// Get valid domain
			resp, err := k.Role(ctx, tc.request)
			if tc.err == nil {
				suite.Require().Nil(err)
				suite.Require().Equal(tc.response, resp)
			} else {
				suite.Require().NotNil(err)
				suite.Require().Equal(tc.err.Error(), err.Error())
			}
		})
	}
}
