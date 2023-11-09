package types

import (
	"testing"
	"time"

	errorsmod "cosmossdk.io/errors"
	"github.com/stretchr/testify/require"
)

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
			expErr: errorsmod.Wrapf(ErrInvalidTopLevelDomainName, ".foo"),
		},
		{domain: TopLevelDomain{Name: ""},
			expErr: errorsmod.Wrapf(ErrInvalidTopLevelDomainName, ""),
		},
		{domain: TopLevelDomain{Name: "bar.foo"},
			expErr: errorsmod.Wrapf(ErrInvalidTopLevelDomainName, "bar.foo"),
		},
		{domain: TopLevelDomain{Name: "."},
			expErr: errorsmod.Wrapf(ErrInvalidTopLevelDomainName, "."),
		},
		{domain: TopLevelDomain{Name: "##"},
			expErr: errorsmod.Wrapf(ErrInvalidTopLevelDomainName, "##"),
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

func TestExtendExpirationDate(t *testing.T) {
	testCases := []struct {
		from                   time.Time
		extensionPeriodInYear  uint64
		expectedExpirationDate time.Time
	}{
		{
			from:                   time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
			extensionPeriodInYear:  1,
			expectedExpirationDate: time.Date(2020, 12, 31, 0, 0, 0, 0, time.UTC),
		},
		{
			from:                   time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
			extensionPeriodInYear:  2,
			expectedExpirationDate: time.Date(2021, 12, 31, 0, 0, 0, 0, time.UTC),
		},
	}

	for _, tc := range testCases {
		domain := TopLevelDomain{
			Name: "myc",
		}
		extendExpirationDate := domain.ExtendExpirationDate(tc.from, tc.extensionPeriodInYear)
		require.Equal(t, tc.expectedExpirationDate, domain.ExpirationDate)
		require.Equal(t, tc.expectedExpirationDate, extendExpirationDate)
	}
}
