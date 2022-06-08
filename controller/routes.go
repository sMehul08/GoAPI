package controller

import (
	"net/http"
)

func Register() *http.ServeMux {
	mux := http.NewServeMux()
	// mux.HandleFunc("/index", data())
	mux.HandleFunc("/create", todo())
	return mux
}
