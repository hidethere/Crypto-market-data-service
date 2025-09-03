package repo

import (
	"encoding/json"
	"log"
	"time"

	"github.com/gorilla/websocket"
	"github.com/hidethere/market-data-service/internal/model"
)

type WSTickerRepo interface {
	StreamTickers(out chan<- model.WSTickerResponse)
}

type wsTickerRepo struct {
	url string
}

func NewWSTickerRepo() WSTickerRepo {
	return &wsTickerRepo{
		url: "wss://stream.binance.com:9443/ws/!ticker@arr",
	}
}

// StreamTickers sends live tickers to the service channel
func (r *wsTickerRepo) StreamTickers(out chan<- model.WSTickerResponse) {
	for {
		conn, _, err := websocket.DefaultDialer.Dial(r.url, nil)
		if err != nil {
			log.Println("Websocket dial error:", err)
			time.Sleep(5 * time.Second)
			continue // retry connecting
		}
		log.Println("Connected to Binance WS")

		// read loop
		for {
			_, msg, err := conn.ReadMessage()
			if err != nil {
				log.Println("Websocket read error:", err)
				conn.Close()
				break // exit inner loop reconnect
			}

			var tickers []model.WSTickerResponse
			if err := json.Unmarshal(msg, &tickers); err != nil {
				log.Println("JSON unmarshal error:", err)
				continue
			}

			for _, ticker := range tickers {
				out <- ticker
			}
		}

		log.Println("Reconnecting in 5s...")
		time.Sleep(5 * time.Second)
	}
}
