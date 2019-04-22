package domain

type Currency struct {
	CurrencyCode string `json:"currency_code"`
}

func NewCurrency(currencyCode string) Currency {
	currency := new(Currency)
	currency.CurrencyCode = currencyCode
	return *currency
}

func (currency *Currency) Code() string {
	return currency.CurrencyCode
}
