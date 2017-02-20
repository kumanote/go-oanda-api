package api

import (
	"github.com/kumanote/go-oanda-api/data"
	"strconv"
)

func (o *Oanda) GetTrades(accountId int64, parameters interface{}) (*data.GetTrades, error) {
	var ret data.GetTrades
	endpoint := "/v1/accounts/" + strconv.FormatInt(accountId, 10) + "/trades"
	if err := o.apiJsonUnmarshal(endpoint, "GET", parameters, nil, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}

func (o *Oanda) GetTrade(accountId int64, tradeId int64) (*data.Trade, error) {
	var ret data.Trade
	endpoint := "/v1/accounts/" + strconv.FormatInt(accountId, 10) + "/trades/" + strconv.FormatInt(tradeId, 10)
	if err := o.apiJsonUnmarshal(endpoint, "GET", nil, nil, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}

func (o *Oanda) PatchTrade(accountId int64, tradeId int64, parameters interface{}) (*data.Trade, error) {
	var ret data.Trade
	endpoint := "/v1/accounts/" + strconv.FormatInt(accountId, 10) + "/trades/" + strconv.FormatInt(tradeId, 10)
	if err := o.apiJsonUnmarshal(endpoint, "PATCH", parameters, nil, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}

func (o *Oanda) CloseTrade(accountId int64, tradeId int64) (*data.CloseTradeResult, error) {
	var ret data.CloseTradeResult
	endpoint := "/v1/accounts/" + strconv.FormatInt(accountId, 10) + "/trades/" + strconv.FormatInt(tradeId, 10)
	if err := o.apiJsonUnmarshal(endpoint, "DELETE", nil, nil, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
