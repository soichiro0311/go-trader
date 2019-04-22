package repository

import (
	"../domain"
)

type BoardInfoRepositoryImpl struct {
	cache map[int]domain.BoardInfo
}

func NewBoardInfoRepositoryImpl() *BoardInfoRepositoryImpl {
	repository := new(BoardInfoRepositoryImpl)
	repository.cache = map[int]domain.BoardInfo{}
	return repository
}

func (repository *BoardInfoRepositoryImpl) Save(info domain.BoardInfo) {
	repository.cache[len(repository.cache)] = info
}

func (repository *BoardInfoRepositoryImpl) FindByCurrencyPair(cur1 domain.Currency, cur2 domain.Currency) []domain.BoardInfo {
	values := []domain.BoardInfo{}
	for _, v := range repository.cache {
		if v.Currency1 == cur1 && v.Currency2 == cur2 {
			values = append(values, v)
		}
	}
	return values
}

func (repository *BoardInfoRepositoryImpl) FindAll() []domain.BoardInfo {
	values := []domain.BoardInfo{}
	for _, v := range repository.cache {
		values = append(values, v)
	}
	return values
}
