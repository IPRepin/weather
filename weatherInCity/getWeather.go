package weatherInCity

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"weather/geoSity"
)

func GetWeather(geo geoSity.GeoData, format int) string {
	baseUrl, err := url.Parse("https://wttr.in/" + geo.City)
	if err != nil {
		log.Println(err)
		return ""
	}
	params := url.Values{}
	params.Add("format", fmt.Sprint(format))
	baseUrl.RawQuery = params.Encode()
	resp, err := http.Get(baseUrl.String())
	if err != nil {
		log.Println(err)
	}
	if resp.StatusCode != 200 {
		log.Println(resp.Status)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	return string(body)
}
