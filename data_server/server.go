package data_server

import (
	"database/sql"
	"net/http"
)

func DataServer(db *sql.DB) {
	routes := Routes{DB: db}

	http.HandleFunc("/alerts", routes.AlertsHandler)

	http.ListenAndServe(":8000", nil)
}
