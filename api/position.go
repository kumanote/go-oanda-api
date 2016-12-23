package api

import (
	"github.com/kumanote/go-oanda-api/data"
	"math/big"
)

func (o *Oanda) GetPositions(accountId big.Int) (*data.GetPositions, error) {
	var ret data.GetPositions
	endpoint := "/v1/accounts/" + accountId.String() + "/positions"
	if err := o.apiJsonUnmarshal(endpoint, "GET", nil, nil, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}

func (o *Oanda) GetPosition(accountId big.Int, instrument string) (*data.Position, error) {
	var ret data.Position
	endpoint := "/v1/accounts/" + accountId.String() + "/positions/" + instrument
	if err := o.apiJsonUnmarshal(endpoint, "GET", nil, nil, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}

func (o *Oanda) ClosePosition(accountId big.Int, instrument string) (*data.ClosePositionResult, error) {
	var ret data.ClosePositionResult
	endpoint := "/v1/accounts/" + accountId.String() + "/positions/" + instrument
	if err := o.apiJsonUnmarshal(endpoint, "DELETE", nil, nil, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
