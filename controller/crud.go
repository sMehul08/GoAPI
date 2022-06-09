package controller

import (
	"GoAPI/model"
	"GoAPI/view"
	"encoding/json"
	"net/http"
)

func todo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			data := view.PostRequest{}
			json.NewDecoder(r.Body).Decode(&data)
			if err := model.CreateToDo(data.Name, data.Todo); err != nil {
				w.Write([]byte("Something went wrong..."))
				return
			}
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(data)
		} else if r.Method == http.MethodGet {
			// data, err := model.ReadAll()
			name := r.URL.Query().Get("name")
			data, err := model.ReadByName(name)
			if err != nil {
				w.Write([]byte(err.Error()))
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(&data)
		} else if r.Method == http.MethodDelete {

			name := r.URL.Query().Get("name")
			if err := model.DeleteToDo(name); err != nil {
				w.Write([]byte("Something wrong"))
				return
			}

			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(struct {
				Status string `json:"status"`
			}{"Item Deleted"})
		}
	}
}
