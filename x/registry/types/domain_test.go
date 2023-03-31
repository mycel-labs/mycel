package types

import (
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

type DomainTest struct {
	Domain       Domain
	DomainLevel  int
	DomainParent Domain
	DomainPrice  sdk.Coins
}

type WalletRecordTest struct {
	WalletRecordType    string
	Address             string
	IsInvalidRecordType bool
	IsInvalidValue      bool
}

func TestDomainValidate(t *testing.T) {
	testCases := []struct {
		domain          Domain
		expDomainLevel  int
		expDomainParent Domain
		expDomainPrice  sdk.Coin
		expErr          string
	}{
		// Valid domains
		{
			domain:          Domain{Name: "foo", Parent: "myc"},
			expDomainLevel:  2,
			expDomainParent: Domain{Name: "myc", Parent: ""},
			expDomainPrice:  sdk.NewCoin("MYCEL", sdk.NewInt(30_300)),
		},
		{
			domain:          Domain{Name: "12345", Parent: ""},
			expDomainLevel:  1,
			expDomainParent: Domain{Name: "", Parent: ""},
			expDomainPrice:  sdk.NewCoin("MYCEL", sdk.NewInt(303)),
			expErr:          "",
		},
		{
			domain:          Domain{Name: "1234", Parent: "foo.myc"},
			expDomainLevel:  3,
			expDomainParent: Domain{Name: "foo", Parent: "myc"},
			expDomainPrice:  sdk.NewCoin("MYCEL", sdk.NewInt(3_030)),
		},
		{
			domain:          Domain{Name: "123", Parent: "foo.myc"},
			expDomainLevel:  3,
			expDomainParent: Domain{Name: "foo", Parent: "myc"},
			expDomainPrice:  sdk.NewCoin("MYCEL", sdk.NewInt(30_300)),
		},
		{
			domain:          Domain{Name: "12", Parent: "foo.myc"},
			expDomainLevel:  3,
			expDomainParent: Domain{Name: "foo", Parent: "myc"},
			expDomainPrice:  sdk.NewCoin("MYCEL", sdk.NewInt(303_000)),
		},
		{
			domain:          Domain{Name: "🍭", Parent: "foo.🍭"},
			expDomainLevel:  3,
			expDomainParent: Domain{Name: "foo", Parent: "🍭"},
			expDomainPrice:  sdk.NewCoin("MYCEL", sdk.NewInt(3_030_000)),
		},
		{
			domain:          Domain{Name: "🍭", Parent: "foo.🍭.myc"},
			expDomainLevel:  4,
			expDomainParent: Domain{Name: "foo.🍭", Parent: "myc"},
			expDomainPrice:  sdk.NewCoin("MYCEL", sdk.NewInt(3_030_000)),
		},
		// Invalid name
		{domain: Domain{Name: ".foo", Parent: "myc"},
			expErr: fmt.Sprintf("invalid name: .foo"),
		},
		{domain: Domain{Name: "", Parent: "myc"},
			expErr: fmt.Sprintf("invalid name: "),
		},
		{domain: Domain{Name: "bar.foo", Parent: "myc"},
			expErr: fmt.Sprintf("invalid name: bar.foo"),
		},
		{domain: Domain{Name: ".", Parent: "myc"},
			expErr: fmt.Sprintf("invalid name: ."),
		},
		{domain: Domain{Name: "##", Parent: "myc"},
			expErr: fmt.Sprintf("invalid name: ##"),
		},
		// Invalid parent
		{
			domain: Domain{Name: "foo", Parent: ".##"},
			expErr: fmt.Sprintf("invalid parent: .##"),
		},
		{
			domain: Domain{Name: "foo", Parent: ".myc"},
			expErr: fmt.Sprintf("invalid parent: .myc"),
		},
		{
			domain: Domain{Name: "foo", Parent: ".foo.myc"},
			expErr: fmt.Sprintf("invalid parent: .foo.myc"),
		},
	}

	for _, tc := range testCases {
		err := tc.domain.Validate()
		if tc.expErr == "" {
			require.Nil(t, err)
			// Check domain level
			require.Equal(t, tc.expDomainLevel, tc.domain.GetDomainLevel())

			// Check domain parent
			name, parent := tc.domain.ParseParent()
			require.Equal(t, tc.expDomainParent.Name, name)
			require.Equal(t, tc.expDomainParent.Parent, parent)

			// Check domain price
			require.Equal(t, tc.expDomainPrice, tc.domain.GetRegistrationFee())

		} else {
			require.EqualError(t, err, tc.expErr)
		}

	}

}

func TestDomainUpdateRecord(t *testing.T) {
	testCases := []struct {
		walletRecordType string
		address          string
		expErr           string
	}{
		// Valid wallet records
		{walletRecordType: "ETHEREUM_MAINNET", address: "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"},
		{walletRecordType: "ETHEREUM_GOERLI", address: "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"},
		{walletRecordType: "POLYGON_MAINNET", address: "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"},
		{walletRecordType: "POLYGON_MUMBAI", address: "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"},

		// Invalid record type
		{
			walletRecordType: "ETHEREUM_MUMBAI", address: "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266",
			expErr: fmt.Sprintf("invalid wallet record type: ETHEREUM_MUMBAI"),
		},
		{
			walletRecordType: "ETHEREUM", address: "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266",
			expErr: fmt.Sprintf("invalid wallet record type: ETHEREUM"),
		},
		// Invalid address
		{
			walletRecordType: "ETHEREUM_GOERLI", address: "0xf9Fd6e51aad88F6F4ce6aB8827279cffFb92266",
			expErr: fmt.Sprintf("invalid wallet address: ETHEREUM 0xf9Fd6e51aad88F6F4ce6aB8827279cffFb92266"),
		},
		{
			walletRecordType: "ETHEREUM_GOERLI", address: "cosmos1jyc4rrtz5f93n80uuj378dq7x3v7z09j0h6dqx",
			expErr: fmt.Sprintf("invalid wallet address: ETHEREUM cosmos1jyc4rrtz5f93n80uuj378dq7x3v7z09j0h6dqx"),
		},
	}
	for _, tc := range testCases {
		domain := Domain{Name: "foo", Parent: "myc"}
		err := domain.UpdateWalletRecord(tc.walletRecordType, tc.address)
		if tc.expErr == "" {
			require.Nil(t, err)
			require.Equal(t, tc.address, domain.WalletRecords[tc.walletRecordType].Value)
		} else {
			require.EqualError(t, err, tc.expErr)
		}
	}
}
