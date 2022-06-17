package data_server

import (
	"database/sql"
	"fmt"
)

type PushedAlert struct {
	Date    int    `json:"date"`
	Content string `json:"content"`
	APIKey  string `json:"api_key"`
}

func GetAllPushedAlertsMatching(api_key string, db *sql.DB) *[]PushedAlert {
	statement := "select * from alerts where api_key=\"" + api_key + "\""
	rows, err := db.Query(statement)
	if err != nil {
		fmt.Println(err)
	}

	pushedAlerts := make([]PushedAlert, 0)

	for rows.Next() {
		pushedAlert := PushedAlert{}
		err = rows.Scan(&pushedAlert.Date, &pushedAlert.Content, &pushedAlert.APIKey)

		if err != nil {
			fmt.Println(err)
		}

		pushedAlerts = append(pushedAlerts, pushedAlert)
	}

	return &pushedAlerts
}
