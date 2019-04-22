package repository

import (
	"../domain"
)

type QuoteInfoRepositoryImpl struct {
	cache map[int]domain.QuoteInfo
}

func NewQuoteInfoRepositoryImpl() *QuoteInfoRepositoryImpl {
	repository := new(QuoteInfoRepositoryImpl)
	repository.cache = map[int]domain.QuoteInfo{}
	return repository
}

func (repository *QuoteInfoRepositoryImpl) Save(info domain.QuoteInfo) {
	repository.cache[len(repository.cache)] = info
}

func (repository *QuoteInfoRepositoryImpl) FindByCurrencyPair(cur1 domain.Currency, cur2 domain.Currency) []domain.QuoteInfo {
	values := []domain.QuoteInfo{}
	for _, v := range repository.cache {
		if v.Currency1 == cur1 && v.Currency2 == cur2 {
			values = append(values, v)
		}
	}
	return values
}

func (repository *QuoteInfoRepositoryImpl) FindAll() []domain.QuoteInfo {
	values := []domain.QuoteInfo{}
	for _, v := range repository.cache {
		values = append(values, v)
	}
	return values
}
