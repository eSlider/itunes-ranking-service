package itunes

import (
	"database/sql/driver"
	"fmt"
)

type PriceLabel string

const (
	Get PriceLabel = "Get"
)

type Price struct {
	Label      PriceLabel `json:"label,omitempty"`
	Attributes struct {
		Amount   string   `json:"amount,omitempty"`
		Currency Currency `json:"currency,omitempty"`
	} `json:"attributes,omitempty"`
}

func (l Price) Value() (driver.Value, error) {
	return fmt.Sprintf("%s %s", l.Attributes.Amount, l.Attributes.Currency), nil
}
