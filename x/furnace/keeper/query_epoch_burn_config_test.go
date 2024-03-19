package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/mycel-domain/mycel/testutil/keeper"
	"github.com/mycel-domain/mycel/testutil/nullify"
	"github.com/mycel-domain/mycel/x/furnace/types"
)

func TestEpochBurnConfigQuery(t *testing.T) {
	keeper, ctx := keepertest.FurnaceKeeper(t)
	item := createTestEpochBurnConfig(keeper, ctx)
	tests := []struct {
		desc     string
		request  *types.QueryGetEpochBurnConfigRequest
		response *types.QueryGetEpochBurnConfigResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetEpochBurnConfigRequest{},
			response: &types.QueryGetEpochBurnConfigResponse{EpochBurnConfig: item},
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.EpochBurnConfig(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}
