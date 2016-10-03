package openweathermap

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

//GetWeather get data from openweathermap
func GetWeather(lat float64, lon float64) []byte {

	latstr := strconv.FormatFloat(lat, 'f', -1, 64)
	lonstr := strconv.FormatFloat(lon, 'f', -1, 64)

	log.Println(latstr, lonstr)
	u := url.URL{}
	u.Scheme = "http"
	u.Host = "api.openweathermap.org"
	u.Path = "/data/2.5/weather"
	q := u.Query()
	q.Set("lat", latstr)
	q.Set("lon", lonstr)
	q.Set("appid", "c9ae06b4a503a56ffcab71c614b81081")
	q.Set("units", "metric")
	u.RawQuery = q.Encode()
	log.Println(u.String())
	resp, err := http.Get(u.String())
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err.Error())
	} else {

		log.Println(string(body))
	}
	return body
}
