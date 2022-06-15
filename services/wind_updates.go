package services

import (
	"math"
	"pws/api"
	"strconv"
)

func WindUpdates(weatherResponse *api.WeatherResponse, prevData int, output chan string) int {

	if math.Abs(float64(prevData)-weatherResponse.Current.WindSpeed) == 5 {
		// TODO: Implement notifications
		output <- "Wind speed/direction has changed. " + strconv.Itoa(int(weatherResponse.Current.WindSpeed)) + " MPH heading " + strconv.Itoa(int(weatherResponse.Current.WindDeg)) + " degrees."
	}

	return prevData
}
