package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"mycel/x/mycel/keeper"
	"mycel/x/mycel/types"
)

func SimulateMsgUpdateWalletRecord(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgUpdateWalletRecord{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the UpdateWalletRecord simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "UpdateWalletRecord simulation not implemented"), nil, nil
	}
}
