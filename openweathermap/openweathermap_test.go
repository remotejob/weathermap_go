package openweathermap

import "testing"

func TestGetWeather(t *testing.T) {

	lat := 60.1708
	lon := 24.9375

	GetWeather(lat, lon)

}
