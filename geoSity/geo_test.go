package geoSity_test

import (
	"testing"
	"weather/geoSity"
)

func TestGetMyLocation(t *testing.T) {
	//	Arrange - подготовка
	city := "Sochi"
	expect := geoSity.GeoData{"Sochi"}
	//	Act - выполнение функции
	got, err := geoSity.GetMyLocation(city)

	//	Asseert - проверка результата
	if err != nil {
		t.Error(err)
	}
	if got.City != expect.City {
		t.Errorf("got %v", got)
	}
}

func TestGetMyLocationNoCity(t *testing.T) {
	city := "Sochii"
	_, err := geoSity.GetMyLocation(city)
	if err != geoSity.ErrorNoCity {
		t.Errorf("got %v", err)
	}
}
