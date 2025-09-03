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
