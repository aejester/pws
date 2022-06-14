package api

import (
	"database/sql"
	"fmt"
)

type Service struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	DisplayName string `json:"display_name"`
}

func GetAllServices(db *sql.DB) *[]Service {
	rows, err := db.Query("select * from services")
	if err != nil {
		fmt.Println(err)
	}

	services := make([]Service, 0)

	for rows.Next() {
		config := Service{}
		err = rows.Scan(&config.Id, &config.Name, &config.DisplayName)

		if err != nil {
			fmt.Println(err)
		}

		services = append(services, config)
	}

	return &services
}
