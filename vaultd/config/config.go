package config

import (
	"encoding/json"
	"io/ioutil"
)

const (
	defAddress       = "127.0.0.1"
	defPort          = 8080
	defDatabaseName  = "vault.db"
	defHMacSecret    = ":°F_WQEùwqeflpùwa.pelfùkepwfùw,koefopwkepfwv"
	defUsername      = "vault"
	defPassword      = "vault"
	defTokenDuration = 60
)

type Configuration struct {
	Address       string `json:"address"`
	Port          int    `json:"port"`
	Database      string `json:"database"`
	Secret        string `json:"secret"`
	Username      string `json:"username"`
	Password      string `json:"password"`
	TokenDuration int    `json:"token_duration"`
}

var Conf = Configuration{
	Address:       defAddress,
	Port:          defPort,
	Database:      defDatabaseName,
	Secret:        defHMacSecret,
	Username:      defUsername,
	Password:      defPassword,
	TokenDuration: defTokenDuration,
}

func Load(filename string) error {
	raw, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	return json.Unmarshal(raw, &Conf)
}