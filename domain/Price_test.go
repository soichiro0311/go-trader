package domain

import (
	"math/big"
	"testing"
)

func TestConvertToCurrency2(t *testing.T) {
	price := NewPrice("BTC", "USD", *big.NewFloat(100))
	quantity := price.converToCurrency2(*big.NewInt(10))
	if quantity.Cmp(big.NewFloat(1000)) != 0 {
		t.Fail()
	}
}
