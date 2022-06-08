package controller

import (
	"GoAPI/view"
	"encoding/json"
	"fmt"
	"net/http"
)

func data() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			fmt.Println("Request Method is GET")

			data := view.Response{
				Code: http.StatusOK,
				Body: "Data",
			}

			json.NewEncoder(w).Encode(data)
		}
	}
}
