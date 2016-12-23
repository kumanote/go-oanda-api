package api

import (
	"github.com/kumanote/go-oanda-api/data"
	"math/big"
)

func (o *Oanda) GetTransactions(accountId big.Int, parameters interface{}) (*data.GetTransactions, error) {
	var ret data.GetTransactions
	endpoint := "/v1/accounts/" + accountId.String() + "/transactions"
	if err := o.apiJsonUnmarshal(endpoint, "GET", parameters, nil, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}

func (o *Oanda) GetTransaction(accountId big.Int, transactionId big.Int) (*data.Transaction, error) {
	var ret data.Transaction
	endpoint := "/v1/accounts/" + accountId.String() + "/transactions/" + transactionId.String()
	if err := o.apiJsonUnmarshal(endpoint, "GET", nil, nil, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
