package domain

import (
	"../dto"
)

type QuoteInfo struct {
	Currency1   Currency `json:"currency1"`
	Currency2   Currency `json:"currency2"`
	Amount      int64    `json:"amount"`
	Price       Price    `json:"price"`
	AccuredTime string   `json:"accured_time"`
}

func NewQuoteInfo(dto dto.QuoteInfoDto) QuoteInfo {
	info := new(QuoteInfo)
	info.Currency1 = NewCurrency(dto.Currency1)
	info.Currency2 = NewCurrency(dto.Currency2)
	info.Amount = dto.Amount
	_, info.Price = NewPrice(info.Currency1, info.Currency2, dto.Price)
	info.AccuredTime = dto.AccuredTime

	return *info
}
