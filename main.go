package main

import (
	"fmt"
	"pws/api"
	"pws/services"
	"time"
	// github.com/sideshow/apns2
)

func main() {
	serverConfiguration := api.LoadServerConfiguration()

	if serverConfiguration.APIKey != "" && len([]rune(serverConfiguration.APIKey)) == 32 {
		fmt.Println()
		fmt.Println("configuration file loaded!")
		fmt.Println()
		fmt.Print("lat/lon: ")
		fmt.Println(serverConfiguration.Coordinate)
		fmt.Println("subscriptions: ")

		for _, sub := range serverConfiguration.Subcriptions {
			fmt.Print("\t" + sub.Name + ": ")
			fmt.Println(sub.Subscribed)
		}

		fmt.Print("update_frequency: ")
		fmt.Print(serverConfiguration.UpdateFrequency)

		if serverConfiguration.MaximizeFetch {
			fmt.Println(" (maximize fetch is enabled, will not use this number)")
		} else {
			fmt.Println()
		}

		fmt.Print("max_daily_requests: ")
		fmt.Println(serverConfiguration.MaxDailyRequests)
		fmt.Println()
	} else {
		panic("api key is invalid/not detected")
	}
	var url string
	var currentWeatherData api.WeatherResponse

	lat := fmt.Sprintf("%f", serverConfiguration.Coordinate.Latitude)
	lon := fmt.Sprintf("%f", serverConfiguration.Coordinate.Longitude)

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

	for _, service := range serverConfiguration.Subcriptions {
		if service.Subscribed {
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
	}

	select {}
}
