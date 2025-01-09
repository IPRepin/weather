package weatherInCity

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"weather/geoSity"
)

var ErrorBadUrl = errors.New("BAD_URL")
var ErrorBadFormat = errors.New("BAD_FORMAT")
var ErrorResponse = errors.New("ERROR_RESPONSE")
var ErrorNo200 = errors.New("NOT_200")
var ErrorReadBody = errors.New("ERROR_READ_BODY")

func GetWeather(geo geoSity.GeoData, format int) (string, error) {
	if format < 1 || format > 4 {
		return "", ErrorBadFormat
	}
	baseUrl, err := url.Parse("https://wttr.in/" + geo.City)
	if err != nil {
		return "", ErrorBadUrl
	}
	params := url.Values{}
	params.Add("format", fmt.Sprint(format))
	baseUrl.RawQuery = params.Encode()
	resp, err := http.Get(baseUrl.String())
	if err != nil {
		return "", ErrorResponse
	}
	if resp.StatusCode != 200 {
		return "", ErrorNo200
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", ErrorReadBody
	}
	return string(body), nil
}
