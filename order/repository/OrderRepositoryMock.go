package repository

import "../domain"

type OrderRepositoryMock struct {
	inmemoryMap map[string]*domain.Order
}

func NewOrderRepositoryMock() *OrderRepositoryMock {
	mock := new(OrderRepositoryMock)
	mock.inmemoryMap = map[string]*domain.Order{}
	return mock
}

func (repository *OrderRepositoryMock) Save(order domain.Order) {
	repository.inmemoryMap[order.Id] = &order

}

func (repository *OrderRepositoryMock) FindAll() []*domain.Order {
	values := []*domain.Order{}
	for _, v := range repository.inmemoryMap {
		values = append(values, v)
	}
	return values
}
