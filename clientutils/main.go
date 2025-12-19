package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"runtime/debug"
	"strconv"
	"strings"
	"time"

	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v2"
	"go.mills.io/bitcask/v2"
)

func Version() string {
	if info, ok := debug.ReadBuildInfo(); ok {
		for _, setting := range info.Settings {
			if setting.Key == "vcs.revision" {
				return setting.Value[:8]
			}
		}
	}
	return "unknown, please build with Go 1.13+ or use Git"
}

// UpdateConf and StopSync are optional since only GUI makes use of them
func Run(configPath string, databasePath string, updateConf *bool, stopSync *bool) error {
	config, err := FindConfigData(configPath)
	if err != nil {
		return cli.Exit(err.Error(), 1)
	}

	db, err := FindDatabase(databasePath)
	if err != nil {
		return cli.Exit(err.Error(), 1)
	}
	defer db.Close()

	// Figure out frequency based on the mode numbers
	duration := time.Duration(config.TestFrequency) * time.Hour

	if db == nil {
		fmt.Println("Database not found.")
		os.Exit(1)
	}

	// If it wasn't run before acc to DB, run it now
	val, err := db.Get([]byte("lastRun"))
	if err != nil {
		if strings.Contains(err.Error(), "key not found") {
			runErr := FetchAndRun(config, db)
			if runErr != nil {
				return cli.Exit(runErr.Error(), 1)
			}
		} else {
			return err
		}
	} else if string(val) == "" {
		runErr := FetchAndRun(config, db)
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
			runErr := FetchAndRun(config, db)
			if runErr != nil {
				return cli.Exit(runErr.Error(), 1)
			}
		} else {
			// Wait till duration is complete and then run
			time.Sleep(durationRemain)
			runErr := FetchAndRun(config, db)
			if runErr != nil {
				return cli.Exit(runErr.Error(), 1)
			}
		}
	}

	// Do the periodic ones based on the determined duration
	for {
		for range time.Tick(duration) {
			if !*stopSync {
				if *updateConf {
					config, err = FindConfigData(configPath)
					if err != nil {
						return cli.Exit(err.Error(), 1)
					}
					// If frequency got changed, reset the clock
					duration2 := time.Duration(config.TestFrequency) * time.Hour
					if duration != duration2 {
						duration = duration2
						break
					}
					*updateConf = false
				}
				runErr := FetchAndRun(config, db)
				if runErr != nil {
					return cli.Exit(runErr.Error(), 1)
				}
			}
		}
	}
}
func FetchAndRun(config Config, db *bitcask.Bitcask) error {
	fmt.Println("Running a check cycle")
	err := db.Put([]byte("lastRun"), []byte(strconv.Itoa(int(time.Now().Unix()))))
	if err != nil {
		return err
	}

	catStr := strings.Join(config.CheckCategories, ",")
	sitesList, err := GetRequest(config.Server + "/sites?category=" + catStr)
	if err != nil {
		return err
	}
	if !gjson.Valid(sitesList) {
		return err
	}
	domains := gjson.Get(sitesList, "#.ping_url").Array()

	for i, domain := range domains {
		_, err := GetRequest(domain.String())
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
			_, err = PostRequest(config.Server+"/blocks", []byte(data), "application/json")
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
