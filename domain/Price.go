package domain

import "math/big"

type Price struct {
	Currency1 string
	Currency2 string
	Value     big.Float
}

func NewPrice(Currency1 string, Currency2 string, Value big.Float) *Price {
	price := new(Price)
	price.Currency1 = Currency1
	price.Currency2 = Currency2
	price.Value = Value
	return price
}

func (price *Price) converToCurrency2(amount big.Int) big.Float {
	quantity := new(big.Float).SetPrec(1024).Mul(new(big.Float).SetInt(&amount), &price.Value)
	return *quantity
}
