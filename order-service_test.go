package sdek

import (
	"fmt"
	"testing"
)

func Test_OrderCreate(t *testing.T) {

	orderRes, err := client.OrderCreate(orderTestOk)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("[D] %v", orderRes)
}

func TestClient_OrderInfoByUUID(t *testing.T) {

	orderRes, err := client.OrderCreate(orderTestOk)
	if err != nil {
		t.Fatal(err)
	}

	order, err := client.OrderInfoByUUID(orderRes.Entity.UUID)
	if err != nil {
		t.FailNow()
	}
	fmt.Printf("[D] %v", order)
}

var orderTestOk = OrderReq{
	Type:    1,
	Number:  "N-001",
	Comment: "коммент 1",
	DeliveryRecipientCost: DeliveryRecipientCost{
		Value: 123,
	},
	DeliveryRecipientCostAdv: []DeliveryRecipientCostAdv{{
		Sum:       112,
		Threshold: 10,
	}},
	FromLocation: LocationOrder{
		Code:        "44",
		FiasGUID:    "",
		PostalCode:  "",
		Longitude:   "",
		Latitude:    "",
		CountryCode: "",
		Region:      "",
		SubRegion:   "",
		City:        "Москва",
		KladrCode:   "",
		Address:     "пр. Ленинградский, д.4",
	},
	ToLocation: LocationOrder{
		Code:        "270",
		FiasGUID:    "",
		PostalCode:  "",
		Longitude:   "",
		Latitude:    "",
		CountryCode: "",
		Region:      "",
		SubRegion:   "",
		City:        "Новосибирск",
		KladrCode:   "",
		Address:     "ул. Блюхера, 32",
	},
	Packages: []Package{{
		Number:  "Упаковка 1",
		Comment: "Коммент упаковки 1",
		Height:  234,
		Items: []Item{{
			WareKey: "ART-001",
			Payment: DeliveryRecipientCost{Value: 123.12},
			Name:    "Товар 1",
			Cost:    100,
			Amount:  1,
			Weight:  700,
			URL:     "https://123.com",
		}},
		Length: 0,
		Weight: 0,
		Width:  0,
	}},
	Recipient: Recipient{
		Name:   "Получатель 1",
		Phones: []Phone{{Number: "+79134637228"}},
	},
	Sender:     Sender{Name: "Отправитель 1"},
	Services:   []Service{{Code: "SECURE_PACKAGE_A2"}},
	TariffCode: 139,
}
