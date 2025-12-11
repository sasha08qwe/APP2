package domain

type Stats struct {
	Currency string  `json:"currency"`
	MinPrice float64 `json:"min_price"`
	MaxPrice float64 `json:"max_price"`
	Change1h float64 `json:"change_1h_percent"`
}
