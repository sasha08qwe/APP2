package cases

import (
	"context"

	"awesomeProject7/internal/entities"
)

type RateRepository interface {
	Latest(ctx context.Context, currency entities.Currency) (entities.Rate, error)
}

type GetRateCase struct {
	repo RateRepository
}

func NewGetRateCase(repo RateRepository) GetRateCase {
	return GetRateCase{repo: repo}
}

func (c GetRateCase) Execute(
	ctx context.Context,
	currency entities.Currency,
) (entities.Rate, error) {
	return c.repo.Latest(ctx, currency)
}
