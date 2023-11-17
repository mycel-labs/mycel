package types

import (
	fmt "fmt"
	"strings"

	errorsmod "cosmossdk.io/errors"
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

	record := &Record{
		Record: &Record_WalletRecord{WalletRecord: walletRecord},
	}

	// Initialize WalletRecords map if it is nil
	if secondLevelDomain.Records == nil {
		secondLevelDomain.Records = make(map[string]*Record)
	}

	secondLevelDomain.Records[walletRecordType] = record

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

	record := &Record{
		Record: &Record_DnsRecord{DnsRecord: dnsRecord},
	}

	// Initialize WalletRecords map if it is nil
	if secondLevelDomain.Records == nil {
		secondLevelDomain.Records = make(map[string]*Record)
	}

	secondLevelDomain.Records[dnsRecordType] = record

	return err
}

func (secondLevelDomain SecondLevelDomain) IsRecordEditable(sender string) (isEditable bool, err error) {
	if secondLevelDomain.AccessControl[sender] == DomainRole_NO_ROLE {
		err = errorsmod.Wrapf(ErrSecondLevelDomainNotEditable, "%s", sender)
	}
	isEditable = secondLevelDomain.AccessControl[sender] == DomainRole_EDITOR || secondLevelDomain.AccessControl[sender] == DomainRole_OWNER
	return isEditable, err
}

func (secondLevelDomain *SecondLevelDomain) GetRole(address string) (role DomainRole) {
	role = secondLevelDomain.AccessControl[address]
	return role
}
