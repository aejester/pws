package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"pws/data_server"
)

func (routes *Routes) GetAlerts(w http.ResponseWriter, req *http.Request) {
	api_key := req.URL.Query().Get("api_key")

	encoded, err := json.Marshal(data_server.GetAllPushedAlertsMatching(api_key, routes.DB))
	if err != nil {
		fmt.Println(err)
	}

	w.Write([]byte(string(encoded)))
}
