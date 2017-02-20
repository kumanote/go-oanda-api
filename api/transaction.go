package api

import (
	"github.com/kumanote/go-oanda-api/data"
	"strconv"
)

func (o *Oanda) GetTransactions(accountId int64, parameters interface{}) (*data.GetTransactions, error) {
	var ret data.GetTransactions
	endpoint := "/v1/accounts/" + strconv.FormatInt(accountId, 10) + "/transactions"
	if err := o.apiJsonUnmarshal(endpoint, "GET", parameters, nil, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}

func (o *Oanda) GetTransaction(accountId int64, transactionId int64) (*data.Transaction, error) {
	var ret data.Transaction
	endpoint := "/v1/accounts/" + strconv.FormatInt(accountId, 10) + "/transactions/" + strconv.FormatInt(transactionId, 10)
	if err := o.apiJsonUnmarshal(endpoint, "GET", nil, nil, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
