package go_sdk_cargo_sdek

import (
	"net/url"
	"reflect"
)

type Phone struct {
	Number string `json:"number"`
}

type WorkTimeException struct {
}

type WorkTimeList struct {
	Day  int    `json:"day"`
	Time string `json:"time"`
}

func (c *Client) GetPVZs(filters map[string]string) ([]PVZ, error) {

	method := "deliverypoints?"

	vs := url.Values{}
	for k, v := range filters {
		vs.Add(k, v)
	}

	var pvzs []PVZ
	var errorsRes *ErrorsSDK
	_, err := c.get(method+vs.Encode(), &pvzs, &errorsRes)
	if err != nil {
		panic(err)
	}

	err = errorsRes
	if reflect.ValueOf(errorsRes).Kind() == reflect.Ptr && reflect.ValueOf(errorsRes).IsNil() {
		err = nil
	}

	return pvzs, err
}
