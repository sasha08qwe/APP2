package cases

import (
	"context"

	"awesomeProject7/internal/entities"
)

type RateProvider interface {
	FetchRates(
		ctx context.Context,
		currencies []entities.Currency,
	) ([]entities.Rate, error)
}

type RateSaver interface {
	Save(ctx context.Context, rate entities.Rate) error
}

type UpdateRatesCase struct {
	provider RateProvider
	saver    RateSaver
}

func NewUpdateRatesCase(
	provider RateProvider,
	saver RateSaver,
) UpdateRatesCase {
	return UpdateRatesCase{
		provider: provider,
		saver:    saver,
	}
}

func (c UpdateRatesCase) Execute(ctx context.Context) error {
	rates, err := c.provider.FetchRates(ctx, []entities.Currency{
		entities.BTC,
		entities.ETH,
	})
	if err != nil {
		return err
	}

	for _, rate := range rates {
		if err := c.saver.Save(ctx, rate); err != nil {
			return err
		}
	}

	return nil
}
