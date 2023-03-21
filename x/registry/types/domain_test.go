package types

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

type DomainTest struct {
	Domain       Domain
	DomainLevel  int
	DomainParent Domain
}

type WalletRecordTest struct {
	WalletRecordType    string
	Address             string
	IsInvalidRecordType bool
	IsInvalidValue      bool
}

func GetValidDomains() []DomainTest {
	return []DomainTest{
		{Domain: Domain{Name: "foo", Parent: "myc"}, DomainLevel: 2, DomainParent: Domain{Name: "myc", Parent: ""}},
		{Domain: Domain{Name: "foo", Parent: ""}, DomainLevel: 1, DomainParent: Domain{Name: "", Parent: ""}},
		{Domain: Domain{Name: "bar", Parent: "foo.myc"}, DomainLevel: 3, DomainParent: Domain{Name: "foo", Parent: "myc"}},
		{Domain: Domain{Name: "🍭", Parent: "foo.🍭"}, DomainLevel: 3, DomainParent: Domain{Name: "foo", Parent: "🍭"}},
		{Domain: Domain{Name: "🍭", Parent: "foo.🍭.myc"}, DomainLevel: 4, DomainParent: Domain{Name: "foo.🍭", Parent: "myc"}},
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

func TestValidateDomainNameSuccess(t *testing.T) {
	for _, v := range GetValidDomains() {
		err := v.Domain.ValidateDomain()
		require.Nil(t, err)
	}
}
func TestValidateDomainNameFailure(t *testing.T) {
	for _, v := range GetInvalidNameDomains() {
		err := v.Domain.ValidateDomainName()
		require.EqualError(t, err, fmt.Sprintf("invalid name: %s", v.Domain.Name))
	}
}

func TestValidateDomainParentSuccess(t *testing.T) {
	for _, v := range GetValidDomains() {
		err := v.Domain.ValidateDomainParent()
		require.Nil(t, err)
	}
}

func TestValidateDomainParentFailure(t *testing.T) {
	for _, v := range GetInvalidParentDomains() {
		err := v.Domain.ValidateDomainParent()
		require.EqualError(t, err, fmt.Sprintf("invalid parent: %s", v.Domain.Parent))
	}
}

func TestValidateDomainSuccess(t *testing.T) {
	for _, v := range GetValidDomains() {
		err := v.Domain.ValidateDomain()
		require.Nil(t, err)
	}
}

func TestValidateDomainFailure(t *testing.T) {
	for _, v := range GetInvalidNameDomains() {
		err := v.Domain.ValidateDomain()
		require.EqualError(t, err, fmt.Sprintf("invalid name: %s", v.Domain.Name))
	}
	for _, v := range GetInvalidParentDomains() {
		err := v.Domain.ValidateDomainParent()
		require.EqualError(t, err, fmt.Sprintf("invalid parent: %s", v.Domain.Parent))
	}
}

func TestGetDomainLevel(t *testing.T) {
	for _, v := range GetValidDomains() {
		domainLevel := v.Domain.GetDomainLevel()
		require.Equal(t, domainLevel, v.DomainLevel)
	}
}

func TestGetDomainParent(t *testing.T) {
	for _, v := range GetValidDomains() {
		name, parent := v.Domain.ParseParent()
		require.Equal(t, v.DomainParent.Name, name)
		require.Equal(t, v.DomainParent.Parent, parent)
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
