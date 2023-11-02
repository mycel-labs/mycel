package keeper_test

import (
	"fmt"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	"github.com/mycel-domain/mycel/testutil"
	"github.com/mycel-domain/mycel/x/registry/types"
)

func (suite *KeeperTestSuite) TestWithdrawRegistrationFee() {
	testCases := []struct {
		withdrawer         string
		topLevelDomainName string
		expErr             error
		fn                 func()
	}{
		{
			withdrawer:         testutil.Alice,
			topLevelDomainName: "bar",
		},
		{
			withdrawer:         testutil.Bob,
			topLevelDomainName: "bar",
			expErr:             errorsmod.Wrapf(types.ErrNoPermissionToWithdraw, "%s", testutil.Bob),
		},
	}

	for i, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %d", i), func() {
			suite.SetupTest()

			// Register top level domain
			registerTopLevelDomainMsg := &types.MsgRegisterTopLevelDomain{
				Creator:                  testutil.Alice,
				Name:                     tc.topLevelDomainName,
				RegistrationPeriodInYear: 1,
			}
			_, err := suite.msgServer.RegisterTopLevelDomain(suite.ctx, registerTopLevelDomainMsg)
			suite.Require().Nil(err)

			// Register second level domain
			registerSecondLevelDomainMsg := &types.MsgRegisterSecondLevelDomain{
				Creator:                  testutil.Bob,
				Name:                     "foo",
				Parent:                   tc.topLevelDomainName,
				RegistrationPeriodInYear: 1,
			}
			_, err = suite.msgServer.RegisterSecondLevelDomain(suite.ctx, registerSecondLevelDomainMsg)
			suite.Require().Nil(err)

			// Before balance
			withdrawerAddress, err := sdk.AccAddressFromBech32(tc.withdrawer)
			suite.Require().Nil(err)
			beforeWithdrawerBalance := suite.app.BankKeeper.GetAllBalances(suite.ctx, withdrawerAddress)
			registryAddress := authtypes.NewModuleAddress(types.ModuleName)
			beforeRegistryAddress := suite.app.BankKeeper.GetAllBalances(suite.ctx, registryAddress)

			// Withdraw registration fee
			withdrawMsg := &types.MsgWithdrawRegistrationFee{
				Creator: tc.withdrawer,
				Name:    tc.topLevelDomainName,
			}
			fees, err := suite.msgServer.WithdrawRegistrationFee(suite.ctx, withdrawMsg)

			if tc.expErr == nil {
				suite.Require().Nil(err)

				// Check event
				events, found := testutil.FindEventsByType(suite.ctx.EventManager().Events(), types.EventTypeWithdrawRegistrationFee)
				suite.Require().True(found)
				for _, event := range events {
					suite.Require().Equal(tc.topLevelDomainName, event.Attributes[0].Value)
					suite.Require().Equal(fees.RegistrationFee.String(), event.Attributes[1].Value)
				}

				// Check balance
				afterWithdrawerBalance := suite.app.BankKeeper.GetAllBalances(suite.ctx, withdrawerAddress)
				suite.Require().Equal(beforeWithdrawerBalance.Add(fees.RegistrationFee...), afterWithdrawerBalance)
				afterRegistyAddress := suite.app.BankKeeper.GetAllBalances(suite.ctx, registryAddress)
				suite.Require().Equal(beforeRegistryAddress.Sub(fees.RegistrationFee...), afterRegistyAddress)

				// Check top level domain
				topLevelDomain, found := suite.app.RegistryKeeper.GetTopLevelDomain(suite.ctx, tc.topLevelDomainName)
				suite.Require().True(found)
				suite.Require().True(topLevelDomain.TotalWithdrawalAmount.IsEqual(sdk.NewCoins()))

			} else {
				suite.Require().EqualError(err, tc.expErr.Error())
			}
		})
	}
}
