package keeper_test

import (
	"fmt"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

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
			beforeBalance := suite.app.BankKeeper.GetAllBalances(suite.ctx, withdrawerAddress)

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
				afterBalance := suite.app.BankKeeper.GetAllBalances(suite.ctx, withdrawerAddress)
				suite.Require().Equal(beforeBalance.Add(fees.RegistrationFee...), afterBalance)
			} else {
				suite.Require().EqualError(err, tc.expErr.Error())
			}
		})
	}
}
