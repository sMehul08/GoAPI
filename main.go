package main

import (
	"GoAPI/controller"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Welcome to Go")

	mux := controller.Register()
	http.ListenAndServe(":3000", mux)
}
