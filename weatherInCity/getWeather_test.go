package weatherInCity_test

import (
	"strings"
	"testing"
	"weather/geoSity"
	"weather/weatherInCity"
)

func TestGetWeather(t *testing.T) {
	expect := "Sochi"
	geo := geoSity.GeoData{expect}
	format := 3

	result, _ := weatherInCity.GetWeather(geo, format)
	if !strings.Contains(result, expect) {
		t.Errorf("Ожидалось %v, пршло %v", result, expect)
	}
}

var TestCasesFormat = []struct {
	name   string
	format int
}{
	{name: "Big format", format: 134},
	{name: "Zero format", format: 0},
	{name: "Minus format", format: -1},
}

func TestGetWeatherWrongFormat(t *testing.T) {
	for _, testCase := range TestCasesFormat {
		t.Run(testCase.name, func(t *testing.T) {
			expect := "Sochi"
			geo := geoSity.GeoData{expect}

			_, err := weatherInCity.GetWeather(geo, testCase.format)
			if err != weatherInCity.ErrorBadFormat {
				t.Errorf("Ожидалось %v пришло %v", weatherInCity.ErrorBadFormat, err)
			}
		})
	}
}
