package domain

type QuoteInfo struct {
	Currency1   Currency `json:"currency1"`
	Currency2   Currency `json:"currency2"`
	Amount      float64  `json:"amount"`
	Price       Price    `json:"price"`
	AccuredTime string   `json:"accured_time"`
}
