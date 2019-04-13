package domain

import (
	"fmt"
	"testing"
)

func TestNewPriceError(t *testing.T) {
	currency1 := NewCurrency("BTC")
	currency2 := NewCurrency("BTC")
	error, price := NewPrice(*currency1, *currency2, 100)
	if error == nil {
		fmt.Print(price)
		t.Fail()
	}
}

func TestNewPrice(t *testing.T) {
	currency1 := NewCurrency("BTC")
	currency2 := NewCurrency("USD")
	error, price := NewPrice(*currency1, *currency2, 100)
	if error != nil {
		fmt.Print(price)
		t.Fail()
	}
}

func TestConvertToCurrency2(t *testing.T) {
	currency1 := NewCurrency("BTC")
	currency2 := NewCurrency("USD")
	error, price := NewPrice(*currency1, *currency2, 100)
	quantity := price.converToCurrency2(10)
	if quantity != 1000 {
		fmt.Print(error)
		t.Fail()
	}
}
