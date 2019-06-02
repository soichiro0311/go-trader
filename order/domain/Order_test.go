package domain

import (
	"fmt"
	"testing"
)

func TestValidateAmount(t *testing.T) {
	currency1 := NewCurrency("BTC")
	currency2 := NewCurrency("USD")
	if error, order := NewOrder(200, currency1, currency2, "BUY"); error != nil {
		fmt.Print(order)
		t.Fatal()
	}
}

func TestValidateAmountError(t *testing.T) {
	currency1 := NewCurrency("BTC")
	currency2 := NewCurrency("USD")
	if error, _ := NewOrder(-1, currency1, currency2, "SELL"); error == nil {
		t.Fatal()
	}
}
