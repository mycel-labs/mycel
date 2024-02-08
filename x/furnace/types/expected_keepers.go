package types

import (
	context "context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	epochstypes "github.com/mycel-domain/mycel/x/epochs/types"
)

type EpochsKeeper interface {
	// Methods imported from epochs should be defined here
	GetEpochInfo(ctx sdk.Context, identifier string) (val epochstypes.EpochInfo, found bool)
}

// AccountKeeper defines the expected account keeper used for simulations (noalias)
type AccountKeeper interface {
	GetAccount(ctx context.Context, addr sdk.AccAddress) sdk.AccountI
	// Methods imported from account should be defined here
}

// BankKeeper defines the expected interface needed to retrieve account balances.
type BankKeeper interface {
	SpendableCoins(ctx context.Context, addr sdk.AccAddress) sdk.Coins
	// Methods imported from bank should be defined here
}
