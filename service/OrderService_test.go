package service

import (
	"testing"

	"../component"
	"../model"
	"../repository"
)

func TestRegisterOrder(t *testing.T) {
	component.GetPriceGenerator().Generate("BTC", "USD")
	var rep repository.OrderRepository = repository.NewOrderRepositoryMock()
	service := NewOrderService(rep)
	req := model.NewRegisterOrderRequest("BTC", "USD", 100)
	service.RegisterOrder(*req)
	if len(rep.FindAll()) != 1 {
		t.Fail()
	}

	order := rep.FindAll()[0]
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
	component.GetPriceGenerator().Generate("BTC", "USD")
	var rep repository.OrderRepository = repository.NewOrderRepositoryMock()
	service := NewOrderService(rep)
	req1 := model.NewRegisterOrderRequest("BTC", "USD", 100)
	service.RegisterOrder(*req1)
	req2 := model.NewRegisterOrderRequest("BTC", "USD", 25)
	service.RegisterOrder(*req2)

	orders := service.Orders()
	if len(orders) != 2 {
		t.Fail()
	}

	for _, order := range orders {
		if order.CurrencyCode1 != "BTC" {
			t.Fail()
		}

		if order.CurrencyCode2 != "USD" {
			t.Fail()
		}
	}

}
