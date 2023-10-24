package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	furnacetypes "github.com/mycel-domain/mycel/x/furnace/types"
)

// AccountKeeper defines the expected account keeper used for simulations (noalias)
type AccountKeeper interface {
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) types.AccountI
	// Methods imported from account should be defined here
	GetModuleAddress(name string) sdk.AccAddress
	GetModuleAccount(ctx sdk.Context, moduleName string) types.ModuleAccountI
}

// BankKeeper defines the expected interface needed to retrieve account balances.
type BankKeeper interface {
	SpendableCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	// Methods imported from bank should be defined here
	SendCoins(ctx sdk.Context, fromAddr sdk.AccAddress, toAddr sdk.AccAddress, amt sdk.Coins) error
	SendCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
	SendCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
	GetSupply(ctx sdk.Context, denom string) sdk.Coin
	GetBalance(ctx sdk.Context, addr sdk.AccAddress, denom string) sdk.Coin
}

type FurnaceKeeper interface {
	NewBurnAmount(ctx sdk.Context, config furnacetypes.EpochBurnConfig, index uint64) (burnAmount furnacetypes.BurnAmount)
	AddToTotalBurnAmount(ctx sdk.Context, index uint64, amount sdk.Coin) (newBurnAmount furnacetypes.BurnAmount)
	GetEpochBurnConfig(ctx sdk.Context) (val furnacetypes.EpochBurnConfig, found bool)
}
