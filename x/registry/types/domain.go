package types

import (
	fmt "fmt"
	math "math"
	"strings"
	"unicode/utf8"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	BaseFee     = 303
)

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
		WalletRecordType:    NetworkName(NetworkName_value[walletRecordType]),
		Value:               address,
	}

	// Initialize WalletRecords map if it is nil
	if domain.WalletRecords == nil {
		domain.WalletRecords = make(map[string]*WalletRecord)
	}

	domain.WalletRecords[walletRecordType] = walletRecord

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


