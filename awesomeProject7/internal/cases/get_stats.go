package cases

import (
	"context"
	"time"

	"awesomeProject7/internal/entities"
)

type StatsRepository interface {
	MinForDay(ctx context.Context, currency entities.Currency, day time.Time) (entities.Rate, error)
	MaxForDay(ctx context.Context, currency entities.Currency, day time.Time) (entities.Rate, error)
	ForPeriod(ctx context.Context, currency entities.Currency, from, to time.Time) ([]entities.Rate, error)
}

type GetStatsCase struct {
	repo StatsRepository
}

func NewGetStatsCase(repo StatsRepository) GetStatsCase {
	return GetStatsCase{repo: repo}
}

func (c GetStatsCase) Execute(
	ctx context.Context,
	currency entities.Currency,
) (entities.Stats, error) {

	today := time.Now().Truncate(24 * time.Hour)

	min, err := c.repo.MinForDay(ctx, currency, today)
	if err != nil {
		return entities.Stats{}, err
	}

	max, err := c.repo.MaxForDay(ctx, currency, today)
	if err != nil {
		return entities.Stats{}, err
	}

	rates, err := c.repo.ForPeriod(
		ctx,
		currency,
		time.Now().Add(-1*time.Hour),
		time.Now(),
	)
	if err != nil || len(rates) < 2 {
		return entities.NewStats(currency, min, max, 0), nil
	}

	change := rates[len(rates)-1].Price() - rates[0].Price()

	return entities.NewStats(currency, min, max, change), nil
}
