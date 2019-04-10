package domain

import (
	"errors"
	"math/rand"
	"time"
)

type Order struct {
	Id       int
	Amount   int64
	Currency string
}

func NewOrder(amount int64, currency string) (error, *Order) {
	order := new(Order)

	order.Amount = amount
	err := order.validateAmount(amount)

	order.Currency = currency

	rand.Seed(time.Now().UnixNano())
	order.Id = rand.Int()
	return err, order
}

func (order *Order) validateAmount(amount int64) error {
	if order.Amount <= 0 {
		return errors.New("Order amount should be greater than 0")
	}
	return nil
}
