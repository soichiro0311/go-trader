package domain

import "errors"

type Price struct {
	Currency1 Currency `json:"currency1"`
	Currency2 Currency `json:"currency2"`
	Value     int64    `json:"value"`
}

func NewPrice(cur1 Currency, cur2 Currency, value int64) ([]error, Price) {
	price := new(Price)
	price.Currency1 = cur1
	price.Currency2 = cur2
	price.Value = value
	valueError := price.validateValue()
	currencyError := price.validateCurrency()

	if valueError == nil && currencyError == nil {
		return nil, *price
	}

	return []error{valueError, currencyError}, Price{}
}

func (price Price) validateValue() error {
	if price.Value <= 0 {
		return errors.New("Price value should be greater than 0")
	}
	return nil
}

func (price Price) validateCurrency() error {
	if price.Currency1.CurrencyCode == price.Currency2.CurrencyCode {
		return errors.New("Currency1,Currency2 should not be equal")
	}
	return nil
}
