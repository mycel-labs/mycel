package keeper_test

import (
	"fmt"
	"strconv"
	"testing"
	"time"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

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
