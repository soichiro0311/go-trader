package service

import (
	"testing"
	"time"

	"../repository"
)

func TestGetByCurrencyPair(t *testing.T) {
	service := NewQuoteInfoService(repository.NewQuoteInfoRepositoryImpl())
	cur1 := "BTC"
	cur2 := "USD"
	cur3 := "ETH"

	service.GenerateInfo(cur1, cur2)
	service.GenerateInfo(cur3, cur2)
	service.GenerateInfo(cur1, cur3)
	service.GenerateInfo(cur1, cur2)

	infos := service.GetByCurrencyPair(cur1, cur2)
	if len(infos) != 2 {
		t.Fail()
	}

	for _, v := range infos {
		if v.Currency1.Code() != cur1 || v.Currency2.Code() != cur2 {
			t.Fail()
		}
	}
}

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
