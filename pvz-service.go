package go_sdk_cargo_sdek

import (
	"net/url"
	"reflect"
)

type PVZ struct {
	Code                string              `json:"code"`
	Name                string              `json:"name"`
	Location            Location            `json:"location"`
	AddressComment      string              `json:"address_comment"`
	NearestStation      string              `json:"nearest_station"`
	NearestMetroStation string              `json:"nearest_metro_station"`
	WorkTime            string              `json:"work_time"`
	Phones              []Phone             `json:"phones"`
	Email               string              `json:"email"`
	Note                string              `json:"note"`
	Type                string              `json:"type"`
	OwnerCode           string              `json:"owner_code"`
	TakeOnly            bool                `json:"take_only"`
	IsDressingRoom      bool                `json:"is_dressing_room"`
	HaveCashless        bool                `json:"have_cashless"`
	HaveCash            bool                `json:"have_cash"`
	AllowedCod          bool                `json:"allowed_cod"`
	WorkTimeList        []WorkTimeList      `json:"work_time_list"`
	WorkTimeExceptions  []WorkTimeException `json:"work_time_exceptions"`
	WeightMin           float32             `json:"weight_min,omitempty"`
	WeightMax           float32             `json:"weight_max,omitempty"`
	Errors              []ErrorSDK          `json:"errors"`
}

type Location struct {
	CountryCode string  `json:"country_code"`
	RegionCode  int     `json:"region_code"`
	Region      string  `json:"region"`
	CityCode    int     `json:"city_code"`
	City        string  `json:"city"`
	PostalCode  string  `json:"postal_code"`
	Longitude   float32 `json:"longitude"`
	Latitude    float32 `json:"latitude"`
	Address     string  `json:"address"`
	AddressFull string  `json:"address_full"`
}

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
