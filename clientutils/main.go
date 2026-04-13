package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"runtime/debug"
	"strconv"
	"strings"
	"time"
	"path/filepath"
	"context"

	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v2"
	"go.mills.io/bitcask/v2"
	"github.com/ooni/probe-engine/pkg/engine"
	"github.com/ooni/probe-engine/pkg/experiment/webconnectivity"
	"github.com/ooni/probe-engine/pkg/kvstore"
	"github.com/ooni/probe-engine/pkg/model"
	"errors"
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
func Run(configPath string, db *bitcask.Bitcask, updateConf *bool, stopSync *bool) error {
	config, err := FindConfigData(configPath)
	if err != nil {
		return cli.Exit(err.Error(), 1)
	}

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

	// create tempdir for KVStore shit
	stateDir, err := os.MkdirTemp("", "internet-chowkidar")
	if err != nil {
		return err
	}
	defer os.RemoveAll(stateDir)

	kvStore, err := kvstore.NewFS(filepath.Join(stateDir, "kvstore"))
	if err != nil {
		return err
	}

	// Create OONI Session
	ctx := context.Background()
	sess, err := engine.NewSession(ctx, engine.SessionConfig{
		KVStore:         kvStore,
		Logger:          model.DiscardLogger,
		SoftwareName:    "InternetChowkidar",
		SoftwareVersion: "1.0.0",
		TempDir:         stateDir,
	})
	if err != nil {
		return err
	}
	defer sess.Close()

	// Bootstrap: discover probe services and fetch the test-helper list.
	if err := sess.MaybeLookupBackendsContext(ctx); err != nil {
		return err
	}
	if err := sess.MaybeLookupLocationContext(ctx); err != nil {
		return err
	}

	if sess.ProbeASNString() != strings.Split(config.ISP, " ")[0]{
		return errors.New("Test is being run from an ISP different from the ISP configured. It will start running when you go back home")
	}

	measurer := webconnectivity.NewExperimentMeasurer(webconnectivity.Config{})

	for i, domain := range domains {
		fmt.Println(domain.String())
		measurement, err := testWebsite(domain.String(), ctx, sess, measurer)
		if err != nil {
			return err
		}

		// perform type assertion on TestKeys
		tk, asserted := measurement.TestKeys.(*webconnectivity.TestKeys)
		if !asserted {
			return errors.New("Unable to type assert TestKeys")
		}
		var blocked bool
		if tk.Blocking == nil {
			blocked = false
			fmt.Println("Site " + domain.String() + " reported as not blocked")
		}
		switch t := tk.Blocking.(type) {
		case bool:
			if t {
				blocked = true
				fmt.Println("Site " + domain.String() + " reported as blocked")
				fmt.Println(*tk.BlockingReason)
			} else {
				blocked = false
				fmt.Println("Site " + domain.String() + " reported as not blocked")
			}
		case *string:
			blocked = true
			fmt.Println("Site " + domain.String() + " reported as blocked")
			fmt.Println(*tk.BlockingReason)
		default:
			return errors.New("Could not parse block status")
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
	fmt.Println("done!")
	return nil
}

func testWebsite(url string, ctx context.Context, sess *engine.Session, measurer model.ExperimentMeasurer) (*model.Measurement, error){
	measurement := &model.Measurement{
		DataFormatVersion:    "0.2.0",
		Input:                model.MeasurementInput(url),
		MeasurementStartTime: time.Now().UTC().Format("2006-01-02 15:04:05"),
		ProbeASN:             sess.ProbeASNString(),
		ProbeCC:              sess.ProbeCC(),
		ProbeIP:              "127.0.0.1", // never store the real IP
		ReportID:             "", // We aren't uploading
		SoftwareName:    "InternetChowkidar",
		SoftwareVersion: "1.0.0",
		TestName:             measurer.ExperimentName(),
		TestVersion:          measurer.ExperimentVersion(),
		TestStartTime:        time.Now().UTC().Format("2006-01-02 15:04:05"),
	}

	args := &model.ExperimentArgs{
		Callbacks:   model.NewPrinterCallbacks(model.DiscardLogger),
		Measurement: measurement,
		Session:     sess,
	}

	if err := measurer.Run(ctx, args); err != nil {
		return &model.Measurement{}, err
	}
	return measurement, nil
}
