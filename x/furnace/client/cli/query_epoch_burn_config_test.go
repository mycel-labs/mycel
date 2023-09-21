package cli_test

import (
	"fmt"
	"testing"

	tmcli "github.com/cometbft/cometbft/libs/cli"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/status"

	"github.com/mycel-domain/mycel/testutil/network"
	"github.com/mycel-domain/mycel/testutil/nullify"
	"github.com/mycel-domain/mycel/x/furnace/client/cli"
	"github.com/mycel-domain/mycel/x/furnace/types"
)

func networkWithEpochBurnConfigObjects(t *testing.T) (*network.Network, types.EpochBurnConfig) {
	t.Helper()
	cfg := network.DefaultConfig()
	state := types.GenesisState{}
	epochBurnConfig := &types.EpochBurnConfig{}
	nullify.Fill(&epochBurnConfig)
	state.EpochBurnConfig = *epochBurnConfig
	buf, err := cfg.Codec.MarshalJSON(&state)
	require.NoError(t, err)
	cfg.GenesisState[types.ModuleName] = buf
	return network.New(t, cfg), state.EpochBurnConfig
}

func TestShowEpochBurnConfig(t *testing.T) {
	net, obj := networkWithEpochBurnConfigObjects(t)

	ctx := net.Validators[0].ClientCtx
	common := []string{
		fmt.Sprintf("--%s=json", tmcli.OutputFlag),
	}
	tests := []struct {
		desc string
		args []string
		err  error
		obj  types.EpochBurnConfig
	}{
		{
			desc: "get",
			args: common,
			obj:  obj,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			var args []string
			args = append(args, tc.args...)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdShowEpochBurnConfig(), args)
			if tc.err != nil {
				stat, ok := status.FromError(tc.err)
				require.True(t, ok)
				require.ErrorIs(t, stat.Err(), tc.err)
			} else {
				require.NoError(t, err)
				var resp types.QueryGetEpochBurnConfigResponse
				require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
				require.NotNil(t, resp.EpochBurnConfig)
				require.Equal(t,
					nullify.Fill(&tc.obj),
					nullify.Fill(&resp.EpochBurnConfig),
				)
			}
		})
	}
}
