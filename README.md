# Oanda Api Wrapper Go
golang implementation of the Oanda Rest API Wrapper.

## Usage

### get via glide

```
glide get github.com/kumanote/go-oanda-api
```

### set Environment Variables

set oanda api access token value to environment variable named "OANDA_PERSONAL_ACCESS_TOKEN"

```
$export OANDA_PERSONAL_ACCESS_TOKEN=your_oanda_api_access_token_value
```

### example

* import package
* create an oanda instance
* create parameters for api(struct or map[string]string are supported)
* call api methods

```
import (
	"fmt"
	"github.com/kumanote/go-oanda-api/api"  // here
	"github.com/kumanote/go-oanda-api/data" // here
)

/********** account **********/
var oanda *api.Oanda
oanda = api.NewOandaPractice()
a, err := oanda.Accounts()
fmt.Printf("%v\n", a)
fmt.Printf("%v\n", err)

ai, err := oanda.AccountInformation(a.Accounts[0].AccountId)
fmt.Printf("%v\n", ai)
fmt.Printf("%v\n", err)

/********** rate **********/
// rate instruments
ip := &data.InstrumentsParam{
	AccountId:   a.Accounts[0].AccountId,
	Instruments: "AUD_CAD,AUD_CHF",
}
i, err := oanda.Instruments(ip)
fmt.Printf("%v\n", i)
fmt.Printf("%v\n", err)

ipmap := map[string]string{
	"accountId":   a.Accounts[0].AccountId.String(),
	"instruments": "AUD_CAD,AUD_CHF",
	"fields":      "instrument,displayName,pip,precision,maxTrailingStop,minTrailingStop,marginRate,halted,maxTradeUnits",
}
i2, err := oanda.Instruments(ipmap)
fmt.Printf("%v\n", i2)
fmt.Printf("%v\n", err)
```

## Official Site
see http://developer.oanda.com/rest-live/introduction/

### for japanese
see http://developer.oanda.com/docs/jp/

note that the documents above might be a little bit behind the latest reference.(which is in English edition.)

## LICENSE
This software is released under the MIT License, see LICENSE.
