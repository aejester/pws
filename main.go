package main

import "pws/push_notifications"

func main() {
	go push_notifications.PushNotificationServer()

	select {}
}
