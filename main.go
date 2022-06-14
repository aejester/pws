package main

import (
	"database/sql"
	"fmt"
	"pws/api"
	"pws/services"
	"strconv"
	"strings"
	"time"

	_ "github.com/mattn/go-sqlite3"
	// "github.com/sideshow/apns2"
)

func main() {
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		fmt.Println(err)
	}

	configs := api.GetAllServerConfigurations(db)
	serviceNames := *api.GetAllServices(db)

	for _, serverConfiguration := range *configs {

		subscriptions := strings.Split(serverConfiguration.Subcriptions, ",")

		if serverConfiguration.APIKey != "" && len([]rune(serverConfiguration.APIKey)) == 32 {
			fmt.Println()
			fmt.Println("configuration file loaded!")
			fmt.Println()
			fmt.Print("lat/lon: ")
			fmt.Println(serverConfiguration.Latitude, serverConfiguration.Longitude)
			fmt.Println("subscriptions: ")

			for _, sub := range subscriptions {
				id, err := strconv.Atoi(sub)
				if err != nil {
					fmt.Println(err)
				}
				fmt.Println(serviceNames[id-1])
			}

			fmt.Print("update_frequency: ")
			fmt.Print(serverConfiguration.UpdateFrequency)

			if serverConfiguration.MaximizeFetch == 1 {
				fmt.Println(" (maximize fetch is enabled, will not use this number)")
			} else {
				fmt.Println()
			}

			fmt.Print("max_daily_requests: ")
			fmt.Println(serverConfiguration.MaxDailyRequests)
			fmt.Println()

			var url string
			var currentWeatherData api.WeatherResponse

			lat := fmt.Sprintf("%f", serverConfiguration.Latitude)
			lon := fmt.Sprintf("%f", serverConfiguration.Longitude)

			url = "https://api.openweathermap.org/data/3.0/onecall?lat=" + lat + "&lon=" + lon + "&units=imperial&appid=" + serverConfiguration.APIKey

			ticker := time.NewTicker(5 * time.Minute)
			quit := make(chan struct{})

			fmt.Println("starting weather data fetch cycle")
			currentWeatherData = *api.FetchWeatherData(url)

			go func() {
				for {
					select {
					case <-ticker.C:
						currentWeatherData = *api.FetchWeatherData(url)
					case <-quit:
						fmt.Println("stopping weather data fetch cycle")
						ticker.Stop()
						return
					}
				}
			}()

			for _, sub := range subscriptions {
				id, err := strconv.Atoi(sub)
				if err != nil {
					fmt.Println(err)
				}

				service := serviceNames[id-1]

				if service.Name == "precipitation_updates" {
					go services.ServiceScheduler(service.Name, &currentWeatherData, 30, services.PrecipitationUpdates)
				} else if service.Name == "hurricane_updates" {
					go services.ServiceScheduler(service.Name, &currentWeatherData, 300, services.HurricaneUpdates)
				} else if service.Name == "wind_updates" {
					go services.ServiceScheduler(service.Name, &currentWeatherData, 10, services.WindUpdates)
				} else if service.Name == "nws_alerts" {
					go services.ServiceScheduler(service.Name, &currentWeatherData, 5, services.NWSAlerts)
				}
			}
		} else {
			fmt.Println("api key is invalid/not detected")
		}
	}

	select {}
}
