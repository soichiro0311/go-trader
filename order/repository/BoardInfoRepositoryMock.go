package repository

import (
	"sort"

	"../domain"
)

type BoardInfoRepositoryMock struct {
	inmemoryMap map[string]domain.BoardInfo
}

func NewBoardInfoRepositoryMock() *BoardInfoRepositoryMock {
	mock := new(BoardInfoRepositoryMock)
	mock.inmemoryMap = map[string]domain.BoardInfo{}
	return mock
}

func (repository *BoardInfoRepositoryMock) Save(info domain.BoardInfo) {
	repository.inmemoryMap[info.AccuredTime] = info

}

func (repository *BoardInfoRepositoryMock) GetLatestInfoByCurPair(currencyCode1 string, currencyCode2 string) domain.BoardInfo {
	values := []domain.BoardInfo{}
	for _, v := range repository.inmemoryMap {
		if v.Currency1 == domain.NewCurrency(currencyCode1) && v.Currency2 == domain.NewCurrency(currencyCode2) {
			values = append(values, v)
		}
	}
	sort.SliceStable(values, func(i, j int) bool { return values[i].AccuredTime > values[j].AccuredTime })
	return values[0]
}
