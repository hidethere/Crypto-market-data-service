package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hidethere/market-data-service/internal/api"
	"github.com/hidethere/market-data-service/internal/application"
	"github.com/hidethere/market-data-service/internal/model"
	repo "github.com/hidethere/market-data-service/internal/repository"
	"github.com/hidethere/market-data-service/internal/service"
)

func main() {
	fmt.Println("CRYPTO PRICE TRACKER MVP")

	tickerRepo := repo.NewTickerRepo()
	tickerService := service.NewTickerService(tickerRepo)
	tickerUseCase := application.NewGetTickerUseCase(tickerService)
	tickerHandler := api.NewTickerHandler(tickerUseCase)

	tickerChan := make(chan model.WSTickerResponse, 100)
	wsTickerRepo := repo.NewWSTickerRepo()
	wsTickerService := service.NewWSTickerService()
	wsTickerUseCase := application.NewWSGetTickerUserCase(wsTickerService)
	wsTickerHandler := api.NewWSTickerHandler(wsTickerUseCase)

	// Start streaming from Binance WS
	go wsTickerRepo.StreamTickers(tickerChan)
	go func() {
		for t := range tickerChan {
			wsTickerService.BroadCast(t)
		}
	}()

	r := mux.NewRouter()

	r.HandleFunc("/api/v1/ticker", tickerHandler.GetTickerHandler).Methods("GET")
	r.HandleFunc("/ws/tickers", wsTickerHandler.WsTickers)

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
