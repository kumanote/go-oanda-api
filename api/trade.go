package api

import (
	"github.com/kumanote/go-oanda-api/data"
	"math/big"
)

func (o *Oanda) GetTrades(accountId big.Int, parameters interface{}) (*data.GetTrades, error) {
	var ret data.GetTrades
	endpoint := "/v1/accounts/" + accountId.String() + "/trades"
	if err := o.apiJsonUnmarshal(endpoint, "GET", parameters, nil, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}

func (o *Oanda) GetTrade(accountId big.Int, tradeId big.Int) (*data.Trade, error) {
	var ret data.Trade
	endpoint := "/v1/accounts/" + accountId.String() + "/trades/" + tradeId.String()
	if err := o.apiJsonUnmarshal(endpoint, "GET", nil, nil, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}

func (o *Oanda) PatchTrade(accountId big.Int, tradeId big.Int, parameters interface{}) (*data.Trade, error) {
	var ret data.Trade
	endpoint := "/v1/accounts/" + accountId.String() + "/trades/" + tradeId.String()
	if err := o.apiJsonUnmarshal(endpoint, "PATCH", parameters, nil, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}

func (o *Oanda) CloseTrade(accountId big.Int, tradeId big.Int) (*data.CloseTradeResult, error) {
	var ret data.CloseTradeResult
	endpoint := "/v1/accounts/" + accountId.String() + "/trades/" + tradeId.String()
	if err := o.apiJsonUnmarshal(endpoint, "DELETE", nil, nil, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
