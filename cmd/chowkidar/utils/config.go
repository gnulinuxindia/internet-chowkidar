package utils

import (
	"encoding/json"
	"errors"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

type Config struct {
	ISPID           int      `json:"id"`
	Server          string   `json:"server"`
	ISP             string   `json:"isp"`
	City            string   `json:"city"`
	Latitude        float64  `json:"lat"`
	Longitude       float64  `json:"lon"`
	ClientID        int      `json:"client_id"`
	CheckCategories []string `json:"categories"`
	TestFrequency   int      `json:"testFrequency"`
}

func DefaultStoragePath() (config, data string) {
	confDir, err := os.UserConfigDir()
	if err != nil {
		config = "./config.json"
		data = "./chowkidar.db"
	} else {
		err := os.Mkdir(confDir+"/chowkidar", 0600)
		if err != nil && !os.IsExist(err) {
			log.Fatal(err)
		}
		config = confDir + "/chowkidar/config.json"
		data = confDir + "/chowkidar/chowkidar.db"
	}
	return config, data
}

func FindConfigData(cCtx *cli.Context) (Config, error) {
	configFile, err := os.ReadFile(cCtx.String("config"))
	if err != nil {
		return Config{}, errors.New("Config " + cCtx.String("config") + " does not exist. Please run `setup` or create the config manually")
	}

	Debug("Config file confirmed to exist.", cCtx)

	var config Config
	err = json.Unmarshal(configFile, &config)
	if err != nil {
		return Config{}, errors.New("Config " + cCtx.String("config") + " is invalid, please run `setup` or fix the config manually")
	}
	Debug("Config file is JSON and has been unmarshalled", cCtx)

	if ValidateConfig(config) == false {
		return Config{}, errors.New("Config " + cCtx.String("config") + " is incomplete, please run `setup` or fix the config manually")
	}
	Debug("Config file has been validated", cCtx)

	return config, nil
}
