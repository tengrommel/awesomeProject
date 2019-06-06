package currency

import (
	"encoding/json"
	"fmt"
)

type Currency string

func (c *Currency)UnmarshalJSON(in []byte) error {
	var s string
	err := json.Unmarshal(in, &s)
	if err != nil {
		return err
	}
	currency, valid := validCurrencies[s]
	if !valid {
		return fmt.Errorf("'%s' is not a valid currency", s)
	}
	*c = currency
	return nil
}

const (
	AUD = Currency("AUD")
	CNY = Currency("CNY")
	EUR = Currency("EUR")
	USD = Currency("USD")
)

var validCurrencies = map[string]Currency{
	string(AUD): AUD,
	string(CNY): CNY,
	string(EUR): EUR,
	string(USD): USD,
}