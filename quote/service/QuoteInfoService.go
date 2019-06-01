package service

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"time"

	"../domain"
	"../enum"
	"../repository"
)

type QuoteInfoService struct {
	repository repository.QuoteInfoRepository
}

func NewQuoteInfoService(repository repository.QuoteInfoRepository) *QuoteInfoService {
	service := new(QuoteInfoService)
	service.repository = repository
	service.publishInfo()
	return service
}

// TODO: 外部サービスからレートを取得できるように修正
func (service *QuoteInfoService) GenerateInfo(cur1 string, cur2 string) domain.QuoteInfo {
	rand.Seed(time.Now().UnixNano())

	value := rand.Int63n(200)*1000 + 1
	error, price := domain.NewPrice(domain.NewCurrency(cur1), domain.NewCurrency(cur2), value)
	for _, v := range error {
		if v != nil {
			panic(v)
		}
	}

	amount := round(rand.Float64(), 3)

	var side enum.SideEnum

	sideRand := rand.Int63n(3)
	if sideRand%2 == 0 {
		side = enum.BUY
	} else {
		side = enum.SELL
	}

	error, info := domain.NewQuoteInfo(domain.NewCurrency(cur1), domain.NewCurrency(cur2), amount, price, side)
	for _, v := range error {
		if v != nil {
			panic(v)
		}
	}
	service.repository.Save(info)
	return info
}

func (service *QuoteInfoService) GetLatestByCurrencyPair(cur1 string, cur2 string) map[enum.SideEnum]domain.QuoteInfo {
	infos := service.repository.FindByCurrencyPair(domain.NewCurrency(cur1), domain.NewCurrency(cur2))
	buyInfos := []domain.QuoteInfo{}
	for _, v := range infos {
		if v.Side == enum.BUY {
			buyInfos = append(buyInfos, v)
		}
	}
	sort.SliceStable(buyInfos, func(i, j int) bool { return buyInfos[i].AccuredTime > buyInfos[j].AccuredTime })

	sellInfos := []domain.QuoteInfo{}
	for _, v := range infos {
		if v.Side == enum.SELL {
			sellInfos = append(sellInfos, v)
		}
	}
	sort.SliceStable(sellInfos, func(i, j int) bool { return sellInfos[i].AccuredTime > sellInfos[j].AccuredTime })

	latestInfos := map[enum.SideEnum]domain.QuoteInfo{}
	latestInfos[enum.BUY] = buyInfos[0]
	latestInfos[enum.SELL] = sellInfos[0]
	return latestInfos
}

func (service *QuoteInfoService) publishInfo() {
	go func() {
		ticker := time.NewTicker(5 * time.Second) // 1秒間隔のTicker

		for {
			select {
			case <-ticker.C:
				info := service.GenerateInfo("BTC", "USD")
				fmt.Println(info.Display())
			}
		}
	}()
}

func round(f float64, places int) float64 {
	shift := math.Pow(10, float64(places))
	return math.Floor(f*shift+.5) / shift
}
