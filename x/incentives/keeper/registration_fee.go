package keeper

import (
	epochstypes "mycel/x/epochs/types"
	"mycel/x/incentives/types"
	registrytypes "mycel/x/registry/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) SetIncentivesOnRegistration(ctx sdk.Context, registrationPeriodInWeek uint, amount sdk.Int) {
	// Get current epoch
	epoch, found := k.epochsKeeper.GetEpochInfo(ctx, epochstypes.WeeklyEpochId)
	if !found {
		panic("current epoch not found")
	}

	amountByEpoch := amount.QuoRaw(int64(registrationPeriodInWeek))
	nextEpoch := epoch.CurrentEpoch + 1

	// Set incentives store
	for i := nextEpoch; i <= nextEpoch+int64(registrationPeriodInWeek); i++ {
		incentive := types.Incentive{
			Epoch:  i,
			Amount: sdk.NewCoins(sdk.NewCoin(registrytypes.MycelDenom, amountByEpoch)),
		}
		k.SetIncentive(ctx, incentive)
	}

}

func (k Keeper) SendRegistrationFeeToIncentiveModule(ctx sdk.Context) {}
