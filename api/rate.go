package api

import (
	"github.com/kumanote/go-oanda-api/data"
)

func (o *Oanda) Instruments(parameters interface{}) (*data.Instruments, error) {
	var ret data.Instruments
	endpoint := "/v1/instruments"
	if err := o.apiJsonUnmarshal(endpoint, "GET", parameters, nil, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}

func (o *Oanda) Prices(parameters interface{}) (*data.Prices, error) {
	var ret data.Prices
	endpoint := "/v1/prices"
	if err := o.apiJsonUnmarshal(endpoint, "GET", parameters, nil, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}

func (o *Oanda) Candles(parameters interface{}) (*data.Candles, error) {
	var ret data.Candles
	endpoint := "/v1/candles"
	if err := o.apiJsonUnmarshal(endpoint, "GET", parameters, nil, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
