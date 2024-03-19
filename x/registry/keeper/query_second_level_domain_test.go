package keeper_test

import (
	"strconv"
	"testing"

	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/mycel-domain/mycel/testutil/keeper"

	"github.com/mycel-domain/mycel/testutil/nullify"
	"github.com/mycel-domain/mycel/x/registry/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestSecondLevelDomainQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.RegistryKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNSecondLevelDomainResponse(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetSecondLevelDomainRequest
		response *types.QueryGetSecondLevelDomainResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetSecondLevelDomainRequest{
				Name:   msgs[0].Name,
				Parent: msgs[0].Parent,
			},
			response: &types.QueryGetSecondLevelDomainResponse{SecondLevelDomain: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetSecondLevelDomainRequest{
				Name:   msgs[1].Name,
				Parent: msgs[1].Parent,
			},
			response: &types.QueryGetSecondLevelDomainResponse{SecondLevelDomain: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetSecondLevelDomainRequest{
				Name:   strconv.Itoa(100000),
				Parent: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.SecondLevelDomain(wctx, tc.request)
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

func TestSecondLevelDomainQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.RegistryKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNSecondLevelDomainResponse(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllSecondLevelDomainRequest {
		return &types.QueryAllSecondLevelDomainRequest{
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
			resp, err := keeper.SecondLevelDomainAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.SecondLevelDomain), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.SecondLevelDomain),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.SecondLevelDomainAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.SecondLevelDomain), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.SecondLevelDomain),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.SecondLevelDomainAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.SecondLevelDomain),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.SecondLevelDomainAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
