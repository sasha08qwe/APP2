package ports

import "project/internal/domain"

type RateProvider interface {
	FetchRate(currency string) (*domain.Rate, error)
}
