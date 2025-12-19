package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"log"
	"time"

	"github.com/getlantern/systray"
	"github.com/gnulinuxindia/internet-chowkidar/cmd/chowkidar/utils"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v2"
	"go.mills.io/bitcask/v2"
)

var updateConf bool
var stopSync bool

func Run(cCtx *cli.Context) error {
	config, err := utils.FindConfigData(cCtx.String("config"))
	if err != nil {
		return cli.Exit(err.Error(), 1)
	}

	db, err := utils.FindDatabase(cCtx.String("database"))
	if err != nil {
		return cli.Exit(err.Error(), 1)
	}
	defer db.Close()

	go func() {
		systray.Run(func() { createSystray(cCtx.String("config"), db) }, nil)
	}()

	// Figure out frequency based on the mode numbers
	duration := time.Duration(config.TestFrequency) * time.Hour
	fmt.Println("EXITING")
	return cli.Exit("abc", 1)

	// If it wasn't run before acc to DB, run it now
	val, err := db.Get([]byte("lastRun"))
	if err != nil {
		if strings.Contains(err.Error(), "key not found") {
			runErr := fetchAndRun(config, db)
			if runErr != nil {
				return cli.Exit(runErr.Error(), 1)
			}
		} else {
			return err
		}
	} else if string(val) == "" {
		runErr := fetchAndRun(config, db)
		if runErr != nil {
			return cli.Exit(runErr.Error(), 1)
		}
	} else {
		// If it was run before, check if it has been more than `duration` since it happened
		timeInt, err := strconv.ParseInt(string(val), 10, 64)
		if err != nil {
			return cli.Exit("Database has invalid data or is corrupted.", 1)
		}

		timeLast := time.Unix(timeInt, 0)
		durationRemain := time.Since(timeLast)
		if durationRemain >= duration {
			runErr := fetchAndRun(config, db)
			if runErr != nil {
				return cli.Exit(runErr.Error(), 1)
			}
		} else {
			// Wait till duration is complete and then run
			time.Sleep(durationRemain)
			runErr := fetchAndRun(config, db)
			if runErr != nil {
				return cli.Exit(runErr.Error(), 1)
			}
		}
	}

	// Do the periodic ones based on the determined duration
	for range time.Tick(duration) {
		if !stopSync {
			if updateConf {
				config, err = utils.FindConfigData(cCtx.String("config"))
				if err != nil {
					return cli.Exit(err.Error(), 1)
				}
				updateConf = false
			}
			runErr := fetchAndRun(config, db)
			if runErr != nil {
				return cli.Exit(runErr.Error(), 1)
			}
		}
	}
	return nil
}

func fetchAndRun(config utils.Config, db *bitcask.Bitcask) error {
	log.Println("Running a check cycle")
	err := db.Put([]byte("lastRun"), []byte(strconv.Itoa(int(time.Now().Unix()))))
	if err != nil {
		return err
	}

	catStr := strings.Join(config.CheckCategories, ",")
	sitesList, err := utils.GetRequest(config.Server + "/sites?category=" + catStr)
	if err != nil {
		return err
	}
	if !gjson.Valid(sitesList) {
		return err
	}
	domains := gjson.Get(sitesList, "#.ping_url").Array()

	for i, domain := range domains {
		_, err := utils.GetRequest(domain.String())
		//fmt.Println(domain.String())
		//fmt.Println(siteData)

		var blocked bool
		// If it times out or has part of the government boilerplate
		//if err != nil || (strings.Contains(siteData, "blocked") && (strings.Contains(siteData, "directions") || strings.Contains(siteData, "order"))) {
		if err != nil {
			blocked = true
			fmt.Println("Site " + domain.String() + " reported as blocked")
		} else {
			blocked = false
			fmt.Println("Site " + domain.String() + " reported as not blocked")
		}
		type BlockStruct struct {
			SiteID    int  `json:"site_id"`
			IspID     int  `json:"isp_id"`
			ClientID  int  `json:"client_id"`
			IsBlocked bool `json:"is_blocked"`
			Changed   bool `json:"changed"`
		}
		iStr := strconv.Itoa(i)
		val, domainErr := db.Get([]byte(domain.String() + "_status"))
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
			blockStruct := BlockStruct{SiteID: int(gjson.Get(sitesList, string(iStr)+".id").Int()), IspID: config.ISPID, IsBlocked: blocked, Changed: changed, ClientID: config.ClientID}
			data, err := json.Marshal(blockStruct)
			if err != nil {
				return err
			}
			_, err = utils.PostRequest(config.Server+"/blocks", []byte(data), "application/json")
			if err != nil {
				return err
			}
		}
		err = db.Put([]byte(domain.String()+"_status"), []byte(block))
		if err != nil {
			return err
		}
	}

	return nil
}
