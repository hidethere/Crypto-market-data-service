package repo

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/hidethere/market-data-service/internal/api/response"
	"github.com/hidethere/market-data-service/internal/model"
)

type TickerRepo interface {
	GetTicker(symbols []string) ([]model.TickerResponse, error)
}

type tickerRepo struct {
}

func NewTickerRepo() TickerRepo {
	return &tickerRepo{}
}

func (t *tickerRepo) GetTicker(symbols []string) ([]model.TickerResponse, error) {
	var tickerResp []model.TickerResponse

	for _, s := range symbols {
		url := fmt.Sprintf("https://api.binance.com/api/v3/ticker/24hr?symbol=%s", s)

		resp, err := http.Get(url)
		if err != nil {
			log.Println("Error", err)
			return nil, err
		}
		if resp.StatusCode == http.StatusBadRequest {
			return nil, response.ErrSymbolNotFound
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println("Error reading response body:", err)
			return nil, err
		}

		var ticker model.TickerResponse
		if err := json.Unmarshal(body, &ticker); err != nil {
			log.Println("Error parsing:", err)
			return nil, err
		}
		log.Println(ticker)
		tickerResp = append(tickerResp, ticker)
	}
	return tickerResp, nil
}
