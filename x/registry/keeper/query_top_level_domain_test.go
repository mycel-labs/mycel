package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/mycel-domain/mycel/testutil/keeper"
	"github.com/mycel-domain/mycel/testutil/nullify"
	"github.com/mycel-domain/mycel/x/registry/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestTopLevelDomainQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.RegistryKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNTopLevelDomain(keeper, ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetTopLevelDomainRequest
		response *types.QueryGetTopLevelDomainResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetTopLevelDomainRequest{
				Name: msgs[0].Name,
			},
			response: &types.QueryGetTopLevelDomainResponse{TopLevelDomain: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetTopLevelDomainRequest{
				Name: msgs[1].Name,
			},
			response: &types.QueryGetTopLevelDomainResponse{TopLevelDomain: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetTopLevelDomainRequest{
				Name: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.TopLevelDomain(wctx, tc.request)
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

func TestTopLevelDomainQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.RegistryKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNTopLevelDomain(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllTopLevelDomainRequest {
		return &types.QueryAllTopLevelDomainRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.TopLevelDomainAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.TopLevelDomain), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.TopLevelDomain),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.TopLevelDomainAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.TopLevelDomain), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.TopLevelDomain),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.TopLevelDomainAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.TopLevelDomain),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.TopLevelDomainAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
