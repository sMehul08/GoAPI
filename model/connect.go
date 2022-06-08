package model

import (
	"database/sql"
	"fmt"
	"log"
)

var con *sql.DB

func Connect() *sql.DB {
	dsn := "root:root@tcp(localhost:3316)/Golang"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to Database.")
	con = db
	return db
}
