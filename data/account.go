package data

import (
	"math/big"
)

type Accounts struct {
	Accounts []Account `json:"accounts"`
}

type Account struct {
	AccountId       big.Int `json:"accountId"`
	AccountName     string  `json:"accountName"`
	AccountCurrency string  `json:"accountCurrency"`
	MarginRate      float64 `json:"marginRate"`
}

type AccountInformation struct {
	AccountId       big.Int `json:"accountId"`
	AccountName     string  `json:"accountName"`
	Balance         float64 `json:"balance"`
	UnrealizedPl    float64 `json:"unrealizedPl"`
	MarginUsed      float64 `json:"marginUsed"`
	MarginAvail     float64 `json:"marginAvail"`
	OpenTrades      float64 `json:"openTrades"`
	OpenOrders      float64 `json:"openOrders"`
	MarginRate      float64 `json:"marginRate"`
	AccountCurrency string  `json:"accountCurrency"`
}
