package config

import (
	"database/sql"
	"log"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("mysql", "root:@tcp(localhost:3306)/project_point_of_sale_api_go")
	if err != nil {
		log.Fatalf("Error connecting to database: %s", err.Error())
	}

	if err := DB.Ping(); err != nil {
		log.Fatalf("Error pinging databse: %s", err.Error())
	}
}
