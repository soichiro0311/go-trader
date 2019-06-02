package model

import (
	"../domain"
)

type RegisterOrderRequest struct {
	CurrencyCode1 string `json:"currencycode1"`
	CurrencyCode2 string `json:"currencycode2"`
	Amount        int64  `json:"amount"`
	Side          string `json:"side"`
}

func NewRegisterOrderRequest(currency1 string, currency2 string, amount int64, side string) *RegisterOrderRequest {
	req := new(RegisterOrderRequest)
	req.CurrencyCode1 = currency1
	req.CurrencyCode2 = currency2
	req.Amount = amount
	req.Side = side
	return req
}

func (req *RegisterOrderRequest) ToOrder() *domain.Order {
	error, order := domain.NewOrder(req.Amount, domain.NewCurrency(req.CurrencyCode1), domain.NewCurrency(req.CurrencyCode2), req.Side)
	if error == nil {
		return order
	}
	panic(error)
}
