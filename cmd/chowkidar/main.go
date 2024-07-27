package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"runtime/debug"
	"strings"

	"github.com/koki-develop/go-fzf"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v2"
)

type Config struct {
	ID              string   `json:"id"`
	Server          string   `json:"server"`
	ISP             string   `json:"isp"`
	City            string   `json:"city"`
	Latitude        float64  `json:"lat"`
	Longitude       float64  `json:"lon"`
	CheckCategories []string `json:"categories"`
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
	app := &cli.App{
		Name:    "Internet Chowkidar",
		Usage:   "Run the chowkidar daemon to report blocked sites",
		Version: Version(),
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "config",
				Aliases: []string{"c"},
				Value:   os.Getenv("XDG_CONFIG_HOME") + "/chowkidar.json",
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

					gjsonArr := gjson.Get(categoriesOut, "categories").Array()

					f, err := fzf.New(fzf.WithNoLimit(true))
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

					vars.ID = gjson.Get(ISPOut, "id").String()

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