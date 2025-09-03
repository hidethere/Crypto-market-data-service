package application

import (
	"github.com/hidethere/market-data-service/internal/model"
	service "github.com/hidethere/market-data-service/internal/service"
)

type GetTickerUseCase interface {
	GetTicker(symbols []string) ([]model.TickerResponse, error)
}

type getTickerUseCase struct {
	service service.TickerService
}

func NewGetTickerUseCase(s service.TickerService) GetTickerUseCase {
	return &getTickerUseCase{service: s}
}

func (u *getTickerUseCase) GetTicker(symbols []string) ([]model.TickerResponse, error) {
	resp, err := u.service.GetTicker(symbols)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
