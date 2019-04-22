package service

import (
	"../domain"
	"../model"
	"../repository"
)

type OrderService struct {
	orderRepository repository.OrderRepository
	quoteRepository repository.QuoteInfoRepository
}

func NewOrderService(orderRep repository.OrderRepository, quoteRep repository.QuoteInfoRepository) *OrderService {
	service := new(OrderService)
	service.orderRepository = orderRep
	service.quoteRepository = quoteRep
	return service
}

func (service *OrderService) RegisterOrder(req model.RegisterOrderRequest) {
	info := service.quoteRepository.GetLatestInfoByCurPair(req.CurrencyCode1, req.CurrencyCode2)
	order := req.ToOrder(info)
	service.orderRepository.Save(*order)
}

func (service *OrderService) Orders() []*domain.Order {
	return service.orderRepository.FindAll()
}
