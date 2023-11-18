package keeper_test

import (
	"fmt"
	"strconv"
	"testing"
	"time"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/mycel-domain/mycel/testutil"
	keepertest "github.com/mycel-domain/mycel/testutil/keeper"
	"github.com/mycel-domain/mycel/testutil/nullify"
	"github.com/mycel-domain/mycel/x/registry/keeper"
	"github.com/mycel-domain/mycel/x/registry/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNTopLevelDomain(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.TopLevelDomain {
	items := make([]types.TopLevelDomain, n)
	for i := range items {
		items[i].Name = strconv.Itoa(i)

		keeper.SetTopLevelDomain(ctx, items[i])
	}
	return items
}

// Register top-level domains with k.RegisterTopLevelDomain()
// Domain name is set to `celn` (n is a incremantal number)
// e.g.) `cel1`, `cel2`, `celn`...
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

func TestTopLevelDomainGet(t *testing.T) {
	keeper, ctx := keepertest.RegistryKeeper(t)
	items := createNTopLevelDomain(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetTopLevelDomain(ctx,
			item.Name,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestTopLevelDomainRemove(t *testing.T) {
	keeper, ctx := keepertest.RegistryKeeper(t)
	items := createNTopLevelDomain(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveTopLevelDomain(ctx,
			item.Name,
		)
		_, found := keeper.GetTopLevelDomain(ctx,
			item.Name,
		)
		require.False(t, found)
	}
}

func TestTopLevelDomainGetAll(t *testing.T) {
	keeper, ctx := keepertest.RegistryKeeper(t)
	items := createNTopLevelDomain(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllTopLevelDomain(ctx)),
	)
}

func (suite *KeeperTestSuite) TestGetValidTopLevelDomain() {
	testCases := []struct {
		topLevelDomain types.TopLevelDomain
		expErr         error
	}{
		{
			topLevelDomain: types.TopLevelDomain{
				Name:           "test",
				ExpirationDate: suite.ctx.BlockTime().AddDate(0, 0, 20),
			},
			expErr: nil,
		},
		{
			topLevelDomain: types.TopLevelDomain{
				Name:           "test",
				ExpirationDate: time.Time{},
			},
			expErr: nil,
		},
		{
			topLevelDomain: types.TopLevelDomain{
				Name:           "test",
				ExpirationDate: suite.ctx.BlockTime().AddDate(0, 0, -20),
			},
			expErr: errorsmod.Wrapf(types.ErrTopLevelDomainExpired, "test"),
		},
	}
	for i, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %d", i), func() {
			suite.SetupTest()

			// Set domain
			suite.app.RegistryKeeper.SetTopLevelDomain(suite.ctx, tc.topLevelDomain)

			// Get valid domain
			topLevelDomain, err := suite.app.RegistryKeeper.GetValidTopLevelDomain(suite.ctx, tc.topLevelDomain.Name)
			if tc.expErr == nil {
				suite.Require().Nil(err)
				suite.Require().Equal(tc.topLevelDomain, topLevelDomain)
			} else {
				suite.Require().NotNil(err)
				suite.Require().Equal(tc.expErr.Error(), err.Error())
			}
		})
	}

}

func (suite *KeeperTestSuite) TestGetTopLevelDomainRole() {
	suite.SetupTest()
	k := suite.app.RegistryKeeper
	ctx := suite.ctx
	creator := testutil.Alice

	tlds, err := registerNTopLevelDomain(&k, ctx, creator, 1)
	if err != nil {
		suite.FailNow(fmt.Sprintf("%v", err))
	}

	type resp struct {
		Role  types.DomainRole
		Found bool
	}

	tcs := []struct {
		desc     string
		request  *types.QueryRoleRequest
		response *resp
	}{
		{
			desc: "Owner",
			request: &types.QueryRoleRequest{
				DomainName: tlds[0].Name,
				Address:    creator,
			},
			response: &resp{
				Role:  types.DomainRole_OWNER,
				Found: true,
			},
		},
		// TODO: Add a test case for EDITOR
		{
			desc: "Not owner of TLD",
			request: &types.QueryRoleRequest{
				DomainName: tlds[0].Name,
				Address:    testutil.Bob,
			},
			response: &resp{
				Role:  types.DomainRole_NO_ROLE,
				Found: true,
			},
		},
		{
			desc: "Domain not found",
			request: &types.QueryRoleRequest{
				DomainName: "notexist",
				Address:    creator,
			},
			response: &resp{
				Role:  types.DomainRole_NO_ROLE,
				Found: false,
			},
		},
	}

	for i, tc := range tcs {
		suite.Run(fmt.Sprintf("Case %d", i), func() {
			role, found := k.GetTopLevelDomainRole(ctx, tc.request.DomainName, tc.request.Address)
			suite.Require().Equal(tc.response.Role, role)
			suite.Require().Equal(tc.response.Found, found)
		})
	}
}
