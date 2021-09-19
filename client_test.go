package sdek

import "testing"

var client = NewClientTest()

func TestNewClient(t *testing.T) {

	err := client.TokenRefresh()
	if err != nil {
		t.Fatal(err)
	}
}
