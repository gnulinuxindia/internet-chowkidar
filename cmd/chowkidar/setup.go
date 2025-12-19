package main

import (
	"encoding/json"
	"fmt"
	"math/rand/v2"
	"os"
	"strings"

	"github.com/gnulinuxindia/internet-chowkidar/cmd/chowkidar/utils"
	"github.com/koki-develop/go-fzf"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v2"
)

func Setup(cCtx *cli.Context) error {
	fmt.Println("Welcome to the Internet Chowkidar Setup Wizard")
	fmt.Println("This wizard will help you setup your Internet Chowkidar Daemon")

	vars := utils.Config{}

	serverValid := false
	for !serverValid {
		fmt.Println("Enter which server you want to use (defaults to https://api.inet.watch): ")
		fmt.Scanln(&vars.Server)
		if vars.Server == "" {
			vars.Server = "https://api.inet.watch"
		}

		if utils.ValidateServer(vars.Server) {
			serverValid = true
		} else {
			vars.Server = ""
			fmt.Println("Invalid server, try again")
		}
	}

	cityValid := false
	for !cityValid {
		fmt.Println("Enter your city: ")
		fmt.Scanln(&vars.City)
		vars.City, vars.Latitude, vars.Longitude, cityValid = utils.ValidateCity(vars.City)
		if !cityValid {
			fmt.Println("Invalid city, try again")
		}
	}

	categoriesOut, err := utils.GetRequest(vars.Server + "/categories")
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
			fmt.Print(gjsonArr[i].String() + ", ")
		}
		fmt.Println()
		var catStr string
		fmt.Println("Input the categories you want to select, separated by spaces: ")
		fmt.Scanln(&catStr)

		vars.CheckCategories = strings.Split(catStr, " ")
	} else {
		catInt, err := f.Find(gjsonArr, func(i int) string { return gjsonArr[i].String() })
		if err != nil {
			return cli.Exit("Failed to select categories: "+err.Error(), 1)
		}
		for _, i := range catInt {
			vars.CheckCategories = append(vars.CheckCategories, gjsonArr[i].String())
		}
	}
	if vars.CheckCategories == nil {
		vars.CheckCategories = []string{"all"}
	}

	freqOptions := []string{
		"Hourly",
		"4 times a day (every 6 hours)",
		"2 times a day (every 12 hours)",
		"Once a day",
		"Once a week",
		"Custom",
	}
	freqValues := []int{1, 6, 12, 24, 24 * 7, 0}

	f2, err := fzf.New(fzf.WithPrompt("Select frequency for checking blocked sites: "))
	if err != nil {
		fmt.Println("Unable to initialize go-fzf, falling back to manual input")
		fmt.Println("Enter frequency you want to check for blocked sites:")
		for i, opt := range freqOptions {
			fmt.Printf("%d: %s\n", i+1, opt)
		}

		freqValid := false
		for !freqValid {
			var freq string
			fmt.Scanln(&freq)
			freqValid = true
			switch freq {
			case "1":
				vars.TestFrequency = 1
			case "2":
				vars.TestFrequency = 6
			case "3":
				vars.TestFrequency = 12
			case "4":
				vars.TestFrequency = 24
			case "5":
				vars.TestFrequency = 24 * 7
			case "6":
				fmt.Println("Input the duration between each check (in hours): ")
				fmt.Scanln(&vars.TestFrequency)
			default:
				fmt.Println("Invalid choice, try again")
				freqValid = false
			}
		}
	} else {
		idxs, err := f2.Find(freqOptions, func(i int) string { return freqOptions[i] })
		if err != nil {
			return cli.Exit("Failed to select frequency: "+err.Error(), 1)
		}
		if len(idxs) == 0 {
			return cli.Exit("No frequency selected", 1)
		}

		selectedIdx := idxs[0]
		if freqValues[selectedIdx] == 0 {
			fmt.Println("Input the duration between each check (in hours): ")
			fmt.Scanln(&vars.TestFrequency)
		} else {
			vars.TestFrequency = freqValues[selectedIdx]
		}
	}

	// TODO: Find a better alternative to using ipinfo.io
	ipInfoOut, err := utils.GetRequest("http://ipinfo.io")
	if err != nil {
		return cli.Exit("Unable to retrieve ISP details from IPInfo.io", 1)
	}
	
	fmt.Println(ipInfoOut)
	
	if !gjson.Valid(ipInfoOut) {
		return cli.Exit("Unable to parse ISP details from IPInfo.io", 1)
	}

	vars.ISP = gjson.Get(ipInfoOut, "org").String()

	type ISPStruct struct {
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
		Name      string  `json:"name"`
		City      string  `json:"city"`
	}
	ispStruct := ISPStruct{Latitude: vars.Latitude, Longitude: vars.Longitude, Name: vars.ISP, City: vars.City}
	data, err := json.Marshal(ispStruct)
	ISPOut, err := utils.PostRequest(vars.Server+"/isps", []byte(data), "application/json")
	if err != nil {
		return cli.Exit("Unable to receive unique node ID from server", 1)
	}
	if !gjson.Valid(ISPOut) {
		return cli.Exit("Unable to parse server response for ISP creation request", 1)
	}

	vars.ISPID = int(gjson.Get(ISPOut, "id").Int())

	vars.ClientID = rand.IntN(999999999)

	data, err = json.Marshal(vars)
	if err != nil {
		return cli.Exit("Unable to marshal into json: "+err.Error(), 1)
	}

	err = os.WriteFile(cCtx.String("config"), data, 0644)
	if err != nil {
		return cli.Exit("Unable to write config: "+err.Error(), 1)
	}

	return nil
}
