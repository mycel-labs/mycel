package cli_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	tmcli "github.com/cometbft/cometbft/libs/cli"

	"github.com/cosmos/cosmos-sdk/client/flags"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"

	"github.com/mycel-domain/mycel/testutil/network"
	"github.com/mycel-domain/mycel/testutil/nullify"
	"github.com/mycel-domain/mycel/x/registry/client/cli"
	"github.com/mycel-domain/mycel/x/registry/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func networkWithDomainOwnershipObjects(t *testing.T, n int) (*network.Network, []types.DomainOwnership) {
	t.Helper()
	cfg := network.DefaultConfig()
	state := types.GenesisState{}
	require.NoError(t, cfg.Codec.UnmarshalJSON(cfg.GenesisState[types.ModuleName], &state))

	for i := 0; i < n; i++ {
		domainOwnership := types.DomainOwnership{
			Owner: strconv.Itoa(i),
		}
		nullify.Fill(&domainOwnership)
		state.DomainOwnerships = append(state.DomainOwnerships, domainOwnership)
	}
	buf, err := cfg.Codec.MarshalJSON(&state)
	require.NoError(t, err)
	cfg.GenesisState[types.ModuleName] = buf
	return network.New(t, cfg), state.DomainOwnerships
}

func TestShowDomainOwnership(t *testing.T) {
	net, objs := networkWithDomainOwnershipObjects(t, 2)

	ctx := net.Validators[0].ClientCtx
	common := []string{
		fmt.Sprintf("--%s=json", tmcli.OutputFlag),
	}
	tests := []struct {
		desc    string
		idOwner string

		args []string
		err  error
		obj  types.DomainOwnership
	}{
		{
			desc:    "found",
			idOwner: objs[0].Owner,

			args: common,
			obj:  objs[0],
		},
		{
			desc:    "not found",
			idOwner: strconv.Itoa(100000),

			args: common,
			err:  status.Error(codes.NotFound, "not found"),
		},
	}
	for i := range tests {
		tc := tests[i]
		t.Run(tc.desc, func(t *testing.T) {
			args := []string{
				tc.idOwner,
			}
			args = append(args, tc.args...)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdShowDomainOwnership(), args)
			if tc.err != nil {
				stat, ok := status.FromError(tc.err)
				require.True(t, ok)
				require.ErrorIs(t, stat.Err(), tc.err)
			} else {
				require.NoError(t, err)
				var resp types.QueryGetDomainOwnershipResponse
				require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
				require.NotNil(t, resp.DomainOwnership)
				require.Equal(t,
					nullify.Fill(&tc.obj),
					nullify.Fill(&resp.DomainOwnership),
				)
			}
		})
	}
}

func TestListDomainOwnership(t *testing.T) {
	net, objs := networkWithDomainOwnershipObjects(t, 5)

	ctx := net.Validators[0].ClientCtx
	request := func(next []byte, offset, limit uint64, total bool) []string {
		args := []string{
			fmt.Sprintf("--%s=json", tmcli.OutputFlag),
		}
		if next == nil {
			args = append(args, fmt.Sprintf("--%s=%d", flags.FlagOffset, offset))
		} else {
			args = append(args, fmt.Sprintf("--%s=%s", flags.FlagPageKey, next))
		}
		args = append(args, fmt.Sprintf("--%s=%d", flags.FlagLimit, limit))
		if total {
			args = append(args, fmt.Sprintf("--%s", flags.FlagCountTotal))
		}
		return args
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(objs); i += step {
			args := request(nil, uint64(i), uint64(step), false)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListDomainOwnership(), args)
			require.NoError(t, err)
			var resp types.QueryAllDomainOwnershipResponse
			require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
			require.LessOrEqual(t, len(resp.DomainOwnership), step)
			require.Subset(t,
				nullify.Fill(objs),
				nullify.Fill(resp.DomainOwnership),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(objs); i += step {
			args := request(next, 0, uint64(step), false)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListDomainOwnership(), args)
			require.NoError(t, err)
			var resp types.QueryAllDomainOwnershipResponse
			require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
			require.LessOrEqual(t, len(resp.DomainOwnership), step)
			require.Subset(t,
				nullify.Fill(objs),
				nullify.Fill(resp.DomainOwnership),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		args := request(nil, 0, uint64(len(objs)), true)
		out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListDomainOwnership(), args)
		require.NoError(t, err)
		var resp types.QueryAllDomainOwnershipResponse
		require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
		require.NoError(t, err)
		require.Equal(t, len(objs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(objs),
			nullify.Fill(resp.DomainOwnership),
		)
	})
}
