package go_sdk_cargo_sdek

import "strings"

// Token
type (
	TokenRes struct {
		AccessToken string `json:"access_token,omitempty"`
		TokenType   string `json:"token_type,omitempty"`
		ExpiresIn   int    `json:"expires_in,omitempty"`
		Scope       string `json:"scope,omitempty"`
		JTI         string `json:"jti,omitempty"`
	}
	TokenErr struct {
		ErrorType        string `json:"error,omitempty"`
		ErrorDescription string `json:"error_description,omitempty"`
	}
)

func (err *TokenErr) Error() string {
	return err.ErrorType + ": " + err.ErrorDescription
}

// Error
type (
	ErrorsSDK struct {
		Errors []ErrorSDK `json:"errors,omitempty"`
	}
	ErrorSDK struct {
		Code    string `json:"code,omitempty"`
		Message string `json:"message,omitempty"`
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
	CountryCode string     `json:"country_code,omitempty"`
	Region      string     `json:"region,omitempty"`
	Country     string     `json:"country,omitempty"`
	RegionCode  *int       `json:"region_code,omitempty"`
	Errors      []ErrorSDK `json:"errors,omitempty"`
}

// City
type City struct {
	Code        int        `json:"code,omitempty"`
	PostalCodes []string   `json:"postal_codes,omitempty"`
	CountryCode string     `json:"country_code,omitempty"`
	FiasGUID    string     `json:"fias_guid,omitempty"`
	Country     string     `json:"country,omitempty"`
	Region      string     `json:"region,omitempty"`
	RegionCode  int        `json:"region_code,omitempty"`
	SubRegion   string     `json:"sub_region,omitempty"`
	City        string     `json:"city,omitempty"`
	Longitude   float64    `json:"longitude,omitempty"`
	Latitude    float64    `json:"latitude,omitempty"`
	TimeZone    string     `json:"time_zone,omitempty"`
	Errors      []ErrorSDK `json:"errors,omitempty"`
}

// Oder
type (
	OrderReq struct {
		Type                     int                        `json:"type,omitempty"`
		Number                   string                     `json:"number,omitempty"`
		Comment                  string                     `json:"comment,omitempty"`
		DeliveryRecipientCost    DeliveryRecipientCost      `json:"delivery_recipient_cost,omitempty"`
		DeliveryRecipientCostAdv []DeliveryRecipientCostAdv `json:"delivery_recipient_cost_adv,omitempty"`
		ShipmentPoint            *string                    `json:"shipment_point,omitempty"`
		DeliveryPoint            *string                    `json:"delivery_point,omitempty"`
		FromLocation             *Location                  `json:"from_location,omitempty"`
		ToLocation               *Location                  `json:"to_location,omitempty"`
		Packages                 []Package                  `json:"packages,omitempty"`
		Recipient                Recipient                  `json:"recipient,omitempty"`
		Sender                   Sender                     `json:"sender,omitempty"`
		Services                 []Service                  `json:"services,omitempty"`
		TariffCode               int                        `json:"tariff_code,omitempty"`
	}
	DeliveryRecipientCost struct {
		Value float32 `json:"value,omitempty"`
	}
	DeliveryRecipientCostAdv struct {
		Sum       float32 `json:"sum,omitempty"`
		Threshold float32 `json:"threshold,omitempty"`
	}

	Package struct {
		Number  string  `json:"number,omitempty"`
		Comment string  `json:"comment,omitempty"`
		Height  float32 `json:"height,omitempty"`
		Items   []Item  `json:"items,omitempty"`
		Length  float32 `json:"length,omitempty"`
		Weight  float32 `json:"weight,omitempty"`
		Width   float32 `json:"width,omitempty"`
	}
	Item struct {
		Brand       string                `json:"brand,omitempty"`
		WareKey     string                `json:"ware_key,omitempty"` // Артикул
		Marking     string                `json:"marking,omitempty"`  // Маркировка товара
		Value       float32               `json:"value,omitempty"`    // Сумма наложенного платежа (в случае предоплаты = 0)
		Payment     DeliveryRecipientCost `json:"payment,omitempty"`
		Name        string                `json:"name,omitempty"`
		Cost        float32               `json:"cost,omitempty"`
		Amount      float32               `json:"amount,omitempty"`
		Weight      float32               `json:"weight,omitempty"`
		URL         string                `json:"url,omitempty"`
		WeightGross float32               `json:"weight_gross,omitempty"`
	}
	Recipient struct {
		Name   string  `json:"name,omitempty"`
		Phones []Phone `json:"phones,omitempty"`
	}
	Sender struct {
		Name string `json:"name,omitempty"`
	}
	Service struct {
		Code string `json:"code,omitempty"`
	}
)

type (
	OrderRes struct {
		Entity          *Order          `json:"entity,omitempty"`
		Requests        []Request       `json:"requests,omitempty"`
		RelatedEntities []RelatedEntity `json:"related_entities,omitempty"` // Возврат
	}

	Request struct {
		RequestUUID string     `json:"request_uuid,omitempty"`
		Type        string     `json:"type,omitempty"`
		State       string     `json:"state,omitempty"`
		DateTime    string     `json:"date_time,omitempty"`
		Errors      []ErrorSDK `json:"errors,omitempty"`
		Warnings    []ErrorSDK `json:"warnings,omitempty"`
	}
	RelatedEntity struct {
		UUID string `json:"uuid,omitempty"`
		Type string `json:"type,omitempty"`
	}
)

type (
	Order struct {
		UUID                     string                     `json:"uuid,omitempty"`
		Type                     int                        `json:"type,omitempty"`
		IsReturn                 bool                       `json:"is_return,omitempty"`
		Number                   string                     `json:"number,omitempty"`
		CdekNumber               string                     `json:"cdek_number,omitempty"`
		TariffCode               int                        `json:"tariff_code,omitempty"`
		Sender                   Recipient                  `json:"sender,omitempty"`
		Recipient                Recipient                  `json:"recipient,omitempty"`
		FromLocation             Location                   `json:"from_location,omitempty"`
		ToLocation               Location                   `json:"to_location,omitempty"`
		Packages                 []Package                  `json:"packages,omitempty"`
		Services                 []Service                  `json:"services,omitempty"`
		DeliveryRecipientCost    DeliveryRecipientCost      `json:"delivery_recipient_cost,omitempty"`
		DeliveryRecipientCostAdv []DeliveryRecipientCostAdv `json:"delivery_recipient_cost_adv,omitempty"`
		RecipientCurrency        string                     `json:"recipient_currency,omitempty"`
		ItemsCostCurrency        string                     `json:"items_cost_currency,omitempty"`
		Comment                  string                     `json:"comment,omitempty"`
		ShopSellerName           string                     `json:"shop_seller_name,omitempty"`
		Statuses                 []Status                   `json:"statuses,omitempty"`
		Seller                   Seller                     `json:"seller,omitempty"`
		ShipmentPoint            string                     `json:"shipment_point,omitempty"`
		DeliveryPoint            string                     `json:"delivery_point,omitempty"`
	}
	Seller struct {
		Name string `json:"name,omitempty"`
	}
	Status struct {
		Code     string `json:"code,omitempty"`
		Name     string `json:"name,omitempty"`
		DateTime string `json:"date_time,omitempty"`
		City     string `json:"city,omitempty"`
	}
)

type PVZ struct {
	Code                string              `json:"code,omitempty"`
	Name                string              `json:"name,omitempty"`
	Location            Location            `json:"location,omitempty"`
	AddressComment      string              `json:"address_comment,omitempty"`
	NearestStation      string              `json:"nearest_station,omitempty"`
	NearestMetroStation string              `json:"nearest_metro_station,omitempty"`
	WorkTime            string              `json:"work_time,omitempty"`
	Phones              []Phone             `json:"phones,omitempty"`
	Email               string              `json:"email,omitempty"`
	Note                string              `json:"note,omitempty"`
	Type                string              `json:"type,omitempty"`
	OwnerCode           string              `json:"owner_code,omitempty"`
	TakeOnly            bool                `json:"take_only,omitempty"`
	IsDressingRoom      bool                `json:"is_dressing_room,omitempty"`
	HaveCashless        bool                `json:"have_cashless,omitempty"`
	HaveCash            bool                `json:"have_cash,omitempty"`
	AllowedCod          bool                `json:"allowed_cod,omitempty"`
	WorkTimeList        []WorkTimeList      `json:"work_time_list,omitempty"`
	WorkTimeExceptions  []WorkTimeException `json:"work_time_exceptions,omitempty"`
	WeightMin           float32             `json:"weight_min,omitempty,omitempty"`
	WeightMax           float32             `json:"weight_max,omitempty,omitempty"`
	Errors              []ErrorSDK          `json:"errors,omitempty"`
}

type Location struct {
	Code        int     `json:"code,omitempty"`
	FiasGUID    string  `json:"fias_guid,omitempty"`
	PostalCode  string  `json:"postal_code,omitempty"`
	Longitude   float32 `json:"longitude,omitempty"`
	Latitude    float32 `json:"latitude,omitempty"`
	CountryCode string  `json:"country_code,omitempty"`
	Country     string  `json:"country,omitempty"`
	Region      string  `json:"region,omitempty"`
	RegionCode  int     `json:"region_code,omitempty"`
	SubRegion   string  `json:"sub_region,omitempty"`
	CityCode    int     `json:"city_code,omitempty"`
	City        string  `json:"city,omitempty"`
	KladrCode   string  `json:"kladr_code,omitempty"`
	Address     string  `json:"address,omitempty"`
	AddressFull string  `json:"address_full,omitempty"`
}
