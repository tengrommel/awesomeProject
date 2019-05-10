package config

import (
	"errors"
	"github.com/tengrommel/awesomeProject/design/12hands_on_dependency_injection/chapter01/02_code_smells/04_tight_coupling/01_circular_dependencies/payment"
)

// Config defines the JSON format of the config file
type Config struct {
	// Address is the host and post to bind to.
	// Default 0.0.0.0:8080
	Address string
	// DefaultCurrency is the default currency of the system
	DefaultCurrency payment.Currency
}

// Load will load the JSON config from the file supplied
func Load(filename string) (*Config, error) {
	// TODO: load currency from title
	return nil, errors.New("not implemented yet")
}