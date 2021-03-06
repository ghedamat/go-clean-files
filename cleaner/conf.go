package cleaner

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

type Settings struct {
	Server          string `json:server`
	Port            int    `json:port`
	Username        string `json:username`
	Password        string `json:password`
	To              string `json:to`
	MailThreshold   int    `json:mailThreshold`
	DeleteThreshold int    `json:deleteThreshold`
	Subject         string `json:subject`
}

func (conf Settings) ToAddresses() []string {
	parts := strings.Split(conf.To, ",")
	for i, v := range parts {
		parts[i] = strings.TrimSpace(v)
	}
	return parts
}

func openConf(path string) *os.File {
	path, _ = filepath.Abs(path)
	configFile, err := os.Open(path)
	if err != nil {
		log.Fatal("opening config file ", err.Error())
	}
	return configFile
}

func ReadConf(configFile io.Reader) Settings {

	var settings Settings

	jsonParser := json.NewDecoder(configFile)
	if err := jsonParser.Decode(&settings); err != nil {
		log.Fatal("parsing config file", err.Error())
	}

	return settings
}

func ParseConf() Settings {
	usr, _ := user.Current()
	dir := usr.HomeDir
	path := dir + "/.go-clean-filesrc"
	return ReadConf(openConf(path))
}
