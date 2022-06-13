package services

import (
	"fmt"
	"pws/api"
	"time"
)

func ServiceScheduler(name string, data *api.WeatherResponse, refresh int, service func(data *api.WeatherResponse)) {
	ticker := time.NewTicker(time.Duration(refresh) * time.Minute)
	quit := make(chan int)
	fmt.Println("starting thread " + name)
	service(data)
	go func() {
		for {
			select {
			case <-ticker.C:
				service(data)
			case <-quit:
				fmt.Println("stopping thread " + name)
				ticker.Stop()
				return
			}
		}
	}()
}
