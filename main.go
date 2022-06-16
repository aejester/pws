package main

import (
	"database/sql"
	"fmt"
	"pws/push_notifications"
)

func main() {
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		fmt.Println(err)
	}

	go push_notifications.PushNotificationServer(db)

	select {}
}
