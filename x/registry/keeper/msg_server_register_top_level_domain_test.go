package keeper_test

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/mycel-domain/mycel/testutil"
	furnacetypes "github.com/mycel-domain/mycel/x/furnace/types"
	"github.com/mycel-domain/mycel/x/registry/types"

	errorsmod "cosmossdk.io/errors"
)

func (suite *KeeperTestSuite) TestRegisterTopLevelDomain() {
	testCases := []struct {
		creator                  string
		name                     string
		registrationPeriodInYear uint64
		expErr                   error
		fn                       func()
	}{
		{
			creator:                  testutil.Alice,
			name:                     "cel0",
			registrationPeriodInYear: 1,
			expErr:                   nil,
			fn:                       func() {},
		},
		{
			creator:                  testutil.Alice,
			name:                     "cel1",
			registrationPeriodInYear: 4,
			expErr:                   nil,
			fn:                       func() {},
		},
		{
			creator:                  testutil.Alice,
			name:                     "cel2",
			registrationPeriodInYear: 1,
			expErr:                   errorsmod.Wrapf(types.ErrTopLevelDomainAlreadyTaken, "cel2"),
			fn: func() {
				// Register domain once
				domain := &types.MsgRegisterTopLevelDomain{
					Creator:                  testutil.Alice,
					Name:                     "cel2",
					RegistrationPeriodInYear: 1,
				}
				_, err := suite.msgServer.RegisterTopLevelDomain(suite.ctx, domain)
				suite.Require().Nil(err)
			},
		},
	}

	for i, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %d", i), func() {
			suite.SetupTest()

			registerMsg := &types.MsgRegisterTopLevelDomain{
				Creator:                  tc.creator,
				Name:                     tc.name,
				RegistrationPeriodInYear: tc.registrationPeriodInYear,
			}

			// Run test case function
			tc.fn()

			// Before balances
			furnaceAddress := authtypes.NewModuleAddress(furnacetypes.ModuleName)
			beforeFurnaceBalance := suite.app.BankKeeper.GetAllBalances(suite.ctx, furnaceAddress)
			beforeTreasuryBalance := suite.app.DistrKeeper.GetFeePool(suite.ctx).CommunityPool

			// Register domain
			_, err := suite.msgServer.RegisterTopLevelDomain(suite.ctx, registerMsg)

			// After balances
			afterFurnaceBalance := suite.app.BankKeeper.GetAllBalances(suite.ctx, furnaceAddress)
			afterTreasuryBalance := suite.app.DistrKeeper.GetFeePool(suite.ctx).CommunityPool

			if tc.expErr == nil {
				// Evaluate if domain is registered
				domain, found := suite.app.RegistryKeeper.GetTopLevelDomain(suite.ctx, tc.name)
				suite.Require().True(found)
				suite.Require().Equal(domain.AccessControl[tc.creator], types.DomainRole_OWNER)

				// Evalute events
				events, found := testutil.FindEventsByType(suite.ctx.EventManager().Events(), types.EventTypeRegisterTopLevelDomain)
				suite.Require().True(found)
				for _, event := range events {
					suite.Require().Equal(tc.name, event.Attributes[0].Value)

					// Check if the registration fee is correct
					total, err := sdk.ParseCoinsNormalized(event.Attributes[3].Value)
					suite.Require().Nil(err)
					toBurn, err := sdk.ParseCoinNormalized(event.Attributes[5].Value)
					suite.Require().Nil(err)
					toTreasury, err := sdk.ParseCoinNormalized(event.Attributes[6].Value)
					suite.Require().Nil(err)

					// Check if the total is equal to the sum of toBurn and toTreasury
					if total.Len() == 1 {
						suite.Require().Equal(total, sdk.NewCoins(toBurn.Add(toTreasury)))

					} else {
						suite.Require().Equal(total, sdk.NewCoins(toBurn, toTreasury))
					}

					// Check if the furnace balance is increased
					expectedFurnaceBalance := beforeFurnaceBalance.Add(toBurn)
					suite.Require().Equal(expectedFurnaceBalance, afterFurnaceBalance)
					expectedTreasuryBalance := beforeTreasuryBalance.Add(sdk.NewDecCoinFromCoin(toTreasury))
					suite.Require().Equal(expectedTreasuryBalance, afterTreasuryBalance)
				}

			} else {
				suite.Require().EqualError(err, tc.expErr.Error())
			}
		})
	}
}
