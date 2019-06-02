package domain

import (
	"errors"
	"math/rand"
	"strconv"
	"time"
)

type Order struct {
	Id            string   `gorm:"primary_key" json:"id"`
	Amount        int64    `json:"amount"`
	Side          string   `json:"side"`
	Currency1     Currency `sql:"-" json:"-"`
	CurrencyCode1 string   `json:"currency1"`
	Currency2     Currency `sql:"-" json:"-"`
	CurrencyCode2 string   `json:"currency2"`
	QuoteIds      []string `sql:"-" json:"quoteIds"`
	OrderDatetime string   `json:"order_datetime"`
}

func NewOrder(amount int64, currency1 Currency, currency2 Currency, side string) (error, *Order) {
	order := new(Order)

	rand.Seed(time.Now().UnixNano())
	order.Id = strconv.Itoa(rand.Int())

	order.Amount = amount
	err := order.validateAmount()

	order.Currency1 = currency1
	order.CurrencyCode1 = currency1.CurrencyCode
	order.Currency2 = currency2
	order.CurrencyCode2 = currency2.CurrencyCode
	order.Side = side

	order.OrderDatetime = time.Now().Format("20060102150405")
	return err, order
}

func (order *Order) validateAmount() error {
	if order.Amount <= 0 {
		return errors.New("Order amount should be greater than 0")
	}
	return nil
}
