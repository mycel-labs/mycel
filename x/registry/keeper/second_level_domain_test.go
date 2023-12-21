package keeper_test

import (
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	errorsmod "cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/mycel-domain/mycel/testutil"
	keepertest "github.com/mycel-domain/mycel/testutil/keeper"
	"github.com/mycel-domain/mycel/testutil/nullify"
	"github.com/mycel-domain/mycel/x/registry/keeper"
	"github.com/mycel-domain/mycel/x/registry/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNSecondLevelDomain(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.SecondLevelDomain {
	items := make([]types.SecondLevelDomain, n)
	for i := range items {
		items[i].Name = strconv.Itoa(i)
		items[i].Parent = strconv.Itoa(i)

		keeper.SetSecondLevelDomain(ctx, items[i])
	}
	return items
}

func createNSecondLevelDomainResponse(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.SecondLevelDomainResponse {
	items := createNSecondLevelDomain(keeper, ctx, n)
	responses := make([]types.SecondLevelDomainResponse, n)
	for i := range responses {
		responses[i].Name = items[i].Name
		responses[i].Parent = items[i].Parent
		responses[i].ExpirationDate = items[i].ExpirationDate
	}
	return responses
}

// Register top-level domains with k.RegisterSecondLevelDomain()
// Domain name is set to `n` (n is a incremantal number)
// e.g.) `1`, `2`, `n`...
func registerNSecondLevelDomain(k *keeper.Keeper, ctx sdk.Context, creatorAddr string, n int) ([]types.SecondLevelDomain, error) {
	items := make([]types.SecondLevelDomain, n)
	for i := range items {
		creator, err := sdk.AccAddressFromBech32(creatorAddr)
		if err != nil {
			return nil, err
		}
		name := strconv.Itoa(i)
		accessControl := types.AccessControl{
			Address: creator.String(),
			Role:    types.DomainRole_OWNER,
		}
		sld := types.SecondLevelDomain{
			Name:           name,
			Parent:         "cel" + name,
			Owner:          creator.String(),
			ExpirationDate: time.Time{},
			Records:        nil,
			AccessControl:  []*types.AccessControl{&accessControl},
		}

		if err := k.RegisterSecondLevelDomain(ctx, sld, creator, 1); err != nil {
			return nil, err
		}
		items[i] = sld
	}
	return items, nil
}

func TestSecondLevelDomainGet(t *testing.T) {
	keeper, ctx := keepertest.RegistryKeeper(t)
	items := createNSecondLevelDomain(keeper, ctx, 10)
	for i := range items {
		item := items[i]
		rst, found := keeper.GetSecondLevelDomain(ctx,
			item.Name,
			item.Parent,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}

func TestSecondLevelDomainRemove(t *testing.T) {
	keeper, ctx := keepertest.RegistryKeeper(t)
	items := createNSecondLevelDomain(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveSecondLevelDomain(ctx,
			item.Name,
			item.Parent,
		)
		_, found := keeper.GetSecondLevelDomain(ctx,
			item.Name,
			item.Parent,
		)
		require.False(t, found)
	}
}

func TestDomainGetAll(t *testing.T) {
	keeper, ctx := keepertest.RegistryKeeper(t)
	items := createNSecondLevelDomain(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllSecondLevelDomain(ctx)),
	)
}

func (suite *KeeperTestSuite) TestGetValidSecondLevelDomain() {
	testCases := []struct {
		secondLevelDomain types.SecondLevelDomain
		expErr            error
	}{
		{
			secondLevelDomain: types.SecondLevelDomain{
				Name:           "test",
				Parent:         "cel",
				ExpirationDate: suite.ctx.BlockTime().AddDate(0, 0, 20),
			},
			expErr: nil,
		},
		{
			secondLevelDomain: types.SecondLevelDomain{
				Name:           "test",
				Parent:         "cel",
				ExpirationDate: time.Time{},
			},
			expErr: nil,
		},
		{
			secondLevelDomain: types.SecondLevelDomain{
				Name:           "test",
				Parent:         "test",
				ExpirationDate: suite.ctx.BlockTime().AddDate(0, 0, -20),
			},
			expErr: errorsmod.Wrapf(types.ErrTopLevelDomainNotFound, "test"),
		},
	}
	for i, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %d", i), func() {
			suite.SetupTest()

			// Set domain
			suite.app.RegistryKeeper.SetSecondLevelDomain(suite.ctx, tc.secondLevelDomain)

			// Get valid domain
			secondLevelDomain, err := suite.app.RegistryKeeper.GetValidSecondLevelDomain(suite.ctx, tc.secondLevelDomain.Name, tc.secondLevelDomain.Parent)
			if tc.expErr == nil {
				suite.Require().Nil(err)
				suite.Require().Equal(tc.secondLevelDomain, secondLevelDomain)
			} else {
				suite.Require().NotNil(err)
				suite.Require().Equal(tc.expErr.Error(), err.Error())
			}
		})
	}
}

func (suite *KeeperTestSuite) TestGetSecondLevelDomainRole() {
	suite.SetupTest()
	k := suite.app.RegistryKeeper
	ctx := suite.ctx
	creator := testutil.Alice

	_, err := registerNTopLevelDomain(&k, ctx, creator, 1)
	if err != nil {
		suite.FailNow(fmt.Sprintf("%v", err))
	}
	slds, err := registerNSecondLevelDomain(&k, ctx, creator, 1)
	if err != nil {
		suite.FailNow(fmt.Sprintf("%v", err))
	}

	type req struct {
		Name    string
		Parent  string
		Address string
	}

	type resp struct {
		Role  types.DomainRole
		Found bool
	}

	tcs := []struct {
		desc     string
		request  *req
		response *resp
	}{
		{
			desc: "Owner",
			request: &req{
				Name:    slds[0].Name,
				Parent:  slds[0].Parent,
				Address: testutil.Alice,
			},
			response: &resp{
				Role:  types.DomainRole_OWNER,
				Found: true,
			},
		},
		// TODO: Add a test case for EDITOR
		{
			desc: "Not owner of sld",
			request: &req{
				Name:    slds[0].Name,
				Parent:  slds[0].Parent,
				Address: testutil.Bob,
			},
			response: &resp{
				Role:  types.DomainRole_NO_ROLE,
				Found: true,
			},
		},
		{
			desc: "Domain not found because the name does not exist",
			request: &req{
				Name:    "notexist",
				Parent:  slds[0].Parent,
				Address: testutil.Alice,
			},
			response: &resp{
				Role:  types.DomainRole_NO_ROLE,
				Found: false,
			},
		},
		{
			desc: "Domain not found because the parent does not exist",
			request: &req{
				Name:    slds[0].Name,
				Parent:  "notexist",
				Address: testutil.Alice,
			},
			response: &resp{
				Role:  types.DomainRole_NO_ROLE,
				Found: false,
			},
		},
	}

	for i, tc := range tcs {
		suite.Run(fmt.Sprintf("Case %d", i), func() {
			role, found := k.GetSecondLevelDomainRole(ctx, tc.request.Name, tc.request.Parent, tc.request.Address)
			suite.Require().Equal(tc.response.Role, role)
			suite.Require().Equal(tc.response.Found, found)
		})
	}
}
