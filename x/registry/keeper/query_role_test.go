package keeper_test

import (
	"fmt"

	errorsmod "cosmossdk.io/errors"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/mycel-domain/mycel/testutil"
	"github.com/mycel-domain/mycel/x/registry/types"
)

func (suite *KeeperTestSuite) TestRole() {
	suite.SetupTest()
	k := suite.app.RegistryKeeper
	ctx := suite.ctx
	creator := testutil.Alice

	tlds, err := registerNTopLevelDomain(&k, ctx, creator, 1)
	if err != nil {
		suite.FailNow(fmt.Sprintf("%v", err))
	}

	slds, err := registerNSecondLevelDomain(&k, ctx, creator, 1)
	if err != nil {
		suite.FailNow(fmt.Sprintf("%v", err))
	}

	tcs := []struct {
		desc     string
		request  *types.QueryRoleRequest
		response *types.QueryRoleResponse
		err      error
	}{
		{
			desc: "TLD",
			request: &types.QueryRoleRequest{
				DomainName: tlds[0].Name,
				Address:    creator,
			},
			response: &types.QueryRoleResponse{
				Role: types.DomainRole.String(1),
			},
		},
		{
			desc: "SLD",
			request: &types.QueryRoleRequest{
				DomainName: fmt.Sprintf("%s.%s", slds[0].Name, tlds[0].Name),
				Address:    creator,
			},
			response: &types.QueryRoleResponse{
				Role: types.DomainRole.String(1),
			},
		},
		{
			desc: "Not owner of TLD",
			request: &types.QueryRoleRequest{
				DomainName: tlds[0].Name,
				Address:    testutil.Bob,
			},
			response: &types.QueryRoleResponse{
				Role: types.DomainRole.String(0),
			},
		},
		{
			desc: "Not owner of SLD",
			request: &types.QueryRoleRequest{
				DomainName: fmt.Sprintf("%s.%s", slds[0].Name, tlds[0].Name),
				Address:    testutil.Bob,
			},
			response: &types.QueryRoleResponse{
				Role: types.DomainRole.String(0),
			},
		},
		{
			desc: "Domain not found",
			request: &types.QueryRoleRequest{
				DomainName: "notexist",
				Address:    creator,
			},
			err: errorsmod.Wrapf(sdkerrors.ErrNotFound, "domain not found"),
		},
		{
			desc: "InvalidRequest",
			err:  errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "invalid request: empty request"),
		},
	}

	for i, tc := range tcs {
		suite.Run(fmt.Sprintf("Case %d", i), func() {
			// Get valid domain
			resp, err := k.Role(ctx, tc.request)
			if tc.err == nil {
				suite.Require().Nil(err)
				suite.Require().Equal(tc.response, resp)
			} else {
				suite.Require().NotNil(err)
				suite.Require().Equal(tc.err.Error(), err.Error())
			}
		})
	}
}
