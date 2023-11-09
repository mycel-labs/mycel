package keeper_test

import (
	"testing"
	"time"

	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"

	mycelapp "github.com/mycel-domain/mycel/app"
	"github.com/mycel-domain/mycel/x/epochs/types"
)

type KeeperTestSuite struct {
	suite.Suite

	ctx         sdk.Context
	app         *mycelapp.App
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
	types.RegisterQueryServer(queryHelper, suite.app.EpochsKeeper)
	suite.queryClient = types.NewQueryClient(queryHelper)
}
