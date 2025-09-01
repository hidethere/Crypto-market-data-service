package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hidethere/market-data-service/internal/api"
	"github.com/hidethere/market-data-service/internal/application"
	repo "github.com/hidethere/market-data-service/internal/repository"
	"github.com/hidethere/market-data-service/internal/service"
)

func main() {
	fmt.Println("CRYPTO PRICE TRACKER MVP")

	tickerRepo := &repo.TickerRepo{}
	tickerService := &service.TickerService{Repository: tickerRepo}
	tickerUserCase := &application.GetTickerUseCase{Service: tickerService}
	tickerHandler := &api.TickerHandler{UserCase: tickerUserCase}

	r := mux.NewRouter()

	r.HandleFunc("/ticker", tickerHandler.GetTickerHandler).Methods("GET")

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
