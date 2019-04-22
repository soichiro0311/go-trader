package service

import (
	"fmt"
	"testing"

	"../domain"
	"../model"
	"../repository"
)

func TestRegisterOrder(t *testing.T) {
	var orderRep repository.OrderRepository = repository.NewOrderRepositoryMock()
	repMock := repository.NewQuoteInfoRepositoryMock()
	cur1 := domain.NewCurrency("BTC")
	cur2 := domain.NewCurrency("USD")
	_, price := domain.NewPrice(cur1, cur2, 2900.555)
	repMock.Save(domain.QuoteInfo{cur1, cur2, 100, price, "20190412123045"})
	var quoteRep repository.QuoteInfoRepository = repMock
	service := NewOrderService(orderRep, quoteRep)
	req := model.NewRegisterOrderRequest("BTC", "USD")
	service.RegisterOrder(*req)
	if len(orderRep.FindAll()) != 1 {
		t.Fail()
	}

	order := orderRep.FindAll()[0]
	fmt.Println(order)
	if order.CurrencyCode1 != "BTC" {
		t.Fail()
	}

	if order.CurrencyCode2 != "USD" {
		t.Fail()
	}

	amount := order.Amount
	if amount != 100 {
		t.Fail()
	}
}

func TestOrders(t *testing.T) {
	var orderRep repository.OrderRepository = repository.NewOrderRepositoryMock()
	repMock := repository.NewQuoteInfoRepositoryMock()
	cur1 := domain.NewCurrency("BTC")
	cur2 := domain.NewCurrency("USD")
	_, price1 := domain.NewPrice(cur1, cur2, 2900.555)
	repMock.Save(domain.QuoteInfo{cur1, cur2, 100, price1, "20190412123045"})
	_, price2 := domain.NewPrice(cur1, cur2, 1200.095)
	var quoteRep repository.QuoteInfoRepository = repMock
	service := NewOrderService(orderRep, quoteRep)
	req1 := model.NewRegisterOrderRequest("BTC", "USD")
	service.RegisterOrder(*req1)
	repMock.Save(domain.QuoteInfo{cur1, cur2, 100, price2, "20190412123050"})
	req2 := model.NewRegisterOrderRequest("BTC", "USD")
	service.RegisterOrder(*req2)

	orders := service.Orders()
	if len(orders) != 2 {
		t.Fail()
	}

	for _, order := range orders {
		fmt.Println(order)
		if order.CurrencyCode1 != "BTC" {
			t.Fail()
		}

		if order.CurrencyCode2 != "USD" {
			t.Fail()
		}
	}

}
