package repository

import "../domain"

type OrderRepository interface {
	Save(domain.Order)
	FindAll() []*domain.Order
}
