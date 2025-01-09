package main

import (
	"flag"
	"fmt"
	"weather/geoSity"
	"weather/weatherInCity"
)

func main() {
	city := flag.String("city", "", "Ваш город")
	formatOutput := flag.Int("format", 4, "Формат вывода погоды")
	geoData, err := geoSity.GetMyLocation(*city)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(geoData)
	weatherData, _ := weatherInCity.GetWeather(*geoData, *formatOutput)
	fmt.Println(weatherData)
}
