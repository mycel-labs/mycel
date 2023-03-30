package keeper_test

import (
	mycelapp "mycel/app"
	"mycel/x/registry/keeper"
	"mycel/x/registry/types"
	"testing"
	"time"

	"mycel/testutil"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/stretchr/testify/suite"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
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
	suite.app.BankKeeper.InitGenesis(suite.ctx, getBankGenesis())

}

func makeBalance(address string, balance int64) banktypes.Balance {
	return banktypes.Balance{
		Address: address,
		Coins: sdk.Coins{
			sdk.Coin{
				Denom:  types.MycelDenom,
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
		[]banktypes.Metadata{})

	return state
}

func (suite *KeeperTestSuite) RequireBankBalance(expected int, atAddress string) {
	sdkAdd, err := sdk.AccAddressFromBech32(atAddress)
	suite.Require().Nil(err, "Failed to parse address: %s", atAddress)
	suite.Require().Equal(
		int64(expected),
		suite.app.BankKeeper.GetBalance(suite.ctx, sdkAdd, sdk.DefaultBondDenom).Amount.Int64())
}
