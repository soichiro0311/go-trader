package service

import (
	"../domain"
	"../model"
	"../repository"
)

type OrderService struct {
	orderRepository repository.OrderRepository
	boardRepository repository.BoardInfoRepository
}

func NewOrderService(orderRep repository.OrderRepository, boardRep repository.BoardInfoRepository) *OrderService {
	service := new(OrderService)
	service.orderRepository = orderRep
	service.boardRepository = boardRep
	return service
}

func (service *OrderService) RegisterOrder(req model.RegisterOrderRequest) {
	info := service.boardRepository.GetLatestInfoByCurPair(req.CurrencyCode1, req.CurrencyCode2)
	order := req.ToOrder(info)
	service.orderRepository.Save(*order)
}

func (service *OrderService) Orders() []*domain.Order {
	return service.orderRepository.FindAll()
}
