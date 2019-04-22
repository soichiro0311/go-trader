package domain

import "testing"

func TestNewPrice(t *testing.T) {
	currency1 := NewCurrency("BTC")
	currency2 := NewCurrency("USD")
	_, price := NewPrice(currency1, currency2, 2500)
	if price.Value != 2500 {
		t.Fail()
	}

	if price.Currency1 != currency1 {
		t.Fail()
	}

	if price.Currency2 != currency2 {
		t.Fail()
	}
}

func TestNewPriceErrorValue(t *testing.T) {
	currency1 := NewCurrency("BTC")
	currency2 := NewCurrency("USD")
	errors, _ := NewPrice(currency1, currency2, -500)
	if errors[0].Error() != "Price value should be greater than 0" {
		t.Fail()
	}
	if errors[1] != nil {
		t.Fail()
	}
}

func TestNewPriceErrorCurrency(t *testing.T) {
	currency1 := NewCurrency("BTC")
	currency2 := NewCurrency("BTC")
	errors, _ := NewPrice(currency1, currency2, 2500)
	if errors[0] != nil {
		t.Fail()
	}
	if errors[1].Error() != "Currency1,Currency2 should not be equal" {
		t.Fail()
	}
}
