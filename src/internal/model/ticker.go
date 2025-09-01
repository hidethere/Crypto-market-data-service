package model

import (
	"encoding/json"
)

type TickerResponse struct {
	Symbol    string      `json:"symbol"`
	Bid       json.Number `json:"bidPrice"`
	Ask       json.Number `json:"askPrice"`
	Last      json.Number `json:"lastPrice"`
	High      json.Number `json:"highPrice"`
	Volume    json.Number `json:"volume"`
	Low       json.Number `json:"lowPrice"`
	OpenTime  json.Number `json:"openTime"`
	CloseTime json.Number `json:"closeTime"`
}

type Ticker struct {
	Symbol    string
	Bid       float64
	Ask       float64
	Last      float64
	High      float64
	Volume    float64
	Low       float64
	Spread    float64
	OpenTime  float64
	CloseTime float64
}
