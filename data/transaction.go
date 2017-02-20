package data

import (
	"time"
)

type GetTransactionsParam struct {
	MaxId      int64  `json:"maxId"`
	MinId      int64  `json:"minId"`
	Count      int    `json:"count"`
	Instrument string `json:"instrument"`
	Ids        string `json:"ids"`
}

type GetTransactions struct {
	Transactions []Transaction `json:"transactions"`
}

type Transaction struct {
	Id                       int64          `json:"id"`
	AccountId                int64          `json:"accountId"`
	Time                     time.Time      `json:"time"`
	Type                     string         `json:"type"`
	Instrument               string         `json:"instrument"`
	Side                     string         `json:"side"`
	Units                    int            `json:"units"`
	Price                    float64        `json:"price"`
	LowerBound               float64        `json:"lowerBound"`
	UpperBound               float64        `json:"upperBound"`
	TakeProfitPrice          float64        `json:"takeProfitPrice"`
	StopLossPrice            float64        `json:"stopLossPrice"`
	TrailingStopLossDistance int            `json:"trailingStopLossDistance"`
	Pl                       float64        `json:"pl"`
	Interest                 float64        `json:"interest"`
	AccountBalance           float64        `json:"accountBalance"`
	OrderId                  int64          `json:"orderId"`
	TradeId                  int64          `json:"tradeId"`
	TradeOpened              TxTradeOpened  `json:"tradeOpened"`
	TradeReduced             TxTradeReduced `json:"tradeReduced"`
	Expiry                   time.Time      `json:"expiry"`
	Reason                   string         `json:"reason"`
	MarginRate               float64        `json:"marginRate"`
	Amount                   float64        `json:"amount"`
}

type TxTradeOpened struct {
	Id    int64 `json:"id"`
	Units int   `json:"units"`
}

type TxTradeReduced struct {
	Id       int64   `json:"id"`
	Units    int     `json:"units"`
	Pl       float64 `json:"pl"`
	Interest float64 `json:"interest"`
}
