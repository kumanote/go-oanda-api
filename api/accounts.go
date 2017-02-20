package api

import (
	"github.com/kumanote/go-oanda-api/data"
	"strconv"
)

func (o *Oanda) Accounts() (*data.Accounts, error) {
	var ret data.Accounts
	endpoint := "/v1/accounts"
	if err := o.apiJsonUnmarshal(endpoint, "GET", nil, nil, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}

func (o *Oanda) AccountInformation(accountId int64) (*data.AccountInformation, error) {
	var ret data.AccountInformation
	endpoint := "/v1/accounts/" + strconv.FormatInt(accountId, 10)
	if err := o.apiJsonUnmarshal(endpoint, "GET", nil, nil, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
