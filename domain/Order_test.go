package domain

import (
	"fmt"
	"math/big"
	"testing"
)

func TestValidateAmount(t *testing.T) {
	price := NewPrice("BTC", "USD", *big.NewFloat(100))
	if error, order := NewOrder(*big.NewInt(200), "BTC", "USD", *price); error != nil {
		fmt.Printf(order.Currency1)
		t.Fatal()
	}
}

func TestValidateAmountError(t *testing.T) {
	price := NewPrice("BTC", "USD", *big.NewFloat(100))
	if error, order := NewOrder(*big.NewInt(-1), "BTC", "USD", *price); error == nil {
		fmt.Printf(order.Amount.String())
		t.Fatal()
	}
}
