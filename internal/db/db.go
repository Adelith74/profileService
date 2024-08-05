package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func init() {
	connStr := "user=pqgotest dbname=pqgotest sslmode=verify-full"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	db.Query("SELECT * FROM public.Data")
}
