package data

import (
	"math/big"
	"time"
)

type GetOrdersParam struct {
	MaxId      big.Int `json:"maxId"`
	Count      int     `json:"count"`
	Instrument string  `json:"instrument"`
	Ids        string  `json:"ids"`
}

type Orders struct {
	Orders []Order `json:"orders"`
}

type Order struct {
	Id           big.Int   `json:"id"`
	Instrument   string    `json:"instrument"`
	Units        int       `json:"units"`
	Side         string    `json:"side"`
	Type         string    `json:"type"`
	Time         time.Time `json:"time"`
	Price        float64   `json:"price"`
	TakeProfit   float64   `json:"takeProfit"`
	StopLoss     float64   `json:"stopLoss"`
	Expiry       time.Time `json:"expiry"`
	UpperBound   float64   `json:"upperBound"`
	LowerBound   float64   `json:"lowerBound"`
	TrailingStop float64   `json:"trailingStop"`
}

type PostOrderParam struct {
	Instrument   string    `json:"instrument"`
	Units        int       `json:"units"`
	Side         string    `json:"side"`
	Type         string    `json:"type"`
	Expiry       time.Time `json:"expiry"`
	Price        float64   `json:"price"`
	LowerBound   float64   `json:"lowerBound"`
	UpperBound   float64   `json:"upperBound"`
	StopLoss     float64   `json:"stopLoss"`
	TakeProfit   float64   `json:"takeProfit"`
	TrailingStop float64   `json:"trailingStop"`
}

type PostOrderResult struct {
	Instrument  string      `json:"instrument"`
	Time        time.Time   `json:"time"`
	Price       float64     `json:"price"`
	TradeOpened TradeOpened `json:"tradeOpened"`
	OrderOpened OrderOpened `json:"orderOpened"`
}

type TradeOpened struct {
	Id           big.Int `json:"id"`
	Units        int     `json:"units"`
	Side         string  `json:"side"`
	TakeProfit   float64 `json:"takeProfit"`
	StopLoss     float64 `json:"stopLoss"`
	TrailingStop float64 `json:"trailingStop"`
}

type OrderOpened struct {
	Id           big.Int   `json:"id"`
	Units        int       `json:"units"`
	Side         string    `json:"side"`
	TakeProfit   float64   `json:"takeProfit"`
	StopLoss     float64   `json:"stopLoss"`
	Expiry       time.Time `json:"expiry"`
	UpperBound   float64   `json:"upperBound"`
	LowerBound   float64   `json:"lowerBound"`
	TrailingStop float64   `json:"trailingStop"`
}

type PatchOrderParam struct {
	Units        int       `json:"units"`
	Price        float64   `json:"price"`
	Expiry       time.Time `json:"expiry"`
	LowerBound   float64   `json:"lowerBound"`
	UpperBound   float64   `json:"upperBound"`
	StopLoss     float64   `json:"stopLoss"`
	TakeProfit   float64   `json:"takeProfit"`
	TrailingStop float64   `json:"trailingStop"`
}

type DeleteOrderResult struct {
	Id         big.Int   `json:"id"`
	Instrument string    `json:"instrument"`
	Units      int       `json:"units"`
	Side       string    `json:"side"`
	Price      float64   `json:"price"`
	Time       time.Time `json:"time"`
}