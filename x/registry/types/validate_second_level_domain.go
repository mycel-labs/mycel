package types

import (
	fmt "fmt"
	"regexp"

	errorsmod "cosmossdk.io/errors"
)

const (
	NamePattern = `-a-z0-9\p{So}\p{Sk}`
)

func ValidateSecondLevelDomainName(name string) (err error) {
	regex := regexp.MustCompile(fmt.Sprintf(`(^[%s]+$)`, NamePattern))
	if !regex.MatchString(name) {
		err = errorsmod.Wrapf(ErrInvalidDomainName, "%s", name)
	}
	return err
}

func (secondLevelDomain SecondLevelDomain) ValidateName() (err error) {
	err = ValidateSecondLevelDomainName(secondLevelDomain.Name)
	return err
}

func ValidateSecondLevelDomainParent(parent string) (err error) {
	regex := regexp.MustCompile(fmt.Sprintf(`(^[%s]+[%[1]s\.]*[%[1]s]$)|^$`, NamePattern))
	if !regex.MatchString(parent) {
		err = errorsmod.Wrapf(ErrInvalidDomainParent, "%s", parent)
	}
	return err

}

func (secondLevelDomain SecondLevelDomain) ValidateParent() (err error) {
	err = ValidateSecondLevelDomainParent(secondLevelDomain.Parent)
	return err
}

func (secondLevelDomain SecondLevelDomain) Validate() (err error) {
	err = secondLevelDomain.ValidateName()
	if err != nil {
		return err
	}
	err = secondLevelDomain.ValidateParent()
	if err != nil {
		return err
	}
	return err
}

func ValidateWalletRecordType(walletRecordType string) (err error) {
	_, isFound := NetworkName_value[walletRecordType]
	if !isFound {
		err = errorsmod.Wrapf(ErrInvalidWalletRecordType, "%s", walletRecordType)
	}
	return err
}

func ValidateDnsRecordValue(dnsRecordFormat string, address string) (err error) {
	dnsRecordRegex, isFound := DnsRecordValueRegex()[dnsRecordFormat]
	if !isFound {
		panic(fmt.Sprintf("Dns record value format %s is not found in DnsRecordValueRegex", dnsRecordFormat))
	}
	regex := regexp.MustCompile(dnsRecordRegex)
	if !regex.MatchString(address) {
		err = errorsmod.Wrapf(ErrInvalidDnsRecordValue, "%s %s", dnsRecordFormat, address)
	}
	return err
}

func ValidateDnsRecordType(dnsRecordType string) (err error) {
	_, isFound := DnsRecordType_value[dnsRecordType]
	if !isFound {
		err = errorsmod.Wrapf(ErrInvalidDnsRecordType, "%s", dnsRecordType)
	}
	return err
}
