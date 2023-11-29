package types

import (
	"fmt"
  "math"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// Default TLD Names
func GetDefaultTLDNames() []string {
	// Define default TLD names here
	return []string{"cel"}
}

// Get default TLDs
func GetDefaultTLDs() (defaultTLDs []TopLevelDomain) {
	defaultRegistrationConfig := GetDefaultSubdomainConfig(3030)
  defaultRegistrationConfig.MaxSubdomainRegistrations = math.MaxInt64
	for _, v := range GetDefaultTLDNames() {
		defaultTLDs = append(defaultTLDs, TopLevelDomain{
			Name:            v,
			SubdomainConfig: defaultRegistrationConfig,
		})
	}
	return defaultTLDs
}

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		TopLevelDomains:    GetDefaultTLDs(),
		SecondLevelDomains: []SecondLevelDomain{},
		DomainOwnerships:   []DomainOwnership{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in topLevelDomain
	topLevelDomainIndexMap := make(map[string]struct{})

	for _, elem := range gs.TopLevelDomains {
		index := string(TopLevelDomainKey(elem.Name))
		if _, ok := topLevelDomainIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for topLevelDomain")
		}
		topLevelDomainIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in secondLevelDomain
	secondLevelDomainIndexMap := make(map[string]struct{})

	for _, elem := range gs.SecondLevelDomains {
		index := string(SecondLevelDomainKey(elem.Name, elem.Parent))
		if _, ok := secondLevelDomainIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for secondLevelDomain")
		}
		secondLevelDomainIndexMap[index] = struct{}{}
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
