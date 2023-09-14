package types

import (
	"errors"
	fmt "fmt"
	"regexp"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TLDNamePattern = `-a-z0-9\p{So}\p{Sk}`
)

func ValidateTopLevelDomainName(name string) (err error) {
	regex := regexp.MustCompile(fmt.Sprintf(`(^[%s]+$)`, TLDNamePattern))
	if !regex.MatchString(name) {
		err = sdkerrors.Wrapf(errors.New(fmt.Sprintf("%s", name)), ErrInvalidDomainName.Error())
	}
	return err

}

func (topLevelDomain TopLevelDomain) ValidateName() (err error) {
	err = ValidateTopLevelDomainName(topLevelDomain.Name)
	return err
}

func (topLevelDomain TopLevelDomain) Validate() (err error) {
	err = topLevelDomain.ValidateName()
	if err != nil {
		return err
	}
	return err
}
