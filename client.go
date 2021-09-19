package sdek

import "os"

const (
	ENDPOINT      = "https://api.cdek.ru/v2/"
	ENDPOINT_TEST = "https://api.edu.cdek.ru/v2/"

	LOG_DIR = "./logs/"
)

type (
	Client struct {
		DEBUG    bool
		auth     *auth
		endPoint string
		token    string
	}
	auth struct {
		clientId     string
		clientSecret string
	}
)

type ServiceСonfigurator interface {
	SetDebugMode(mode bool)
	SetAuth(clientId, clientSecret string)
}

type ServiceProvider interface {
	ServiceСonfigurator

	TokenRefresh() error

	// Catalogs
	GetRegions(filters map[string]string, size, page int) ([]Region, error)
	GetRegionsAll(filters map[string]string) ([]Region, error)
	GetCities(filter map[string]string, size, page int) ([]City, error)
	GetCitiesAll(filter map[string]string) ([]City, error)
	GetPVZs(filters map[string]string) ([]PVZ, error)

	// Order
	OrderCreate(orderReq OrderReq) (*OrderRes, error)
	OrderInfoByUUID(uuid string) (*Order, error)
	OrderInfoByN(n string) (*Order, error)
	OrderInfoByIM(im string) (*Order, error)
}

func NewClient(clientId, clientSecret string) ServiceProvider {
	client := newClient(ENDPOINT, clientId, clientSecret)
	return client
}

func NewClientTest() ServiceProvider {

	clientId, clientSecret := "epT5FMOa7IwjjlwTc1gUjO1GZDH1M1rE", "cYxOu9iAMZYQ1suEqfEvsHld4YQzjY0X"

	client := newClient(ENDPOINT_TEST, clientId, clientSecret)
	client.SetDebugMode(true)

	err := client.TokenRefresh()
	if err != nil {
		panic(err)
	}

	return client
}

func newClient(ENDPOINT, clientId, clientSecret string) ServiceProvider {

	client := &Client{
		endPoint: ENDPOINT,
		auth: &auth{
			clientId:     clientId,
			clientSecret: clientSecret,
		},
	}

	err := client.TokenRefresh()
	if err != nil {
		panic(err)
	}

	return client
}

func (c *Client) SetDebugMode(mode bool) {
	c.DEBUG = mode

	if c.DEBUG {
		err := os.MkdirAll(LOG_DIR, 0666)
		if err != nil {
			panic(err)
		}
	}
}

func (c *Client) SetAuth(clientId, clientSecret string) {
	c.auth.clientId = clientId
	c.auth.clientSecret = clientSecret
}
