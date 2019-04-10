package domain

import (
	"fmt"
	"testing"

	"../../domain"
)

func TestValidateAmount(t *testing.T) {
	if error, order := domain.NewOrder(200, "BTC"); error != nil {
		fmt.Printf(order.Currency)
		t.Fatal()
	}
}

func TestValidateAmountError(t *testing.T) {
	if error, order := domain.NewOrder(-1, "BTC"); error == nil {
		fmt.Printf(order.Currency)
		t.Fatal()
	}
}
