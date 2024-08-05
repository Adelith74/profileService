package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func init() {
	connStr := "user=pqgotest dbname=pqgotest sslmode=verify-full"
	db_init := `CREATE TABLE IF NOT EXISTS profiles (
		id SERIAL PRIMARY KEY
		sex BOOLEAN NOT NULL
		birth_year SERIAL NOT NULL
		first_name VARCHAR (50) NOT NULL
		second_name VARCHAR (50)
		last_name VARCHAR (50) NOT NULL
		video_id BIGINT	NOT NULL
	)`
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec(db_init)
	if err != nil {
		log.Fatal(err)
	}
}
