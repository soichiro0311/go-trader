package repository

import "../domain"

type QuoteInfoRepository interface {
	Save(info domain.QuoteInfo)
	FindByCurrencyPair(cur1 domain.Currency, cur2 domain.Currency) []domain.QuoteInfo
	FindAll() []domain.QuoteInfo
}
