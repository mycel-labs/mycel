package types

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

type DomainTest struct {
	Domain       Domain
	IsTLD        bool
	IsRootDomain bool
	IsSubDomain  bool
}

func GetValidDomains() []DomainTest {
	return []DomainTest{
		{Domain: Domain{Name: "foo", Parent: "myc"}, IsTLD: false, IsRootDomain: true, IsSubDomain: false},
		{Domain: Domain{Name: "foo", Parent: ""}, IsTLD: true, IsRootDomain: false, IsSubDomain: false},
		{Domain: Domain{Name: "bar", Parent: "foo.myc"}, IsTLD: false, IsRootDomain: false, IsSubDomain: true},
		{Domain: Domain{Name: "🍭", Parent: "foo.🍭"}, IsTLD: false, IsRootDomain: false, IsSubDomain: true},
	}
}

// Name is invalid
func GetInvalidNameDomains() []DomainTest {
	return []DomainTest{
		{Domain: Domain{Name: ".foo", Parent: "myc"}},
		{Domain: Domain{Name: "", Parent: "myc"}},
		{Domain: Domain{Name: "bar.foo", Parent: "myc"}},
		{Domain: Domain{Name: ".", Parent: "myc"}},
		{Domain: Domain{Name: "##", Parent: "myc"}},
	}
}

// Parent is invalid
func GetInvalidParentDomains() []DomainTest {
	return []DomainTest{
		{Domain: Domain{Name: "foo", Parent: ".##"}},
		{Domain: Domain{Name: "foo", Parent: ".myc"}},
		{Domain: Domain{Name: "foo", Parent: ".foo.myc"}},
	}
}

func TestValidateDomainNameSuccess(t *testing.T) {
	for _, v := range GetValidDomains() {
		err := v.Domain.ValidateDomain()
		require.Nil(t, err)
	}
}
func TestValidateDomainNameFailure(t *testing.T) {
	for _, v := range GetInvalidNameDomains() {
		err := v.Domain.ValidateDomainName()
		require.EqualError(t, err, fmt.Sprintf("name is invalid: %s", v.Domain.Name))
	}
}

func TestValidateDomainParentSuccess(t *testing.T) {
	for _, v := range GetValidDomains() {
		err := v.Domain.ValidateDomainParent()
		require.Nil(t, err)
	}
}

func TestValidateDomainParentFailure(t *testing.T) {
	for _, v := range GetInvalidParentDomains() {
		err := v.Domain.ValidateDomainParent()
		require.EqualError(t, err, fmt.Sprintf("parent is invalid: %s", v.Domain.Parent))
	}
}

func TestGetIsRootDomain(t *testing.T) {
	for _, v := range GetValidDomains() {
		isRootDomain := v.Domain.GetIsRootDomain()
		require.Equal(t, isRootDomain, v.IsRootDomain)
	}
}

func TestGetIsTLD(t *testing.T) {
	for _, v := range GetValidDomains() {
		isTLD := v.Domain.GetIsTLD()
		require.Equal(t, isTLD, v.IsTLD)
	}
}

func TestGetIsSubDomain(t *testing.T) {
	for _, v := range GetValidDomains() {
		isSubDomain := v.Domain.GetIsSubDomain()
		require.Equal(t, isSubDomain, v.IsSubDomain)
	}

}

func TestValidateDomainSuccess(t *testing.T) {
	for _, v := range GetValidDomains() {
		err := v.Domain.ValidateDomain()
		require.Nil(t, err)
	}
}

func TestValidateDomainFailure(t *testing.T) {
	for _, v := range GetInvalidNameDomains() {
		err := v.Domain.ValidateDomain()
		require.EqualError(t, err, fmt.Sprintf("name is invalid: %s", v.Domain.Name))
	}
	for _, v := range GetInvalidParentDomains() {
		err := v.Domain.ValidateDomainParent()
		require.EqualError(t, err, fmt.Sprintf("parent is invalid: %s", v.Domain.Parent))
	}
}
