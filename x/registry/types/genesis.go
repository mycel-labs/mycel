package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// Default TLD Names
func GetDefaultTLDNames() []string {
	// Define default TLD names here
	return []string{"cel"}
}

// Get default TLDs
func GetDefaultTLDs() (defaultTLDs []Domain) {
	defaultRegistrationConfig := GetDefaultSubdomainRegistrationConfig(3030)
	for _, v := range GetDefaultTLDNames() {
		defaultTLDs = append(defaultTLDs, Domain{
			Name: v,
			SubdomainRegistrationConfig: &defaultRegistrationConfig,
		})
	}
	return defaultTLDs
}

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		Domains:          GetDefaultTLDs(),
		DomainOwnerships: []DomainOwnership{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in domain
	domainIndexMap := make(map[string]struct{})

	for _, elem := range gs.Domains {
		index := string(DomainKey(elem.Name, elem.Parent))
		if _, ok := domainIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for domain")
		}
		domainIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in domainOwnership
	domainOwnershipIndexMap := make(map[string]struct{})

	for _, elem := range gs.DomainOwnerships {
		index := string(DomainOwnershipKey(elem.Owner))
		if _, ok := domainOwnershipIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for domainOwnership")
		}
		domainOwnershipIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
