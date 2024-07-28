package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"time"

	"github.com/tidwall/gjson"
	"go.mills.io/bitcask/v2"
)

func fetchAndRun(config Config, db *bitcask.Bitcask) error {
	err := db.Put([]byte("lastRun"), []byte(strconv.Itoa(int(time.Now().Unix()))))
	if err != nil {
		return err
	}

	catStr := ""
	for i := range config.CheckCategories {
		catStr = catStr + "," + config.CheckCategories[i]
	}
	catStr = strings.TrimPrefix(strings.TrimSuffix(catStr, ","), ",")
	sitesList, err := getRequest(config.Server + "/sites?category=" + catStr)
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
		if err != nil || (strings.Contains(siteData, "blocked") && (strings.Contains(siteData, "directions") || strings.Contains(siteData, "order"))) {
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
			Changed   bool `json:"changed"`
		}
		iStr := strconv.Itoa(i)
		val, domainErr := db.Get([]byte(domains[i].String() + "_status"))
		block := ""
		switch blocked {
		case true:
			block = "blocked"
		case false:
			block = "unblocked"
		}
		var changed bool
		if domainErr != nil {
			changed = false
		} else if string(val) != block {
			changed = true
		}
		if string(val) != block || domainErr != nil {
			blockStruct := BlockStruct{SiteID: int(gjson.Get(sitesList, string(iStr)+".id").Int()), IspID: config.ID, IsBlocked: blocked, Changed: changed}
			data, err := json.Marshal(blockStruct)
			if err != nil {
				return err
			}
			_, err = postRequest(config.Server+"/blocks", []byte(data), "application/json")
			if err != nil {
				return err
			}
		}
		err = db.Put([]byte(domains[i].String()+"_status"), []byte(block))
		if err != nil {
			return err
		}
	}

	return nil
}
