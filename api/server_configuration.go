package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type ServerConfiguration struct {
	Id               int     `json:"id"`
	APIKey           string  `json:"api_key"`
	Latitude         float64 `json:"latitude"`
	Longitude        float64 `json:"longitude"`
	Subcriptions     string  `json:"subcriptions"`
	UpdateFrequency  int     `json:"update_frequency"`
	MaxDailyRequests int     `json:"max_daily_requests"`
	MaximizeFetch    int     `json:"maximize_fetch"`
	DeviceToken      string  `json:"device_token"`
}

func AddNewServerConfigutation(db *sql.DB, config ServerConfiguration) {
	statement, _ := db.Prepare("insert into stations (api_key, latitude, longitude, subscriptions, update_frequency, max_daily_requests, maximize_fetch, device_token) values (?, ?, ?, ?, ?, ?, ?, ?)")
	statement.Exec(config.APIKey, config.Latitude, config.Longitude, config.Subcriptions, config.UpdateFrequency, config.MaxDailyRequests, config.MaximizeFetch, config.DeviceToken)

	defer statement.Close()
}

func GetAllServerConfigurations(db *sql.DB) *[]ServerConfiguration {
	rows, err := db.Query("select * from stations")
	if err != nil {
		fmt.Println(err)
	}

	configs := make([]ServerConfiguration, 0)

	for rows.Next() {
		config := ServerConfiguration{}
		err = rows.Scan(&config.Id, &config.APIKey, &config.Latitude, &config.Longitude, &config.Subcriptions, &config.UpdateFrequency, &config.MaxDailyRequests, &config.MaximizeFetch, &config.DeviceToken)

		if err != nil {
			fmt.Println(err)
		}

		configs = append(configs, config)
	}

	return &configs
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
