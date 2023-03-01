package types

import (
	"errors"
	fmt "fmt"
	"regexp"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	NamePattern = `-a-z0-9\p{So}\p{Sk}`
)

func (domain Domain) ValidateDomainName() (err error) {
	regex := regexp.MustCompile(fmt.Sprintf(`(^[%s]+$)|(^$)`, NamePattern))
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

func (domain Domain) GetIsTLD() (isTLD bool) {
	if domain.Parent == "" {
		isTLD = true
	}
	return isTLD
}

func (domain Domain) GetIsRootDomain() (isRootDomain bool) {
	regex := regexp.MustCompile(fmt.Sprintf(`^[%s]+$`, NamePattern))
	if regex.MatchString(domain.Parent) {
		isRootDomain = true
	}
	return isRootDomain
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
