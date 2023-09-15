package types

import (
	"errors"
	fmt "fmt"
	"regexp"

	errorsmod "cosmossdk.io/errors"
)

const (
	TLDNamePattern = `-a-z0-9\p{So}\p{Sk}`
)

func (topLevelDomain TopLevelDomain) ValidateName() (err error) {
	regex := regexp.MustCompile(fmt.Sprintf(`(^[%s]+$)`, TLDNamePattern))
	if !regex.MatchString(topLevelDomain.Name) {
		err = errorsmod.Wrapf(errors.New(fmt.Sprintf("%s", topLevelDomain.Name)), ErrInvalidDomainName.Error())
	}
	return err
}

func (topLevelDomain TopLevelDomain) Validate() (err error) {
	err = topLevelDomain.ValidateName()
	if err != nil {
		return err
	}
	return err
}
