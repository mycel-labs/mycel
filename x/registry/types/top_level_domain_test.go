package types

import (
	"testing"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

type TopLevelDomainTest struct {
	Domain      TopLevelDomain
	DomainPrice sdk.Coins
}

func TestTopLevelDomainValidate(t *testing.T) {
	testCases := []struct {
		domain TopLevelDomain
		expErr error
	}{
		// Valid domains
		{
			domain: TopLevelDomain{Name: "myc"},
		},
		// Invalid name
		{domain: TopLevelDomain{Name: ".foo"},
			expErr: errorsmod.Wrapf(ErrInvalidDomainName, ".foo"),
		},
		{domain: TopLevelDomain{Name: ""},
			expErr: errorsmod.Wrapf(ErrInvalidDomainName, ""),
		},
		{domain: TopLevelDomain{Name: "bar.foo"},
			expErr: errorsmod.Wrapf(ErrInvalidDomainName, "bar.foo"),
		},
		{domain: TopLevelDomain{Name: "."},
			expErr: errorsmod.Wrapf(ErrInvalidDomainName, "."),
		},
		{domain: TopLevelDomain{Name: "##"},
			expErr: errorsmod.Wrapf(ErrInvalidDomainName, "##"),
		},
	}

	for _, tc := range testCases {
		err := tc.domain.Validate()
		if tc.expErr == nil {
			require.Nil(t, err)
		} else {
			require.EqualError(t, err, tc.expErr.Error())
		}
	}
}
