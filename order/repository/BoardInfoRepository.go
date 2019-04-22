package repository

import "../domain"

type BoardInfoRepository interface {
	GetLatestInfoByCurPair(currencyCode1 string, currencyCode2 string) domain.BoardInfo
}
