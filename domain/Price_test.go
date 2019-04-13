package domain

import (
	"testing"
)

func TestNewPriceError(t *testing.T) {
	currency1 := NewCurrency("BTC")
	currency2 := NewCurrency("BTC")
	error, _ := NewPrice(*currency1, *currency2, 100)
	if error == nil {
		t.Fail()
	}
}

func TestNewPrice(t *testing.T) {
	currency1 := NewCurrency("BTC")
	currency2 := NewCurrency("USD")
	error, _ := NewPrice(*currency1, *currency2, 100)
	if error != nil {
		t.Fail()
	}
}

func TestConvertToCurrency2(t *testing.T) {
	currency1 := NewCurrency("BTC")
	currency2 := NewCurrency("USD")
	_, price := NewPrice(*currency1, *currency2, 100)
	quantity := price.converToCurrency2(10)
	if quantity != 1000 {
		t.Fail()
	}
}
