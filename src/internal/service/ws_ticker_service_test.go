package service_test

import (
	"testing"

	"github.com/hidethere/market-data-service/internal/model"
	"github.com/hidethere/market-data-service/internal/service"
)

func TestBroadcast_FilteredSubscription(t *testing.T) {
	s := service.NewWSTickerService()
	c := s.Subscribe()

	got := make(chan model.WSTickerResponse, 1)
	c.Ch = got

	s.UpdateSubscription(c, []string{"BTCUSDT"})

	// Broadcast ETH ticker (should NOT be sent)
	s.BroadCast(model.WSTickerResponse{Symbol: "ETHUSDT"})
	select {
	case <-got:
		t.Error("unexpected ETHUSDT receivedd")
	default:
	}

	// Broadcast BTC ticker (should be sent)
	s.BroadCast(model.WSTickerResponse{Symbol: "BTCUSDT"})
	select {
	case msg := <-got:
		if msg.Symbol != "BTCUSDT" {
			t.Errorf("expected BTCUSDT, got %s", msg.Symbol)

		}

	default:
		t.Error("expected BTCUSDT but got nothing")
	}

}

func TestBroadCast_DefaultAll(t *testing.T) {
	s := service.NewWSTickerService()
	c := s.Subscribe()

	got := make(chan model.WSTickerResponse, 1)
	c.Ch = got

	s.BroadCast(model.WSTickerResponse{Symbol: "DOGEUSDT"})

	select {
	case msg := <-got:
		if msg.Symbol != "DOGEUSDT" {
			t.Errorf("expected DOGEUSDT, got %s", msg.Symbol)

		}
	default:
		t.Error("expected DOGEUSDT but got nothing")
	}
}
