package model

import (
	"../domain"
)

type RegisterOrderRequest struct {
	CurrencyCode1 string `json:"currencycode1"`
	CurrencyCode2 string `json:"currencycode2"`
	Amount        int64  `json:"amount"`
}

func NewRegisterOrderRequest(currency1 string, currency2 string, amount int64) *RegisterOrderRequest {
	req := new(RegisterOrderRequest)
	req.CurrencyCode1 = currency1
	req.CurrencyCode2 = currency2
	req.Amount = amount
	return req
}

func (req *RegisterOrderRequest) ToOrder(price domain.Price) *domain.Order {
	error, order := domain.NewOrder(req.Amount, *domain.NewCurrency(req.CurrencyCode1), *domain.NewCurrency(req.CurrencyCode2), price)
	if error == nil {
		return order
	}
	panic(error)
}
