package repository

import (
	"encoding/json"
	"net/http"

	"../domain"
	"../dto"
)

type quoteInfoRepositoryImpl struct {
	cache map[string][]domain.QuoteInfo
}

func NewQuoteInfoRepositoryImpl() *quoteInfoRepositoryImpl {
	repository := new(quoteInfoRepositoryImpl)
	repository.cache = map[string][]domain.QuoteInfo{}
	//TODO: 初期ロードする対象の銘柄情報をどこかに記録しておく。今はベタ書き。
	resp, err := http.Get("http://localhost:8091/quote/all/BTC/USD")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	infoMap := map[int]dto.QuoteInfoDto{}
	json.NewDecoder(resp.Body).Decode(&infoMap)
	for _, v := range infoMap {
		repository.cache[v.Side] = append(repository.cache[v.Side], domain.NewQuoteInfo(v))
	}
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
