package api

import (
	"net/http"
	"strings"

	"github.com/hidethere/market-data-service/internal/api/response"
	"github.com/hidethere/market-data-service/internal/application"
)

type TickerHandler struct {
	UserCase *application.GetTickerUseCase
}

func (h *TickerHandler) GetTickerHandler(w http.ResponseWriter, r *http.Request) {
	symbolsParam := r.URL.Query().Get("symbol")
	symbols := strings.Split(symbolsParam, ",")

	resp, err := h.UserCase.GetTicker(symbols)
	if err != nil {
		response.Error(w, err)
	} else {
		response.Success(w, resp)
	}
}
