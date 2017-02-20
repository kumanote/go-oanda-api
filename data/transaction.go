package data

import (
	"math/big"
	"time"
)

type GetTransactionsParam struct {
	MaxId      big.Int `json:"maxId"`
	MinId      big.Int `json:"minId"`
	Count      int     `json:"count"`
	Instrument string  `json:"instrument"`
	Ids        string  `json:"ids"`
}

type GetTransactions struct {
	Transactions []Transaction `json:"transactions"`
}

type Transaction struct {
	Id                       big.Int        `json:"id"`
	AccountId                big.Int        `json:"accountId"`
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
	OrderId                  big.Int        `json:"orderId"`
	TradeId                  big.Int        `json:"tradeId"`
	TradeOpened              TxTradeOpened  `json:"tradeOpened"`
	TradeReduced             TxTradeReduced `json:"tradeReduced"`
	Expiry                   time.Time      `json:"expiry"`
	Reason                   string         `json:"reason"`
	MarginRate               float64        `json:"marginRate"`
	Amount                   float64        `json:"amount"`
}

type TxTradeOpened struct {
	Id    big.Int `json:"id"`
	Units int     `json:"units"`
}

type TxTradeReduced struct {
	Id       big.Int `json:"id"`
	Units    int     `json:"units"`
	Pl       float64 `json:"pl"`
	Interest float64 `json:"interest"`
}
