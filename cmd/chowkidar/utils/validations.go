package utils

import (
	"strconv"

	"github.com/tidwall/gjson"
)

func ValidateServer(server string) bool {
	healthcheck, err := GetRequest(server + "/health")
	if err != nil {
		return false
	}

	if healthcheck == `"ok"` {
		return true
	}
	return false
}
func ValidateConfig(config Config) bool {
	if config.ISPID < 0 {
		return false
	}

	if config.ISP == "" {
		return false
	}

	_, _, _, cityValid := ValidateCity(config.City)
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
func ValidateCity(city string) (newCity string, lat float64, lon float64, valid bool) {
	osmOut, err := GetRequest("https://nominatim.openstreetmap.org/search?q=" + city + "&format=json&polygon=1&addressdetails=1&limit=1")
	if err != nil {
		return "", 0.0, 0.0, false
	}
	if !gjson.Valid(osmOut) {
		return "", 0.0, 0.0, false
	}

	gjsonArr := gjson.Get(osmOut, "0.address.city").String()
	newCityOut, err := GetRequest("https://nominatim.openstreetmap.org/search?q=" + gjsonArr + "&format=json&polygon=1&addressdetails=1&limit=1")
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
