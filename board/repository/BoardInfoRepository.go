package repository

import "../domain"

type BoardInfoRepository interface {
	Save(info domain.BoardInfo)
	FindByCurrencyPair(cur1 domain.Currency, cur2 domain.Currency) []domain.BoardInfo
	FindAll() []domain.BoardInfo
}
