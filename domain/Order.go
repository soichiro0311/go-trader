package domain

import (
	"errors"
	"math/big"
	"math/rand"
	"time"
)

type Order struct {
	Id        int
	Amount    big.Int
	Currency1 string
	Currency2 string
	Price     Price
	Quantity  big.Float
}

func NewOrder(amount big.Int, currency1 string, currency2 string, price Price) (error, *Order) {
	order := new(Order)
	order.Price = price
	order.Quantity = price.converToCurrency2(amount)

	order.Amount = amount
	err := order.validateAmount(amount)

	order.Currency1 = currency1
	order.Currency1 = currency2

	rand.Seed(time.Now().UnixNano())
	order.Id = rand.Int()
	return err, order
}

func (order *Order) validateAmount(amount big.Int) error {
	if order.Amount.Cmp(big.NewInt(0)) <= 0 {
		return errors.New("Order amount should be greater than 0")
	}
	return nil
}
