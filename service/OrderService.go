package service

import (
	"../component"
	"../domain"
	"../model"
	"../repository"
)

type OrderService struct {
	repository repository.OrderRepository
}

func NewOrderService(rep repository.OrderRepository) *OrderService {
	service := new(OrderService)
	service.repository = rep
	return service
}

func (service *OrderService) RegisterOrder(req model.RegisterOrderRequest) {
	priceGenerator := component.GetPriceGenerator()
	price := priceGenerator.GetLatestPrice()
	order := req.ToOrder(*price)
	service.repository.Save(*order)
}

func (service *OrderService) Orders() []*domain.Order {
	return service.repository.FindAll()
}
