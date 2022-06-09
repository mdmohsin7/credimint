package types

import (
	"fmt"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		UserCredibilityList: []UserCredibility{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ID in userCredibility
	userCredibilityIdMap := make(map[uint64]bool)
	userCredibilityCount := gs.GetUserCredibilityCount()
	for _, elem := range gs.UserCredibilityList {
		if _, ok := userCredibilityIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for userCredibility")
		}
		if elem.Id >= userCredibilityCount {
			return fmt.Errorf("userCredibility id should be lower or equal than the last id")
		}
		userCredibilityIdMap[elem.Id] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
