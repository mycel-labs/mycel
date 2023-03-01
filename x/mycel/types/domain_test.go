package types

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

type DomainTest struct {
	Domain      Domain
	DomainLevel int
}

func GetValidDomains() []DomainTest {
	return []DomainTest{
		{Domain: Domain{Name: "foo", Parent: "myc"}, DomainLevel: 2},
		{Domain: Domain{Name: "foo", Parent: ""}, DomainLevel: 1},
		{Domain: Domain{Name: "bar", Parent: "foo.myc"}, DomainLevel: 3},
		{Domain: Domain{Name: "🍭", Parent: "foo.🍭"}, DomainLevel: 3},
		{Domain: Domain{Name: "🍭", Parent: "foo.🍭.myc"}, DomainLevel: 4},
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

func TestGetDomainLevel(t *testing.T) {
	for _, v := range GetValidDomains() {
		domainLevel := v.Domain.GetDomainLevel()
		require.Equal(t, domainLevel, v.DomainLevel)
	}
}
