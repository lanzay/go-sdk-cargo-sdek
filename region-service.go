package go_sdk_cargo_sdek

import (
	"net/url"
	"reflect"
	"strconv"
)

func (c *Client) GetRegionsAll(filters map[string]string) ([]Region, error) {

	size := 1000
	page := 0

	var err error
	var res []Region
	for {
		regions, err := c.GetRegions(filters, size, page)
		res = append(res, regions...)
		if err != nil || len(regions) < size {
			err = err
			break
		}
		if page <= 100 {
			break
		}
		page++
	}
	return res, err
}

func (c *Client) GetRegions(filters map[string]string, size, page int) ([]Region, error) {

	method := "location/regions?"

	vs := url.Values{}
	vs.Add("size", strconv.Itoa(size))
	vs.Add("page", strconv.Itoa(page))
	for k, v := range filters {
		vs.Add(k, v)
	}

	var regions []Region
	var errorsRes *ErrorsSDK
	_, err := c.get(method+vs.Encode(), &regions, &errorsRes)
	if err != nil {
		panic(err)
	}

	err = errorsRes
	if reflect.ValueOf(errorsRes).Kind() == reflect.Ptr && reflect.ValueOf(errorsRes).IsNil() {
		err = nil
	}

	return regions, err
}
