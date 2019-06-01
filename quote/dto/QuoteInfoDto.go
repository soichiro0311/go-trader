package dto

import (
	"../domain"
)

type QuoteInfoDto struct {
	Currency1   string  `json:"currency1"`
	Currency2   string  `json:"currency2"`
	Amount      float64 `json:"amount"`
	Price       int64   `json:"price"`
	Side        string  `json:"side"`
	AccuredTime string  `json:"accured_time"`
}

func NewQuoteInfoDto(info domain.QuoteInfo) *QuoteInfoDto {
	dto := new(QuoteInfoDto)
	dto.Currency1 = info.Currency1.CurrencyCode
	dto.Currency2 = info.Currency2.CurrencyCode
	dto.Amount = info.Amount
	dto.Price = info.Price.Value
	dto.Side = info.Side.String()
	dto.AccuredTime = info.AccuredTime

	return dto
}
