package model

import (
	"../domain"
)

type RegisterOrderRequest struct {
	CurrencyCode1 string `json:"currencycode1"`
	CurrencyCode2 string `json:"currencycode2"`
}

func NewRegisterOrderRequest(currency1 string, currency2 string) *RegisterOrderRequest {
	req := new(RegisterOrderRequest)
	req.CurrencyCode1 = currency1
	req.CurrencyCode2 = currency2
	return req
}

func (req *RegisterOrderRequest) ToOrder(info domain.BoardInfo) *domain.Order {
	error, order := domain.NewOrder(info.Amount, domain.NewCurrency(req.CurrencyCode1), domain.NewCurrency(req.CurrencyCode2), info.Price)
	if error == nil {
		return order
	}
	panic(error)
}
