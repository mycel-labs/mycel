package keeper_test

import (
	"errors"
	"fmt"

	"github.com/mycel-domain/mycel/testutil"
	"github.com/mycel-domain/mycel/x/registry/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (suite *KeeperTestSuite) TestUpdateDnsRecord() {
	testCases := []struct {
		creator       string
		name          string
		parent        string
		dnsRecordType string
		value         string
		expErr        error
		fn            func()
	}{
		{
			creator:       testutil.Alice,
			name:          "foo",
			parent:        "cel",
			dnsRecordType: "A",
			value:         "192.168.0.1",
			expErr:        nil,
			fn:            func() {},
		},
		{
			creator:       testutil.Alice,
			name:          "foo",
			parent:        "cel",
			dnsRecordType: "AAAA",
			value:         "2001:0db8:85a3:0000:0000:8a2e:0370:7334",
			expErr:        nil,
			fn:            func() {},
		},
		{
			creator:       testutil.Alice,
			name:          "foo",
			parent:        "cel",
			dnsRecordType: "CNAME",
			value:         "example.com.",
			expErr:        nil,
			fn:            func() {},
		},
		{
			creator:       testutil.Alice,
			name:          "hoge",
			parent:        "fuga",
			dnsRecordType: "A",
			value:         "192.168.0.1",
			expErr:        errorsmod.Wrapf(errors.New(fmt.Sprintf("hoge.fuga")), types.ErrDomainNotFound.Error()),
			fn:            func() {},
		},
		{
			creator:       testutil.Bob,
			name:          "foo",
			parent:        "cel",
			dnsRecordType: "A",
			value:         "192.168.0.1",
			expErr:        errorsmod.Wrapf(errors.New(fmt.Sprintf(testutil.Bob)), types.ErrDomainNotEditable.Error()),
			fn:            func() {},
		},
	}
	for i, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %d", i), func() {
			suite.SetupTest()

			// Register domain
			domain := &types.MsgRegisterSecondLevelDomain{
				Creator:                  testutil.Alice,
				Name:                     "foo",
				Parent:                   "cel",
				RegistrationPeriodInYear: 1,
			}
			_, err := suite.msgServer.RegisterSecondLevelDomain(suite.ctx, domain)
			suite.Require().Nil(err)
			// Run test case function
			tc.fn()

			// Update dns record
			msgUpdateRecord := &types.MsgUpdateDnsRecord{
				Creator:       tc.creator,
				Name:          tc.name,
				Parent:        tc.parent,
				DnsRecordType: tc.dnsRecordType,
				Value:         tc.value,
			}
			_, err = suite.msgServer.UpdateDnsRecord(suite.ctx, msgUpdateRecord)

			if tc.expErr == nil {
				// Evalute events
				suite.Require().Nil(err)
				res, _ := suite.app.RegistryKeeper.GetSecondLevelDomain(suite.ctx, domain.Name, domain.Parent)
				suite.Require().Equal(tc.value, res.Records[tc.dnsRecordType].GetDnsRecord().GetValue())

				// Event check
				events := sdk.StringifyEvents(suite.ctx.EventManager().ABCIEvents())
				eventIndex := len(events) - 1
				suite.Require().EqualValues(sdk.StringEvent{
					Type: types.EventTypeUpdateDnsRecord,
					Attributes: []sdk.Attribute{
						{Key: types.AttributeUpdateDnsRecordEventDomainName, Value: tc.name},
						{Key: types.AttributeUpdateDnsRecordEventDomainParent, Value: tc.parent},
						{Key: types.AttributeUpdateDnsRecordEventDnsRecordType, Value: tc.dnsRecordType},
						{Key: types.AttributeUpdateDnsRecordEventValue, Value: tc.value},
					},
				},
					events[eventIndex])

			} else {
				suite.Require().EqualError(err, tc.expErr.Error())
			}

		})
	}

}
