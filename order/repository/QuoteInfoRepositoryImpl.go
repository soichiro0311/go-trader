package repository

import (
	"encoding/json"
	"net/http"

	"../domain"
	"../dto"
)

type quoteInfoRepositoryImpl struct {
}

func NewQuoteInfoRepositoryImpl() *quoteInfoRepositoryImpl {
	repository := new(quoteInfoRepositoryImpl)
	return repository
}

func (repository *quoteInfoRepositoryImpl) GetLatestInfoByCurPair(currencyCode1 string, currencyCode2 string, side string) domain.QuoteInfo {
	resp, err := http.Get("http://localhost:8091/quote/latest/" + currencyCode1 + "/" + currencyCode2)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	infoMap := map[int]dto.QuoteInfoDto{}
	json.NewDecoder(resp.Body).Decode(&infoMap)
	for _, v := range infoMap {
		if side == v.Side {
			return domain.NewQuoteInfo(v)
		}
	}

	panic("Quote Not Exist")
}
