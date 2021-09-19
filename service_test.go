package sdek

import (
	"fmt"
	"testing"
)

func Test_GetRegionsAll(t *testing.T) {

	regions, err := client.GetRegionsAll(map[string]string{
		"country_codes": "RU",
	})
	if err != nil {
		t.Fatal(err)
	}

	if len(regions) == 0 {
		t.FailNow()
	}
	fmt.Println("[T]", regions)
}
func Test_GetCitiesAll(t *testing.T) {

	cities, err := client.GetCitiesAll(map[string]string{
		"country_codes": "RU",
		"region_code":   "81",
		//"city":"Москва",
	})
	if err != nil {
		t.Fatal(err)
	}

	if len(cities) == 0 {
		t.FailNow()
	}
	fmt.Println("[T]", cities)
}

func Test_GetPVZs(t *testing.T) {

	pvzs, err := client.GetPVZs(map[string]string{
		"country_codes": "RU",
		"region_code":   "81", // Москва
		//"postal_code":     "", // Москва
		//"city_code":     "5663234234", // Москва
		"allowed_cod": "1", // Разрешен наложенный платеж, может принимать значения
		"is_handout":  "1", // Является пунктом выдачи, может принимать значения
	})
	if err != nil {
		t.Fatal(err)
	}

	if len(pvzs) == 0 {
		t.FailNow()
	}
	fmt.Println("[T]", pvzs)
}
