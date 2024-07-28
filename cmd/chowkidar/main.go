package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"runtime/debug"
	"strconv"
	"strings"
	"time"

	"github.com/koki-develop/go-fzf"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v2"
	"go.mills.io/bitcask/v2"
)

type Config struct {
	ID              int      `json:"id"`
	Server          string   `json:"server"`
	ISP             string   `json:"isp"`
	City            string   `json:"city"`
	Latitude        float64  `json:"lat"`
	Longitude       float64  `json:"lon"`
	CheckCategories []string `json:"categories"`
	TestFrequency   int      `json:"testFrequency"`
}

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

func main() {
	conf := os.Getenv("XDG_CONFIG_HOME") + "/chowkidar.json"
	if conf == "/chowkidar.json" {
		conf = "./chowkidar.json"
	}
	data := os.Getenv("XDG_DATA_HOME") + "/chowkidar.db"
	if data == "/chowkidar.db" {
		data = "./chowkidar.db"
	}
	app := &cli.App{
		Name:    "Internet Chowkidar",
		Usage:   "Run the chowkidar daemon to report blocked sites",
		Version: Version(),
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "config",
				Aliases: []string{"c"},
				Value:   conf,
				Usage:   "config file to read from",
			},
			&cli.StringFlag{
				Name:    "database",
				Aliases: []string{"d"},
				Value:   data,
				Usage:   "config file to read from",
			},
			&cli.BoolFlag{
				Name:  "verbose",
				Usage: "Change verbosity of the application",
			},
		},
		Action: func(cCtx *cli.Context) error {
			configFile, err := os.ReadFile(cCtx.String("config"))
			if err != nil {
				return cli.Exit("Config "+cCtx.String("config")+" does not exist. Please run `setup` or create the config manually", 1)
			}

			var config Config
			err = json.Unmarshal(configFile, &config)
			if err != nil {
				return cli.Exit("Config "+cCtx.String("config")+" is invalid, please run `setup` or fix the config manually", 1)
			}

			if validateConfig(config) == false {
				return cli.Exit("Config "+cCtx.String("config")+" is incomplete, please run `setup` or fix the config manually", 1)
			}
			log.Println("Starting the daemon")

			// Figure out frequency based on the mode numbers
			var duration time.Duration
			switch config.TestFrequency {
			case 1:
				duration = 1 * time.Hour
			case 2:
				duration = 6 * time.Hour
			case 3:
				duration = 12 * time.Hour
			case 4:
				duration = 24 * time.Hour
			case 5:
				duration = 24 * 7 * time.Hour
			}

			// If it wasn't run before acc to DB, run it now
			db, _ := bitcask.Open(cCtx.String("database"))
			defer db.Close()
			val, err := db.Get([]byte("lastRun"))
			if err != nil {
				if strings.Contains(err.Error(), "key not found") {
					fetchAndRun(config,db)
				} else {
					return err
				}
			}
			if string(val) == "" {
					fetchAndRun(config, db)
			} else {
				// If it was run before, check if it has been more than `duration` since it happened
				timeInt, err := strconv.ParseInt(string(val), 10, 64)
				if err != nil {
					return cli.Exit("Database has invalid data or is corrupted.", 1)
				}

				timeLast := time.Unix(timeInt, 0)
				durationRemain := time.Now().Sub(timeLast)
				if durationRemain >= duration {
					fetchAndRun(config, db)
				} else {
					// Wait till duration is complete and then run
					time.Tick(durationRemain)
					fetchAndRun(config, db)
				}
			}
			// Do the periodic ones based on the determined duration
			for range time.Tick(duration) {
				fetchAndRun(config, db)
			}
			return nil
		},
		Commands: []*cli.Command{
			{
				Name:  "setup",
				Usage: "setup the chowkidar daemon (interactive)",
				Action: func(cCtx *cli.Context) error {
					fmt.Println("Welcome to the Internet Chowkidar Setup Wizard")
					fmt.Println("This wizard will help you setup your Internet Chowkidar Daemon")

					vars := Config{}

					fmt.Println("Enter which server you want to use (defaults to https://inetc.kat.directory): ")
					fmt.Scanln(&vars.Server)
					if vars.Server == "" {
						vars.Server = "https://inetc.kat.directory"
					}

					if validateServer(vars.Server) == false {
						return cli.Exit("Server "+vars.Server+" is invalid.", 1)
					}

					ipInfoOut, err := getRequest("http://ipinfo.io")
					if err != nil {
						return cli.Exit("Unable to retrieve ISP details from IPInfo.io", 1)
					}
					if !gjson.Valid(ipInfoOut) {
						return cli.Exit("Unable to parse ISP details from IPInfo.io", 1)
					}

					vars.ISP = gjson.Get(ipInfoOut, "org").String()

					cityValid := false
					for cityValid == false {
						fmt.Println("Enter your city (leave blank for autodetect): ")
						fmt.Scanln(&vars.City)
						if vars.City == "" {
							vars.City = gjson.Get(ipInfoOut, "city").String()
							cityValid = true
						}
						vars.City, vars.Latitude, vars.Longitude, cityValid = validateCity(vars.City)
						if cityValid == false {
							log.Println("Invalid City. Please try again...")
						}
					}

					categoriesOut, err := getRequest(vars.Server + "/categories")
					if err != nil {
						return cli.Exit("Unable to fetch categories from server", 1)
					}
					if !gjson.Valid(categoriesOut) {
						return cli.Exit("Unable to parse categories from server", 1)
					}

					gjsonArr := gjson.Get(categoriesOut, "#.name").Array()

					f, err := fzf.New(fzf.WithNoLimit(true), fzf.WithPrompt("Enter categories you want to use (tab to select; enter to submit): "))
					if err != nil {
						fmt.Println("Unable to initialize go-fzf, printing all categories as is:")
						for i := range gjsonArr {
							fmt.Printf(gjsonArr[i].String() + ", ")
						}
						var catStr string
						fmt.Println("Input the categories you want to select, separated by spaces: ")
						fmt.Scanln(&catStr)

						vars.CheckCategories = strings.Split(catStr, " ")
					} else {
						catInt, err := f.Find(gjsonArr, func(i int) string { return gjsonArr[i].String() })
						if err != nil {
							log.Fatal(err)
						}
						for _, i := range catInt {
							vars.CheckCategories = append(vars.CheckCategories, gjsonArr[i].String())
						}
					}
					if vars.CheckCategories == nil {
						vars.CheckCategories = []string{"all"}
					}

					fmt.Println("Enter frequency you want to check for blocked sites:")
					fmt.Println("1: hourly")
					fmt.Println("2: 4 times a day (every 6 hours)")
					fmt.Println("3: 2 times a day (every 12 hours)")
					fmt.Println("4: once a day")
					fmt.Println("5: once a week")
					fmt.Scanln(&vars.TestFrequency)

					type ISPStruct struct {
						Latitude  float64 `json:"latitude"`
						Longitude float64 `json:"longitude"`
						Name      string  `json:"name"`
					}
					ispStruct := ISPStruct{Latitude: vars.Latitude, Longitude: vars.Longitude, Name: vars.ISP}
					data, err := json.Marshal(ispStruct)
					ISPOut, err := postRequest(vars.Server+"/isps", []byte(data), "application/json")
					if err != nil {
						return cli.Exit("Unable to receive unique node ID from server", 1)
					}
					if !gjson.Valid(ISPOut) {
						return cli.Exit("Unable to parse server response for ISP creation request", 1)
					}

					vars.ID, err = strconv.Atoi(gjson.Get(ISPOut, "id").String())
					if err != nil {
						return cli.Exit("Unable to retrieve the ID"+err.Error(), 1)
					}

					data, err = json.Marshal(vars)
					if err != nil {
						return cli.Exit("Unable to marshal into json: "+err.Error(), 1)
					}

					err = os.WriteFile(cCtx.String("config"), data, 0644)
					if err != nil {
						return cli.Exit("Unable to write config: "+err.Error(), 1)
					}

					return nil
				},
			},
		},
	}

	app.Suggest = true

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
