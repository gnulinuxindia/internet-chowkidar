package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/tidwall/gjson"
)

func fetchAndRun(config Config) error {
	catStr := ""
	for i := range config.CheckCategories {
		catStr = catStr + "," + config.CheckCategories[i]
	}
	catStr = strings.TrimSuffix(catStr, ",")
	sitesList, err := getRequest(config.Server + "/isps?category=" + catStr)
	if err != nil {
		return err
	}
	if !gjson.Valid(sitesList) {
		return err
	}
	domains := gjson.Get(sitesList, "#.domain").Array()

	for i := range domains {
		siteData, err := getRequest("https://" + domains[i].String())

		var blocked bool
		// If it times out or has part of the government boilerplate
		if err != nil || (strings.Contains(siteData, "blocked") && strings.Contains(siteData, "directions")) {
			blocked = true
			fmt.Println("Site " + domains[i].String() + " reported as blocked")
		} else {
			blocked = false
			fmt.Println("Site " + domains[i].String() + " reported as not blocked")
		}
		type BlockStruct struct {
			SiteID    int  `json:"site_id"`
			IspID     int  `json:"isp_id"`
			IsBlocked bool `json:"is_blocked"`
		}
		blockStruct := BlockStruct{SiteID: i, IspID: config.ID, IsBlocked: blocked}
		data, err := json.Marshal(blockStruct)
		if err != nil {
			return err
		}
		_, err = postRequest(config.Server+"/blocks", []byte(data), "application/json")
		if err != nil {
			return err
		}
	}

	return nil
}
