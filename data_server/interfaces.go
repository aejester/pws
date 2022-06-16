package data_server

import (
	"database/sql"
	"fmt"
)

type PushedAlert struct {
	Date     int    `json:"date"`
	Category string `json:"category"`
	Content  string `json:"content"`
}

func GetAllPushedAlertsMatching(api_key string, db *sql.DB) *[]PushedAlert {
	rows, err := db.Query("select * from stations where api_key=`" + api_key + "`")
	if err != nil {
		fmt.Println(err)
	}

	pushedAlerts := make([]PushedAlert, 0)

	for rows.Next() {
		pushedAlert := PushedAlert{}
		err = rows.Scan(&pushedAlert.Date, &pushedAlert.Category, &pushedAlert.Content)

		if err != nil {
			fmt.Println(err)
		}

		pushedAlerts = append(pushedAlerts, pushedAlert)
	}

	return &pushedAlerts
}
