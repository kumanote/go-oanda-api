package data

import (
	"math/big"
	"time"
)

type GetTradesParam struct {
	MaxId      big.Int `json:"maxId"`
	Count      int     `json:"count"`
	Instrument string  `json:"instrument"`
	Ids        string  `json:"ids"`
}

type GetTrades struct {
	Trades []Trade `json:"trades"`
}

type Trade struct {
	Id             big.Int   `json:"id"`
	Units          int       `json:"units"`
	Side           string    `json:"side"`
	Instrument     string    `json:"instrument"`
	Time           time.Time `json:"time"`
	Price          float64   `json:"price"`
	TakeProfit     float64   `json:"takeProfit"`
	StopLoss       float64   `json:"stopLoss"`
	TrailingStop   int       `json:"trailingStop"`
	TrailingAmount float64   `json:"trailingAmount"`
}

type PatchTradeParam struct {
	StopLoss     float64 `json:"stopLoss"`
	TakeProfit   float64 `json:"takeProfit"`
	TrailingStop int     `json:"trailingStop"`
}

type CloseTradeResult struct {
	Id         big.Int   `json:"id"`
	Price      float64   `json:"price"`
	Instrument string    `json:"instrument"`
	Profit     float64   `json:"profit"`
	Side       string    `json:"side"`
	Time       time.Time `json:"time"`
}
