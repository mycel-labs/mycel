package types

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func GetValidDomains() []Domain {
	return []Domain{
		{Name: "foo", Parent: "myc"},
		{Name: "", Parent: "myc"},
		{Name: "bar", Parent: "foo.myc"},
		{Name: "🍭", Parent: "foo.🍭"},
	}
}

// Name is invalid
func GetInvalidNameDomains() []Domain {
	return []Domain{
		{Name: ".foo", Parent: "myc"},
		{Name: "bar.foo", Parent: "myc"},
		{Name: ".", Parent: "myc"},
		{Name: "##", Parent: "myc"},
	}
}

// Parent is invalid
func GetInvalidParentDomains() []Domain {
	return []Domain{
		{Name: "foo", Parent: ""},
		{Name: "foo", Parent: ".##"},
		{Name: "foo", Parent: ".myc"},
		{Name: "foo", Parent: ".foo.myc"},
	}
}

func TestValidateDomainNameSuccess(t *testing.T) {
	for _, v := range GetValidDomains() {
		err := v.ValidateDomainName()
		require.Nil(t, err)
	}
}
func TestValidateDomainNameFailure(t *testing.T) {
	for _, v := range GetInvalidNameDomains() {
		err := v.ValidateDomainName()
		require.EqualError(t, err, fmt.Sprintf("name is invalid: %s", v.Name))
	}
}

func TestValidateDomainParentSuccess(t *testing.T) {
	for _, v := range GetValidDomains() {
		err := v.ValidateDomainParent()
		require.Nil(t, err)
	}
}

func TestValidateDomainParentFailure(t *testing.T) {
	for _, v := range GetInvalidParentDomains() {
		err := v.ValidateDomainParent()
		require.EqualError(t, err, fmt.Sprintf("parent is invalid: %s", v.Parent))
	}
}

func TestValidateDomainSuccess(t *testing.T) {
	for _, v := range GetValidDomains() {
		err := v.ValidateDomain()
		require.Nil(t, err)
	}
}

func TestValidateDomainFailure(t *testing.T) {
	for _, v := range GetInvalidNameDomains() {
		err := v.ValidateDomain()
		require.EqualError(t, err, fmt.Sprintf("name is invalid: %s", v.Name))
	}
	for _, v := range GetInvalidParentDomains() {
		err := v.ValidateDomainParent()
		require.EqualError(t, err, fmt.Sprintf("parent is invalid: %s", v.Parent))
	}
}

func TestGetIsRootDomain(t *testing.T) {
	for i, v := range GetValidDomains() {
		isRootDomain, err := v.GetIsRootDomain()
		require.Nil(t, err)
		if i < 2 {
			require.Equal(t, isRootDomain, true)
		} else {
			require.Equal(t, isRootDomain, false)
		}
	}
}
