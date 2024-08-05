package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"profileService/internal/db/model"

	"github.com/jackc/pgx/v5"
)

type DataBase struct {
	Db *pgx.Conn
}

func InitDb() *DataBase {
	urlExample := "postgres://postgres:root@localhost:5432/profiles"
	db, err := pgx.Connect(context.Background(), urlExample)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	db_init := `CREATE TABLE IF NOT EXISTS profiles (
		id SERIAL PRIMARY KEY,
		sex BOOLEAN NOT NULL,
		birth_year SERIAL NOT NULL,
		first_name VARCHAR (50) NOT NULL,
		second_name VARCHAR (50),
		last_name VARCHAR (50) NOT NULL,
		video_id BIGINT	NOT NULL
	);`
	_, err = db.Exec(context.Background(), db_init)
	if err != nil {
		log.Fatal(err)
	}
	return &DataBase{Db: db}
}

func (db *DataBase) GetProfiles() ([]model.Profile, error) {
	profiles := []model.Profile{}
	rows, err := db.Db.Query(context.Background(), "SELECT * FROM profiles")
	if err != nil {
		return profiles, err
	}
	for rows.Next() {
		profile := model.Profile{}
		rows.Scan(&profile.Id, &profile.Sex, &profile.BirthYear, &profile.FirstName, &profile.SecondName, &profile.LastName, &profile.VideoID)
		profiles = append(profiles, profile)
	}
	return profiles, nil
}

func (db *DataBase) GetProfile(id int) (model.Profile, error) {
	return model.Profile{}, nil
}

func (db *DataBase) InsertProfile() {

}

func (db *DataBase) DeleteProfile(id int) {

}
