package services

import (
	"fmt"
	"pws/api"
	"strconv"
)

func PrecipitationUpdates(weatherResponse *api.WeatherResponse, prevData *[]api.Minute) *[]api.Minute {

	oldTimeUntil, willRainOld := minutesUntilRain(prevData)
	newTimeUntil, willRain := minutesUntilRain(&weatherResponse.Minutely)

	if oldTimeUntil != newTimeUntil {
		var message string
		if willRain {
			message = "Rain expected to start in " + strconv.Itoa(newTimeUntil) + " minutes."
		} else if willRainOld {
			message = "It is no longer going to rain."
		}

		// TODO: Implement push notifications
		fmt.Println(message)
	}

	return &weatherResponse.Minutely
}

func minutesUntilRain(minutes *[]api.Minute) (int, bool) {
	if minutes != nil {
		count := 0
		didFindRain := false
		for _, min := range *minutes {
			if min.Precipitation > 0 {
				count++
				didFindRain = true
			}
		}
		return count, didFindRain
	}
	return 0, false
}
