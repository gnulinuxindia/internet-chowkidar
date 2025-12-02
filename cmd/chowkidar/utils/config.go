package utils

import (
	"encoding/json"
	"errors"
	"log"
	"os"

	"go.mills.io/bitcask/v2"
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

func FindConfigData(configPath string) (Config, error) {
	configFile, err := os.ReadFile(configPath)
	if err != nil {
		return Config{}, errors.New("Config " + configPath + " does not exist. Please run `setup` or create the config manually")
	}

	var config Config
	err = json.Unmarshal(configFile, &config)
	if err != nil {
		return Config{}, errors.New("Config " + configPath + " is invalid, please run `setup` or fix the config manually")
	}

	if ValidateConfig(config) == false {
		return Config{}, errors.New("Config " + configPath + " is incomplete, please run `setup` or fix the config manually")
	}

	return config, nil
}
func FindDatabase(databasePath string) (*bitcask.Bitcask, error) {
	return bitcask.Open(databasePath)
}
