package types

import (
	fmt "fmt"
	"testing"
	"time"

	errorsmod "cosmossdk.io/errors"
	"github.com/stretchr/testify/require"

	"github.com/mycel-domain/mycel/testutil"
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

func TestGetRoleTLD(t *testing.T) {
	testCases := []struct {
		domain TopLevelDomain
		req    string
		exp    DomainRole
	}{
		// Valid domains
		{
			domain: TopLevelDomain{
				Name:          "myc",
				AccessControl: map[string]DomainRole{testutil.Alice: DomainRole_NO_ROLE},
			},
			req: testutil.Alice,
			exp: DomainRole_NO_ROLE,
		},
		{
			domain: TopLevelDomain{
				Name:          "myc",
				AccessControl: map[string]DomainRole{testutil.Alice: DomainRole_OWNER},
			},
			req: testutil.Alice,
			exp: DomainRole_OWNER,
		},
		{
			domain: TopLevelDomain{
				Name:          "myc",
				AccessControl: map[string]DomainRole{testutil.Alice: DomainRole_EDITOR},
			},
			req: testutil.Alice,
			exp: DomainRole_EDITOR,
		},
		{
			domain: TopLevelDomain{
				Name:          "myc",
				AccessControl: map[string]DomainRole{testutil.Alice: DomainRole_OWNER},
			},
			req: testutil.Bob,
			exp: DomainRole_NO_ROLE,
		},
		{
			domain: TopLevelDomain{
				Name: "myc",
			},
			req: testutil.Alice,
			exp: DomainRole_NO_ROLE,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("Case %d", i), func(t *testing.T) {
			r := tc.domain.GetRole(tc.req)
			require.Equal(t, tc.exp, r)
		})
	}
}
