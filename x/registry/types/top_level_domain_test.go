package types

import (
	"fmt"
	"testing"

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
		expErr string
	}{
		// Valid domains
		{
			domain: TopLevelDomain{Name: "myc"},
		},
		// Invalid name
		{domain: TopLevelDomain{Name: ".foo"},
			expErr: fmt.Sprintf("invalid name: .foo"),
		},
		{domain: TopLevelDomain{Name: ""},
			expErr: fmt.Sprintf("invalid name: "),
		},
		{domain: TopLevelDomain{Name: "bar.foo"},
			expErr: fmt.Sprintf("invalid name: bar.foo"),
		},
		{domain: TopLevelDomain{Name: "."},
			expErr: fmt.Sprintf("invalid name: ."),
		},
		{domain: TopLevelDomain{Name: "##"},
			expErr: fmt.Sprintf("invalid name: ##"),
		},
	}

	for _, tc := range testCases {
		err := tc.domain.Validate()
		if tc.expErr == "" {
			require.Nil(t, err)
		} else {
			require.EqualError(t, err, tc.expErr)
		}
	}
}
