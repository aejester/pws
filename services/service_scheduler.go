package services

import (
	"fmt"
	"pws/api"
	"time"
)

func ServiceScheduler[T any](name string, data *api.WeatherResponse, refresh int, service func(data *api.WeatherResponse, prev T) T) {
	ticker := time.NewTicker(time.Duration(refresh) * time.Minute)
	quit := make(chan int)
	fmt.Println("starting thread " + name)

	var prev T

	prev = service(data, prev)

	go func() {
		for {
			select {
			case <-ticker.C:
				prev = service(data, prev)
			case <-quit:
				fmt.Println("stopping thread " + name)
				ticker.Stop()
				return
			}
		}
	}()
}
