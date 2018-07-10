package main

import (
	"encoding/json"

	lib "github.com/niklucky/go-lib"
	"github.com/syndicatedb/vodka"
)

/*
Config - service configuration
*/
type Config struct {
	HTTPConfig vodka.HTTPConfig `json:"http_server"`
}

/*
NewConfig - config constructors
*/
func NewConfig(configFileName string) (Config, error) {
	config := Config{}

	fileData, err := lib.ReadFile(configFileName)
	if err != nil {
		return config, err
	}
	err = json.Unmarshal(fileData, &config)
	if err != nil {
		return config, err
	}
	return config, nil
}
