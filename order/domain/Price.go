package domain

import (
	"errors"
)

type Price struct {
	Currency1 Currency `json:"currency1"`
	Currency2 Currency `json:"currency2"`
	Value     int64    `json:"value"`
}

func NewPrice(currency1 Currency, currency2 Currency, value int64) (error, Price) {
	price := new(Price)
	price.Currency1 = currency1
	price.Currency2 = currency2
	err := price.validateCurrency()
	price.Value = value
	return err, *price
}

func (price *Price) converToCurrency2(amount int64) int64 {
	quantity := price.Value * amount
	return quantity
}

func (price *Price) validateCurrency() error {
	if price.Currency1.CurrencyCode == price.Currency2.CurrencyCode {
		return errors.New("Currency1,Currency2 should not be equal")
	}
	return nil
}
