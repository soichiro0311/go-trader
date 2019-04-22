package repository

import (
	"sort"

	"../domain"
)

type QuoteInfoRepositoryMock struct {
	inmemoryMap map[string]domain.QuoteInfo
}

func NewQuoteInfoRepositoryMock() *QuoteInfoRepositoryMock {
	mock := new(QuoteInfoRepositoryMock)
	mock.inmemoryMap = map[string]domain.QuoteInfo{}
	return mock
}

func (repository *QuoteInfoRepositoryMock) Save(info domain.QuoteInfo) {
	repository.inmemoryMap[info.AccuredTime] = info

}

func (repository *QuoteInfoRepositoryMock) GetLatestInfoByCurPair(currencyCode1 string, currencyCode2 string) domain.QuoteInfo {
	values := []domain.QuoteInfo{}
	for _, v := range repository.inmemoryMap {
		if v.Currency1 == domain.NewCurrency(currencyCode1) && v.Currency2 == domain.NewCurrency(currencyCode2) {
			values = append(values, v)
		}
	}
	sort.SliceStable(values, func(i, j int) bool { return values[i].AccuredTime > values[j].AccuredTime })
	return values[0]
}
