package model

import "encoding/json"

type WSTickerResponse struct {
	Symbol    string      `json:"s"`
	Bid       json.Number `json:"b"`
	Ask       json.Number `json:"a"`
	Last      json.Number `json:"c"`
	High      json.Number `json:"h"`
	Volume    json.Number `json:"v"`
	Low       json.Number `json:"l"`
	OpenTime  json.Number `json:"O"`
	CloseTime json.Number `json:"C"`
}
