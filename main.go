package main

import (
	"database/sql"
	"fmt"
	"pws/data_server"
	"pws/push_notifications"
)

func main() {
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		fmt.Println(err)
	}

	go push_notifications.PushNotificationServer(db)
	go data_server.DataServer(db)

	select {}
}
