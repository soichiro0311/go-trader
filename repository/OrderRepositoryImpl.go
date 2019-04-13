package repository

import (
	"fmt"

	"../db"
	"../domain"
	"github.com/jinzhu/gorm"
)

type orderRepositoryImpl struct {
	db *gorm.DB
}

func NewOrderRepositoryImpl() *orderRepositoryImpl {
	repository := new(orderRepositoryImpl)
	repository.db = db.GetDB()
	return repository
}

func (repository *orderRepositoryImpl) Save(order domain.Order) {
	fmt.Print(order)
	repository.db.Save(&order)
}

func (repository *orderRepositoryImpl) FindAll() []*domain.Order {
	result := []*domain.Order{}
	repository.db.Find(result)
	return result
}
