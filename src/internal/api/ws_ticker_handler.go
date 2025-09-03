package api

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/hidethere/market-data-service/internal/application"
)

type wsTickerHandler struct {
	wsGetTickerUseCase application.WSGetTickerUseCase
}

type SubscribeMessage struct {
	Subscribe []string `json:"subscribe"`
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func NewWSTickerHandler(wsGetTickerUseCase application.WSGetTickerUseCase) *wsTickerHandler {
	return &wsTickerHandler{wsGetTickerUseCase: wsGetTickerUseCase}

}

func (h *wsTickerHandler) WsTickers(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Websocket upgrade error:", err)
		return
	}
	defer conn.Close()

	client := h.wsGetTickerUseCase.SubscribeClient()
	defer h.wsGetTickerUseCase.UnsubscribeClient(client)

	// Ping Pong HeartBeat
	conn.SetReadDeadline(time.Now().Add(60 * time.Second)) // conn will timeout if no read after 60s
	conn.SetPongHandler(func(appData string) error {       // Extend read deadline
		conn.SetReadDeadline(time.Now().Add(60 * time.Second))
		return nil
	})
	ticker := time.NewTicker(30 * time.Second) // send ping every 30s

	done := make(chan struct{})

	go func() {
		for {
			select {
			case <-ticker.C:
				if err := conn.WriteMessage(websocket.PingMessage, nil); err != nil {
					close(done)
					return
				}
			case <-done:
				return

			}
		}
	}()

	// Handle incoming subscription messags
	go func() {
		for {
			_, msg, err := conn.ReadMessage()
			if err != nil {
				return
			}
			var subMsg SubscribeMessage
			if err := json.Unmarshal(msg, &subMsg); err != nil {
				log.Println("Invalid subscription message:", err)
				continue
			}
			h.wsGetTickerUseCase.UpdateSubscription(client, subMsg.Subscribe)
		}
	}()

	for ticker := range client.Ch {
		if err := conn.WriteJSON(ticker); err != nil {
			return
		}
	}

}
