package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

// connect to db
func GetDBConnection() *sql.DB {
	//supabase postgres database (password: dbpass2024!!!)
	var connStr string = "postgres://postgres.msebrbkfgjlrrqhadcoh:dbpass2024!!!@aws-0-us-west-1.pooler.supabase.com:5432/postgres"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
