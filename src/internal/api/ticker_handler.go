package api

import (
	"net/http"
	"strings"

	"github.com/hidethere/market-data-service/internal/api/response"
	"github.com/hidethere/market-data-service/internal/application"
)

type tickerHandler struct {
	getTickerUseCase application.GetTickerUseCase
}

func NewTickerHandler(getUserUC application.GetTickerUseCase) *tickerHandler {
	return &tickerHandler{getTickerUseCase: getUserUC}
}

func (h *tickerHandler) GetTickerHandler(w http.ResponseWriter, r *http.Request) {
	symbolsParam := r.URL.Query().Get("symbol")
	symbols := strings.Split(symbolsParam, ",")

	resp, err := h.getTickerUseCase.GetTicker(symbols)
	if err != nil {
		response.Error(w, err)
	} else {
		response.Success(w, resp)
	}
}
