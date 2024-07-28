package main

import (
	"net/url"
	"strconv"

	"github.com/tidwall/gjson"
)

func validateServer(server string) bool {
	healthcheck, err := getRequest(server + "/health")
	if err != nil {
		return false
	}

	if healthcheck == `"ok"` {
		return true
	}
	return false
}
func validateConfig(config Config) bool {
	if config.ID < 1 {
		return false
	}

	if config.ISP == "" {
		return false
	}

	_, _, _, cityValid := validateCity(config.City)
	if cityValid == false {
		return false
	}

	if config.Longitude > 180 || config.Longitude < -180 {
		return false
	}

	if config.Latitude > 90 || config.Latitude < -90 {
		return false
	}
	if config.TestFrequency > 5 || config.TestFrequency < 1 {
		return false
	}
	return true
}
func validateCity(city string) (newCity string, lat float64, lon float64, valid bool) {
	osmOut, err := getRequest("https://nominatim.openstreetmap.org/search?q=" + city + "&format=json&polygon=1&addressdetails=1&limit=1")
	if err != nil {
		return "", 0.0, 0.0, false
	}
	if !gjson.Valid(osmOut) {
		return "", 0.0, 0.0, false
	}

	gjsonArr := gjson.Get(osmOut, "0.address.city").String()
	newCityOut, err := getRequest("https://nominatim.openstreetmap.org/search?q=" + url.QueryEscape(gjsonArr) + "&format=json&polygon=1&addressdetails=1&limit=1")
	latStr := gjson.Get(newCityOut, "0.lat").String()
	lonStr := gjson.Get(newCityOut, "0.lon").String()
	lat, err = strconv.ParseFloat(latStr, 64)
	if err != nil {
		return "", 0.0, 0.0, false
	}
	lon, err = strconv.ParseFloat(lonStr, 64)
	if err != nil {
		return "", 0.0, 0.0, false
	}
	if gjsonArr != "" {
		return gjsonArr, lat, lon, true
	}
	return "", 0.0, 0.0, false
}
