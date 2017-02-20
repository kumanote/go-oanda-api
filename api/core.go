package api

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"math/big"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

type Oanda struct {
	baseURL string
	Token   string
}

func NewOandaProd() *Oanda {
	return &Oanda{
		baseURL: "https://api-fxtrade.oanda.com",
	}
}

func NewOandaPractice() *Oanda {
	return &Oanda{
		baseURL: "https://api-fxpractice.oanda.com",
	}
}

func (o *Oanda) getToken() string {
	if o.Token != "" {
		return o.Token
	}
	return os.Getenv("OANDA_PERSONAL_ACCESS_TOKEN")
}

func (o *Oanda) apiJsonUnmarshal(endpoint string, method string, parameters interface{}, headers interface{}, val interface{}) error {

	responseBody, err := o.api(endpoint, method, struct2map(parameters), struct2map(headers))
	if err != nil {
		return err
	}
	if len(responseBody) < 1 {
		return nil
	}
	err = json.Unmarshal(responseBody, val)
	if err != nil {
		return err
	}
	return nil
}

func (o *Oanda) api(endpoint string, method string, parameters map[string]string, headers map[string]string) ([]byte, error) {

	switch method {
	case "GET":
		return o.get(endpoint, parameters, headers)
	case "POST":
		return o.post(endpoint, parameters, headers)
	case "PATCH":
		return o.patch(endpoint, parameters, headers)
	case "DELETE":
		return o.delete(endpoint, parameters, headers)
	default:
		return nil, fmt.Errorf("unexpected api method: %s", method)
	}
}

func (o *Oanda) get(endpoint string, parameters map[string]string, headers map[string]string) ([]byte, error) {

	calluri := o.baseURL + endpoint
	if parameters != nil && 0 < len(parameters) {
		values := url.Values{}
		for key, value := range parameters {
			values.Add(key, value)
		}
		calluri += ("?" + values.Encode())
	}

	return o.apiCore("GET", calluri, nil, headers)
}

func (o *Oanda) post(endpoint string, parameters map[string]string, headers map[string]string) ([]byte, error) {
	calluri := o.baseURL + endpoint
	values := url.Values{}
	if parameters != nil && 0 < len(parameters) {
		for key, value := range parameters {
			values.Add(key, value)
		}
	}
	return o.apiCore("POST", calluri, strings.NewReader(values.Encode()), headers)
}

func (o *Oanda) patch(endpoint string, parameters map[string]string, headers map[string]string) ([]byte, error) {
	calluri := o.baseURL + endpoint
	values := url.Values{}
	if parameters != nil && 0 < len(parameters) {
		for key, value := range parameters {
			values.Add(key, value)
		}
	}
	return o.apiCore("PATCH", calluri, strings.NewReader(values.Encode()), headers)
}

func (o *Oanda) delete(endpoint string, parameters map[string]string, headers map[string]string) ([]byte, error) {
	calluri := o.baseURL + endpoint
	values := url.Values{}
	if parameters != nil && 0 < len(parameters) {
		for key, value := range parameters {
			values.Add(key, value)
		}
	}
	return o.apiCore("DELETE", calluri, strings.NewReader(values.Encode()), headers)
}

func (o *Oanda) apiCore(method, urlStr string, body io.Reader, headers map[string]string) ([]byte, error) {

	client := &http.Client{}

	req, err := http.NewRequest(method, urlStr, body)
	if err != nil {
		return nil, err
	}
	if headers != nil && 0 < len(headers) {
		for key, value := range headers {
			req.Header.Add(key, value)
		}
	}
	if req.Header.Get("Content-Type") == "" {
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	}
	if req.Header.Get("Authorization") == "" {
		req.Header.Add("Authorization", "Bearer "+o.getToken())
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	rb, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode < http.StatusOK || http.StatusIMUsed < res.StatusCode {
		return nil, fmt.Errorf("http status is not ok: %s, details: %s", res.StatusCode, string(rb))
	}

	return rb, nil
}

func struct2map(i interface{}) map[string]string {
	ret := make(map[string]string)
	if i == nil {
		return ret
	}
	if m, ok := i.(map[string]string); ok {
		return m
	}
	elem := reflect.ValueOf(i).Elem()
	size := elem.NumField()
	for idx := 0; idx < size; idx++ {
		var value string
		k := ""
		includeDefault := elem.Type().Field(idx).Tag.Get("includeDefault") == "true"
		if jsonTagName := elem.Type().Field(idx).Tag.Get("json"); jsonTagName != "" {
			k = jsonTagName
		} else {
			k = elem.Type().Field(idx).Name
		}
		switch v := elem.Field(idx).Interface(); vt := v.(type) {
		case string:
			value = vt
		case bool:
			if includeDefault || vt != false {
				value = strconv.FormatBool(vt)
			}
		case int:
			if includeDefault || vt != 0 {
				value = strconv.Itoa(vt)
			}
		case int64:
			if includeDefault || vt != 0 {
				value = strconv.FormatInt(vt, 10)
			}
		case float64:
			if includeDefault || vt != 0.0 {
				value = strconv.FormatFloat(vt, 'f', 2, 64)
			}
		case big.Int:
			if includeDefault || vt.Cmp(big.NewInt(0)) != 0 {
				value = vt.String()
			}
		case time.Time:
			if includeDefault || !vt.IsZero() {
				value = vt.Format(time.RFC3339)
			}
		default:
			panic(fmt.Sprintf("unexpected parameter field: %v\n", value))
		}
		if includeDefault || value != "" {
			ret[k] = value
		}
	}
	return ret
}
