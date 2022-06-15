package services

import (
	"pws/api"
	"time"
)

func ServiceScheduler[T any](name string, data *api.WeatherResponse, refresh int, service func(data *api.WeatherResponse, prev T, out chan string) T, output chan string) {
	ticker := time.NewTicker(time.Duration(refresh) * time.Minute)
	quit := make(chan int)
	msg := "starting thread " + name
	output <- msg

	var prev T

	prev = service(data, prev, output)

	go func() {
		for {
			select {
			case <-ticker.C:
				prev = service(data, prev, output)
			case <-quit:
				msg = "stopping thread " + name
				output <- msg
				ticker.Stop()
				return
			}
		}
	}()
}
