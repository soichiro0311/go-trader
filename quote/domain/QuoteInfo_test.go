package domain

import "testing"

func TestNewQuoteInfo(t *testing.T) {
	currency1 := NewCurrency("BTC")
	currency2 := NewCurrency("USD")
	_, price := NewPrice(currency1, currency2, 2500)
	_, quoteInfo := NewQuoteInfo(currency1, currency2, 100, price)
	if quoteInfo.Amount != 100 {
		t.Fail()
	}

	if quoteInfo.Currency1 != currency1 {
		t.Fail()
	}

	if quoteInfo.Currency2 != currency2 {
		t.Fail()
	}
}

func TestNewQuoteInfoErrorAmount(t *testing.T) {
	currency1 := NewCurrency("BTC")
	currency2 := NewCurrency("USD")
	_, price := NewPrice(currency1, currency2, 2500)
	errors, _ := NewQuoteInfo(currency1, currency2, -500, price)
	if errors[0].Error() != "QuoteInfo amount should be greater than 0" {
		t.Fail()
	}
	if errors[1] != nil {
		t.Fail()
	}
}

func TestNewQuoteInfoErrorCurrency(t *testing.T) {
	currency1 := NewCurrency("BTC")
	currency2 := NewCurrency("BTC")
	_, price := NewPrice(currency1, currency2, 2500)
	errors, _ := NewQuoteInfo(currency1, currency2, 2500, price)
	if errors[0] != nil {
		t.Fail()
	}
	if errors[1].Error() != "Currency1,Currency2 should not be equal" {
		t.Fail()
	}
}
