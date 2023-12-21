package keeper_test

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"

	keepertest "github.com/mycel-domain/mycel/testutil/keeper"
	"github.com/mycel-domain/mycel/testutil/nullify"
	"github.com/mycel-domain/mycel/x/furnace/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestBurnAmountQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.FurnaceKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNBurnAmount(keeper, ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetBurnAmountRequest
		response *types.QueryGetBurnAmountResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetBurnAmountRequest{
				Index: msgs[0].Index,
			},
			response: &types.QueryGetBurnAmountResponse{BurnAmount: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetBurnAmountRequest{
				Index: msgs[1].Index,
			},
			response: &types.QueryGetBurnAmountResponse{BurnAmount: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetBurnAmountRequest{
				Index: 100000,
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
			response, err := keeper.BurnAmount(wctx, tc.request)
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

func TestBurnAmountQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.FurnaceKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNBurnAmount(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllBurnAmountRequest {
		return &types.QueryAllBurnAmountRequest{
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
			resp, err := keeper.BurnAmountAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.BurnAmount), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.BurnAmount),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.BurnAmountAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.BurnAmount), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.BurnAmount),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.BurnAmountAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.BurnAmount),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.BurnAmountAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
