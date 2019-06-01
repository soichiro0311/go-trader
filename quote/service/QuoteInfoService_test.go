package service

import (
	"testing"
	"time"

	"../repository"
)

func TestGetLatestByCurrencyPair(t *testing.T) {
	service := NewQuoteInfoService(repository.NewQuoteInfoRepositoryImpl())
	cur1 := "BTC"
	cur2 := "USD"

	service.GenerateInfo(cur1, cur2)
	time.Sleep(2 * time.Second)
	service.GenerateInfo(cur1, cur2)
	time.Sleep(2 * time.Second)
	service.GenerateInfo(cur1, cur2)
	time.Sleep(2 * time.Second)
	latestInfo := service.GenerateInfo(cur1, cur2)

	info := service.GetLatestByCurrencyPair(cur1, cur2)

	if latestInfo != info {
		t.Fail()
	}
}
