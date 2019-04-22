package domain

import (
	"errors"
	"math/rand"
	"strconv"
	"time"
)

type Order struct {
	Id            string   `gorm:"primary_key" json:"id"`
	Amount        float64  `json:"amount"`
	Currency1     Currency `sql:"-" json:"-"`
	CurrencyCode1 string   `json:"currency1"`
	Currency2     Currency `sql:"-" json:"-"`
	CurrencyCode2 string   `json:"currency2"`
	Quantity      float64  `json:"quantity"`
	OrderDatetime string   `json:"order_datetime"`
}

func NewOrder(amount float64, currency1 Currency, currency2 Currency, price Price) (error, *Order) {
	order := new(Order)

	rand.Seed(time.Now().UnixNano())
	order.Id = strconv.Itoa(rand.Int())

	order.Amount = amount
	err := order.validateAmount()
	order.Quantity = price.converToCurrency2(amount)

	order.Currency1 = currency1
	order.CurrencyCode1 = currency1.CurrencyCode
	order.Currency2 = currency2
	order.CurrencyCode2 = currency2.CurrencyCode

	order.OrderDatetime = time.Now().Format("20060102150405")
	return err, order
}

func (order *Order) validateAmount() error {
	if order.Amount <= 0 {
		return errors.New("Order amount should be greater than 0")
	}
	return nil
}

func (order *Order) ToString() string {
	return "id:" + order.Id +
		" currency1:" + order.Currency1.Code() +
		" currency2:" + order.Currency2.Code() +
		" amount:" + strconv.FormatFloat(order.Amount, 'e', 10, 64) +
		" quantity:" + strconv.FormatFloat(order.Quantity, 'e', 10, 64) +
		" orderDatetime:" + order.OrderDatetime
}
