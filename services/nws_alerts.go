package services

import (
	"fmt"
	"pws/api"
)

func NWSAlerts(weatherResponse *api.WeatherResponse, previousData *[]api.Alert) *[]api.Alert {

	alertable := api.AlertArrayDifference(previousData, &weatherResponse.Alerts)

	if len(*alertable) != 0 {
		for _, alert := range *alertable {
			// TODO: Implement Push Notification Service
			fmt.Println(alert)

		}
	}

	return &weatherResponse.Alerts

}
