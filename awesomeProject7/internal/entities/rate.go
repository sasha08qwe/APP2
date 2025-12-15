package entities

import "time"

type Currency string

const (
	BTC Currency = "BTC"
	ETH Currency = "ETH"
)

type Rate struct {
	currency  Currency
	price     int64
	timestamp time.Time
}

func NewRate(currency Currency, price int64, timestamp time.Time) Rate {
	return Rate{
		currency:  currency,
		price:     price,
		timestamp: timestamp,
	}
}

func (r Rate) Currency() Currency {
	return r.currency
}

func (r Rate) Price() int64 {
	return r.price
}

func (r Rate) Timestamp() time.Time {
	return r.timestamp
}
