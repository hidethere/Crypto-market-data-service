package service

import (
	"github.com/hidethere/market-data-service/internal/model"
	repo "github.com/hidethere/market-data-service/internal/repository"
)

type TickerService interface {
	GetTicker(symbols []string) ([]model.TickerResponse, error)
}

type tickerService struct {
	repository repo.TickerRepo
}

func NewTickerService(r repo.TickerRepo) TickerService {
	return &tickerService{repository: r}
}

func (t *tickerService) GetTicker(symbols []string) ([]model.TickerResponse, error) {
	resp, err := t.repository.GetTicker(symbols)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
