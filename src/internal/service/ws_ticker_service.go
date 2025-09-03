package service

import (
	"sync"

	"github.com/hidethere/market-data-service/internal/model"
)

type WSTickerService interface {
	BroadCast(t model.WSTickerResponse)
	Subscribe() *Client
	Unsubscribe(c *Client)
	UpdateSubscription(c *Client, symbols []string)
}

type Client struct {
	Ch      chan model.WSTickerResponse
	Symbols map[string]struct{}
	Filter  bool
}

type wsTickerService struct {
	clients map[*Client]struct{}
	mutex   sync.RWMutex
}

func NewWSTickerService() WSTickerService {
	return &wsTickerService{
		clients: make(map[*Client]struct{}),
	}
}

func (s *wsTickerService) BroadCast(t model.WSTickerResponse) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	for c := range s.clients {
		if !c.Filter || containsSymbol(c, t.Symbol) {
			select {
			case c.Ch <- t:
			default:
			}
		}
	}
}

func (s *wsTickerService) Subscribe() *Client {
	c := &Client{
		Ch:      make(chan model.WSTickerResponse, 100), //Buffered
		Symbols: make(map[string]struct{}),
		Filter:  false,
	}
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.clients[c] = struct{}{}
	return c

}

func (s *wsTickerService) Unsubscribe(c *Client) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	delete(s.clients, c)
	close(c.Ch)
}

func (s *wsTickerService) UpdateSubscription(c *Client, symbols []string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	c.Filter = true
	c.Symbols = make(map[string]struct{})

	if len(symbols) == 0 {
		c.Filter = false
		c.Symbols = make(map[string]struct{})
	}

	for _, symb := range symbols {
		c.Symbols[symb] = struct{}{}
	}
}

func containsSymbol(c *Client, symbol string) bool {
	_, ok := c.Symbols[symbol]
	return ok
}
