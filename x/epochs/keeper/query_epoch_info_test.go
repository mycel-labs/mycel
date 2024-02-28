package keeper_test

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	sdk "github.com/cosmos/cosmos-sdk/types"
	// "github.com/cosmos/cosmos-sdk/types/query"

	keepertest "github.com/mycel-domain/mycel/testutil/keeper"
	"github.com/mycel-domain/mycel/testutil/nullify"
	"github.com/mycel-domain/mycel/x/epochs/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestEpochInfoQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.EpochsKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNEpochInfo(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryCurrentEpochRequest
		response *types.QueryCurrentEpochResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryCurrentEpochRequest{
				Identifier: msgs[0].Identifier,
			},
			response: &types.QueryCurrentEpochResponse{EpochInfo: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryCurrentEpochRequest{
				Identifier: msgs[1].Identifier,
			},
			response: &types.QueryCurrentEpochResponse{EpochInfo: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryCurrentEpochRequest{
				Identifier: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.CurrentEpoch(wctx, tc.request)
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

//
// func TestEpochInfoQueryPaginated(t *testing.T) {
// keeper, ctx := keepertest.EpochsKeeper(t)
// wctx := sdk.WrapSDKContext(ctx)
// msgs := createNEpochInfo(keeper, ctx, 5)
//
// request := func(next []byte, offset, limit uint64, total bool) *types.QueryEpochsInfoRequest {
// return &types.QueryEpochsInfoRequest{
// Pagination: &query.PageRequest{
// Key:        next,
// Offset:     offset,
// Limit:      limit,
// CountTotal: total,
// },
// }
// }
// t.Run("ByOffset", func(t *testing.T) {
// step := 2
// for i := 0; i < len(msgs); i += step {
// resp, err := keeper.EpochInfos(wctx, request(nil, uint64(i), uint64(step), false))
// require.NoError(t, err)
// require.LessOrEqual(t, len(resp.EpochInfo), step)
// require.Subset(t,
// nullify.Fill(msgs),
// nullify.Fill(resp.EpochInfo),
// )
// }
// })
// t.Run("ByKey", func(t *testing.T) {
// step := 2
// var next []byte
// for i := 0; i < len(msgs); i += step {
// resp, err := keeper.EpochInfos(wctx, request(next, 0, uint64(step), false))
// require.NoError(t, err)
// require.LessOrEqual(t, len(resp.Epochs), step)
// require.Subset(t,
// nullify.Fill(msgs),
// nullify.Fill(resp.Epochs),
// )
// next = resp.Pagination.NextKey
// }
// })
// t.Run("Total", func(t *testing.T) {
// resp, err := keeper.EpochInfos(wctx, request(nil, 0, 0, true))
// require.NoError(t, err)
// require.Equal(t, len(msgs), int(resp.Pagination.Total))
// require.ElementsMatch(t,
// nullify.Fill(msgs),
// nullify.Fill(resp.Epochs),
// )
// })
// t.Run("InvalidRequest", func(t *testing.T) {
// _, err := keeper.EpochInfos(wctx, nil)
// require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
// })
// }
