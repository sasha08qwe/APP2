package main

import "awesomeProject7/internal/usecases"

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	db := postgres.NewPostgresConn(cfg)

	repo := postgres.NewPostgresRateRepository(db)
	provider := coingecko.NewCoinGeckoProvider(cfg)

	getRateUseCase := usecases.NewGetRateUseCase(repo)
	getStatsUseCase := usecases.NewGetStatsUseCase(repo)

	httpServer := http.NewServer(getRateUseCase, getStatsUseCase)
	httpServer.Start()
}
