package domain

import (
	"errors"
	"strconv"
	"time"
)

type BoardInfo struct {
	Currency1   Currency `json:"currency1"`
	Currency2   Currency `json:"currency2"`
	Amount      float64  `json:"amount"`
	Price       Price    `json:"price"`
	AccuredTime string   `json:"accured_time"`
}

func NewBoardInfo(cur1 Currency, cur2 Currency, amount float64, price Price) ([]error, BoardInfo) {
	info := new(BoardInfo)
	info.Currency1 = cur1
	info.Currency2 = cur2
	info.Amount = amount
	info.Price = price
	info.AccuredTime = time.Now().Format("20060102150405")
	amountError := info.validateAmount()
	currencyError := info.validateCurrency()

	if amountError == nil && currencyError == nil {
		return nil, *info
	}

	return []error{amountError, currencyError}, BoardInfo{}
}

func (info BoardInfo) validateAmount() error {
	if info.Amount <= 0 {
		return errors.New("BoardInfo amount should be greater than 0")
	}
	return nil
}

func (info BoardInfo) validateCurrency() error {
	if info.Currency1.CurrencyCode == info.Currency2.CurrencyCode {
		return errors.New("Currency1,Currency2 should not be equal")
	}
	return nil
}

func (info BoardInfo) Display() string {
	return "currency1:" + info.Currency1.Code() +
		" currency2:" + info.Currency2.Code() +
		" amount:" + strconv.FormatFloat(info.Amount, 'e', 5, 64) +
		" price: " + strconv.FormatInt(info.Price.Value, 10) +
		" accured_datetime: " + info.AccuredTime

}
