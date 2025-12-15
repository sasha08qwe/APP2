package entities

type Stats struct {
	currency Currency
	min      Rate
	max      Rate
	change1h int64
}

func NewStats(currency Currency, min Rate, max Rate, change1h int64) Stats {
	return Stats{
		currency: currency,
		min:      min,
		max:      max,
		change1h: change1h,
	}
}

func (s Stats) Currency() Currency {
	return s.currency
}

func (s Stats) Min() Rate {
	return s.min
}

func (s Stats) Max() Rate {
	return s.max
}

func (s Stats) Change1h() int64 {
	return s.change1h
}
