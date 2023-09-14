package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	registrytypes "github.com/mycel-domain/mycel/x/registry/types"
)

type RegistryKeeper interface {
	// Methods imported from registry should be defined here
	GetSecondLevelDomain(ctx sdk.Context, name string, parent string) (val registrytypes.SecondLevelDomain, found bool)
	GetTopLevelDomain(ctx sdk.Context, name string) (val registrytypes.TopLevelDomain, found bool)
	GetValidSecondLevelDomain(ctx sdk.Context, name string, parent string) (secondLevelDomain registrytypes.SecondLevelDomain, err error)
	GetValidTopLevelDomain(ctx sdk.Context, name string) (topLevelDomain registrytypes.TopLevelDomain, err error)
}

// AccountKeeper defines the expected account keeper used for simulations (noalias)
type AccountKeeper interface {
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) types.AccountI
	// Methods imported from account should be defined here
}

// BankKeeper defines the expected interface needed to retrieve account balances.
type BankKeeper interface {
	SpendableCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	// Methods imported from bank should be defined here
}
