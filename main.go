package main

import (
	"GoAPI/controller"
	"GoAPI/model"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Welcome to Go")

	mux := controller.Register()

	db := model.Connect()
	fmt.Println("Serving....")
	defer db.Close()

	http.ListenAndServe(":3000", mux)
}
