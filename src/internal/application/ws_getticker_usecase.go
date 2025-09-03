package application

import (
	"github.com/hidethere/market-data-service/internal/service"
)

type WSGetTickerUseCase interface {
	SubscribeClient() *service.Client
	UnsubscribeClient(c *service.Client)
	UpdateSubscription(c *service.Client, symbols []string)
}

type wsGetTickerUseCase struct {
	service service.WSTickerService
}

func NewWSGetTickerUserCase(s service.WSTickerService) WSGetTickerUseCase {
	return &wsGetTickerUseCase{service: s}
}

func (uc *wsGetTickerUseCase) SubscribeClient() *service.Client {
	return uc.service.Subscribe()
}

func (uc *wsGetTickerUseCase) UnsubscribeClient(c *service.Client) {
	uc.service.Unsubscribe(c)
}

func (uc *wsGetTickerUseCase) UpdateSubscription(c *service.Client, symbols []string) {
	uc.service.UpdateSubscription(c, symbols)
}
