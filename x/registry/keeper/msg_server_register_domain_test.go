package keeper_test

import (
	"fmt"
	"testing"

	"mycel/testutil"
	"mycel/x/registry/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func GetMsgRegisterDomain() *types.MsgRegisterDomain {
	return &types.MsgRegisterDomain{
		Creator:                  testutil.Alice,
		Name:                     "foo",
		Parent:                   "cel",
		RegistrationPeriodInYear: 1,
	}
}

func TestRegisterDomainSuccess(t *testing.T) {
	msgServer, _, context := setupMsgServer(t)
	domain := GetMsgRegisterDomain()
	_, err := msgServer.RegisterDomain(context, domain)
	require.Nil(t, err)

	// Event emitted
	ctx := sdk.UnwrapSDKContext(context)
	require.NotNil(t, ctx)
	events := sdk.StringifyEvents(ctx.EventManager().ABCIEvents())
	require.Len(t, events, 1)
	require.EqualValues(t, sdk.StringEvent{
		Type: types.EventTypeRegsterDomain,
		Attributes: []sdk.Attribute{
			{Key: types.AttributeRegisterDomainEventName, Value: domain.Name},
			{Key: types.AttributeRegisterDomainEventParent, Value: domain.Parent},
			{Key: types.AttributeRegisterDomainEventRegistrationPeriodInYear, Value: fmt.Sprintf("%d", domain.RegistrationPeriodInYear)},
			{Key: types.AttributeRegisterDomainEventExpirationDate, Value: events[0].Attributes[3].Value},
			{Key: types.AttributeRegisterDomainLevel, Value: "2"},
		},
	}, events[0])

}

func TestRegisterDomainIsDomainAlreadyTakenFailure(t *testing.T) {
	msgServer, _, context := setupMsgServer(t)
	domain := GetMsgRegisterDomain()
	_, err1 := msgServer.RegisterDomain(context, domain)
	require.Nil(t, err1)
	_, err2 := msgServer.RegisterDomain(context, domain)
	require.EqualError(t, err2, fmt.Sprintf("domain is already taken: %s.%s", domain.Name, domain.Parent))
}
