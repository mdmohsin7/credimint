package types

import (
	"fmt"
	host "github.com/cosmos/ibc-go/v5/modules/core/24-host"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		PortId:   PortID,
		LoanList: []Loan{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	if err := host.PortIdentifierValidator(gs.PortId); err != nil {
		return err
	}
	// Check for duplicated ID in loan
	loanIdMap := make(map[uint64]bool)
	loanCount := gs.GetLoanCount()
	for _, elem := range gs.LoanList {
		if _, ok := loanIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for loan")
		}
		if elem.Id >= loanCount {
			return fmt.Errorf("loan id should be lower or equal than the last id")
		}
		loanIdMap[elem.Id] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
