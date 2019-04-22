package domain

import (
	"errors"
	"math"
)

type Price struct {
	Currency1 Currency `json:"currency1"`
	Currency2 Currency `json:"currency2"`
	Value     float64  `json:"value"`
}

func NewPrice(currency1 Currency, currency2 Currency, value float64) (error, Price) {
	price := new(Price)
	price.Currency1 = currency1
	price.Currency2 = currency2
	err := price.validateCurrency()
	price.Value = value
	return err, *price
}

func (price *Price) converToCurrency2(amount float64) float64 {
	quantity := price.Value * float64(amount)
	return round(quantity, .5, 5)
}

func (price *Price) validateCurrency() error {
	if price.Currency1.CurrencyCode == price.Currency2.CurrencyCode {
		return errors.New("Currency1,Currency2 should not be equal")
	}
	return nil
}

func round(val float64, roundOn float64, places int) (newVal float64) {
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * val
	_, div := math.Modf(digit)
	if div >= roundOn {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	newVal = round / pow
	return newVal
}
