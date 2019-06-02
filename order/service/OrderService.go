package service

import (
	"fmt"

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
	info := service.quoteRepository.GetLatestInfoByCurPair(req.CurrencyCode1, req.CurrencyCode2, req.Side)
	order := req.ToOrder()
	service.orderRepository.Save(*order)
}

func (service *OrderService) Orders() []*domain.Order {
	return service.orderRepository.FindAll()
}

