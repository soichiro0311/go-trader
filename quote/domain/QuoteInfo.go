package domain

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"../enum"
)

type QuoteInfo struct {
	Currency1   Currency
	Currency2   Currency
	Amount      float64
	Price       Price
	Side        enum.SideEnum
	AccuredTime string
}

func NewQuoteInfo(cur1 Currency, cur2 Currency, amount float64, price Price, side enum.SideEnum) ([]error, QuoteInfo) {
	info := new(QuoteInfo)
	info.Currency1 = cur1
	info.Currency2 = cur2
	info.Amount = amount
	info.Price = price
	info.Side = side
	info.AccuredTime = time.Now().Format("20060102150405")
	amountError := info.validateAmount()
	currencyError := info.validateCurrency()

	if amountError == nil && currencyError == nil {
		return nil, *info
	}

	return []error{amountError, currencyError}, QuoteInfo{}
}

func (info QuoteInfo) validateAmount() error {
	if info.Amount <= 0 {
		return errors.New("QuoteInfo amount should be greater than 0")
	}
	return nil
}

func (info QuoteInfo) validateCurrency() error {
	if info.Currency1.CurrencyCode == info.Currency2.CurrencyCode {
		return errors.New("Currency1,Currency2 should not be equal")
	}
	return nil
}

func (info QuoteInfo) Display() string {
	return "currency1:" + info.Currency1.Code() +
		" currency2:" + info.Currency2.Code() +
		" amount:" + fmt.Sprintf("%.3f", info.Amount) +
		" price: " + strconv.FormatInt(info.Price.Value, 10) +
		" accured_datetime: " + info.AccuredTime

}
