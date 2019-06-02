package dto

type QuoteInfoDto struct {
	Currency1   string `json:"currency1"`
	Currency2   string `json:"currency2"`
	Amount      int64  `json:"amount"`
	Price       int64  `json:"price"`
	Side        string `json:"side"`
	AccuredTime string `json:"accured_time"`
}
