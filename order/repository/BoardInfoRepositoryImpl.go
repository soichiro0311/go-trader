package repository

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../domain"
)

type boardInfoRepositoryImpl struct {
}

func NewBoardInfoRepositoryImpl() *boardInfoRepositoryImpl {
	repository := new(boardInfoRepositoryImpl)
	return repository
}

func (repository *boardInfoRepositoryImpl) GetLatestInfoByCurPair(currencyCode1 string, currencyCode2 string) domain.BoardInfo {
	resp, err := http.Get("http://localhost:8080/board/latest/" + currencyCode1 + "/" + currencyCode2)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	info := domain.BoardInfo{}
	json.NewDecoder(resp.Body).Decode(&info)
	fmt.Println(info)
	return info
}
