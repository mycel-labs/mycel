package types

import (
	fmt "fmt"
	"strings"
)

const (
	BaseFee = 303
)

func (secondLevelDomain SecondLevelDomain) ParseParent() (parent string) {
	if secondLevelDomain.Parent != "" {
		split := strings.Split(secondLevelDomain.Parent, ".")
		parent = split[len(split)-1]
	}
	return parent
}

func GetWalletAddressFormat(walletRecordType string) (walletAddressFormat string, err error) {
	// Validate wallet record type
	err = ValidateWalletRecordType(walletRecordType)
	if err != nil {
		return "", err
	}
	// Get wallet address format from wallet record type
	walletAddressFormat, isFound := WalletRecordFormats()[walletRecordType]
	if !isFound {
		panic(fmt.Sprintf("Wallet record type %s is not found in WalletRecordFormats", walletRecordType))
	}
	return walletAddressFormat, err
}

func (secondLevelDomain *SecondLevelDomain) UpdateWalletRecord(walletRecordType string, address string) (err error) {

	// Get wallet address format from wallet record type
	walletAddressFormat, err := GetWalletAddressFormat(walletRecordType)
	if err != nil {
		return err
	}

	err = ValidateWalletAddress(walletAddressFormat, address)
	if err != nil {
		return err
	}

	walletRecord := &WalletRecord{
		WalletRecordType: NetworkName(NetworkName_value[walletRecordType]),
		Value:            address,
	}

	// Initialize WalletRecords map if it is nil
	if secondLevelDomain.WalletRecords == nil {
		secondLevelDomain.WalletRecords = make(map[string]*WalletRecord)
	}

	secondLevelDomain.WalletRecords[walletRecordType] = walletRecord

	return err
}

func GetDnsRecordValueFormat(dnsRecordType string) (dnsRecordTypeFormat string, err error) {
	err = ValidateDnsRecordType(dnsRecordType)
	if err != nil {
		return "", err
	}
	dnsRecordTypeFormat, isFound := DnsRecordTypeFormats()[dnsRecordType]
	if !isFound {
		panic(fmt.Sprintf("Dns record type %s is not found in DnsRecordFormats", dnsRecordType))
	}
	return dnsRecordTypeFormat, err
}

func (secondLevelDomain *SecondLevelDomain) UpdateDnsRecord(dnsRecordType string, value string) (err error) {

	// Get wallet address format from dns record type
	dnsRecordFormat, err := GetDnsRecordValueFormat(dnsRecordType)
	if err != nil {
		return err
	}

	err = ValidateDnsRecordValue(dnsRecordFormat, value)
	if err != nil {
		return err
	}

	dnsRecord := &DnsRecord{
		DnsRecordType: DnsRecordType(DnsRecordType_value[dnsRecordType]),
		Value:         value,
	}

	// Initialize WalletRecords map if it is nil
	if secondLevelDomain.DnsRecords == nil {
		secondLevelDomain.DnsRecords = make(map[string]*DnsRecord)
	}

	secondLevelDomain.DnsRecords[dnsRecordType] = dnsRecord

	return err
}
