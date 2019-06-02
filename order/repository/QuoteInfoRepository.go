package repository

import "../domain"

type QuoteInfoRepository interface {
	GetLatestInfoByCurPair(currencyCode1 string, currencyCode2 string, side string) domain.QuoteInfo
}
