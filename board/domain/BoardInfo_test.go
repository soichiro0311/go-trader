package domain

import "testing"

func TestNewBoardInfo(t *testing.T) {
	currency1 := NewCurrency("BTC")
	currency2 := NewCurrency("USD")
	_, price := NewPrice(currency1, currency2, 2500)
	_, boardInfo := NewBoardInfo(currency1, currency2, 100, price)
	if boardInfo.Amount != 100 {
		t.Fail()
	}

	if boardInfo.Currency1 != currency1 {
		t.Fail()
	}

	if boardInfo.Currency2 != currency2 {
		t.Fail()
	}
}

func TestNewBoardInfoErrorAmount(t *testing.T) {
	currency1 := NewCurrency("BTC")
	currency2 := NewCurrency("USD")
	_, price := NewPrice(currency1, currency2, 2500)
	errors, _ := NewBoardInfo(currency1, currency2, -500, price)
	if errors[0].Error() != "BoardInfo amount should be greater than 0" {
		t.Fail()
	}
	if errors[1] != nil {
		t.Fail()
	}
}

func TestNewBoardInfoErrorCurrency(t *testing.T) {
	currency1 := NewCurrency("BTC")
	currency2 := NewCurrency("BTC")
	_, price := NewPrice(currency1, currency2, 2500)
	errors, _ := NewBoardInfo(currency1, currency2, 2500, price)
	if errors[0] != nil {
		t.Fail()
	}
	if errors[1].Error() != "Currency1,Currency2 should not be equal" {
		t.Fail()
	}
}
