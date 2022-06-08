package main

import (
	"GoAPI/structs"
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Welcome to Go")

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			fmt.Println("Request Method is GET")

			data := structs.Response{
				Code: http.StatusOK,
				Body: "Data",
			}

			json.NewEncoder(w).Encode(data)
		}
	})

	http.ListenAndServe("localhost:3000", mux)
}
