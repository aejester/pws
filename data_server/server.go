package data_server

import (
	"database/sql"
	"net/http"
)

func DataServer(db *sql.DB) {
	http.HandleFunc("/alerts", nil)

	http.ListenAndServe(":8000", nil)
}
