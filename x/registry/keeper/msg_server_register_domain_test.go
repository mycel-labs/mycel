package keeper_test

import (
	"fmt"
	"testing"

	"mycel/testutil"
	"mycel/x/registry/types"

	"github.com/stretchr/testify/require"
)

func GetMsgRegisterDomain() *types.MsgRegisterDomain {
	return &types.MsgRegisterDomain{
		Creator:                  testutil.Alice,
		Name:                     "poyo",
		Parent:                   "ninniku",
		RegistrationPeriodInYear: 1,
	}
}

func TestRegisterDomainSuccess(t *testing.T) {
	msgServer, _, context := setupMsgServer(t)
	_, err := msgServer.RegisterDomain(context, GetMsgRegisterDomain())
	require.Nil(t, err)

}

func TestRegisterDomainIsDomainAlreadyTakenFailure(t *testing.T) {
	msgServer, _, context := setupMsgServer(t)
	domain := GetMsgRegisterDomain()
	_, err1 := msgServer.RegisterDomain(context, domain)
	require.Nil(t, err1)
	_, err2 := msgServer.RegisterDomain(context, domain)
	require.EqualError(t, err2, fmt.Sprintf("domain is already taken: %s.%s", domain.Name, domain.Parent))
}
