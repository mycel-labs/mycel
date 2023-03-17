package types

import (
	"errors"
	fmt "fmt"
	"regexp"
	"strings"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	NamePattern = `-a-z0-9\p{So}\p{Sk}`
)

func (domain Domain) ValidateDomainName() (err error) {
	regex := regexp.MustCompile(fmt.Sprintf(`(^[%s]+$)`, NamePattern))
	if !regex.MatchString(domain.Name) {
		err = sdkerrors.Wrapf(errors.New(fmt.Sprintf("%s", domain.Name)), ErrInvalidDomainName.Error())
	}
	return err
}

func (domain Domain) ValidateDomainParent() (err error) {
	regex := regexp.MustCompile(fmt.Sprintf(`(^[%s]+[%[1]s\.]*[%[1]s]$)|^$`, NamePattern))
	if !regex.MatchString(domain.Parent) {
		err = sdkerrors.Wrapf(errors.New(fmt.Sprintf("%s", domain.Parent)), ErrInvalidDomainParent.Error())
	}
	return err
}

func (domain Domain) ValidateDomain() (err error) {
	err = domain.ValidateDomainName()
	if err != nil {
		return err
	}
	err = domain.ValidateDomainParent()
	if err != nil {
		return err
	}
	return err
}

func (domain Domain) GetDomainLevel() (domainLevel int) {
	if domain.Parent == "" {
		domainLevel = 1
	} else {
		domainLevel = len(strings.Split(domain.Parent, ".")) + 1
	}
	return domainLevel
}

func ValidateWalletAddress(walletAddressFormat string, address string) (err error) {
	walletAddressRegex, isFound := WalletAddressRegex()[walletAddressFormat]
	if !isFound {
		panic(fmt.Sprintf("Wallet address format %s is not found in WalletAddressRegex", walletAddressFormat))
	}
	regex := regexp.MustCompile(walletAddressRegex)
	if !regex.MatchString(address) {
		err = sdkerrors.Wrapf(errors.New(fmt.Sprintf("%s %s", walletAddressFormat, address)), ErrInvalidWalletAddress.Error())
	}
	return err
}

func ValidateWalletRecordType(walletRecordType string) (err error) {
	_, isFound := WalletRecordType_value[walletRecordType]
	if !isFound {
		err = sdkerrors.Wrapf(errors.New(fmt.Sprintf("%s", walletRecordType)), ErrInvalidWalletRecordType.Error())
	}
	return err
}

func GetWalletAddressFormat(walletRecordType string) (walletAddressFormat string, err error) {
	err = ValidateWalletRecordType(walletRecordType)
	if err != nil {
		return "", err
	}
	walletAddressFormat, isFound := WalletRecordFormats()[walletRecordType]
	if !isFound {
		panic(fmt.Sprintf("Wallet record type %s is not found in WalletRecordFormats", walletRecordType))
	}
	return walletAddressFormat, err
}

func (domain Domain) updateWalletRecord(walletRecordType string, address string) (err error) {

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
		WalletRecordType:    WalletRecordType(WalletRecordType_value[walletRecordType]),
		Value:               address,
		WalletAddressFormat: WalletAddressFormat(WalletAddressFormat_value[walletAddressFormat]),
	}

	// Initialize WalletRecords map if it is nil
	if domain.WalletRecords == nil {
		domain.WalletRecords = make(map[string]*WalletRecord)
	}

	domain.WalletRecords[walletRecordType] = walletRecord

	return err
}
