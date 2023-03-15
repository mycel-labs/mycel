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
		return sdkerrors.Wrapf(errors.New(fmt.Sprintf("%s", domain.Name)), ErrorDomainNameIsInvalid.Error())
	}
	return err
}

func (domain Domain) ValidateDomainParent() (err error) {
	regex := regexp.MustCompile(fmt.Sprintf(`(^[%s]+[%[1]s\.]*[%[1]s]$)|^$`, NamePattern))
	if !regex.MatchString(domain.Parent) {
		return sdkerrors.Wrapf(errors.New(fmt.Sprintf("%s", domain.Parent)), ErrorDomainParentIsInvalid.Error())
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
