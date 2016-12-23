package api

import (
	"github.com/kumanote/go-oanda-api/data"
	"math/big"
)

func (o *Oanda) Accounts() (*data.Accounts, error) {
	var ret data.Accounts
	endpoint := "/v1/accounts"
	if err := o.apiJsonUnmarshal(endpoint, "GET", nil, nil, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}

func (o *Oanda) AccountInformation(accountId big.Int) (*data.AccountInformation, error) {
	var ret data.AccountInformation
	endpoint := "/v1/accounts/" + accountId.String()
	if err := o.apiJsonUnmarshal(endpoint, "GET", nil, nil, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
