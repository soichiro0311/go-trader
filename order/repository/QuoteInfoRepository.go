package repository

import "../domain"

type QuoteInfoRepository interface {
	GetLatestInfoByCurPair(currencyCode1 string, currencyCode2 string) domain.QuoteInfo
}
