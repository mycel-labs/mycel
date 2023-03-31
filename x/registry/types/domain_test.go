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
			expErr:          "",
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
			expErr:          "",
		},
		{
			domain:          Domain{Name: "123", Parent: "foo.myc"},
			expDomainLevel:  3,
			expDomainParent: Domain{Name: "foo", Parent: "myc"},
			expDomainPrice:  sdk.NewCoin("MYCEL", sdk.NewInt(30_300)),
			expErr:          "",
		},
		{
			domain:          Domain{Name: "12", Parent: "foo.myc"},
			expDomainLevel:  3,
			expDomainParent: Domain{Name: "foo", Parent: "myc"},
			expDomainPrice:  sdk.NewCoin("MYCEL", sdk.NewInt(303_000)),
			expErr:          "",
		},
		{
			domain:          Domain{Name: "🍭", Parent: "foo.🍭"},
			expDomainLevel:  3,
			expDomainParent: Domain{Name: "foo", Parent: "🍭"},
			expDomainPrice:  sdk.NewCoin("MYCEL", sdk.NewInt(3_030_000)),
			expErr:          "",
		},
		{
			domain:          Domain{Name: "🍭", Parent: "foo.🍭.myc"},
			expDomainLevel:  4,
			expDomainParent: Domain{Name: "foo.🍭", Parent: "myc"},
			expDomainPrice:  sdk.NewCoin("MYCEL", sdk.NewInt(3_030_000)),
			expErr:          "",
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
		err := tc.domain.ValidateDomain()
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

func GetValidDomains() []DomainTest {
	return []DomainTest{
		{Domain: Domain{Name: "foo", Parent: "myc"},
			DomainLevel:  2,
			DomainParent: Domain{Name: "myc", Parent: ""},
			DomainPrice:  sdk.NewCoins(sdk.NewCoin("MYCEL", sdk.NewInt(30300)))},
		{Domain: Domain{Name: "12345", Parent: ""},
			DomainLevel:  1,
			DomainParent: Domain{Name: "", Parent: ""},
			DomainPrice:  sdk.NewCoins(sdk.NewCoin("MYCEL", sdk.NewInt(303)))},
		{Domain: Domain{Name: "1234", Parent: "foo.myc"},
			DomainLevel:  3,
			DomainParent: Domain{Name: "foo", Parent: "myc"},
			DomainPrice:  sdk.NewCoins(sdk.NewCoin("MYCEL", sdk.NewInt(3030)))},
		{Domain: Domain{Name: "123", Parent: "foo.myc"},
			DomainLevel:  3,
			DomainParent: Domain{Name: "foo", Parent: "myc"},
			DomainPrice:  sdk.NewCoins(sdk.NewCoin("MYCEL", sdk.NewInt(30300)))},
		{Domain: Domain{Name: "12", Parent: "foo.myc"},
			DomainLevel:  3,
			DomainParent: Domain{Name: "foo", Parent: "myc"},
			DomainPrice:  sdk.NewCoins(sdk.NewCoin("MYCEL", sdk.NewInt(303000)))},
		{Domain: Domain{Name: "🍭", Parent: "foo.🍭"},
			DomainLevel:  3,
			DomainParent: Domain{Name: "foo", Parent: "🍭"},
			DomainPrice:  sdk.NewCoins(sdk.NewCoin("MYCEL", sdk.NewInt(3030000)))},
		{Domain: Domain{Name: "🍭", Parent: "foo.🍭.myc"},
			DomainLevel:  4,
			DomainParent: Domain{Name: "foo.🍭", Parent: "myc"},
			DomainPrice:  sdk.NewCoins(sdk.NewCoin("MYCEL", sdk.NewInt(3030000)))},
	}
}

// Name is invalid
func GetInvalidNameDomains() []DomainTest {
	return []DomainTest{
		{Domain: Domain{Name: ".foo", Parent: "myc"}},
		{Domain: Domain{Name: "", Parent: "myc"}},
		{Domain: Domain{Name: "bar.foo", Parent: "myc"}},
		{Domain: Domain{Name: ".", Parent: "myc"}},
		{Domain: Domain{Name: "##", Parent: "myc"}},
	}
}

// Parent is invalid
func GetInvalidParentDomains() []DomainTest {
	return []DomainTest{
		{Domain: Domain{Name: "foo", Parent: ".##"}},
		{Domain: Domain{Name: "foo", Parent: ".myc"}},
		{Domain: Domain{Name: "foo", Parent: ".foo.myc"}},
	}
}

func GetValidUpdateWalletRecords() []WalletRecordTest {
	return []WalletRecordTest{
		{WalletRecordType: "ETHEREUM_MAINNET", Address: "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"},
		{WalletRecordType: "ETHEREUM_GOERLI", Address: "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"},
		{WalletRecordType: "POLYGON_MAINNET", Address: "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"},
		{WalletRecordType: "POLYGON_MUMBAI", Address: "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"},
	}
}
func GetInvalidUpdateWalletRecords() []WalletRecordTest {
	return []WalletRecordTest{
		{WalletRecordType: "ETHEREUM_MUMBAI", Address: "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266", IsInvalidRecordType: true},
		{WalletRecordType: "ETHEREUM_GOERLI", Address: "0xf9Fd6e51aad88F6F4ce6aB8827279cffFb92266", IsInvalidValue: true},
		{WalletRecordType: "ETHEREUM_GOERLI", Address: "cosmos1jyc4rrtz5f93n80uuj378dq7x3v7z09j0h6dqx", IsInvalidValue: true},
	}
}

func TestValidateWalletAddressSuccess(t *testing.T) {
	for _, walletRecord := range GetValidUpdateWalletRecords() {
		addressFormat, errAddressFormat := GetWalletAddressFormat(walletRecord.WalletRecordType)
		require.Nil(t, errAddressFormat)
		err := ValidateWalletAddress(addressFormat, walletRecord.Address)
		require.Nil(t, err)
	}
}

func TestValidateWalletAddressFailure(t *testing.T) {
	for _, walletRecord := range GetInvalidUpdateWalletRecords() {
		addressFormat, errAddressFormat := GetWalletAddressFormat(walletRecord.WalletRecordType)
		if walletRecord.IsInvalidRecordType {
			require.EqualError(t, errAddressFormat, fmt.Sprintf("invalid wallet record type: %s", walletRecord.WalletRecordType))
		} else if walletRecord.IsInvalidValue {
			err := ValidateWalletAddress(addressFormat, walletRecord.Address)
			require.EqualError(t, err, fmt.Sprintf("invalid wallet address: %s %s", addressFormat, walletRecord.Address))
		}
	}
}

func TestValidateWalletRecordTypeSuccess(t *testing.T) {
	for _, walletRecord := range GetValidUpdateWalletRecords() {
		err := ValidateWalletRecordType(walletRecord.WalletRecordType)
		require.Nil(t, err)
	}
}

func TestValidateWalletRecordTypeFailure(t *testing.T) {
	for _, walletRecord := range GetInvalidUpdateWalletRecords() {
		err := ValidateWalletRecordType(walletRecord.WalletRecordType)
		if walletRecord.IsInvalidRecordType {
			require.EqualError(t, err, fmt.Sprintf("invalid wallet record type: %s", walletRecord.WalletRecordType))
		}
	}
}

func TestUpdateWalletRecordSuccess(t *testing.T) {
	domains := GetValidDomains()
	for _, walletRecord := range GetValidUpdateWalletRecords() {
		err := domains[0].Domain.UpdateWalletRecord(walletRecord.WalletRecordType, walletRecord.Address)
		require.Equal(t, walletRecord.Address, domains[0].Domain.WalletRecords[walletRecord.WalletRecordType].Value)
		require.Nil(t, err)
	}
}

func TestUpdateWalletRecordFailure(t *testing.T) {
	domains := GetValidDomains()
	for _, walletRecord := range GetInvalidUpdateWalletRecords() {
		err := domains[0].Domain.UpdateWalletRecord(walletRecord.WalletRecordType, walletRecord.Address)
		if walletRecord.IsInvalidRecordType {
			require.EqualError(t, err, fmt.Sprintf("invalid wallet record type: %s", walletRecord.WalletRecordType))
		} else if walletRecord.IsInvalidValue {
			addressFormat, _ := GetWalletAddressFormat(walletRecord.WalletRecordType)
			require.EqualError(t, err, fmt.Sprintf("invalid wallet address: %s %s", addressFormat, walletRecord.Address))
		}
	}
}
