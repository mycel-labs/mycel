package keeper_test

import (
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/mycel-domain/mycel/app/params"
	"github.com/mycel-domain/mycel/testutil"
	furnacetypes "github.com/mycel-domain/mycel/x/furnace/types"
	"github.com/mycel-domain/mycel/x/registry/types"

	errorsmod "cosmossdk.io/errors"
)

func (suite *KeeperTestSuite) TestExtendTopLevelDomain() {
	name := "hoge"

	testCases := []struct {
		creator               string
		extensionPeriodInYear uint64
		expErr                error
		fn                    func()
	}{
		{
			creator:               testutil.Alice,
			extensionPeriodInYear: 1,
			expErr:                nil,
			fn:                    func() {},
		},
		{
			creator:               testutil.Bob,
			extensionPeriodInYear: 1,
			expErr:                errorsmod.Wrapf(types.ErrDomainNotEditable, "%s", testutil.Bob),
			fn:                    func() {},
		},
	}

	for i, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %d", i), func() {
			suite.SetupTest()

			// Register domain
			registerMsg := &types.MsgRegisterTopLevelDomain{
				Creator:                  testutil.Alice,
				Name:                     name,
				RegistrationPeriodInYear: 1,
			}
			registerRsp, err := suite.msgServer.RegisterTopLevelDomain(suite.ctx, registerMsg)
			suite.Require().Nil(err)
			beforeDomain, found := suite.app.RegistryKeeper.GetTopLevelDomain(suite.ctx, name)
			suite.Require().True(found)
			beforeExpirationDate := time.Unix(0, beforeDomain.ExpirationDate)
			_ = beforeExpirationDate

			// Run test case function
			tc.fn()

			// Before balances
			furnaceAddress := authtypes.NewModuleAddress(furnacetypes.ModuleName)
			beforeFurnaceBalance := suite.app.BankKeeper.GetAllBalances(suite.ctx, furnaceAddress)
			beforeTreasuryBalance := suite.app.DistrKeeper.GetFeePool(suite.ctx).CommunityPool

			// Extend domain
			extendMsg := &types.MsgExtendTopLevelDomainExpirationDate{
				Creator:               tc.creator,
				Name:                  name,
				ExtensionPeriodInYear: tc.extensionPeriodInYear,
			}
			extendRsp, err := suite.msgServer.ExtendTopLevelDomainExpirationDate(suite.ctx, extendMsg)

			// After balances
			afterFurnaceBalance := suite.app.BankKeeper.GetAllBalances(suite.ctx, furnaceAddress)
			afterTreasuryBalance := suite.app.DistrKeeper.GetFeePool(suite.ctx).CommunityPool

			if tc.expErr == nil {
				suite.Require().Nil(err)
				// Evaluate if domain is extended
				// Response
				registerExpirationDate := time.Unix(0, registerRsp.TopLevelDomain.ExpirationDate)
				extendExpirationDate := time.Unix(0, extendRsp.TopLevelDomain.ExpirationDate)
				suite.Require().Equal(registerExpirationDate.AddDate(0, 0, int(tc.extensionPeriodInYear)*params.OneYearInDays), extendExpirationDate)
				// Store
				afterDomain, found := suite.app.RegistryKeeper.GetTopLevelDomain(suite.ctx, name)
				suite.Require().True(found)
				expAfterExpirationDate := beforeExpirationDate.AddDate(0, 0, int(tc.extensionPeriodInYear)*params.OneYearInDays)
				afterExpirationDate := time.Unix(0, afterDomain.ExpirationDate)
				suite.Require().Equal(expAfterExpirationDate, afterExpirationDate)

				// Evalute events
				events, found := testutil.FindEventsByType(suite.ctx.EventManager().Events(), types.EventTypeExtendTopLevelDomainExpirationDate)
				suite.Require().True(found)
				for _, event := range events {
					suite.Require().Equal(name, event.Attributes[0].Value)

					// Check if the extension fee is correct
					total, err := sdk.ParseCoinsNormalized(event.Attributes[2].Value)
					suite.Require().Nil(err)
					toBurn, err := sdk.ParseCoinNormalized(event.Attributes[4].Value)
					suite.Require().Nil(err)
					toTreasury, err := sdk.ParseCoinNormalized(event.Attributes[5].Value)
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
