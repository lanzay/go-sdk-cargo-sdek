package sdek

import (
	"net/url"
	"reflect"
	"strconv"
)

func (c *Client) GetCitiesAll(filters map[string]string) ([]City, error) {

	size := 500
	page := 0

	var err error
	var res []City
	for {
		cities, err := c.GetCities(filters, size, page)
		res = append(res, cities...)
		if err != nil || len(cities) < size {
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

func (c *Client) GetCities(filters map[string]string, size, page int) ([]City, error) {

	method := "location/cities?"

	vs := url.Values{}
	vs.Add("size", strconv.Itoa(size))
	vs.Add("page", strconv.Itoa(page))
	for k, v := range filters {
		vs.Add(k, v)
	}

	var cities []City
	var errorsRes *ErrorsSDK
	_, err := c.get(method+vs.Encode(), &cities, &errorsRes)
	if err != nil {
		panic(err)
	}

	err = errorsRes
	if reflect.ValueOf(errorsRes).Kind() == reflect.Ptr && reflect.ValueOf(errorsRes).IsNil() {
		err = nil
	}

	return cities, err
}
