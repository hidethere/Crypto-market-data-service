package service

import (
	"github.com/hidethere/market-data-service/internal/model"
	repo "github.com/hidethere/market-data-service/internal/repository"
)

type Service interface {
	GetTicker(symbols []string) ([]model.TickerResponse, error)
}

type TickerService struct {
	Repository *repo.TickerRepo
}

func (t *TickerService) GetTicker(symbols []string) ([]model.TickerResponse, error) {
	resp, err := t.Repository.GetTicker(symbols)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
