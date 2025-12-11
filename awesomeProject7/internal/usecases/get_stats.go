package usecases

import (
	"time"

	"project/internal/domain"
	"project/internal/ports"
)

type GetStatsUseCase struct {
	repo ports.RateRepository
}

func NewGetStatsUseCase(repo ports.RateRepository) *GetStatsUseCase {
	return &GetStatsUseCase{repo: repo}
}

func (uc *GetStatsUseCase) Execute(currency string) (*domain.Stats, error) {
	today := time.Now().Truncate(24 * time.Hour)

	min, max, err := uc.repo.GetMinMaxForDay(currency, today)
	if err != nil {
		return nil, err
	}

	// change % for last hour
	end := time.Now()
	start := end.Add(-1 * time.Hour)

	records, err := uc.repo.GetForPeriod(currency, start, end)
	if err != nil || len(records) == 0 {
		return &domain.Stats{Currency: currency, MinPrice: min, MaxPrice: max, Change1h: 0}, nil
	}

	first := records[0].Price
	last := records[len(records)-1].Price

	change := ((last - first) / first) * 100

	return &domain.Stats{
		Currency: currency,
		MinPrice: min,
		MaxPrice: max,
		Change1h: change,
	}, nil
}
