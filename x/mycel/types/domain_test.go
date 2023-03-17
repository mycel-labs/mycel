package types

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

type DomainTest struct {
	Domain      Domain
	DomainLevel int
}

type WalletRecordTest struct {
	walletRecordType    string
	address             string
	isInvalidRecordType bool
	isInvalidValue      bool
}

func GetValidDomains() []DomainTest {
	return []DomainTest{
		{Domain: Domain{Name: "foo", Parent: "myc"}, DomainLevel: 2},
		{Domain: Domain{Name: "foo", Parent: ""}, DomainLevel: 1},
		{Domain: Domain{Name: "bar", Parent: "foo.myc"}, DomainLevel: 3},
		{Domain: Domain{Name: "🍭", Parent: "foo.🍭"}, DomainLevel: 3},
		{Domain: Domain{Name: "🍭", Parent: "foo.🍭.myc"}, DomainLevel: 4},
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
		{walletRecordType: "ETHEREUM_MAINNET", address: "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"},
		{walletRecordType: "ETHEREUM_GOERLI", address: "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"},
		{walletRecordType: "POLYGON_MAINNET", address: "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"},
		{walletRecordType: "POLYGON_MUMBAI", address: "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"},
	}
}
func GetInvalidUpdateWalletRecords() []WalletRecordTest {
	return []WalletRecordTest{
		{walletRecordType: "ETHEREUM_MUMBAI", address: "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266", isInvalidRecordType: true},
		{walletRecordType: "ETHEREUM_GOERLI", address: "0xf9Fd6e51aad88F6F4ce6aB8827279cffFb92266", isInvalidValue: true},
		{walletRecordType: "ETHEREUM_GOERLI", address: "cosmos1jyc4rrtz5f93n80uuj378dq7x3v7z09j0h6dqx", isInvalidValue: true},
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

func TestValidateWalletAddressSuccess(t *testing.T) {
	for _, walletRecord := range GetValidUpdateWalletRecords() {
		addressFormat, errAddressFormat := GetWalletAddressFormat(walletRecord.walletRecordType)
		require.Nil(t, errAddressFormat)
		err := ValidateWalletAddress(addressFormat, walletRecord.address)
		require.Nil(t, err)
	}
}

func TestValidateWalletAddressFailure(t *testing.T) {
	for _, walletRecord := range GetInvalidUpdateWalletRecords() {
		addressFormat, errAddressFormat := GetWalletAddressFormat(walletRecord.walletRecordType)
		if walletRecord.isInvalidRecordType {
			require.EqualError(t, errAddressFormat, fmt.Sprintf("invalid wallet record type: %s", walletRecord.walletRecordType))
		} else if walletRecord.isInvalidValue {
			err := ValidateWalletAddress(addressFormat, walletRecord.address)
			require.EqualError(t, err, fmt.Sprintf("invalid wallet address: %s %s", addressFormat, walletRecord.address))
		}
	}
}

func TestValidateWalletRecordTypeSuccess(t *testing.T) {
	for _, walletRecord := range GetValidUpdateWalletRecords() {
		err := ValidateWalletRecordType(walletRecord.walletRecordType)
		require.Nil(t, err)
	}
}

func TestValidateWalletRecordTypeFailure(t *testing.T) {
	for _, walletRecord := range GetInvalidUpdateWalletRecords() {
		err := ValidateWalletRecordType(walletRecord.walletRecordType)
		if walletRecord.isInvalidRecordType {
			require.EqualError(t, err, fmt.Sprintf("invalid wallet record type: %s", walletRecord.walletRecordType))
		}
	}
}

func TestUpdateWalletRecordSuccess(t *testing.T) {
	domains := GetValidDomains()
	for _, walletRecord := range GetValidUpdateWalletRecords() {
		err := domains[0].Domain.updateWalletRecord(walletRecord.walletRecordType, walletRecord.address)
		require.Nil(t, err)
	}
}

func TestUpdateWalletRecordFailure(t *testing.T) {
	domains := GetValidDomains()
	for _, walletRecord := range GetInvalidUpdateWalletRecords() {
		err := domains[0].Domain.updateWalletRecord(walletRecord.walletRecordType, walletRecord.address)
		if walletRecord.isInvalidRecordType {
			require.EqualError(t, err, fmt.Sprintf("invalid wallet record type: %s", walletRecord.walletRecordType))
		} else if walletRecord.isInvalidValue {
			addressFormat, _ := GetWalletAddressFormat(walletRecord.walletRecordType)
			require.EqualError(t, err, fmt.Sprintf("invalid wallet address: %s %s", addressFormat, walletRecord.address))
		}
	}
}
