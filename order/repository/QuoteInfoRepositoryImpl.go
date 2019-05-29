package repository

import (
	"encoding/json"
	"net/http"

	"../domain"
)

type quoteInfoRepositoryImpl struct {
}

func NewQuoteInfoRepositoryImpl() *quoteInfoRepositoryImpl {
	repository := new(quoteInfoRepositoryImpl)
	return repository
}

func (repository *quoteInfoRepositoryImpl) GetLatestInfoByCurPair(currencyCode1 string, currencyCode2 string) domain.QuoteInfo {
	resp, err := http.Get("http://localhost:8091/quote/latest/" + currencyCode1 + "/" + currencyCode2)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	info := domain.QuoteInfo{}
	json.NewDecoder(resp.Body).Decode(&info)
	return info
}
