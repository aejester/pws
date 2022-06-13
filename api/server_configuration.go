package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type ServerConfiguration struct {
	APIKey     string `json:"api_key"`
	Coordinate struct {
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	} `json:"coordinate"`
	Subcriptions []struct {
		ID          int    `json:"id"`
		Name        string `json:"name"`
		Subscribed  bool   `json:"subscribed"`
		DisplayName string `json:"display_name,omitempty"`
	} `json:"subcriptions"`
	UpdateFrequency  int  `json:"update_frequency"`
	MaxDailyRequests int  `json:"max_daily_requests"`
	MaximizeFetch    bool `json:"maximize_fetch"`
}

func LoadServerConfiguration(config ...string) *ServerConfiguration {
	var serverConfiguration ServerConfiguration
	var configName string

	if len(config) == 0 {
		configName = "./config.json"
	} else if len(config) == 1 {
		configName = config[0]
	} else {
		panic("too many server configuration files were provided")
	}

	f, fErr := os.Open(configName)

	if fErr != nil {
		fmt.Println(fErr)
	}

	defer f.Close()

	byteArr, bErr := ioutil.ReadAll(f)
	if bErr != nil {
		fmt.Println(bErr)
	}

	json.Unmarshal(byteArr, &serverConfiguration)

	return &serverConfiguration
}
