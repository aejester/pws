package data_server

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

type Routes struct {
	DB *sql.DB
}

func (routes Routes) AlertsHandler(w http.ResponseWriter, req *http.Request) {
	api_key := req.URL.Query().Get("api_key")

	encoded, err := json.Marshal(GetAllPushedAlertsMatching(api_key, routes.DB))
	if err != nil {
		fmt.Println(err)
	}

	w.Write([]byte(string(encoded)))
}
