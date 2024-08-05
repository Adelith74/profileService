package db

type Profile struct {
	Id         int64  `db:"id"`
	Sex        bool   `db:"sex"`
	BirthYear  int    `db:"birth_year"`
	FirstName  string `db:"first_name"`
	SecondName string `db:"second_name"`
	LastName   string `db:"last_name"`
	VideoID    int64  `db:"video_id"`
}
