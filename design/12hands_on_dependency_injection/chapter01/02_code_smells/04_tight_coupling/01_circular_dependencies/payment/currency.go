package payment

import (
	"errors"
	"github.com/tengrommel/awesomeProject/design/12hands_on_dependency_injection/chapter01/02_code_smells/04_tight_coupling/01_circular_dependencies/config"
)

// Currency is custom type for currency
type Currency string

// Processor processes payments
type Processor struct {
	Config *config.Config
}

// Pay makes a payment in the default currency
func (p *Processor)Pay(amount float64) error {
	// TODO: implement me
	return errors.New("not implemented yet")
}


