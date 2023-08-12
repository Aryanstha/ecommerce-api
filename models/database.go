package models

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var DB *sql.DB

func ConnectDatabase() {
	db, err := sql.Open("mysql", "root@tcp(Localhost:3307/ecommerce")
	if err != nil {
		log.Fatal(err)
	}
	DB = db

}
