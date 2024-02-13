package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

// connect to db
func GetDBConnection() *sql.DB {
	var connStr string = "user=postgres password=dbpass dbname=SWE_II_ECommerceDB sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
