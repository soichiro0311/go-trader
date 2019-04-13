package domain

import (
	"errors"
	"math"
)

type Price struct {
	currency1 Currency
	currency2 Currency
	value     float64
}

func NewPrice(currency1 Currency, currency2 Currency, value float64) (error, *Price) {
	price := new(Price)
	price.currency1 = currency1
	price.currency2 = currency2
	err := price.validateCurrency()
	price.value = value
	return err, price
}

func (price *Price) converToCurrency2(amount int64) float64 {
	quantity := price.value * float64(amount)
	return round(quantity, .5, 5)
}

func (price *Price) validateCurrency() error {
	if price.currency1.CurrencyCode == price.currency2.CurrencyCode {
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
