package keeper_test

import (
	"errors"
	"fmt"
	"strconv"
	"testing"

	keepertest "github.com/mycel-domain/mycel/testutil/keeper"
	"github.com/mycel-domain/mycel/testutil/nullify"
	"github.com/mycel-domain/mycel/x/registry/keeper"
	"github.com/mycel-domain/mycel/x/registry/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNDomain(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.SecondLevelDomain {
	items := make([]types.SecondLevelDomain, n)
	for i := range items {
		items[i].Name = strconv.Itoa(i)
		items[i].Parent = strconv.Itoa(i)

		keeper.SetSecondLevelDomain(ctx, items[i])
	}
	return items
}

func TestDomainGet(t *testing.T) {
	keeper, ctx := keepertest.RegistryKeeper(t)
	items := createNDomain(keeper, ctx, 10)
	for _, item := range items {
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
func TestDomainRemove(t *testing.T) {
	keeper, ctx := keepertest.RegistryKeeper(t)
	items := createNDomain(keeper, ctx, 10)
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
	items := createNDomain(keeper, ctx, 10)
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
				Parent:         "test",
				ExpirationDate: suite.ctx.BlockTime().AddDate(0, 0, 20).UnixNano(),
			},
			expErr: nil,
		},
		{
			secondLevelDomain: types.SecondLevelDomain{
				Name:           "test",
				Parent:         "test",
				ExpirationDate: 0,
			},
			expErr: nil,
		},
		{
			secondLevelDomain: types.SecondLevelDomain{
				Name:           "test",
				Parent:         "test",
				ExpirationDate: suite.ctx.BlockTime().AddDate(0, 0, -20).UnixNano(),
			},
			expErr: sdkerrors.Wrapf(errors.New(fmt.Sprintf("test")), types.ErrDomainExpired.Error()),
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
