package service

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"time"

	"../domain"
	"../repository"
)

type BoardInfoService struct {
	repository repository.BoardInfoRepository
}

func NewBoardInfoService(repository repository.BoardInfoRepository) *BoardInfoService {
	service := new(BoardInfoService)
	service.repository = repository
	service.publishInfo()
	return service
}

// TODO: 外部サービスからレートを取得できるように修正
func (service *BoardInfoService) GenerateInfo(cur1 string, cur2 string) domain.BoardInfo {
	rand.Seed(time.Now().UnixNano())

	value := rand.Int63n(2000000) * 1000
	error, price := domain.NewPrice(domain.NewCurrency(cur1), domain.NewCurrency(cur2), value)
	for _, v := range error {
		if v != nil {
			panic(v)
		}
	}

	shift := math.Pow(10, 5)
	amount := math.Trunc(rand.Float64()*shift) / shift
	error, info := domain.NewBoardInfo(domain.NewCurrency(cur1), domain.NewCurrency(cur2), amount, price)
	for _, v := range error {
		if v != nil {
			panic(v)
		}
	}
	service.repository.Save(info)
	return info
}

func (service *BoardInfoService) GetByCurrencyPair(cur1 string, cur2 string) []domain.BoardInfo {
	return service.repository.FindByCurrencyPair(domain.NewCurrency(cur1), domain.NewCurrency(cur2))
}

func (service *BoardInfoService) GetLatestByCurrencyPair(cur1 string, cur2 string) domain.BoardInfo {
	infos := service.repository.FindByCurrencyPair(domain.NewCurrency(cur1), domain.NewCurrency(cur2))
	sort.SliceStable(infos, func(i, j int) bool { return infos[i].AccuredTime > infos[j].AccuredTime })
	return infos[0]
}

func (service *BoardInfoService) publishInfo() {
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
