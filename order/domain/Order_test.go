package domain

import (
	"fmt"
	"testing"
)

func TestValidateAmount(t *testing.T) {
	currency1 := NewCurrency("BTC")
	currency2 := NewCurrency("USD")
	error, price := NewPrice(currency1, currency2, 100)
	if error != nil {
		t.Fail()
	}
	if error, order := NewOrder(200, currency1, currency2, price); error != nil {
		fmt.Print(order.Currency1)
		t.Fatal()
	}
}

func TestValidateAmountError(t *testing.T) {
	currency1 := NewCurrency("BTC")
	currency2 := NewCurrency("USD")
	error, price := NewPrice(currency1, currency2, 100)
	if error != nil {
		t.Fail()
	}
	if error, _ := NewOrder(-1, currency1, currency2, price); error == nil {
		t.Fatal()
	}
}
