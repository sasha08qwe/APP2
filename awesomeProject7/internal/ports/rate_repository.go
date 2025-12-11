package ports

import (
	"time"

	"project/internal/domain"
)

type RateRepository interface {
	Save(rate *domain.Rate) error
	GetLatest(currency string) (*domain.Rate, error)
	GetForPeriod(currency string, start, end time.Time) ([]domain.Rate, error)
	GetMinMaxForDay(currency string, day time.Time) (min float64, max float64, err error)
}
