package go_sdk_cargo_sdek

import "strings"

// Token
type (
	TokenRes struct {
		AccessToken string `json:"access_token"`
		TokenType   string `json:"token_type"`
		ExpiresIn   int    `json:"expires_in"`
		Scope       string `json:"scope"`
		JTI         string `json:"jti"`
	}
	TokenErr struct {
		ErrorType        string `json:"error"`
		ErrorDescription string `json:"error_description"`
	}
)

func (err *TokenErr) Error() string {
	return err.ErrorType + ": " + err.ErrorDescription
}

// Error
type (
	ErrorsSDK struct {
		Errors []ErrorSDK `json:"errors"`
	}
	ErrorSDK struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}
)

func (err ErrorsSDK) Error() string {

	var errs []string
	for _, sdk := range err.Errors {
		errs = append(errs, sdk.Error())
	}
	return strings.Join(errs, "; ")
}

func (err *ErrorSDK) Error() string {
	return err.Message
}

// Region
type Region struct {
	CountryCode string     `json:"country_code"`
	Region      string     `json:"region"`
	Country     string     `json:"country"`
	RegionCode  *int       `json:"region_code,omitempty"`
	Errors      []ErrorSDK `json:"errors"`
}

// City
type City struct {
	Code        int        `json:"code"`
	PostalCodes []string   `json:"postal_codes,omitempty"`
	CountryCode string     `json:"country_code"`
	FiasGUID    string     `json:"fias_guid"`
	Country     string     `json:"country"`
	Region      string     `json:"region"`
	RegionCode  int        `json:"region_code"`
	SubRegion   string     `json:"sub_region"`
	City        string     `json:"city"`
	Longitude   float64    `json:"longitude"`
	Latitude    float64    `json:"latitude"`
	TimeZone    string     `json:"time_zone"`
	Errors      []ErrorSDK `json:"errors"`
}

// Oder
type (
	OrderReq struct {
		Type                     int                        `json:"type"`
		Number                   string                     `json:"number"`
		Comment                  string                     `json:"comment"`
		DeliveryRecipientCost    DeliveryRecipientCost      `json:"delivery_recipient_cost"`
		DeliveryRecipientCostAdv []DeliveryRecipientCostAdv `json:"delivery_recipient_cost_adv"`
		FromLocation             LocationOrder              `json:"from_location"`
		ToLocation               LocationOrder              `json:"to_location"`
		Packages                 []Package                  `json:"packages"`
		Recipient                Recipient                  `json:"recipient"`
		Sender                   Sender                     `json:"sender"`
		Services                 []Service                  `json:"services"`
		TariffCode               int64                      `json:"tariff_code"`
	}
	DeliveryRecipientCost struct {
		Value float32 `json:"value"`
	}
	DeliveryRecipientCostAdv struct {
		Sum       int64 `json:"sum"`
		Threshold int64 `json:"threshold"`
	}
	LocationOrder struct {
		Code        string `json:"code"`
		FiasGUID    string `json:"fias_guid"`
		PostalCode  string `json:"postal_code"`
		Longitude   string `json:"longitude"`
		Latitude    string `json:"latitude"`
		CountryCode string `json:"country_code"`
		Region      string `json:"region"`
		SubRegion   string `json:"sub_region"`
		City        string `json:"city"`
		KladrCode   string `json:"kladr_code"`
		Address     string `json:"address"`
		Country     string `json:"country"`
		RegionCode  string `json:"region_code"`
	}
	Package struct {
		Number  string  `json:"number"`
		Comment string  `json:"comment"`
		Height  float32 `json:"height"`
		Items   []Item  `json:"items"`
		Length  float32 `json:"length"`
		Weight  float32 `json:"weight"`
		Width   float32 `json:"width"`
	}
	Item struct {
		Brand       string                `json:"brand"`
		WareKey     string                `json:"ware_key"` // Артикул
		Marking     string                `json:"marking"`  // Маркировка товара
		Value       float32               `json:"value"`    // Сумма наложенного платежа (в случае предоплаты = 0)
		Payment     DeliveryRecipientCost `json:"payment"`
		Name        string                `json:"name"`
		Cost        float32               `json:"cost"`
		Amount      float32               `json:"amount"`
		Weight      float32               `json:"weight"`
		URL         string                `json:"url"`
		WeightGross float32               `json:"weight_gross"`
	}
	Recipient struct {
		Name   string  `json:"name"`
		Phones []Phone `json:"phones"`
	}
	Sender struct {
		Name string `json:"name"`
	}
	Service struct {
		Code string `json:"code"`
	}
)

type (
	OrderRes struct {
		Entity          *Order          `json:"entity"`
		Requests        []Request       `json:"requests"`
		RelatedEntities []RelatedEntity `json:"related_entities"` // Возврат
	}

	Request struct {
		RequestUUID string     `json:"request_uuid"`
		Type        string     `json:"type"`
		State       string     `json:"state"`
		DateTime    string     `json:"date_time"`
		Errors      []ErrorSDK `json:"errors"`
		Warnings    []ErrorSDK `json:"warnings"`
	}
	RelatedEntity struct {
		UUID string `json:"uuid"`
		Type string `json:"type"`
	}
)

type (
	Order struct {
		UUID                     string                     `json:"uuid"`
		Type                     int                        `json:"type"`
		IsReturn                 bool                       `json:"is_return"`
		Number                   string                     `json:"number"`
		CdekNumber               string                     `json:"cdek_number"`
		TariffCode               int                        `json:"tariff_code"`
		Sender                   Recipient                  `json:"sender"`
		Recipient                Recipient                  `json:"recipient"`
		FromLocation             Location                   `json:"from_location"`
		ToLocation               Location                   `json:"to_location"`
		Packages                 []Package                  `json:"packages"`
		Services                 []Service                  `json:"services"`
		DeliveryRecipientCost    DeliveryRecipientCost      `json:"delivery_recipient_cost"`
		DeliveryRecipientCostAdv []DeliveryRecipientCostAdv `json:"delivery_recipient_cost_adv"`
		RecipientCurrency        string                     `json:"recipient_currency"`
		ItemsCostCurrency        string                     `json:"items_cost_currency"`
		Comment                  string                     `json:"comment"`
		ShopSellerName           string                     `json:"shop_seller_name"`
		Statuses                 []Status                   `json:"statuses"`
		Seller                   Seller                     `json:"seller"`
		ShipmentPoint            string                     `json:"shipment_point"`
		DeliveryPoint            string                     `json:"delivery_point"`
	}
	Seller struct {
		Name string `json:"name"`
	}
	Status struct {
		Code     string `json:"code"`
		Name     string `json:"name"`
		DateTime string `json:"date_time"`
		City     string `json:"city"`
	}
)
