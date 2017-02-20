package api

import (
	"github.com/kumanote/go-oanda-api/data"
	"strconv"
)

func (o *Oanda) GetOrders(accountId int64, parameters interface{}) (*data.Orders, error) {
	var ret data.Orders
	endpoint := "/v1/accounts/" + strconv.FormatInt(accountId, 10) + "/orders"
	if err := o.apiJsonUnmarshal(endpoint, "GET", parameters, nil, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}

func (o *Oanda) PostOrders(accountId int64, parameters interface{}) (*data.PostOrderResult, error) {
	var ret data.PostOrderResult
	endpoint := "/v1/accounts/" + strconv.FormatInt(accountId, 10) + "/orders"
	if err := o.apiJsonUnmarshal(endpoint, "POST", parameters, nil, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}

func (o *Oanda) GetOrder(accountId int64, orderId int64) (*data.Order, error) {
	var ret data.Order
	endpoint := "/v1/accounts/" + strconv.FormatInt(accountId, 10) + "/orders/" + strconv.FormatInt(orderId, 10)
	if err := o.apiJsonUnmarshal(endpoint, "GET", nil, nil, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}

func (o *Oanda) PatchOrder(accountId int64, orderId int64, parameters interface{}) (*data.Order, error) {
	var ret data.Order
	endpoint := "/v1/accounts/" + strconv.FormatInt(accountId, 10) + "/orders/" + strconv.FormatInt(orderId, 10)
	if err := o.apiJsonUnmarshal(endpoint, "PATCH", parameters, nil, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}

func (o *Oanda) DeleteOrder(accountId int64, orderId int64) (*data.DeleteOrderResult, error) {
	var ret data.DeleteOrderResult
	endpoint := "/v1/accounts/" + strconv.FormatInt(accountId, 10) + "/orders/" + strconv.FormatInt(orderId, 10)
	if err := o.apiJsonUnmarshal(endpoint, "DELETE", nil, nil, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
