package cleaner

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Settings struct {
	Server   string `json:server`
	Port     int    `json:port`
	Username string `json:username`
	Password string `json:password`
	To       string `json:to`
}

func ReadConf() Settings {
	configFile, err := os.Open("config.json")
	if err != nil {
		log.Fatal("opening config file", err.Error())
	}

	var settings Settings

	jsonParser := json.NewDecoder(configFile)
	if err = jsonParser.Decode(&settings); err != nil {
		log.Fatal("parsing config file", err.Error())
	}

	fmt.Printf("%s %d %s %s\n", settings.Server, settings.Port, settings.Username, settings.Password)
	return settings
}
