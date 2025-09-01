package application

import (
	"github.com/hidethere/market-data-service/internal/model"
	service "github.com/hidethere/market-data-service/internal/service"
)

type UseCase interface {
	GetTicker(symbols []string) ([]model.TickerResponse, error)
}

type GetTickerUseCase struct {
	Service *service.TickerService
}

func (u *GetTickerUseCase) GetTicker(symbols []string) ([]model.TickerResponse, error) {
	resp, err := u.Service.GetTicker(symbols)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
