package controller

import (
	"GoAPI/model"
	"net/http"
)

func todo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			if err := model.CreateToDo(); err != nil {
				w.Write([]byte("Something went wrong..."))
				return
			}
		}
	}
}
