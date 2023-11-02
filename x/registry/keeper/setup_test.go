package keeper_test

import (
	mycelapp "github.com/mycel-domain/mycel/app"
	"github.com/mycel-domain/mycel/x/registry/keeper"
	"github.com/mycel-domain/mycel/x/registry/types"
	"testing"
	"time"

	"github.com/mycel-domain/mycel/testutil"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/mycel-domain/mycel/app/params"
	"github.com/stretchr/testify/suite"
)

type KeeperTestSuite struct {
	suite.Suite

	ctx         sdk.Context
	app         *mycelapp.App
	msgServer   types.MsgServer
	queryClient types.QueryClient
	consAddress sdk.ConsAddress
}

var s *KeeperTestSuite

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

func (suite *KeeperTestSuite) SetupTest() {
	// init app
	app := mycelapp.Setup(suite.T(), false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{Time: time.Now().UTC()})

	suite.app = app
	suite.ctx = ctx

	queryHelper := baseapp.NewQueryServerTestHelper(suite.ctx, suite.app.InterfaceRegistry())
	types.RegisterQueryServer(queryHelper, suite.app.RegistryKeeper)
	suite.queryClient = types.NewQueryClient(queryHelper)

	suite.msgServer = keeper.NewMsgServerImpl(suite.app.RegistryKeeper)

	// Init bank keeper
	suite.app.BankKeeper.InitGenesis(suite.ctx, getBankGenesis())

}

func makeBalance(address string, balance int64) banktypes.Balance {
	return banktypes.Balance{
		Address: address,
		Coins: sdk.Coins{
			sdk.Coin{
				Denom:  params.DefaultBondDenom,
				Amount: sdk.NewInt(balance),
			},
		},
	}
}

func getBankGenesis() *banktypes.GenesisState {
	coins := []banktypes.Balance{
		makeBalance(testutil.Alice, testutil.BalAlice),
		makeBalance(testutil.Bob, testutil.BalBob),
		makeBalance(testutil.Carol, testutil.BalCarol),
	}
	supply := banktypes.Supply{
		Total: coins[0].Coins.Add(coins[1].Coins...).Add(coins[2].Coins...),
	}

	state := banktypes.NewGenesisState(
		banktypes.DefaultParams(),
		coins,
		supply.Total,
		[]banktypes.Metadata{},
		[]banktypes.SendEnabled{},
	)

	return state
}

func (suite *KeeperTestSuite) RequireBankBalance(expected int, atAddress string) {
	sdkAdd, err := sdk.AccAddressFromBech32(atAddress)
	suite.Require().Nil(err, "Failed to parse address: %s", atAddress)
	suite.Require().Equal(
		int64(expected),
		suite.app.BankKeeper.GetBalance(suite.ctx, sdkAdd, sdk.DefaultBondDenom).Amount.Int64())
}
