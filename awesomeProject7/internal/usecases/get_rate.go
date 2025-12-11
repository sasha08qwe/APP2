package usecases

import (
	"project/internal/domain"
	"project/internal/ports"
)

type GetRateUseCase struct {
	repo ports.RateRepository
}

func NewGetRateUseCase(repo ports.RateRepository) *GetRateUseCase {
	return &GetRateUseCase{repo: repo}
}

func (uc *GetRateUseCase) Execute(currency string) (*domain.Rate, error) {
	return uc.repo.GetLatest(currency)
}
