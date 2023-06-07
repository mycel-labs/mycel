package types

import (
	"errors"
	fmt "fmt"
	math "math"
	"regexp"
	"strings"
	"unicode/utf8"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	NamePattern = `-a-z0-9\p{So}\p{Sk}`
	BaseFee     = 303
)

func (domain Domain) ValidateName() (err error) {
	regex := regexp.MustCompile(fmt.Sprintf(`(^[%s]+$)`, NamePattern))
	if !regex.MatchString(domain.Name) {
		err = sdkerrors.Wrapf(errors.New(fmt.Sprintf("%s", domain.Name)), ErrInvalidDomainName.Error())
	}
	return err
}

func (domain Domain) ValidateParent() (err error) {
	regex := regexp.MustCompile(fmt.Sprintf(`(^[%s]+[%[1]s\.]*[%[1]s]$)|^$`, NamePattern))
	if !regex.MatchString(domain.Parent) {
		err = sdkerrors.Wrapf(errors.New(fmt.Sprintf("%s", domain.Parent)), ErrInvalidDomainParent.Error())
	}
	return err
}

func (domain Domain) Validate() (err error) {
	err = domain.ValidateName()
	if err != nil {
		return err
	}
	err = domain.ValidateParent()
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

func (domain Domain) ParseParent() (name string, parent string) {
	if domain.Parent != "" {
		split := strings.Split(domain.Parent, ".")
		if len(split) == 1 {
			name = split[0]
		} else {
			parent = split[len(split)-1]
			name = strings.Join(split[:len(split)-1], ".")
		}
	}
	return name, parent
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

func (domain *Domain) UpdateWalletRecord(walletRecordType string, address string) (err error) {

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

func ValidateDnsRecordValue(dnsRecordFormat string, address string) (err error) {
	dnsRecordRegex, isFound := DnsRecordValueRegex()[dnsRecordFormat]
	if !isFound {
		panic(fmt.Sprintf("Dns record value format %s is not found in DnsRecordValueRegex", dnsRecordFormat))
	}
	regex := regexp.MustCompile(dnsRecordRegex)
	if !regex.MatchString(address) {
		err = sdkerrors.Wrapf(errors.New(fmt.Sprintf("%s %s", dnsRecordFormat, address)), ErrInvalidDnsRecordValue.Error())
	}
	return err
}

func ValidateDnsRecordType(dnsRecordType string) (err error) {
	_, isFound := DnsRecordType_value[dnsRecordType]
	if !isFound {
		err = sdkerrors.Wrapf(errors.New(fmt.Sprintf("%s", dnsRecordType)), ErrInvalidDnsRecordType.Error())
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

func (domain *Domain) UpdateDnsRecord(dnsRecordType string, value string) (err error) {

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
		DnsRecordType:   DnsRecordType(DnsRecordType_value[dnsRecordType]),
		DnsRecordFormat: DnsRecordFormat(DnsRecordFormat_value[dnsRecordFormat]),
		Value:           value,
	}

	// Initialize WalletRecords map if it is nil
	if domain.DnsRecords == nil {
		domain.DnsRecords = make(map[string]*DnsRecord)
	}

	domain.DnsRecords[dnsRecordType] = dnsRecord

	return err
}

func (domain *Domain) GetRegistrationFee() (fee sdk.Coin) {
	nameLen := utf8.RuneCountInString(domain.Name)
	amount := 0
	if nameLen >= 5 {
		amount = BaseFee
	} else {
		amount = BaseFee * int(math.Pow(10, float64((5-nameLen))))
	}
	fee = sdk.NewCoin(MycelDenom, sdk.NewInt(int64(amount)))

	return fee
}
