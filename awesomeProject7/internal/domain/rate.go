package domain

import "time"

type Rate struct {
	Currency  string    `json:"currency"`
	Price     float64   `json:"price"`
	Timestamp time.Time `json:"timestamp"`
}
