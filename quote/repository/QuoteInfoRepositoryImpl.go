package repository

import (
	"../domain"
	"../enum"
)

type QuoteInfoRepositoryImpl struct {
	cache map[QuoteInfoKey]domain.QuoteInfo
}

type QuoteInfoKey struct {
	id   int
	side enum.SideEnum
}

func NewQuoteInfoRepositoryImpl() *QuoteInfoRepositoryImpl {
	repository := new(QuoteInfoRepositoryImpl)
	repository.cache = map[QuoteInfoKey]domain.QuoteInfo{}
	return repository
}

func (repository *QuoteInfoRepositoryImpl) Save(info domain.QuoteInfo) {
	key := new(QuoteInfoKey)
	key.id = len(repository.cache)
	key.side = info.Side
	repository.cache[*key] = info
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
