package services

import (
	"fmt"
	"pws/api"
)

func NWSAlerts(weatherResponse *api.WeatherResponse, previousData *[]api.Alert, output chan string) *[]api.Alert {

	fmt.Println(weatherResponse.Alerts)

	alertable := api.AlertArrayDifference(previousData, &weatherResponse.Alerts)

	if len(*alertable) != 0 {
		for _, alert := range *alertable {
			// TODO: Implement Push Notification Service
			output <- alert.Description

		}
	}

	return &weatherResponse.Alerts

}
