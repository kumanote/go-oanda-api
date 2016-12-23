package api

import (
	"github.com/kumanote/go-oanda-api/data"
	"math/big"
)

func (o *Oanda) GetOrders(accountId big.Int, parameters interface{}) (*data.Orders, error) {
	var ret data.Orders
	endpoint := "/v1/accounts/" + accountId.String() + "/orders"
	if err := o.apiJsonUnmarshal(endpoint, "GET", parameters, nil, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}

func (o *Oanda) PostOrders(accountId big.Int, parameters interface{}) (*data.PostOrderResult, error) {
	var ret data.PostOrderResult
	endpoint := "/v1/accounts/" + accountId.String() + "/orders"
	if err := o.apiJsonUnmarshal(endpoint, "POST", parameters, nil, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}

func (o *Oanda) GetOrder(accountId big.Int, orderId big.Int) (*data.Order, error) {
	var ret data.Order
	endpoint := "/v1/accounts/" + accountId.String() + "/orders/" + orderId.String()
	if err := o.apiJsonUnmarshal(endpoint, "GET", nil, nil, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}

func (o *Oanda) PatchOrder(accountId big.Int, orderId big.Int, parameters interface{}) (*data.Order, error) {
	var ret data.Order
	endpoint := "/v1/accounts/" + accountId.String() + "/orders/" + orderId.String()
	if err := o.apiJsonUnmarshal(endpoint, "PATCH", parameters, nil, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}

func (o *Oanda) DeleteOrder(accountId big.Int, orderId big.Int) (*data.DeleteOrderResult, error) {
	var ret data.DeleteOrderResult
	endpoint := "/v1/accounts/" + accountId.String() + "/orders/" + orderId.String()
	if err := o.apiJsonUnmarshal(endpoint, "DELETE", nil, nil, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
