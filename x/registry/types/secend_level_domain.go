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

func (secondLevelDomain *SecondLevelDomain) GetWalletRecord(walletRecordType string) string {
	for _, rec := range secondLevelDomain.Records {
		if rec.GetWalletRecord() != nil && rec.GetWalletRecord().WalletRecordType.String() == walletRecordType {
			return rec.GetWalletRecord().Value
		}
	}
	return ""
}

func (secondLevelDomain *SecondLevelDomain) GetDnsRecord(dnsRecordType string) string {
	for _, rec := range secondLevelDomain.Records {
		if rec.GetDnsRecord() != nil && rec.GetDnsRecord().DnsRecordType.String() == dnsRecordType {
			return rec.GetDnsRecord().Value
		}
	}
	return ""
}

func (secondLevelDomain *SecondLevelDomain) GetTextRecord(key string) string {
	for _, rec := range secondLevelDomain.Records {
		if rec.GetTextRecord() != nil && rec.GetTextRecord().Key == key {
			return rec.GetTextRecord().Value
		}
	}
	return ""
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

	// Update or add the new record
	updated := false
	for i, rec := range secondLevelDomain.Records {
		if rec.GetWalletRecord() != nil && rec.GetWalletRecord().WalletRecordType.String() == walletRecordType {
			secondLevelDomain.Records[i] = record
			updated = true
			break
		}
	}
	if !updated {
		secondLevelDomain.Records = append(secondLevelDomain.Records, record)
	}

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

	updated := false
	for i, rec := range secondLevelDomain.Records {
		if rec.GetDnsRecord() != nil && rec.GetDnsRecord().DnsRecordType.String() == dnsRecordType {
			secondLevelDomain.Records[i] = record
			updated = true
			break
		}
	}
	if !updated {
		secondLevelDomain.Records = append(secondLevelDomain.Records, record)
	}
	return err
}

func (secondLevelDomain *SecondLevelDomain) UpdateTextRecord(key string, value string) (err error) {
	err = ValidateTextRecordKey(key)
	if err != nil {
		return err
	}
	textRecord := &TextRecord{
		Key:   key,
		Value: value,
	}

	record := &Record{
		Record: &Record_TextRecord{TextRecord: textRecord},
	}

	updated := false
	for i, rec := range secondLevelDomain.Records {
		if rec.GetTextRecord() != nil && rec.GetTextRecord().Key == key {
			secondLevelDomain.Records[i] = record
			updated = true
			break
		}
	}
	if !updated {
		secondLevelDomain.Records = append(secondLevelDomain.Records, record)
	}
	return err
}

func (secondLevelDomain SecondLevelDomain) IsRecordEditable(sender string) (isEditable bool, err error) {
	role := secondLevelDomain.GetRole(sender)
	if role == DomainRole_NO_ROLE {
		err = errorsmod.Wrapf(ErrSecondLevelDomainNotEditable, "%s", sender)
	}
	isEditable = role == DomainRole_EDITOR || role == DomainRole_OWNER
	return isEditable, err
}

func (secondLevelDomain *SecondLevelDomain) GetRole(address string) (role DomainRole) {
	for _, accessControl := range secondLevelDomain.AccessControl {
		if accessControl.Address == address {
			return accessControl.Role
		}
	}
	return DomainRole_NO_ROLE
}
