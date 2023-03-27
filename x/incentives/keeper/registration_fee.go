package keeper

import sdk "github.com/cosmos/cosmos-sdk/types"

func (k Keeper) DivideRegistrationFeeToIncentiveStoreForEachEpoch(ctx sdk.Context, registrationPeriodInWeek uint, amount sdk.Int) {

}

func (k Keeper) SendRegistrationFeeToIncentiveModule(ctx sdk.Context) {}
